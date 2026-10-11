package server

import (
	"context"
	"strings"
	"time"

	fixturev1 "github.com/Larry-Systems-Inc/larry-grpc-fixture/gen/fixturev1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type FixtureServer struct {
	fixturev1.UnimplementedFixtureServiceServer
}

func (s *FixtureServer) Unary(
	ctx context.Context,
	request *fixturev1.UnaryRequest,
) (*fixturev1.UnaryResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	incomingMetadata, _ := metadata.FromIncomingContext(ctx)

	requestID := "not-provided"

	if values := incomingMetadata.Get("x-client-request-id"); len(values) > 0 {
		requestID = values[0]
	}

	if err := grpc.SetHeader(
		ctx,
		metadata.Pairs(
			"x-fixture-service", "fixture",
			"x-fixture-request-id", requestID,
		),
	); err != nil {
		return nil, err
	}

	grpc.SetTrailer(
		ctx,
		metadata.Pairs(
			"x-fixture-result", "completed",
		),
	)

	message := request.GetMessage()

	if message == "" {
		message = "hello from Larry gRPC fixture"
	}

	return &fixturev1.UnaryResponse{
		Message:      message,
		ServerUnixMs: time.Now().UnixMilli(),
	}, nil
}

func (s *FixtureServer) Fail(
	ctx context.Context,
	request *fixturev1.FailureRequest,
) (*fixturev1.UnaryResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if request == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"failure request is required",
		)
	}

	message := request.GetMessage()

	if message == "" {
		message = "fixture error"
	}

	return nil, status.Error(
		failureCode(request.GetCode()),
		message,
	)
}

func failureCode(value string) codes.Code {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "INVALID_ARGUMENT":
		return codes.InvalidArgument
	case "UNAUTHENTICATED":
		return codes.Unauthenticated
	case "NOT_FOUND":
		return codes.NotFound
	case "DEADLINE_EXCEEDED":
		return codes.DeadlineExceeded
	case "INTERNAL":
		return codes.Internal
	default:
		return codes.Internal
	}
}
