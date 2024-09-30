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
	msg := transformMessage(req.Msg, req.GetTransformations())
	respMsg := fmt.Sprintf("got: \"%s\"", msg)
	if req.ClientId != nil {
		respMsg += fmt.Sprintf(", client ID: %s", req.ClientId.Id)
	}
	return &pbModel.HelloResponse{
		Msg: &pbModel.HelloMessage{
			Msg: respMsg,
		},
	}, nil
}

func transformMessage(msg *pbModel.HelloMessage, tfs []*pbModel.HelloTransformation) string {
	if msg == nil {
		return ""
	}
	innerMsg := msg.Msg
	for _, tf := range tfs {
		switch tf.Type {
		case pbModel.HelloTransformationType_HELLO_TRANSFOMRATION_TYPE_TO_UPPER:
			innerMsg = strings.ToUpper(innerMsg)
		case pbModel.HelloTransformationType_HELLO_TRANSFOMRATION_TYPE_REVERSE:
			innerMsg = reverse(innerMsg)
		}
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
