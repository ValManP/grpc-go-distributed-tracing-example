package main

import (
	"context"
	"contrib.go.opencensus.io/exporter/jaeger"
	"fmt"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"grpc-go-service/api"
	"log"
	"math"
	"net"
)

const (
	sqrAddress = "localhost:50052"
	port       = ":50051"
)

type server struct {
	api.UnimplementedCircleServer
}

func (s *server) Area(ctx context.Context, req *api.AreaRequest) (*api.AreaResponse, error) {
	fmt.Printf("calls Circle.Area with radius = %f\n", req.GetRadius())

	ctx, span := trace.StartSpan(ctx, "Area.Circle")
	span.AddAttributes(trace.Float64Attribute("radius", req.Radius))
	span.AddAttributes(trace.StringAttribute("lang", "golang"))
	defer span.End()

	sqrRadius, err := sqr(ctx, req.GetRadius())
	if err != nil {
		return nil, err
	}

	return &api.AreaResponse{Area: math.Pi * sqrRadius}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register jaeger exporter
	err = initExporter()
	if err != nil {
		log.Fatalf("can't register exporter: %v", err)
	}

	// create server
	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	api.RegisterCircleServer(s, &server{})
	fmt.Println("Circle Service started...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func sqr(ctx context.Context, value float64) (float64, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(sqrAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewMathClient(conn)

	r, err := c.Sqr(ctx, &api.SqrRequest{Value: value})
	if err != nil {
		log.Fatalf("could not call sqr: %v", err)
	}

	return r.GetResult(), nil
}

func initExporter() error {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: "http://localhost:14268/api/traces",
		Process: jaeger.Process{
			ServiceName: "circle-service",
		},
	})
	if err != nil {
		return fmt.Errorf("can't register jaeger: %w", err)
	}

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return nil
}
