package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/monmaru/myftp/proto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Listen ...
func Listen(cfg Config) (func(), error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return nil, err
	}

	grpcOpts := []grpc.ServerOption{}
	if cfg.Certificate != "" && cfg.Key != "" {
		grpcCreds, err := credentials.NewServerTLSFromFile(cfg.Certificate, cfg.Key)
		if err != nil {
			return nil, err
		}

		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	grpc_zap.ReplaceGrpcLogger(zapLogger)
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_zap.StreamServerInterceptor(zapLogger),
		),
	)

	ftpServer := &ftpServer{
		destDir: cfg.DestDir,
		logger:  zapLogger,
	}
	proto.RegisterFtpServer(grpcServer, ftpServer)

	if err := grpcServer.Serve(listener); err != nil {
		return nil, err
	}

	return func() {
		grpcServer.Stop()
		zapLogger.Sync()
	}, nil
}

type ftpServer struct {
	destDir string
	logger  *zap.Logger
}

// Upload ...
func (s *ftpServer) Upload(stream proto.Ftp_UploadServer) error {
	s.logger.Info("-------- Start upload function --------")
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
			return errors.Wrap(err, "Failed reading chunks from stream")
		}

		if len(fileData.FileName) == 0 {
			stream.SendAndClose(&proto.UploadResponse{
				Message: "FileName is empty",
				Status:  proto.UploadStatus_FAILED,
			})
			return errors.New("FileName is empty")
		}

		if f == nil {
			f, err = initFile(path.Join(s.destDir, filepath.Base(fileData.FileName)))
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

	s.logger.Info("-------- Finished upload function --------")
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
