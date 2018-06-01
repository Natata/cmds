package client

import (
	server "cmds/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Client cmds client
type Client struct {
	addr   string
	client server.CommandServiceClient
}

// InitClient initial client
func InitClient(addr string, opts ...grpc.DialOption) (*Client, error) {
	client := &Client{
		addr: addr,
	}
	if len(opts) == 0 {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	client.client = server.NewCommandServiceClient(conn)
	return client, nil
}

func (c *Client) Send(code int32, param string) error {
	r, err := c.client.Send(context.Background(), &server.Request{
		Code:  code,
		Param: param,
	})

	if err != nil {
		return err
	}

	if r.Code != 0 {
		return fmt.Errorf("code: %v, description: %v", r.Code, r.Description)
	}

	return nil
}

// CreateCred create credential dial option
func CreateCred(certFile string, serverName string) (grpc.DialOption, error) {
	creds, err := credentials.NewClientTLSFromFile(certFile, serverName)
	if err != nil {
		return nil, err
	}
	return grpc.WithTransportCredentials(creds), nil
}
