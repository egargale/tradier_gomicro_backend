module tstream

go 1.15

require (
	github.com/gofiber/fiber/v2 v2.24.0
	github.com/golang/protobuf v1.4.3
	github.com/micro/micro/v3 v3.0.0
	github.com/pkg/errors v0.9.1
	github.com/valyala/fasthttp v1.32.0
	google.golang.org/protobuf v1.25.0
	nhooyr.io/websocket v1.8.7
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
