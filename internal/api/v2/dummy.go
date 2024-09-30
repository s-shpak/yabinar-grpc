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
	respMsg := fmt.Sprintf("got: \"%s\"", msg)
	if req.ClientId != nil {
		respMsg += fmt.Sprintf(", client ID: %d", req.ClientId.Id)
	}
	return &pbModel.HelloResponse{
		Msg: &pbModel.HelloMessage{
			Msg: respMsg,
		},
	}, nil
}

func transformMessage(msg *pbModel.HelloMessage, tf pbModel.HelloTransformation) string {
	if msg == nil {
		return ""
	}
	innerMsg := msg.Msg
	switch tf {
	case pbModel.HelloTransformation_HELLO_TRANSFOMRATION_TO_UPPER:
		return strings.ToUpper(innerMsg)
	case pbModel.HelloTransformation_HELLO_TRANSFOMRATION_REVERSE:
		return reverse(innerMsg)
	}
	return innerMsg
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
