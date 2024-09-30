package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/resolver"

	pb "webinar-service/internal/protos/v1/server_old"
)

const (
	dummyScheme      = "dummy"
	dummyServiceName = "dummy.service"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) <= 1 {
		return fmt.Errorf("pass a command to run")
	}
	cmd := os.Args[1]

	ctx := context.Background()
	switch cmd {
	case "hello":
		if err := SayHello(ctx); err != nil {
			return fmt.Errorf("failed to say hello: %w", err)
		}
	default:
		return fmt.Errorf("unknown command \"%s\"", cmd)
	}
	return nil
}

func SayHello(ctx context.Context) error {
	client, err := initClient()
	if err != nil {
		return fmt.Errorf("failed to initialize a client: %w", err)
	}

	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Msg:      []string{"Hello", "world!"},
		ClientID: []int64{42},
	}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		return fmt.Errorf("failed to send a message to the server: %v", err)
	}

	log.Printf("response: \"%v\"", resp.Msg)
	return nil
}

type DummyClient struct {
	pb.DummyClient
	conn *grpc.ClientConn
}

func (dc *DummyClient) Close() {
	dc.conn.Close()
}

func initClient() (*DummyClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(fmt.Sprintf("%s:///%s", dummyScheme, dummyServiceName), opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new client: %w", err)
	}

	client := pb.NewDummyClient(conn)

	return &DummyClient{
		conn:        conn,
		DummyClient: client,
	}, err
}

type dummyResolverBuilder struct{}

func (*dummyResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	r := &dummyResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			dummyServiceName: {"localhost:8081"},
		},
	}
	r.start()
	return r, nil
}

func (*dummyResolverBuilder) Scheme() string { return dummyScheme }

type dummyResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *dummyResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint()]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*dummyResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (*dummyResolver) Close()                                {}

func init() {
	resolver.Register(&dummyResolverBuilder{})
}
