# grpc-app

Simple gRPC server/client using generated code from `proto/gen/go/test/v1`.

## Run

In one terminal:

```bash
cd grpc/grpc-app

go run ./cmd/server
```

In another terminal:

```bash
cd grpc/grpc-app

go run ./cmd/client -message "hello"
```
