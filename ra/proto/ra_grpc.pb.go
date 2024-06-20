// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: ra.proto

package proto

import (
	context "context"
	proto1 "github.com/letsencrypt/boulder/ca/proto"
	proto "github.com/letsencrypt/boulder/core/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RegistrationAuthority_NewRegistration_FullMethodName                   = "/ra.RegistrationAuthority/NewRegistration"
	RegistrationAuthority_UpdateRegistration_FullMethodName                = "/ra.RegistrationAuthority/UpdateRegistration"
	RegistrationAuthority_PerformValidation_FullMethodName                 = "/ra.RegistrationAuthority/PerformValidation"
	RegistrationAuthority_DeactivateRegistration_FullMethodName            = "/ra.RegistrationAuthority/DeactivateRegistration"
	RegistrationAuthority_DeactivateAuthorization_FullMethodName           = "/ra.RegistrationAuthority/DeactivateAuthorization"
	RegistrationAuthority_RevokeCertByApplicant_FullMethodName             = "/ra.RegistrationAuthority/RevokeCertByApplicant"
	RegistrationAuthority_RevokeCertByKey_FullMethodName                   = "/ra.RegistrationAuthority/RevokeCertByKey"
	RegistrationAuthority_AdministrativelyRevokeCertificate_FullMethodName = "/ra.RegistrationAuthority/AdministrativelyRevokeCertificate"
	RegistrationAuthority_NewOrder_FullMethodName                          = "/ra.RegistrationAuthority/NewOrder"
	RegistrationAuthority_FinalizeOrder_FullMethodName                     = "/ra.RegistrationAuthority/FinalizeOrder"
	RegistrationAuthority_GenerateOCSP_FullMethodName                      = "/ra.RegistrationAuthority/GenerateOCSP"
	RegistrationAuthority_UnpauseAccount_FullMethodName                    = "/ra.RegistrationAuthority/UnpauseAccount"
)

