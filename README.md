# aws-lambda-go-otel-poc

Proof of concept using [aws-lambda-go][aws-lambda-go] and [OpenTelemetry][otel].

[aws-lambda-go]: https://github.com/aws/aws-lambda-go
[otel]: https://opentelemetry.io/

## Getting started

### Prerequisites

All you need is [Docker][docker] and [pkgx][pkgx].

[docker]: https://www.docker.com/
[pkgx]: https://pkgx.sh/

### Run everything

```shell
pkgx task@latest run
```

### Invoking the Lambda

#### CLI

```shell
pkgx awslocal@latest \
  lambda invoke \
    --function-name my-lambda \
    --cli-binary-format raw-in-base64-out \
    --payload '{"body": "{\"id\": \"10\"}" }' \
    output.txt
```

#### HTTP

```shell
pkgx http \
  http://localhost:4566/2015-03-31/functions/my-lambda/invocations \
  id="4"
```

## Remaining Questions

- **Trace export**: How does the Go app export to the Collector? Do we need to configure a TracerProvider?
- **ADOT**: Do we need to use the AWS Distro of OTel Collector?
- **AppSync Emulation with [Tailcall](https://tailcall.run/)**