// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/attraction/attraction.proto

package go_micro_srv_attraction

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Attraction struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attraction) Reset()         { *m = Attraction{} }
func (m *Attraction) String() string { return proto.CompactTextString(m) }
func (*Attraction) ProtoMessage()    {}
func (*Attraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_attraction_58af7ff5af594b58, []int{0}
}
func (m *Attraction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attraction.Unmarshal(m, b)
}
func (m *Attraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attraction.Marshal(b, m, deterministic)
}
func (dst *Attraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attraction.Merge(dst, src)
}
func (m *Attraction) XXX_Size() int {
	return xxx_messageInfo_Attraction.Size(m)
}
func (m *Attraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Attraction.DiscardUnknown(m)
}

var xxx_messageInfo_Attraction proto.InternalMessageInfo

func (m *Attraction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Attraction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Attraction) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Attraction) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Attraction) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Attraction) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_attraction_58af7ff5af594b58, []int{1}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool          `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Attraction           *Attraction   `protobuf:"bytes,2,opt,name=attraction,proto3" json:"attraction,omitempty"`
	Attractions          []*Attraction `protobuf:"bytes,3,rep,name=attractions,proto3" json:"attractions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_attraction_58af7ff5af594b58, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetAttraction() *Attraction {
	if m != nil {
		return m.Attraction
	}
	return nil
}

func (m *Response) GetAttractions() []*Attraction {
	if m != nil {
		return m.Attractions
	}
	return nil
}

func init() {
	proto.RegisterType((*Attraction)(nil), "go.micro.srv.attraction.Attraction")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.attraction.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.attraction.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AttractionService service

type AttractionServiceClient interface {
	CreateAttraction(ctx context.Context, in *Attraction, opts ...client.CallOption) (*Response, error)
	GetAttractions(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type attractionServiceClient struct {
	c           client.Client
	serviceName string
}

func NewAttractionServiceClient(serviceName string, c client.Client) AttractionServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.attraction"
	}
	return &attractionServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *attractionServiceClient) CreateAttraction(ctx context.Context, in *Attraction, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AttractionService.CreateAttraction", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionServiceClient) GetAttractions(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AttractionService.GetAttractions", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AttractionService service

type AttractionServiceHandler interface {
	CreateAttraction(context.Context, *Attraction, *Response) error
	GetAttractions(context.Context, *GetRequest, *Response) error
}

func RegisterAttractionServiceHandler(s server.Server, hdlr AttractionServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AttractionService{hdlr}, opts...))
}

type AttractionService struct {
	AttractionServiceHandler
}

func (h *AttractionService) CreateAttraction(ctx context.Context, in *Attraction, out *Response) error {
	return h.AttractionServiceHandler.CreateAttraction(ctx, in, out)
}

func (h *AttractionService) GetAttractions(ctx context.Context, in *GetRequest, out *Response) error {
	return h.AttractionServiceHandler.GetAttractions(ctx, in, out)
}

func init() {
	proto.RegisterFile("proto/attraction/attraction.proto", fileDescriptor_attraction_58af7ff5af594b58)
}

var fileDescriptor_attraction_58af7ff5af594b58 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x24, 0xe9, 0x83, 0xb2, 0x41, 0x15, 0xac, 0x90, 0xb0, 0x38, 0xb5, 0xe1, 0xc2, 0xc9, 0x48,
	0xe5, 0x0b, 0x50, 0x85, 0x7a, 0x0f, 0x37, 0xc4, 0x25, 0x38, 0x2b, 0xe4, 0x03, 0x71, 0xb0, 0xb7,
	0x95, 0xfa, 0x2f, 0x7c, 0x04, 0x3f, 0xc2, 0x3f, 0xa1, 0x38, 0x04, 0xfb, 0xd2, 0x2a, 0xb7, 0x9d,
	0x1d, 0xcf, 0x68, 0x67, 0x64, 0x58, 0x36, 0xd6, 0xb0, 0xb9, 0x2f, 0x99, 0x6d, 0xa9, 0x58, 0x9b,
	0x3a, 0x1a, 0xa5, 0xe7, 0xf0, 0xfa, 0xdd, 0xc8, 0x0f, 0xad, 0xac, 0x91, 0xce, 0xee, 0x64, 0xa0,
	0xf3, 0xaf, 0x04, 0xe0, 0xf1, 0x1f, 0xe2, 0x1c, 0x52, 0x5d, 0x89, 0x64, 0x91, 0xdc, 0x9d, 0x15,
	0xa9, 0xae, 0x70, 0x01, 0x59, 0x45, 0x4e, 0x59, 0xdd, 0xb4, 0xb4, 0x48, 0x3d, 0x11, 0xaf, 0x50,
	0xc0, 0x69, 0x59, 0x55, 0x96, 0x9c, 0x13, 0x23, 0xcf, 0xf6, 0x10, 0x11, 0xc6, 0x4a, 0xf3, 0x5e,
	0x8c, 0xfd, 0xda, 0xcf, 0x78, 0x05, 0x13, 0xc7, 0x25, 0x93, 0x98, 0xf8, 0x65, 0x07, 0x5a, 0x0f,
	0x65, 0xb6, 0x35, 0xdb, 0xbd, 0x98, 0x76, 0x1e, 0x7f, 0x30, 0x3f, 0x07, 0xd8, 0x10, 0x17, 0xf4,
	0xb9, 0x25, 0xc7, 0xf9, 0x77, 0x02, 0xb3, 0x82, 0x5c, 0x63, 0x6a, 0xd7, 0x89, 0x2c, 0x95, 0x4c,
	0xdd, 0xbd, 0xb3, 0xa2, 0x87, 0xb8, 0x06, 0x08, 0x09, 0xfd, 0xcd, 0xd9, 0xea, 0x56, 0x1e, 0x68,
	0x40, 0x86, 0xf4, 0x45, 0x24, 0xc3, 0x27, 0xc8, 0x02, 0x6a, 0xb3, 0x8d, 0x86, 0xba, 0xc4, 0xba,
	0xd5, 0x4f, 0x02, 0x97, 0x81, 0x7b, 0x26, 0xbb, 0xd3, 0x8a, 0xf0, 0x15, 0x2e, 0xd6, 0xfe, 0xd8,
	0xa8, 0xfa, 0x21, 0xde, 0x37, 0xcb, 0x83, 0x8f, 0xfa, 0x5e, 0xf2, 0x13, 0x7c, 0x81, 0xf9, 0x86,
	0x38, 0xa8, 0xdc, 0x11, 0xef, 0xd0, 0xee, 0x20, 0xef, 0xb7, 0xa9, 0xff, 0x4f, 0x0f, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x52, 0xbf, 0xc5, 0x11, 0x74, 0x02, 0x00, 0x00,
}
