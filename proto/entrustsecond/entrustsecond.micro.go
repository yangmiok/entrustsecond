// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/entrustsecond/entrustsecond.proto

/*
Package go_micro_api_entrust is a generated protocol buffer package.

It is generated from these files:
	proto/entrustsecond/entrustsecond.proto

It has these top-level messages:
	Pair
	Request
	Response
	TradeRequest
	TradeOrderResponse
	TradeOrder
	UnfinishedRequest
	BusinessRequest
	CurrencyEntrustRequest
	CancelRequest
	InnerRequest
	InnerCancelRequest
*/
package go_micro_api_entrust

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Entrust service

type EntrustService interface {
	Cancel(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetUnfinishedBuyList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetUnfinishedSellList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	EntrustBuy(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	EntrustSell(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetEntrustList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	InnerBuySecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	InnerSellSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	InnerCancelSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	InnerEntrustListSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	InnerEntrustDetailSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type entrustService struct {
	c    client.Client
	name string
}

func NewEntrustService(name string, c client.Client) EntrustService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.api.entrust"
	}
	return &entrustService{
		c:    c,
		name: name,
	}
}

func (c *entrustService) Cancel(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.Cancel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) GetUnfinishedBuyList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.GetUnfinishedBuyList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) GetUnfinishedSellList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.GetUnfinishedSellList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) EntrustBuy(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.EntrustBuy", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) EntrustSell(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.EntrustSell", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) GetEntrustList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.GetEntrustList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) InnerBuySecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.InnerBuySecondary", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) InnerSellSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.InnerSellSecondary", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) InnerCancelSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.InnerCancelSecondary", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) InnerEntrustListSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.InnerEntrustListSecondary", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrustService) InnerEntrustDetailSecondary(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Entrust.InnerEntrustDetailSecondary", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Entrust service

type EntrustHandler interface {
	Cancel(context.Context, *Request, *Response) error
	GetUnfinishedBuyList(context.Context, *Request, *Response) error
	GetUnfinishedSellList(context.Context, *Request, *Response) error
	EntrustBuy(context.Context, *Request, *Response) error
	EntrustSell(context.Context, *Request, *Response) error
	GetEntrustList(context.Context, *Request, *Response) error
	InnerBuySecondary(context.Context, *Request, *Response) error
	InnerSellSecondary(context.Context, *Request, *Response) error
	InnerCancelSecondary(context.Context, *Request, *Response) error
	InnerEntrustListSecondary(context.Context, *Request, *Response) error
	InnerEntrustDetailSecondary(context.Context, *Request, *Response) error
}

func RegisterEntrustHandler(s server.Server, hdlr EntrustHandler, opts ...server.HandlerOption) error {
	type entrust interface {
		Cancel(ctx context.Context, in *Request, out *Response) error
		GetUnfinishedBuyList(ctx context.Context, in *Request, out *Response) error
		GetUnfinishedSellList(ctx context.Context, in *Request, out *Response) error
		EntrustBuy(ctx context.Context, in *Request, out *Response) error
		EntrustSell(ctx context.Context, in *Request, out *Response) error
		GetEntrustList(ctx context.Context, in *Request, out *Response) error
		InnerBuySecondary(ctx context.Context, in *Request, out *Response) error
		InnerSellSecondary(ctx context.Context, in *Request, out *Response) error
		InnerCancelSecondary(ctx context.Context, in *Request, out *Response) error
		InnerEntrustListSecondary(ctx context.Context, in *Request, out *Response) error
		InnerEntrustDetailSecondary(ctx context.Context, in *Request, out *Response) error
	}
	type Entrust struct {
		entrust
	}
	h := &entrustHandler{hdlr}
	return s.Handle(s.NewHandler(&Entrust{h}, opts...))
}

type entrustHandler struct {
	EntrustHandler
}

func (h *entrustHandler) Cancel(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.Cancel(ctx, in, out)
}

func (h *entrustHandler) GetUnfinishedBuyList(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.GetUnfinishedBuyList(ctx, in, out)
}

func (h *entrustHandler) GetUnfinishedSellList(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.GetUnfinishedSellList(ctx, in, out)
}

func (h *entrustHandler) EntrustBuy(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.EntrustBuy(ctx, in, out)
}

func (h *entrustHandler) EntrustSell(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.EntrustSell(ctx, in, out)
}

func (h *entrustHandler) GetEntrustList(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.GetEntrustList(ctx, in, out)
}

func (h *entrustHandler) InnerBuySecondary(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.InnerBuySecondary(ctx, in, out)
}

func (h *entrustHandler) InnerSellSecondary(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.InnerSellSecondary(ctx, in, out)
}

func (h *entrustHandler) InnerCancelSecondary(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.InnerCancelSecondary(ctx, in, out)
}

func (h *entrustHandler) InnerEntrustListSecondary(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.InnerEntrustListSecondary(ctx, in, out)
}

func (h *entrustHandler) InnerEntrustDetailSecondary(ctx context.Context, in *Request, out *Response) error {
	return h.EntrustHandler.InnerEntrustDetailSecondary(ctx, in, out)
}
