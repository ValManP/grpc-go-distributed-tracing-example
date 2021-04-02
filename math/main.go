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
	"net"
)

const (
	port = ":50052"
)

type server struct {
	api.UnimplementedMathServer
}

func (s *server) Sqr(ctx context.Context, req *api.SqrRequest) (*api.SqrResponse, error) {
	fmt.Printf("calls Math.Sqr with value = %f\n", req.Value)
	ctx, span := trace.StartSpan(ctx, "Math.Sqr")
	span.AddAttributes(trace.StringAttribute("lang", "golang"))
	span.AddAttributes(trace.Float64Attribute("value", req.Value))
	defer span.End()
	return &api.SqrResponse{Result: req.Value * req.Value}, nil
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

	// start server
	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	api.RegisterMathServer(s, &server{})
	fmt.Println("Math Service started...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initExporter() error {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: "http://localhost:14268/api/traces",
		Process: jaeger.Process{
			ServiceName: "math-service",
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
