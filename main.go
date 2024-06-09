package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/charmbracelet/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

func main() {
	log.Info("Starting Lambda...")
	
	lambda.StartWithOptions(
		otellambda.InstrumentHandler(
			Handle,
		),
	)
	
	log.Info("Started Lambda.")
}


type Request struct {
	ID string `json:"id"`
}

type Response struct {
	ID string `json:"id"`
}

func Handle(
	ctx context.Context,
	req Request,
) (Response, error) {
	log.Info("Hello there", "request", req)

	return Response{ID: req.ID}, nil
}
