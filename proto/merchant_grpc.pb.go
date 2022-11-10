// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package merchant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MerchantServiceClient is the client API for MerchantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantServiceClient interface {
	AddPartner(ctx context.Context, in *AddPartnerRequest, opts ...grpc.CallOption) (*Partner, error)
	DeletePartner(ctx context.Context, in *DeletePartnerRequest, opts ...grpc.CallOption) (*Partner, error)
	ListPartners(ctx context.Context, in *ListPartnersRequest, opts ...grpc.CallOption) (MerchantService_ListPartnersClient, error)
	AddMerchant(ctx context.Context, in *AddMerchantRequest, opts ...grpc.CallOption) (*Merchant, error)
	DeleteMerchant(ctx context.Context, in *DeleteMerchantRequest, opts ...grpc.CallOption) (*Merchant, error)
	ListMerchants(ctx context.Context, in *ListMerchantsRequest, opts ...grpc.CallOption) (MerchantService_ListMerchantsClient, error)
}

type merchantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantServiceClient(cc grpc.ClientConnInterface) MerchantServiceClient {
	return &merchantServiceClient{cc}
}

func (c *merchantServiceClient) AddPartner(ctx context.Context, in *AddPartnerRequest, opts ...grpc.CallOption) (*Partner, error) {
	out := new(Partner)
	err := c.cc.Invoke(ctx, "/merchant.MerchantService/AddPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) DeletePartner(ctx context.Context, in *DeletePartnerRequest, opts ...grpc.CallOption) (*Partner, error) {
	out := new(Partner)
	err := c.cc.Invoke(ctx, "/merchant.MerchantService/DeletePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) ListPartners(ctx context.Context, in *ListPartnersRequest, opts ...grpc.CallOption) (MerchantService_ListPartnersClient, error) {
	stream, err := c.cc.NewStream(ctx, &MerchantService_ServiceDesc.Streams[0], "/merchant.MerchantService/ListPartners", opts...)
	if err != nil {
		return nil, err
	}
	x := &merchantServiceListPartnersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MerchantService_ListPartnersClient interface {
	Recv() (*Partner, error)
	grpc.ClientStream
}

type merchantServiceListPartnersClient struct {
	grpc.ClientStream
}

func (x *merchantServiceListPartnersClient) Recv() (*Partner, error) {
	m := new(Partner)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *merchantServiceClient) AddMerchant(ctx context.Context, in *AddMerchantRequest, opts ...grpc.CallOption) (*Merchant, error) {
	out := new(Merchant)
	err := c.cc.Invoke(ctx, "/merchant.MerchantService/AddMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) DeleteMerchant(ctx context.Context, in *DeleteMerchantRequest, opts ...grpc.CallOption) (*Merchant, error) {
	out := new(Merchant)
	err := c.cc.Invoke(ctx, "/merchant.MerchantService/DeleteMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) ListMerchants(ctx context.Context, in *ListMerchantsRequest, opts ...grpc.CallOption) (MerchantService_ListMerchantsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MerchantService_ServiceDesc.Streams[1], "/merchant.MerchantService/ListMerchants", opts...)
	if err != nil {
		return nil, err
	}
	x := &merchantServiceListMerchantsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MerchantService_ListMerchantsClient interface {
	Recv() (*Merchant, error)
	grpc.ClientStream
}

type merchantServiceListMerchantsClient struct {
	grpc.ClientStream
}

func (x *merchantServiceListMerchantsClient) Recv() (*Merchant, error) {
	m := new(Merchant)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MerchantServiceServer is the server API for MerchantService service.
// All implementations should embed UnimplementedMerchantServiceServer
// for forward compatibility
type MerchantServiceServer interface {
	AddPartner(context.Context, *AddPartnerRequest) (*Partner, error)
	DeletePartner(context.Context, *DeletePartnerRequest) (*Partner, error)
	ListPartners(*ListPartnersRequest, MerchantService_ListPartnersServer) error
	AddMerchant(context.Context, *AddMerchantRequest) (*Merchant, error)
	DeleteMerchant(context.Context, *DeleteMerchantRequest) (*Merchant, error)
	ListMerchants(*ListMerchantsRequest, MerchantService_ListMerchantsServer) error
}

// UnimplementedMerchantServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMerchantServiceServer struct {
}

func (UnimplementedMerchantServiceServer) AddPartner(context.Context, *AddPartnerRequest) (*Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPartner not implemented")
}
func (UnimplementedMerchantServiceServer) DeletePartner(context.Context, *DeletePartnerRequest) (*Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePartner not implemented")
}
func (UnimplementedMerchantServiceServer) ListPartners(*ListPartnersRequest, MerchantService_ListPartnersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPartners not implemented")
}
func (UnimplementedMerchantServiceServer) AddMerchant(context.Context, *AddMerchantRequest) (*Merchant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMerchant not implemented")
}
func (UnimplementedMerchantServiceServer) DeleteMerchant(context.Context, *DeleteMerchantRequest) (*Merchant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMerchant not implemented")
}
func (UnimplementedMerchantServiceServer) ListMerchants(*ListMerchantsRequest, MerchantService_ListMerchantsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMerchants not implemented")
}

// UnsafeMerchantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServiceServer will
// result in compilation errors.
type UnsafeMerchantServiceServer interface {
	mustEmbedUnimplementedMerchantServiceServer()
}

func RegisterMerchantServiceServer(s grpc.ServiceRegistrar, srv MerchantServiceServer) {
	s.RegisterService(&MerchantService_ServiceDesc, srv)
}

func _MerchantService_AddPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).AddPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.MerchantService/AddPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).AddPartner(ctx, req.(*AddPartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_DeletePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).DeletePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.MerchantService/DeletePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).DeletePartner(ctx, req.(*DeletePartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_ListPartners_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPartnersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MerchantServiceServer).ListPartners(m, &merchantServiceListPartnersServer{stream})
}

type MerchantService_ListPartnersServer interface {
	Send(*Partner) error
	grpc.ServerStream
}

type merchantServiceListPartnersServer struct {
	grpc.ServerStream
}

func (x *merchantServiceListPartnersServer) Send(m *Partner) error {
	return x.ServerStream.SendMsg(m)
}

func _MerchantService_AddMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).AddMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.MerchantService/AddMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).AddMerchant(ctx, req.(*AddMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_DeleteMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).DeleteMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchant.MerchantService/DeleteMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).DeleteMerchant(ctx, req.(*DeleteMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_ListMerchants_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMerchantsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MerchantServiceServer).ListMerchants(m, &merchantServiceListMerchantsServer{stream})
}

type MerchantService_ListMerchantsServer interface {
	Send(*Merchant) error
	grpc.ServerStream
}

type merchantServiceListMerchantsServer struct {
	grpc.ServerStream
}

func (x *merchantServiceListMerchantsServer) Send(m *Merchant) error {
	return x.ServerStream.SendMsg(m)
}

// MerchantService_ServiceDesc is the grpc.ServiceDesc for MerchantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "merchant.MerchantService",
	HandlerType: (*MerchantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPartner",
			Handler:    _MerchantService_AddPartner_Handler,
		},
		{
			MethodName: "DeletePartner",
			Handler:    _MerchantService_DeletePartner_Handler,
		},
		{
			MethodName: "AddMerchant",
			Handler:    _MerchantService_AddMerchant_Handler,
		},
		{
			MethodName: "DeleteMerchant",
			Handler:    _MerchantService_DeleteMerchant_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPartners",
			Handler:       _MerchantService_ListPartners_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListMerchants",
			Handler:       _MerchantService_ListMerchants_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/merchant.proto",
}