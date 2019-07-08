package client

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/hashicorp/go-multierror"
	"github.com/monmaru/myftp/proto"
	"github.com/pkg/errors"
)

type Downloader interface {
	DownloadFiles(ctx context.Context) error
}

type downloaderImpl struct {
	dir         string
	parallelism int
	cli         proto.FtpClient
	pool        *pb.Pool
}

func NewDownloader(cli proto.FtpClient, dir string, parallelism int) Downloader {
	return &downloaderImpl{
		cli:         cli,
		dir:         dir,
		parallelism: parallelism,
	}
}

func (d *downloaderImpl) DownloadFiles(ctx context.Context) error {
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

	d.pool = pool
	var errs error
	for err := range d.downloadFilesInParallel(ctx) {

		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	return errs
}

func (d *downloaderImpl) downloadFilesInParallel(ctx context.Context) <-chan error {
	ret := make(chan error)
	sem := make(chan struct{}, d.parallelism)

	go func() {
		defer func() {
			close(ret)
			close(sem)
		}()

		resp, err := d.cli.ListFiles(ctx, &proto.ListRequest{})
		if err != nil {
			ret <- err
			return
		}

		wg := sync.WaitGroup{}
		for _, fi := range resp.Files {
			sem <- struct{}{}
			wg.Add(1)

			go func(f *proto.FileInfo) {
				defer func() {
					wg.Done()
					<-sem
				}()

				ret <- d.do(ctx, f)
			}(fi)
		}
		wg.Wait()
	}()

	return ret
}

func (d *downloaderImpl) do(ctx context.Context, fi *proto.FileInfo) error {
	fPath := filepath.Join(d.dir, filepath.FromSlash(fi.Name))

	if os.FileMode(fi.Mode).IsDir() {
		if err := os.MkdirAll(fPath, os.FileMode(fi.Mode)); err != nil {
			return errors.Wrapf(err, "Failed MkdirAll: %s", fi.Name)
		}

		t := time.Unix(fi.UpdatedAt.Seconds, int64(fi.UpdatedAt.Nanos))
		if err := os.Chtimes(fPath, t, t); err != nil {
			log.Printf("%s: %v", fi.Name, err)
		}
		return nil
	}

	req := &proto.DownloadRequest{Name: fi.Name}
	stream, err := d.cli.Download(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "Failed DownloadRequest: %s", fi.Name)
	}

	f, err := os.Create(fPath)
	if err != nil {
		if err := stream.CloseSend(); err != nil {
			log.Printf("%s: %v", fi.Name, err)
		}
		return errors.Wrapf(err, "Failed Create file: %s", fi.Name)
	}

	bar := pb.New64(fi.Size).Postfix(" " + fi.Name)
	bar.Units = pb.U_BYTES
	d.pool.Add(bar)
	defer bar.Finish()

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("%s: %v", fi.Name, err)
			break
		}
		n, err := f.Write(res.Content)
		if err != nil {
			log.Printf("%s: %v", fi.Name, err)
			break
		}
		bar.Add64(int64(n))
	}

	if err := f.Close(); err != nil {
		log.Printf("%s: %v", fi.Name, err)
	}

	if err := stream.CloseSend(); err != nil {
		log.Printf("%s: %v", fi.Name, err)
	}

	if runtime.GOOS != "windows" {
		if err := os.Chmod(fPath, os.FileMode(fi.Mode)); err != nil {
			return errors.Wrapf(err, "Failed Chmod on windows: %s", fi.Name)
		}
	}

	t := time.Unix(fi.UpdatedAt.Seconds, int64(fi.UpdatedAt.Nanos))
	if err := os.Chtimes(fPath, t, t); err != nil {
		log.Printf("%s: %v", fi.Name, err)
	}

	return nil
}
