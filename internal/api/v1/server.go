package v1

import (
	"context"
	"fmt"
	"strings"

	pb "webinar-service/internal/protos/v1/server_new"
)

type Server struct {
	pb.UnimplementedDummyServer
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := transformMessage(req.Msg, req.GetTransformation())
	return &pb.HelloResponse{
		Msg: fmt.Sprintf("got: \"%s\", client ID: %d", msg, req.ClientID),
	}, nil
}

func transformMessage(msg string, tf pb.HelloTransformation) string {
	switch tf {
	case pb.HelloTransformation_HELLO_TRANSFOMRATION_TO_UPPER:
		return strings.ToUpper(msg)
	case pb.HelloTransformation_HELLO_TRANSFOMRATION_REVERSE:
		return reverse(msg)
	}
	return msg
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
