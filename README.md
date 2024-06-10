# aws-lambda-go-otel-poc

Proof of concept using [aws-lambda-go][aws-lambda-go] and [OpenTelemetry][otel].

[aws-lambda-go]: https://github.com/aws/aws-lambda-go
[otel]: https://opentelemetry.io/

## Getting started

### Step 0: Prerequisites

All you need is [Docker][docker] and [pkgx][pkgx].

[docker]: https://www.docker.com/
[pkgx]: https://pkgx.sh/

### Step 1: Run everything

```shell
pkgx task@latest run
```

### Step 2: Invoking the Lambda

#### (With the CLI)

```shell
pkgx awslocal@latest \
  lambda invoke \
    --function-name my-lambda \
    --cli-binary-format raw-in-base64-out \
    --payload '{"body": "{\"id\": \"10\"}" }' \
    output.txt
```

#### (With HTTP)

```shell
pkgx http \
  http://localhost:4566/2015-03-31/functions/my-lambda/invocations \
  id="4"
```

### Step 3: Viewing the trace

Open the [Jaeger UI][jaeger-ui].

[jaeger-ui]: http://localhost:16686

## Miscellaneous

### Remaining Questions

- **Trace export**: `traces export: Post "http://localhost:4318/v1/traces": dial tcp 127.0.0.1:4318: connect: connection refused`
- **ADOT**: Do we need to use the AWS Distro of OTel Collector?
- **AppSync Emulation with [Tailcall](https://tailcall.run/)**

## Environment Variables

The default OTel env vars should be fine.

See their [“OTLP Exporter Configuration” guide][guide].

[guide]: https://opentelemetry.io/docs/languages/sdk-configuration/otlp-exporter/