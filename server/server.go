package cmds

import (
	"cmds/proto"
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Code for an action
type Code = int32

// Action do something
type Action = func(s string) error

// Set is the set list for code and action
type Set = map[Code]Action

var (
	// ErrRegistered is the code has been registered
	ErrRegistered = errors.New("this code has beed registered")
)

// CMDServer interface
type CMDServer interface {
	//Send(code Code, param string) error
	server.CommandServiceServer
	Register(code Code, action Action, force bool) error
	Run(addr string, opts ...grpc.ServerOption)
}

type srvImpl struct {
	actSet Set
}

// InitCMDS initial a commander service
func InitCMDS(set Set) CMDServer {
	return &srvImpl{
		actSet: set,
	}
}

func (s *srvImpl) Send(ctx context.Context, req *server.Request) (*server.Response, error) {
	act := s.actSet[req.GetCode()]
	err := act(req.GetParam())
	if err != nil {
		return &server.Response{
			Code:        1,
			Description: err.Error(),
		}, nil
	}
	return &server.Response{
		Code:        0,
		Description: "Success",
	}, nil
}

func (s *srvImpl) Register(code Code, action Action, force bool) error {
	if force {
		s.actSet[code] = action
		return nil
	}

	_, ok := s.actSet[code]
	if ok {
		return ErrRegistered
	}

	s.actSet[code] = action
	return nil
}

func (s *srvImpl) Run(addr string, opts ...grpc.ServerOption) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
	log.Printf("listen at: %v", addr)
	ss := s
	gsrv := grpc.NewServer(opts...)
	server.RegisterCommandServiceServer(gsrv, ss)
	gsrv.Serve(lis)
}

// CreateCred is helper to create creadential for grpc server
func CreateCred(certFile, keyFile string) (grpc.ServerOption, error) {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return grpc.ServerOption(grpc.Creds(creds)), nil
}
