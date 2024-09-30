package api

import (
	"fmt"
	"net"
	server "webinar-service/internal/api/v1"
	dummy "webinar-service/internal/api/v2"

	pb "webinar-service/internal/protos/v1/server_new"
	pbV2Dummy "webinar-service/internal/protos/v2/dummy"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() error {
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		return fmt.Errorf("failed to run gRPC server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDummyServer(grpcServer, server.NewServer())
	pbV2Dummy.RegisterDummyServer(grpcServer, dummy.NewServer())

	reflection.Register(grpcServer)

	return grpcServer.Serve(lis)
}
