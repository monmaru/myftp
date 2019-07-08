package client

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/golang/protobuf/ptypes"
	"github.com/monmaru/myftp/proto"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	Address     string
	Certificate string
}

func New(address, certificate string) (*Client, error) {
	if address == "" {
		return nil, errors.New("Address must be specified")
	}

	return &Client{Address: address, Certificate: certificate}, nil
}

// Upload ...
func (c *Client) Upload(dir string, parallelism int) error {
	conn, err := c.makeGRPCConn()
	if err != nil {
		return errors.Wrap(err, "cannot connect")
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	ctx := context.Background()
	u := NewUploader(
		proto.NewFtpClient(conn),
		dir,
		parallelism)
	return u.UploadFiles(ctx)
}

// List ...
func (c *Client) List() error {
	conn, err := c.makeGRPCConn()
	if err != nil {
		return errors.Wrap(err, "cannot connect")
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()

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
func (c *Client) Download(dir string, parallelism int) error {
	conn, err := c.makeGRPCConn()
	if err != nil {
		return errors.Wrap(err, "cannot connect")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	d := NewDownloader(
		proto.NewFtpClient(conn),
		dir,
		parallelism)
	return d.DownloadFiles(context.Background())
}

func (c *Client) makeGRPCConn() (*grpc.ClientConn, error) {
	var options []grpc.DialOption
	if c.Certificate != "" {
		creds, err := credentials.NewClientTLSFromFile(c.Certificate, "")
		if err != nil {
			return nil, err
		}
		options = append(options, grpc.WithTransportCredentials(creds))
	} else {
		options = append(options, grpc.WithInsecure())
	}

	return grpc.Dial(c.Address, options...)
}
