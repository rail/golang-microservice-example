// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sum.proto

package sum

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Adder service

type AdderService interface {
	Sum(ctx context.Context, in *SumRequest, opts ...client.CallOption) (*SumResponse, error)
}

type adderService struct {
	c    client.Client
	name string
}

func NewAdderService(name string, c client.Client) AdderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "adder"
	}
	return &adderService{
		c:    c,
		name: name,
	}
}

func (c *adderService) Sum(ctx context.Context, in *SumRequest, opts ...client.CallOption) (*SumResponse, error) {
	req := c.c.NewRequest(c.name, "Adder.Sum", in)
	out := new(SumResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Adder service

type AdderHandler interface {
	Sum(context.Context, *SumRequest, *SumResponse) error
}

func RegisterAdderHandler(s server.Server, hdlr AdderHandler, opts ...server.HandlerOption) error {
	type adder interface {
		Sum(ctx context.Context, in *SumRequest, out *SumResponse) error
	}
	type Adder struct {
		adder
	}
	h := &adderHandler{hdlr}
	return s.Handle(s.NewHandler(&Adder{h}, opts...))
}

type adderHandler struct {
	AdderHandler
}

func (h *adderHandler) Sum(ctx context.Context, in *SumRequest, out *SumResponse) error {
	return h.AdderHandler.Sum(ctx, in, out)
}
