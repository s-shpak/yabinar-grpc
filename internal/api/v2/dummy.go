package v2

import (
	"context"
	"fmt"
	"strings"

	pb "webinar-service/internal/protos/v2/dummy"
	pbModel "webinar-service/internal/protos/v2/dummy/model"
)

type Server struct {
	pb.UnimplementedDummyServer
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) SayHello(ctx context.Context, req *pbModel.HelloRequest) (*pbModel.HelloResponse, error) {
	msg := transformMessage(req.Msg, req.GetTransformation())
	return &pbModel.HelloResponse{
		Msg: fmt.Sprintf("got: \"%s\", client ID: %d", msg, req.ClientID),
	}, nil
}

func transformMessage(msg string, tf pbModel.HelloTransformation) string {
	switch tf {
	case pbModel.HelloTransformation_HELLO_TRANSFOMRATION_TO_UPPER:
		return strings.ToUpper(msg)
	case pbModel.HelloTransformation_HELLO_TRANSFOMRATION_REVERSE:
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
