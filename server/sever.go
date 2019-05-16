package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"

	"github.com/monmaru/myftp/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Start ...
func Start(cfg Config) (func(), error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen on  %d", cfg.Address)
	}

	grpcOpts := []grpc.ServerOption{}
	if cfg.Certificate != "" && cfg.Key != "" {
		grpcCreds, err := credentials.NewServerTLSFromFile(cfg.Certificate, cfg.Key)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"failed to create tls grpc server using cert %s and key %s",
				cfg.Certificate,
				cfg.Key)
		}

		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	grpcServer := grpc.NewServer(grpcOpts...)
	proto.RegisterFtpServer(grpcServer, &FtpServer{DestDir: cfg.DestDir})

	if err := grpcServer.Serve(listener); err != nil {
		return nil, errors.Wrap(err, "errored listening for grpc connections")
	}

	return func() {
		grpcServer.Stop()
	}, nil
}

// FtpServer ...
type FtpServer struct {
	DestDir string
}

// Upload ...
func (s *FtpServer) Upload(stream proto.Ftp_UploadServer) error {
	fmt.Println("-------- Start upload function --------")
	var f *os.File
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	for {
		fileData, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return errors.Wrap(err, "failed reading chunks from stream")
		}

		if len(fileData.FileName) == 0 {
			stream.SendAndClose(&proto.UploadResponse{
				Message: "FileName is empty",
				Status:  proto.UploadStatus_FAILED,
			})
			return errors.New("FileName is empty")
		}

		if f == nil {
			f, err = initFile(path.Join(s.DestDir, filepath.Base(fileData.FileName)))
			if err != nil {
				stream.SendAndClose(&proto.UploadResponse{
					Message: fmt.Sprintf("Failed to create file : %s", fileData.FileName),
					Status:  proto.UploadStatus_FAILED,
				})
				return err
			}
		}

		if _, err := f.Write(fileData.Content); err != nil {
			stream.SendAndClose(&proto.UploadResponse{
				Message: fmt.Sprintf("Failed to write file : %s", fileData.FileName),
				Status:  proto.UploadStatus_FAILED,
			})
			return err
		}
	}

	if err := stream.SendAndClose(&proto.UploadResponse{
		Message: "Upload received with success",
		Status:  proto.UploadStatus_OK,
	}); err != nil {
		return err
	}

	fmt.Println("-------- Finished upload function --------")
	return nil
}

func initFile(filePath string) (*os.File, error) {
	if exists(filePath) {
		return os.OpenFile(filePath, os.O_WRONLY, 0666)
	}
	return os.Create(filePath)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