// RegistrationAuthorityClient is the client API for RegistrationAuthority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationAuthorityClient interface {
	NewRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*proto.Registration, error)
	UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*proto.Registration, error)
	PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*proto.Authorization, error)
	DeactivateRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeactivateAuthorization(ctx context.Context, in *proto.Authorization, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeCertByApplicant(ctx context.Context, in *RevokeCertByApplicantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeCertByKey(ctx context.Context, in *RevokeCertByKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AdministrativelyRevokeCertificate(ctx context.Context, in *AdministrativelyRevokeCertificateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*proto.Order, error)
	FinalizeOrder(ctx context.Context, in *FinalizeOrderRequest, opts ...grpc.CallOption) (*proto.Order, error)
	// Generate an OCSP response based on the DB's current status and reason code.
	GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*proto1.OCSPResponse, error)
	UnpauseAccount(ctx context.Context, in *UnpauseAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type registrationAuthorityClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationAuthorityClient(cc grpc.ClientConnInterface) RegistrationAuthorityClient {
	return &registrationAuthorityClient{cc}
}

func (c *registrationAuthorityClient) NewRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*proto.Registration, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(proto.Registration)
	err := c.cc.Invoke(ctx, RegistrationAuthority_NewRegistration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*proto.Registration, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(proto.Registration)
	err := c.cc.Invoke(ctx, RegistrationAuthority_UpdateRegistration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*proto.Authorization, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(proto.Authorization)
	err := c.cc.Invoke(ctx, RegistrationAuthority_PerformValidation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) DeactivateRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RegistrationAuthority_DeactivateRegistration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) DeactivateAuthorization(ctx context.Context, in *proto.Authorization, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RegistrationAuthority_DeactivateAuthorization_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) RevokeCertByApplicant(ctx context.Context, in *RevokeCertByApplicantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RegistrationAuthority_RevokeCertByApplicant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) RevokeCertByKey(ctx context.Context, in *RevokeCertByKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RegistrationAuthority_RevokeCertByKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) AdministrativelyRevokeCertificate(ctx context.Context, in *AdministrativelyRevokeCertificateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RegistrationAuthority_AdministrativelyRevokeCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*proto.Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(proto.Order)
	err := c.cc.Invoke(ctx, RegistrationAuthority_NewOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) FinalizeOrder(ctx context.Context, in *FinalizeOrderRequest, opts ...grpc.CallOption) (*proto.Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(proto.Order)
	err := c.cc.Invoke(ctx, RegistrationAuthority_FinalizeOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*proto1.OCSPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(proto1.OCSPResponse)
	err := c.cc.Invoke(ctx, RegistrationAuthority_GenerateOCSP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) UnpauseAccount(ctx context.Context, in *UnpauseAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RegistrationAuthority_UnpauseAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationAuthorityServer is the server API for RegistrationAuthority service.
// All implementations must embed UnimplementedRegistrationAuthorityServer
// for forward compatibility
type RegistrationAuthorityServer interface {
	NewRegistration(context.Context, *proto.Registration) (*proto.Registration, error)
	UpdateRegistration(context.Context, *UpdateRegistrationRequest) (*proto.Registration, error)
	PerformValidation(context.Context, *PerformValidationRequest) (*proto.Authorization, error)
	DeactivateRegistration(context.Context, *proto.Registration) (*emptypb.Empty, error)
	DeactivateAuthorization(context.Context, *proto.Authorization) (*emptypb.Empty, error)
	RevokeCertByApplicant(context.Context, *RevokeCertByApplicantRequest) (*emptypb.Empty, error)
	RevokeCertByKey(context.Context, *RevokeCertByKeyRequest) (*emptypb.Empty, error)
	AdministrativelyRevokeCertificate(context.Context, *AdministrativelyRevokeCertificateRequest) (*emptypb.Empty, error)
	NewOrder(context.Context, *NewOrderRequest) (*proto.Order, error)
	FinalizeOrder(context.Context, *FinalizeOrderRequest) (*proto.Order, error)
	// Generate an OCSP response based on the DB's current status and reason code.
	GenerateOCSP(context.Context, *GenerateOCSPRequest) (*proto1.OCSPResponse, error)
	UnpauseAccount(context.Context, *UnpauseAccountRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRegistrationAuthorityServer()
}

// UnimplementedRegistrationAuthorityServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationAuthorityServer struct {
}

func (UnimplementedRegistrationAuthorityServer) NewRegistration(context.Context, *proto.Registration) (*proto.Registration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewRegistration not implemented")
}
func (UnimplementedRegistrationAuthorityServer) UpdateRegistration(context.Context, *UpdateRegistrationRequest) (*proto.Registration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegistration not implemented")
}
func (UnimplementedRegistrationAuthorityServer) PerformValidation(context.Context, *PerformValidationRequest) (*proto.Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformValidation not implemented")
}
func (UnimplementedRegistrationAuthorityServer) DeactivateRegistration(context.Context, *proto.Registration) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateRegistration not implemented")
}
func (UnimplementedRegistrationAuthorityServer) DeactivateAuthorization(context.Context, *proto.Authorization) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateAuthorization not implemented")
}
func (UnimplementedRegistrationAuthorityServer) RevokeCertByApplicant(context.Context, *RevokeCertByApplicantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertByApplicant not implemented")
}
func (UnimplementedRegistrationAuthorityServer) RevokeCertByKey(context.Context, *RevokeCertByKeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertByKey not implemented")
}
func (UnimplementedRegistrationAuthorityServer) AdministrativelyRevokeCertificate(context.Context, *AdministrativelyRevokeCertificateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdministrativelyRevokeCertificate not implemented")
}
func (UnimplementedRegistrationAuthorityServer) NewOrder(context.Context, *NewOrderRequest) (*proto.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (UnimplementedRegistrationAuthorityServer) FinalizeOrder(context.Context, *FinalizeOrderRequest) (*proto.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeOrder not implemented")
}
func (UnimplementedRegistrationAuthorityServer) GenerateOCSP(context.Context, *GenerateOCSPRequest) (*proto1.OCSPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateOCSP not implemented")
}
func (UnimplementedRegistrationAuthorityServer) UnpauseAccount(context.Context, *UnpauseAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpauseAccount not implemented")
}
func (UnimplementedRegistrationAuthorityServer) mustEmbedUnimplementedRegistrationAuthorityServer() {}

// UnsafeRegistrationAuthorityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationAuthorityServer will
// result in compilation errors.
type UnsafeRegistrationAuthorityServer interface {
	mustEmbedUnimplementedRegistrationAuthorityServer()
}

func RegisterRegistrationAuthorityServer(s grpc.ServiceRegistrar, srv RegistrationAuthorityServer) {
	s.RegisterService(&RegistrationAuthority_ServiceDesc, srv)
}

func _RegistrationAuthority_NewRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_NewRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewRegistration(ctx, req.(*proto.Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_UpdateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).UpdateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_UpdateRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).UpdateRegistration(ctx, req.(*UpdateRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_PerformValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).PerformValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_PerformValidation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).PerformValidation(ctx, req.(*PerformValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_DeactivateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).DeactivateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_DeactivateRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).DeactivateRegistration(ctx, req.(*proto.Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_DeactivateAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Authorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).DeactivateAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_DeactivateAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).DeactivateAuthorization(ctx, req.(*proto.Authorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_RevokeCertByApplicant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertByApplicantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).RevokeCertByApplicant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_RevokeCertByApplicant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).RevokeCertByApplicant(ctx, req.(*RevokeCertByApplicantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_RevokeCertByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).RevokeCertByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_RevokeCertByKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).RevokeCertByKey(ctx, req.(*RevokeCertByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_AdministrativelyRevokeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministrativelyRevokeCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).AdministrativelyRevokeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_AdministrativelyRevokeCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).AdministrativelyRevokeCertificate(ctx, req.(*AdministrativelyRevokeCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_NewOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_FinalizeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).FinalizeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_FinalizeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).FinalizeOrder(ctx, req.(*FinalizeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_GenerateOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).GenerateOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_GenerateOCSP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).GenerateOCSP(ctx, req.(*GenerateOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_UnpauseAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpauseAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).UnpauseAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationAuthority_UnpauseAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).UnpauseAccount(ctx, req.(*UnpauseAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistrationAuthority_ServiceDesc is the grpc.ServiceDesc for RegistrationAuthority service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistrationAuthority_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ra.RegistrationAuthority",
	HandlerType: (*RegistrationAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewRegistration",
			Handler:    _RegistrationAuthority_NewRegistration_Handler,
		},
		{
			MethodName: "UpdateRegistration",
			Handler:    _RegistrationAuthority_UpdateRegistration_Handler,
		},
		{
			MethodName: "PerformValidation",
			Handler:    _RegistrationAuthority_PerformValidation_Handler,
		},
		{
			MethodName: "DeactivateRegistration",
			Handler:    _RegistrationAuthority_DeactivateRegistration_Handler,
		},
		{
			MethodName: "DeactivateAuthorization",
			Handler:    _RegistrationAuthority_DeactivateAuthorization_Handler,
		},
		{
			MethodName: "RevokeCertByApplicant",
			Handler:    _RegistrationAuthority_RevokeCertByApplicant_Handler,
		},
		{
			MethodName: "RevokeCertByKey",
			Handler:    _RegistrationAuthority_RevokeCertByKey_Handler,
		},
		{
			MethodName: "AdministrativelyRevokeCertificate",
			Handler:    _RegistrationAuthority_AdministrativelyRevokeCertificate_Handler,
		},
		{
			MethodName: "NewOrder",
			Handler:    _RegistrationAuthority_NewOrder_Handler,
		},
		{
			MethodName: "FinalizeOrder",
			Handler:    _RegistrationAuthority_FinalizeOrder_Handler,
		},
		{
			MethodName: "GenerateOCSP",
			Handler:    _RegistrationAuthority_GenerateOCSP_Handler,
		},
		{
			MethodName: "UnpauseAccount",
			Handler:    _RegistrationAuthority_UnpauseAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ra.proto",
}
