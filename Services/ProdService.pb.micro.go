// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ProdService.proto

package Services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for ProdService service

type ProdService interface {
	GetProdList(ctx context.Context, in *ProdRequest, opts ...client.CallOption) (*ProdListResponse, error)
	GetProdDetail(ctx context.Context, in *ProdRequest, opts ...client.CallOption) (*ProdDetailResponse, error)
}

type prodService struct {
	c    client.Client
	name string
}

func NewProdService(name string, c client.Client) ProdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "Services"
	}
	return &prodService{
		c:    c,
		name: name,
	}
}

func (c *prodService) GetProdList(ctx context.Context, in *ProdRequest, opts ...client.CallOption) (*ProdListResponse, error) {
	req := c.c.NewRequest(c.name, "ProdService.GetProdList", in)
	out := new(ProdListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodService) GetProdDetail(ctx context.Context, in *ProdRequest, opts ...client.CallOption) (*ProdDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProdService.GetProdDetail", in)
	out := new(ProdDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProdService service

type ProdServiceHandler interface {
	GetProdList(context.Context, *ProdRequest, *ProdListResponse) error
	GetProdDetail(context.Context, *ProdRequest, *ProdDetailResponse) error
}

func RegisterProdServiceHandler(s server.Server, hdlr ProdServiceHandler, opts ...server.HandlerOption) error {
	type prodService interface {
		GetProdList(ctx context.Context, in *ProdRequest, out *ProdListResponse) error
		GetProdDetail(ctx context.Context, in *ProdRequest, out *ProdDetailResponse) error
	}
	type ProdService struct {
		prodService
	}
	h := &prodServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProdService{h}, opts...))
}

type prodServiceHandler struct {
	ProdServiceHandler
}

func (h *prodServiceHandler) GetProdList(ctx context.Context, in *ProdRequest, out *ProdListResponse) error {
	return h.ProdServiceHandler.GetProdList(ctx, in, out)
}

func (h *prodServiceHandler) GetProdDetail(ctx context.Context, in *ProdRequest, out *ProdDetailResponse) error {
	return h.ProdServiceHandler.GetProdDetail(ctx, in, out)
}
