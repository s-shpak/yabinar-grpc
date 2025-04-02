package api

import (
	"fmt"
	"net"
	server "webinar-service/internal/api/v1"
	dummyV3 "webinar-service/internal/api/v3"

	pb "webinar-service/internal/protos/v1/server_new"
	pbV3Dummy "webinar-service/internal/protos/v3/dummy"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

func Serve() error {
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		return fmt.Errorf("failed to run gRPC server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDummyServer(grpcServer, server.NewServer())
	//pbV2Dummy.RegisterDummyServer(grpcServer, dummy.NewServer())
	pbV3Dummy.RegisterDummyServer(grpcServer, dummyV3.NewServer())

	reflection.Register(grpcServer)

	return grpcServer.Serve(lis)
}
