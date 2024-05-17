// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aib/cron/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cac7f98e937b79a, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cac7f98e937b79a, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryWhitelistedContractsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty" yaml:"pagination"`
}

func (m *QueryWhitelistedContractsRequest) Reset()         { *m = QueryWhitelistedContractsRequest{} }
func (m *QueryWhitelistedContractsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWhitelistedContractsRequest) ProtoMessage()    {}
func (*QueryWhitelistedContractsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cac7f98e937b79a, []int{2}
}
func (m *QueryWhitelistedContractsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhitelistedContractsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhitelistedContractsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhitelistedContractsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhitelistedContractsRequest.Merge(m, src)
}
func (m *QueryWhitelistedContractsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhitelistedContractsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhitelistedContractsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhitelistedContractsRequest proto.InternalMessageInfo

func (m *QueryWhitelistedContractsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryWhitelistedContractsResponse struct {
	WhilistedContracts []WhitelistedContract `protobuf:"bytes,1,rep,name=whilisted_contracts,json=whilistedContracts,proto3" json:"whilisted_contracts"`
	Pagination         *query.PageResponse   `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty" yaml:"pagination"`
}

func (m *QueryWhitelistedContractsResponse) Reset()         { *m = QueryWhitelistedContractsResponse{} }
func (m *QueryWhitelistedContractsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWhitelistedContractsResponse) ProtoMessage()    {}
func (*QueryWhitelistedContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cac7f98e937b79a, []int{3}
}
func (m *QueryWhitelistedContractsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhitelistedContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhitelistedContractsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhitelistedContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhitelistedContractsResponse.Merge(m, src)
}
func (m *QueryWhitelistedContractsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhitelistedContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhitelistedContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhitelistedContractsResponse proto.InternalMessageInfo

func (m *QueryWhitelistedContractsResponse) GetWhilistedContracts() []WhitelistedContract {
	if m != nil {
		return m.WhilistedContracts
	}
	return nil
}

func (m *QueryWhitelistedContractsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "aib.cron.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "aib.cron.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryWhitelistedContractsRequest)(nil), "aib.cron.v1beta1.QueryWhitelistedContractsRequest")
	proto.RegisterType((*QueryWhitelistedContractsResponse)(nil), "aib.cron.v1beta1.QueryWhitelistedContractsResponse")
}

func init() { proto.RegisterFile("aib/cron/v1beta1/query.proto", fileDescriptor_2cac7f98e937b79a) }

var fileDescriptor_2cac7f98e937b79a = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0xaf, 0x0b, 0xf4, 0xe0, 0x5d, 0xc0, 0x1b, 0x52, 0x09, 0x23, 0x2b, 0x11, 0x63, 0x03, 0x09,
	0x5b, 0xed, 0x24, 0x0e, 0x1c, 0xcb, 0x11, 0x21, 0x41, 0x2f, 0x48, 0x88, 0x3f, 0x72, 0x32, 0x2b,
	0xb5, 0x94, 0xd8, 0x5e, 0xec, 0xd2, 0x85, 0x1b, 0x3c, 0x01, 0x12, 0xcf, 0x82, 0xc4, 0x23, 0xec,
	0x38, 0x89, 0x0b, 0xa7, 0x09, 0x5a, 0x9e, 0x80, 0x27, 0x40, 0xb1, 0xdd, 0x86, 0x11, 0xba, 0xc1,
	0xad, 0xea, 0xef, 0xfb, 0xfd, 0xfb, 0x3e, 0x07, 0x6e, 0x52, 0x1e, 0x93, 0xa4, 0x90, 0x82, 0xbc,
	0xe9, 0xc7, 0xcc, 0xd0, 0x3e, 0x39, 0x98, 0xb0, 0xa2, 0xc4, 0xaa, 0x90, 0x46, 0xa2, 0xcb, 0x94,
	0xc7, 0xb8, 0x42, 0xb1, 0x47, 0x83, 0x8d, 0x54, 0xa6, 0xd2, 0x82, 0xa4, 0xfa, 0xe5, 0xe6, 0x82,
	0xcd, 0x54, 0xca, 0x34, 0x63, 0x84, 0x2a, 0x4e, 0xa8, 0x10, 0xd2, 0x50, 0xc3, 0xa5, 0xd0, 0x1e,
	0xbd, 0x9b, 0x48, 0x9d, 0x4b, 0x4d, 0x62, 0xaa, 0x99, 0x93, 0x5f, 0x9a, 0x29, 0x9a, 0x72, 0x61,
	0x87, 0xfd, 0xec, 0x8d, 0x46, 0x1e, 0x45, 0x0b, 0x9a, 0x2f, 0xa4, 0xae, 0x37, 0x60, 0x9b, 0xce,
	0x82, 0xd1, 0x06, 0x44, 0x4f, 0x2b, 0xf5, 0x27, 0x96, 0x31, 0x62, 0x07, 0x13, 0xa6, 0x4d, 0xf4,
	0x18, 0xae, 0x9f, 0xfa, 0x57, 0x2b, 0x29, 0x34, 0x43, 0xf7, 0x61, 0xc7, 0x29, 0x77, 0x41, 0x0f,
	0xec, 0xae, 0x0d, 0xba, 0xf8, 0xcf, 0xae, 0xd8, 0x31, 0x86, 0x17, 0x8f, 0x4e, 0xb6, 0x5a, 0x23,
	0x3f, 0x1d, 0xbd, 0x03, 0xb0, 0x67, 0xf5, 0x9e, 0x8d, 0xb9, 0x61, 0x19, 0xd7, 0x86, 0xed, 0x3f,
	0x94, 0xc2, 0x14, 0x34, 0x31, 0x0b, 0x4f, 0xf4, 0x12, 0xc2, 0xba, 0x99, 0x37, 0xb8, 0x8d, 0xdd,
	0x1a, 0x70, 0xb5, 0x06, 0xec, 0xb6, 0x5c, 0x3b, 0xa5, 0xcc, 0x73, 0x87, 0x57, 0x7f, 0x9e, 0x6c,
	0x5d, 0x29, 0x69, 0x9e, 0x3d, 0x88, 0x6a, 0x8d, 0x68, 0xf4, 0x9b, 0x60, 0xf4, 0x1d, 0xc0, 0x9b,
	0x67, 0x64, 0xf0, 0x0d, 0x5f, 0xc0, 0xf5, 0xe9, 0x98, 0x3b, 0xf4, 0x75, 0xb2, 0x80, 0xbb, 0xa0,
	0x77, 0x61, 0x77, 0x6d, 0xb0, 0xdd, 0xac, 0xfb, 0x17, 0x31, 0xdf, 0x1d, 0x2d, 0x75, 0x96, 0x2e,
	0xe8, 0xd5, 0xa9, 0x8a, 0x6d, 0x5b, 0x71, 0xe7, 0xdc, 0x8a, 0x2e, 0xda, 0x3f, 0x74, 0x1c, 0x7c,
	0x6e, 0xc3, 0x4b, 0xb6, 0x23, 0x9a, 0xc2, 0x8e, 0xbb, 0x04, 0xba, 0xd5, 0x0c, 0xdd, 0x3c, 0x78,
	0xb0, 0x7d, 0xce, 0x94, 0xcb, 0x10, 0xf5, 0xde, 0x7f, 0xf9, 0xf1, 0xb1, 0x1d, 0xa0, 0x2e, 0x59,
	0xf1, 0xe4, 0xd0, 0x27, 0x00, 0xaf, 0xad, 0x5c, 0x33, 0x1a, 0xac, 0xb0, 0x39, 0xe3, 0x5d, 0x04,
	0x7b, 0xff, 0xc5, 0xf1, 0x41, 0x89, 0x0d, 0x7a, 0x07, 0xed, 0x34, 0x83, 0x4e, 0x6b, 0x5e, 0x7d,
	0xe1, 0xe1, 0xa3, 0xa3, 0x59, 0x08, 0x8e, 0x67, 0x21, 0xf8, 0x36, 0x0b, 0xc1, 0x87, 0x79, 0xd8,
	0x3a, 0x9e, 0x87, 0xad, 0xaf, 0xf3, 0xb0, 0xf5, 0xbc, 0x9f, 0x72, 0x33, 0x9e, 0xc4, 0x38, 0x91,
	0x39, 0xd9, 0x2f, 0x73, 0x26, 0x34, 0x97, 0xe2, 0xb0, 0x7c, 0x4b, 0x0a, 0x99, 0x65, 0x54, 0xa9,
	0x7b, 0x53, 0xaa, 0x73, 0x72, 0xe8, 0x4c, 0x4c, 0xa9, 0x98, 0x8e, 0x3b, 0xf6, 0xdb, 0xda, 0xfb,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x88, 0xed, 0x9a, 0xac, 0x29, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	QueryWhitelistedContracts(ctx context.Context, in *QueryWhitelistedContractsRequest, opts ...grpc.CallOption) (*QueryWhitelistedContractsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/aib.cron.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryWhitelistedContracts(ctx context.Context, in *QueryWhitelistedContractsRequest, opts ...grpc.CallOption) (*QueryWhitelistedContractsResponse, error) {
	out := new(QueryWhitelistedContractsResponse)
	err := c.cc.Invoke(ctx, "/aib.cron.v1beta1.Query/QueryWhitelistedContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	QueryWhitelistedContracts(context.Context, *QueryWhitelistedContractsRequest) (*QueryWhitelistedContractsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) QueryWhitelistedContracts(ctx context.Context, req *QueryWhitelistedContractsRequest) (*QueryWhitelistedContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWhitelistedContracts not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aib.cron.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryWhitelistedContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWhitelistedContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryWhitelistedContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aib.cron.v1beta1.Query/QueryWhitelistedContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryWhitelistedContracts(ctx, req.(*QueryWhitelistedContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aib.cron.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QueryWhitelistedContracts",
			Handler:    _Query_QueryWhitelistedContracts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aib/cron/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryWhitelistedContractsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhitelistedContractsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhitelistedContractsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryWhitelistedContractsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhitelistedContractsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhitelistedContractsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.WhilistedContracts) > 0 {
		for iNdEx := len(m.WhilistedContracts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhilistedContracts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryWhitelistedContractsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryWhitelistedContractsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WhilistedContracts) > 0 {
		for _, e := range m.WhilistedContracts {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryWhitelistedContractsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryWhitelistedContractsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhitelistedContractsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryWhitelistedContractsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryWhitelistedContractsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhitelistedContractsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhilistedContracts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhilistedContracts = append(m.WhilistedContracts, WhitelistedContract{})
			if err := m.WhilistedContracts[len(m.WhilistedContracts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
