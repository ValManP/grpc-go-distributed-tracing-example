package main

import (
	"bufio"
	"context"
	"contrib.go.opencensus.io/exporter/jaeger"
	"fmt"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"grpc-go-service/api"
	"log"
	"os"
	"strconv"
)

const (
	circleAddress = "localhost:50051"
)

func main() {
	// set up a connection to the server.
	conn, err := grpc.Dial(circleAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// register jaeger exporter
	err = initExporter()
	if err != nil {
		log.Fatalf("can't register exporter: %v", err)
	}

	c := api.NewCircleClient(conn)

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _, err := stdin.ReadLine()
		if err != nil {
			fmt.Printf("failed to read a line in: %v", err)
			return
		}

		radius, err := strconv.ParseFloat(string(line), 64)
		if err != nil {
			fmt.Printf("failed to convert value to double: %v", err)
			return
		}

		fmt.Println("Call Circle.Area...")
		ctx, span := trace.StartSpan(context.Background(), "ClientSpan")
		span.AddAttributes(trace.StringAttribute("lang", "golang"))
		resp, err := c.Area(ctx, &api.AreaRequest{Radius: radius})
		if err != nil {
			span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
			fmt.Printf("gRPC client got error from server: %v", err)
		} else {
			fmt.Printf("Area: %f\n", resp.GetArea())
		}
		span.End()
	}
}

func initExporter() error {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: "http://localhost:14268/api/traces",
		Process: jaeger.Process{
			ServiceName: "client-service",
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
