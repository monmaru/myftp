package client

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/golang/protobuf/ptypes"
	"github.com/monmaru/myftp/proto"
	"github.com/olekukonko/tablewriter"
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
	u := NewUploader(
		proto.NewFtpClient(conn),
		cfg.SrcDir,
		cfg.Parallelism)
	return u.UploadFiles(ctx)
}

// List ...
func List(cfg Config) error {
	if err := cfg.validate(); err != nil {
		return err
	}

	conn, err := makeGRPCConn(cfg.Address, cfg.Certificate)
	if err != nil {
		return errors.Wrap(err, "cannot connect")
	}
	defer conn.Close()

	ctx := context.Background()
	cli := proto.NewFtpClient(conn)
	resp, err := cli.ListFiles(ctx, &proto.ListRequest{})
	if err != nil {
		return err
	}

	header := []string{
		"FileName",
		"Size",
		"UpdatedAt",
	}

	var data [][]string
	for _, fi := range resp.Files {
		updateAt, err := ptypes.Timestamp(fi.UpdatedAt)
		if err != nil {
			return err
		}

		data = append(data, []string{
			fi.Name,
			humanize.Bytes(uint64(fi.Size)),
			updateAt.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("2006/1/2 15:04:05"),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
	return nil
}

// Download ...
func Download(cfg Config, name string) error {
	if err := cfg.validate(); err != nil {
		return err
	}

	conn, err := makeGRPCConn(cfg.Address, cfg.Certificate)
	if err != nil {
		return errors.Wrap(err, "cannot connect")
	}
	defer conn.Close()

	ctx := context.Background()
	cli := proto.NewFtpClient(conn)

	ftpClient, err := cli.Download(ctx, &proto.DownloadRequest{Name: name})
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(cfg.SrcDir, name))
	if err != nil {
		return err
	}

	defer file.Close()

	for {
		resp, err := ftpClient.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		file.Write(resp.Content)
	}

	return nil
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
