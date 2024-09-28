package v1

import (
	"context"
	"fmt"

	pb "webinar-service/internal/protos/v1/server_new"
)

type Server struct {
	pb.UnimplementedDummyServer
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: fmt.Sprintf("got: \"%s\"", req.Msg),
	}, nil
}
