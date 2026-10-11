package main

import (
	"context"
	"log"
	"time"

	fixturev1 "github.com/Larry-Systems-Inc/larry-grpc-fixture/gen/fixturev1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	connection, err := grpc.NewClient(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	defer connection.Close()

	client := fixturev1.NewFixtureServiceClient(connection)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	ctx = metadata.AppendToOutgoingContext(
		ctx,
		"x-client-request-id",
		"metadata-check",
	)

	var responseHeaders metadata.MD
	var responseTrailers metadata.MD

	response, err := client.Unary(
		ctx,
		&fixturev1.UnaryRequest{
			Message: "Hello from client",
		},
		grpc.Header(&responseHeaders),
		grpc.Trailer(&responseTrailers),
	)

	failureContext, failureCancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer failureCancel()

	_, failureErr := client.Fail(
		failureContext,
		&fixturev1.FailureRequest{
			Code:    "UNAUTHENTICATED",
			Message: "authentication required",
		},
	)

	if failureErr == nil {
		log.Fatal("expected grpc failure, but request succeeded")
	}

	log.Printf(
		"failure code=%s message=%q raw=%v",
		status.Code(failureErr),
		status.Convert(failureErr).Message(),
		failureErr,
	)

	if err != nil {
		log.Fatalf("unary request failed: %v", err)
	}

	log.Printf("response message=%q", response.GetMessage())
	log.Printf("response headers=%v", responseHeaders)
	log.Printf("response trailers=%v", responseTrailers)

}
