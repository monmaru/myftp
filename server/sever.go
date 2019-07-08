package server

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/ptypes"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/monmaru/myftp/proto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// Listen ...
func Listen(cfg Config) (func(), error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	var grpcOpts []grpc.ServerOption
	if cfg.Certificate != "" && cfg.Key != "" {
		grpcCreds, err := credentials.NewServerTLSFromFile(cfg.Certificate, cfg.Key)
		if err != nil {
			return nil, err
		}

		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	zapLogger := newFileRotateLogger(filepath.Join(cfg.LogDir, "ftp.log"))
	grpczap.ReplaceGrpcLogger(zapLogger)

	grpcOpts = append(grpcOpts, grpc.StreamInterceptor(
		grpcmiddleware.ChainStreamServer(
			grpczap.StreamServerInterceptor(zapLogger),
			grpcrecovery.StreamServerInterceptor(
				grpcrecovery.WithRecoveryHandler(func(p interface{}) (err error) {
					zapLogger.Error(fmt.Sprintf("panic: %+v\n", p))
					return status.Errorf(codes.Internal, "Unexpected error")
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
		if err := zapLogger.Sync(); err != nil {
			log.Println(err)
		}
	}, nil
}

type ftpServer struct {
	destDir string
	logger  *zap.Logger
}

// Upload ...
func (s *ftpServer) Upload(stream proto.Ftp_UploadServer) error {
	s.logger.Info("-------- Start Upload --------")
	defer s.logger.Info("-------- Finished Upload --------")
	var f *os.File
	defer func() {
		if f != nil {
			if err := f.Close(); err != nil {
				s.logger.Error(fmt.Sprintf("Failed close file: %v", err))
			}
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
			if err := stream.SendAndClose(&proto.UploadResponse{
				Message: "FileName is empty",
				Status:  proto.UploadResponse_FAILED,
			}); err != nil {
				s.logger.Error(fmt.Sprintf("Failed SendAndClose: %v", err))
			}
			return errors.New("FileName is empty")
		}

		if f == nil {
			f, err = initFile(filepath.Join(s.destDir, filepath.Base(fileData.FileName)))
			if err != nil {
				if err := stream.SendAndClose(&proto.UploadResponse{
					Message: fmt.Sprintf("Failed to create file : %s", fileData.FileName),
					Status:  proto.UploadResponse_FAILED,
				}); err != nil {
					s.logger.Error(fmt.Sprintf("Failed SendAndClose: %v", err))
				}
				return err
			}
		}

		if _, err := f.Write(fileData.Content); err != nil {
			if err := stream.SendAndClose(&proto.UploadResponse{
				Message: fmt.Sprintf("Failed to write file : %s", fileData.FileName),
				Status:  proto.UploadResponse_FAILED,
			}); err != nil {
				s.logger.Error(fmt.Sprintf("Failed SendAndClose: %v", err))
			}
			return err
		}
	}

	if err := stream.SendAndClose(&proto.UploadResponse{
		Message: "Upload received with success",
		Status:  proto.UploadResponse_OK,
	}); err != nil {
		return err
	}

	return nil
}

func (s *ftpServer) ListFiles(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	s.logger.Info("-------- Start ListFiles --------")
	defer s.logger.Info("-------- Finished ListFiles --------")

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
			Mode:      uint32(fi.Mode()),
		})
	}
	return &proto.ListResponse{Files: files}, nil
}

func (s *ftpServer) Download(r *proto.DownloadRequest, stream proto.Ftp_DownloadServer) error {
	s.logger.Info("-------- Start Download --------")
	defer s.logger.Info("-------- Finished Download --------")

	f, err := os.Open(filepath.Join(s.destDir, r.Name))
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			s.logger.Error(fmt.Sprintf("Failed close file: %v", err))
		}
	}()

	var buf [4096 * 1000]byte
	for {
		n, err := f.Read(buf[:])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		if err := stream.Send(&proto.DownloadResponse{
			Content: buf[:n],
		}); err != nil {
			return err
		}
	}

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
