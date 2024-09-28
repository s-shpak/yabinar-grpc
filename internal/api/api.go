package api

import (
	"fmt"
	"net"
	server "webinar-service/internal/api/v1"

	pb "webinar-service/internal/protos/v1/server_new"

	"google.golang.org/grpc"
)

func Serve() error {
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		return fmt.Errorf("failed to run gRPC server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDummyServer(grpcServer, server.NewServer())

	return grpcServer.Serve(lis)
}
