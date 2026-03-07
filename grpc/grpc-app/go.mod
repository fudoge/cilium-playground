module github.com/fudoge/cilium-playground/grpc/grpc-app

go 1.21

require (
	github.com/fudoge/cilium-playground/proto/gen/go/test v0.0.0
	google.golang.org/grpc v1.64.0
)

require (
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/fudoge/cilium-playground/proto/gen/go/test => ../../proto/gen/go/test
