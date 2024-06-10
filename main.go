package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/charmbracelet/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"strings"
)

func main() {
	ctx := context.Background()

	log.Info("Starting Lambda...")

	exporter, err := newSpanExporter()
	if err != nil {
		panic(err)
	}
	defer func() {
		log.Info("Shutting down SpanExporter...")

		_ = exporter.Shutdown(ctx)
	}()

	tp, err := newTracerProvider(exporter)
	if err != nil {
		panic(err)
	}
	defer func() {
		log.Info("Force flushing spans...")
		_ = tp.ForceFlush(ctx)

		log.Info("Shutting down TracerProvider...")
		_ = tp.Shutdown(ctx)
	}()

	lambda.StartWithOptions(
		otellambda.InstrumentHandler(
			Handle,
			otellambda.WithTracerProvider(tp),
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

func newTracerProvider(
	exporter sdktrace.SpanExporter,
) (*sdktrace.TracerProvider, error) {
	res, err := resource.New(
		context.Background(),
		resource.WithFromEnv(),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create OTel resource: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(exporter)

	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSpanProcessor(bsp),
	), nil
}

func newSpanExporter() (sdktrace.SpanExporter, error) {
	//exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())

	var opts []otlptracehttp.Option

	endpoint := `http://localhost:4318`

	switch {
	case strings.HasPrefix(endpoint, "https://"):
		endpoint = endpoint[len("https://"):]
	case strings.HasPrefix(endpoint, "http://"):
		endpoint = endpoint[len("http://"):]
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	return otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			append(opts,
				otlptracehttp.WithEndpoint(endpoint),
			)...,
		),
	)
}
