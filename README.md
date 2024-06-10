# aws-lambda-go-otel-poc

Proof of concept using [aws-lambda-go][aws-lambda-go] and [OpenTelemetry][otel].

This may interest you if you use AWS, Lambdas, Golang, and OpenTelemetry … or 
if you're looking to emulate [AppSync Serverless GraphQL][appsync] locally and 
do not wish to pay for a [LocalStack][localstack] license.

[appsync]: https://aws.amazon.com/appsync/
[aws-lambda-go]: https://github.com/aws/aws-lambda-go
[localstack]: https://docs.localstack.cloud/user-guide/aws/appsync/
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

#### (With GraphQL)

```graphql
mutation {
  invokeLambda(input: {id: "1"}) {
    id
  }
}
```

### Step 3: Viewing the trace

Open the [Jaeger UI][jaeger-ui].

[jaeger-ui]: http://localhost:16686

## Miscellaneous

### Remaining Questions

- **ADOT**: Do we need to use the AWS Distro of OTel Collector?

## Environment Variables

The default OTel env vars should be fine.

See their [“OTLP Exporter Configuration” guide][guide].

[guide]: https://opentelemetry.io/docs/languages/sdk-configuration/otlp-exporter/