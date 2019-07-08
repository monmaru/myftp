package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/hashicorp/go-multierror"
	"github.com/monmaru/myftp/proto"
	"github.com/pkg/errors"
	"gopkg.in/cheggaaa/pb.v1"
)

type Uploader interface {
	UploadFiles(ctx context.Context) error
}

type uploaderImpl struct {
	dir         string
	parallelism int
	cli         proto.FtpClient
	pool        *pb.Pool
}

//NewUploader returns Uploader
func NewUploader(cli proto.FtpClient, dir string, parallelism int) Uploader {
	return &uploaderImpl{
		cli:         cli,
		dir:         dir,
		parallelism: parallelism,
	}
}

func (u *uploaderImpl) UploadFiles(ctx context.Context) error {
	pool, err := pb.StartPool()
	if err != nil {
		return err
	}

	defer func() {
		pool.RefreshRate = 500 * time.Millisecond
		if err := pool.Stop(); err != nil {
			log.Println(err)
		}
	}()

	u.pool = pool
	var errs error
	for err := range u.uploadFilesInParallel(ctx) {
		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	return errs
}

func (u *uploaderImpl) uploadFilesInParallel(ctx context.Context) <-chan error {
	ret := make(chan error)
	sem := make(chan struct{}, u.parallelism)

	go func() {
		defer func() {
			close(ret)
			close(sem)
		}()

		fis, err := ioutil.ReadDir(u.dir)
		if err != nil {
			ret <- err
			return
		}

		wg := sync.WaitGroup{}
		for _, fi := range fis {
			if fi.IsDir() {
				continue
			}

			sem <- struct{}{}
			wg.Add(1)

			go func(s string) {
				defer func() {
					wg.Done()
					<-sem
				}()

				ret <- u.do(ctx, s)
			}(filepath.Join(u.dir, fi.Name()))
		}
		wg.Wait()
	}()

	return ret
}

func (u *uploaderImpl) do(ctx context.Context, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	stream, err := u.cli.Upload(ctx, grpcretry.WithMax(3))
	if err != nil {
		return err
	}
	defer func() {
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	bar := pb.New64(stat.Size()).Postfix(" " + filepath.Base(f.Name())).SetUnits(pb.U_BYTES)
	u.pool.Add(bar)

	buf := make([]byte, 64*1024 /* 64KiB */)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			bar.FinishPrint(fmt.Sprintf("Failed to read bytes from %s", path))
			return err
		}

		if err := stream.Send(&proto.UploadRequest{
			Content:  buf[:n],
			FileName: path,
		}); err != nil {
			bar.FinishPrint(fmt.Sprintf("Failed to send chunk via stream : %s", path))
			return err
		}

		bar.Add64(int64(n))
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		bar.FinishPrint(fmt.Sprintf("Failed uploading file : %v", err))
		return err
	}

	if resp.Status != proto.UploadResponse_OK {
		bar.FinishPrint(fmt.Sprintf("Failed uploading %s : %s", path, resp.Message))
		return errors.New(resp.Message)
	}

	bar.Finish()
	return nil
}
