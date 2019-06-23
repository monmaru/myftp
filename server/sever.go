package server

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/ptypes"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/monmaru/myftp/proto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

// Listen ...
func Listen(cfg Config) (func(), error) {
	if err := cfg.validate(); err != nil {
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

	zapLogger := newFileRotateLogger(filepath.Join(cfg.LogDir, "ftp.log"))
	grpc_zap.ReplaceGrpcLogger(zapLogger)

	grpcOpts = append(grpcOpts, grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
					zapLogger.Error(fmt.Sprintf("panic: %+v\n", p))
					return grpc.Errorf(codes.Internal, "Unexpected error")
				}),
			),
		),
	))

	grpcServer := grpc.NewServer(grpcOpts...)

	ftpServer := &ftpServer{
		destDir: cfg.DestDir,
		logger:  zapLogger,
	}
	proto.RegisterFtpServer(grpcServer, ftpServer)

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return nil, err
	}

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
				goto END
			}
			return errors.Wrap(err, "Failed reading chunks from stream")
		}

		if len(fileData.FileName) == 0 {
			stream.SendAndClose(&proto.UploadResponse{
				Message: "FileName is empty",
				Status:  proto.UploadResponse_FAILED,
			})
			return errors.New("FileName is empty")
		}

		if f == nil {
			f, err = initFile(filepath.Join(s.destDir, filepath.Base(fileData.FileName)))
			if err != nil {
				stream.SendAndClose(&proto.UploadResponse{
					Message: fmt.Sprintf("Failed to create file : %s", fileData.FileName),
					Status:  proto.UploadResponse_FAILED,
				})
				return err
			}
		}

		if _, err := f.Write(fileData.Content); err != nil {
			stream.SendAndClose(&proto.UploadResponse{
				Message: fmt.Sprintf("Failed to write file : %s", fileData.FileName),
				Status:  proto.UploadResponse_FAILED,
			})
			return err
		}
	}

END:
	if err := stream.SendAndClose(&proto.UploadResponse{
		Message: "Upload received with success",
		Status:  proto.UploadResponse_OK,
	}); err != nil {
		return err
	}

	s.logger.Info("-------- Finished upload function --------")
	return nil
}

func (s *ftpServer) ListFiles(context.Context, *proto.ListRequest) (*proto.ListResponse, error) {
	fis, err := ioutil.ReadDir(s.destDir)
	if err != nil {
		return nil, err
	}

	var files []*proto.FileInfo
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}

		ts, err := ptypes.TimestampProto(fi.ModTime())
		if err != nil {
			return nil, err
		}

		files = append(files, &proto.FileInfo{
			Name:      fi.Name(),
			Size:      fi.Size(),
			UpdatedAt: ts,
			Type:      proto.FileInfo_FILE,
		})
	}
	return &proto.ListResponse{Files: files}, nil
}

func (s *ftpServer) Download(*proto.DownloadRequest, proto.Ftp_DownloadServer) error {
	return nil
}

func initFile(path string) (*os.File, error) {
	if exists(path) {
		return os.OpenFile(path, os.O_WRONLY, 0666)
	}
	return os.Create(path)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
