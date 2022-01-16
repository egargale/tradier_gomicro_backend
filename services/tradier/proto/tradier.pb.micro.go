// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/tradier.proto

package tradier

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Tradier service

func NewTradierEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Tradier service

type TradierService interface {
	Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (*StreamResponse, error)
}

type tradierService struct {
	c    client.Client
	name string
}

func NewTradierService(name string, c client.Client) TradierService {
	return &tradierService{
		c:    c,
		name: name,
	}
}

func (c *tradierService) Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (*StreamResponse, error) {
	req := c.c.NewRequest(c.name, "Tradier.Stream", in)
	out := new(StreamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tradier service

type TradierHandler interface {
	Stream(context.Context, *StreamRequest, *StreamResponse) error
}

func RegisterTradierHandler(s server.Server, hdlr TradierHandler, opts ...server.HandlerOption) error {
	type tradier interface {
		Stream(ctx context.Context, in *StreamRequest, out *StreamResponse) error
	}
	type Tradier struct {
		tradier
	}
	h := &tradierHandler{hdlr}
	return s.Handle(s.NewHandler(&Tradier{h}, opts...))
}

type tradierHandler struct {
	TradierHandler
}

func (h *tradierHandler) Stream(ctx context.Context, in *StreamRequest, out *StreamResponse) error {
	return h.TradierHandler.Stream(ctx, in, out)
}
