package v2

import (
	"context"
	"fmt"
	"strings"

	pb "webinar-service/internal/protos/v2/dummy"
	pbModel "webinar-service/internal/protos/v2/dummy/model"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
			Msg: &wrapperspb.StringValue{
				Value: respMsg,
			},
		},
	}, nil
}

func transformMessage(msg *pbModel.HelloMessage, tfs []*pbModel.HelloTransformation) string {
	if msg == nil || msg.Msg == nil {
		return ""
	}
	innerMsg := msg.Msg.Value
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

func (srv *Server) GetSomethingFromDB(ctx context.Context, req *pbModel.GetSomethingFromDBRequest) (*pbModel.GetSomethingFromDBResponse, error) {
	var ct *pbModel.ContinuationTokenInternals
	if req.Ct != nil && req.Ct.Token != nil {
		m := &pbModel.ContinuationTokenInternals{}
		proto.Unmarshal(req.Ct.Token.Value, m)
		ct = m
	}

	// business logic goes here
	// ...

	ct = &pbModel.ContinuationTokenInternals{
		Ts: timestamppb.Now(),
	}
	ctMarshalled, err := proto.Marshal(ct)
	if err != nil {
		panic(err)
	}
	return &pbModel.GetSomethingFromDBResponse{
		Ct: &pbModel.ContinuationToken{
			Token: wrapperspb.Bytes(ctMarshalled),
		},
	}, nil
}
