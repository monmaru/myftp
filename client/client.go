package client

import (
	"context"

	"github.com/monmaru/myftp/client/uploader"
	"github.com/monmaru/myftp/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Upload ...
func Upload(cfg Config) error {
	if err := cfg.validate(); err != nil {
		return err
	}

	conn, err := makeGRPCConn(cfg.Address, cfg.Certificate)
	if err != nil {
		return errors.Wrap(err, "cannot connect")
	}
	defer conn.Close()

	ctx := context.Background()
	u := uploader.New(
		proto.NewFtpClient(conn),
		cfg.SrcDir,
		cfg.Parallelism)
	return u.UploadFiles(ctx)
}

func makeGRPCConn(address, certificate string) (*grpc.ClientConn, error) {
	options := []grpc.DialOption{}
	if certificate != "" {
		creds, err := credentials.NewClientTLSFromFile(certificate, "")
		if err != nil {
			return nil, err
		}
		options = append(options, grpc.WithTransportCredentials(creds))
	} else {
		options = append(options, grpc.WithInsecure())
	}

	return grpc.Dial(address, options...)
}
