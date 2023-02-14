package currency

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion3

type RateRequest struct {
	Base string `protobuf:"bytes,1,opt,name=Base,json=base,proto3" json:"Base,omitempty"`

	Destination          string   `protobuf:"bytes,2,opt,name=Destination,json=destination,proto3" json:"Destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateRequest) Reset()         { *m = RateRequest{} }
func (m *RateRequest) String() string { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()    {}
func (*RateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

func (m *RateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateRequest.Unmarshal(m, b)
}
func (m *RateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateRequest.Marshal(b, m, deterministic)
}
func (m *RateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateRequest.Merge(m, src)
}
func (m *RateRequest) XXX_Size() int {
	return xxx_messageInfo_RateRequest.Size(m)
}
func (m *RateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateRequest proto.InternalMessageInfo

func (m *RateRequest) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *RateRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type RateResponse struct {
	Rate                 float32  `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateResponse) Reset()         { *m = RateResponse{} }
func (m *RateResponse) String() string { return proto.CompactTextString(m) }
func (*RateResponse) ProtoMessage()    {}
func (*RateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{1}
}

func (m *RateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateResponse.Unmarshal(m, b)
}
func (m *RateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateResponse.Marshal(b, m, deterministic)
}
func (m *RateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateResponse.Merge(m, src)
}
func (m *RateResponse) XXX_Size() int {
	return xxx_messageInfo_RateResponse.Size(m)
}
func (m *RateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateResponse proto.InternalMessageInfo

func (m *RateResponse) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func init() {
	proto.RegisterType((*RateRequest)(nil), "RateRequest")
	proto.RegisterType((*RateResponse)(nil), "RateResponse")
}

func init() {
	proto.RegisterFile("currency.proto", fileDescriptor_d3dc60ed002193ea)
}

var fileDescriptor_d3dc60ed002193ea = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2e, 0x2d, 0x2a,
	0x4a, 0xcd, 0x4b, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe6, 0xe2, 0x0e, 0x4a,
	0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x71, 0x4a, 0x2c,
	0x4e, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x49, 0x4a, 0x2c, 0x4e, 0x15, 0x52, 0xe0,
	0xe2, 0x76, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x02, 0x4b,
	0x71, 0xa7, 0x20, 0x84, 0x94, 0x94, 0xb8, 0x78, 0x20, 0x86, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x82, 0x4c, 0x29, 0x4a, 0x2c, 0x81, 0x98, 0xc2, 0x14, 0x04, 0x66, 0x1b, 0x19, 0x71, 0x71, 0x38,
	0x43, 0xad, 0x16, 0x52, 0xe3, 0x62, 0x77, 0x4f, 0x2d, 0x01, 0x69, 0x11, 0xe2, 0xd1, 0x43, 0xb2,
	0x5e, 0x8a, 0x57, 0x0f, 0xd9, 0x9c, 0x24, 0x36, 0xb0, 0x1b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x22, 0xcb, 0xaf, 0x3b, 0xb5, 0x00, 0x00, 0x00,
}

var _ context.Context
var _ grpc.ClientConnInterface

const _ = grpc.SupportPackageIsVersion6

type CurrencyClient interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type CurrencyServer interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
}

type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetRate(ctx context.Context, req *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv CurrencyServer) {
	s.RegisterService(&_Currency_serviceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Currency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}
