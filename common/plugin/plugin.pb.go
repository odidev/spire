// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

/*
Package sriplugin is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	ConfigureRequest
	ConfigureResponse
	GetPluginInfoRequest
	GetPluginInfoResponse
	PluginInfoRequest
	PluginInfoReply
	StopRequest
	StopReply
*/
package sriplugin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// * Represents the plugin-specific configuration string.
type ConfigureRequest struct {
	Configuration string `protobuf:"bytes,1,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ConfigureRequest) Reset()                    { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()               {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConfigureRequest) GetConfiguration() string {
	if m != nil {
		return m.Configuration
	}
	return ""
}

// * Represents a list of configuration problems found in the configuration string.
type ConfigureResponse struct {
	ErrorList []string `protobuf:"bytes,1,rep,name=errorList" json:"errorList,omitempty"`
}

func (m *ConfigureResponse) Reset()                    { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()               {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConfigureResponse) GetErrorList() []string {
	if m != nil {
		return m.ErrorList
	}
	return nil
}

// * Represents an empty request.
type GetPluginInfoRequest struct {
}

func (m *GetPluginInfoRequest) Reset()                    { *m = GetPluginInfoRequest{} }
func (m *GetPluginInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPluginInfoRequest) ProtoMessage()               {}
func (*GetPluginInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// * Represents the plugin metadata.
type GetPluginInfoResponse struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Category    string `protobuf:"bytes,2,opt,name=category" json:"category,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	DateCreated string `protobuf:"bytes,5,opt,name=dateCreated" json:"dateCreated,omitempty"`
	Location    string `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
	Version     string `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
	Author      string `protobuf:"bytes,8,opt,name=author" json:"author,omitempty"`
	Company     string `protobuf:"bytes,9,opt,name=company" json:"company,omitempty"`
}

func (m *GetPluginInfoResponse) Reset()                    { *m = GetPluginInfoResponse{} }
func (m *GetPluginInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPluginInfoResponse) ProtoMessage()               {}
func (*GetPluginInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetPluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPluginInfoResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetPluginInfoResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetPluginInfoResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetPluginInfoResponse) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *GetPluginInfoResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *GetPluginInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetPluginInfoResponse) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *GetPluginInfoResponse) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

type PluginInfoRequest struct {
}

func (m *PluginInfoRequest) Reset()                    { *m = PluginInfoRequest{} }
func (m *PluginInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*PluginInfoRequest) ProtoMessage()               {}
func (*PluginInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PluginInfoReply struct {
	PluginInfo []*GetPluginInfoResponse `protobuf:"bytes,1,rep,name=pluginInfo" json:"pluginInfo,omitempty"`
}

func (m *PluginInfoReply) Reset()                    { *m = PluginInfoReply{} }
func (m *PluginInfoReply) String() string            { return proto.CompactTextString(m) }
func (*PluginInfoReply) ProtoMessage()               {}
func (*PluginInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PluginInfoReply) GetPluginInfo() []*GetPluginInfoResponse {
	if m != nil {
		return m.PluginInfo
	}
	return nil
}

type StopRequest struct {
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StopReply struct {
}

func (m *StopReply) Reset()                    { *m = StopReply{} }
func (m *StopReply) String() string            { return proto.CompactTextString(m) }
func (*StopReply) ProtoMessage()               {}
func (*StopReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*ConfigureRequest)(nil), "sriplugin.ConfigureRequest")
	proto.RegisterType((*ConfigureResponse)(nil), "sriplugin.ConfigureResponse")
	proto.RegisterType((*GetPluginInfoRequest)(nil), "sriplugin.GetPluginInfoRequest")
	proto.RegisterType((*GetPluginInfoResponse)(nil), "sriplugin.GetPluginInfoResponse")
	proto.RegisterType((*PluginInfoRequest)(nil), "sriplugin.PluginInfoRequest")
	proto.RegisterType((*PluginInfoReply)(nil), "sriplugin.PluginInfoReply")
	proto.RegisterType((*StopRequest)(nil), "sriplugin.StopRequest")
	proto.RegisterType((*StopReply)(nil), "sriplugin.StopReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Server service

type ServerClient interface {
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error)
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoReply, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := grpc.Invoke(ctx, "/sriplugin.Server/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoReply, error) {
	out := new(PluginInfoReply)
	err := grpc.Invoke(ctx, "/sriplugin.Server/PluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerServer interface {
	Stop(context.Context, *StopRequest) (*StopReply, error)
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoReply, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sriplugin.Server/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sriplugin.Server/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sriplugin.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stop",
			Handler:    _Server_Stop_Handler,
		},
		{
			MethodName: "PluginInfo",
			Handler:    _Server_PluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x86, 0x15, 0x60, 0x03, 0x99, 0x2c, 0xda, 0xc5, 0xcb, 0x22, 0x2b, 0xe2, 0x10, 0x45, 0x3d,
	0x70, 0x42, 0x2a, 0xed, 0xa1, 0xc7, 0x4a, 0x48, 0xad, 0x2a, 0xf5, 0x50, 0x85, 0x27, 0x48, 0xc3,
	0x40, 0x23, 0x85, 0xd8, 0xb5, 0x1d, 0xa4, 0xbc, 0x40, 0xfb, 0xda, 0x95, 0x1d, 0x03, 0x6e, 0x69,
	0x6f, 0xfe, 0xbf, 0x7f, 0xc6, 0xbf, 0x67, 0x12, 0xf8, 0xcd, 0xcb, 0x7a, 0x5b, 0x54, 0x73, 0x2e,
	0x98, 0x62, 0x24, 0x90, 0xa2, 0x68, 0x41, 0x72, 0x03, 0x7f, 0x97, 0xac, 0xda, 0x14, 0xdb, 0x5a,
	0x60, 0x8a, 0xaf, 0x35, 0x4a, 0x45, 0x2e, 0x60, 0x98, 0x5b, 0x96, 0xa9, 0x82, 0x55, 0xd4, 0x8b,
	0xbd, 0x59, 0x90, 0x7e, 0x86, 0xc9, 0x25, 0x8c, 0x9c, 0x4e, 0xc9, 0x59, 0x25, 0x91, 0x4c, 0x21,
	0x40, 0x21, 0x98, 0x78, 0x2c, 0xa4, 0xa2, 0x5e, 0xdc, 0x9d, 0x05, 0xe9, 0x09, 0x24, 0x13, 0x18,
	0xdf, 0xa3, 0x7a, 0x32, 0xc9, 0x0f, 0xd5, 0x86, 0xd9, 0xc0, 0xe4, 0xbd, 0x03, 0xff, 0xbf, 0x18,
	0xf6, 0x3e, 0x02, 0xbd, 0x2a, 0xdb, 0xa1, 0x7d, 0x81, 0x39, 0x93, 0x08, 0x06, 0x79, 0xa6, 0x70,
	0xcb, 0x44, 0x43, 0x3b, 0x86, 0x1f, 0xb5, 0xae, 0x57, 0x0d, 0x47, 0xda, 0x6d, 0xeb, 0xf5, 0x99,
	0xc4, 0x10, 0xae, 0x51, 0xe6, 0xa2, 0xe0, 0x66, 0x98, 0x9e, 0xb1, 0x5c, 0x64, 0x2a, 0x32, 0x85,
	0x4b, 0x81, 0x99, 0xc2, 0x35, 0xfd, 0x65, 0x2b, 0x4e, 0x48, 0x67, 0x96, 0x2c, 0x6f, 0xb7, 0xe1,
	0xb7, 0x99, 0x07, 0x4d, 0x28, 0xf4, 0xf7, 0x28, 0xa4, 0xb6, 0xfa, 0xc6, 0x3a, 0x48, 0x32, 0x01,
	0x3f, 0xab, 0xd5, 0x0b, 0x13, 0x74, 0x60, 0x0c, 0xab, 0x74, 0x47, 0xce, 0x76, 0x3c, 0xab, 0x1a,
	0x1a, 0xb4, 0x1d, 0x56, 0x26, 0xff, 0x60, 0x74, 0xbe, 0x9e, 0x15, 0xfc, 0x71, 0x21, 0x2f, 0x1b,
	0x72, 0x0b, 0xc0, 0x8f, 0xc8, 0x2c, 0x3a, 0x5c, 0xc4, 0xf3, 0xe3, 0x67, 0x9d, 0x7f, 0xbb, 0xcd,
	0xd4, 0xe9, 0x49, 0x86, 0x10, 0xae, 0x14, 0xe3, 0x87, 0x8c, 0x10, 0x82, 0x56, 0xf2, 0xb2, 0x59,
	0xbc, 0x79, 0xe0, 0xaf, 0x50, 0xec, 0x51, 0x90, 0x6b, 0xe8, 0x69, 0x4e, 0x26, 0xce, 0xe5, 0x4e,
	0x5f, 0x34, 0x3e, 0xe3, 0xfa, 0x79, 0x77, 0x00, 0xa7, 0x78, 0x32, 0x75, 0x6a, 0xce, 0xa6, 0x8b,
	0xa2, 0x1f, 0x5c, 0x5e, 0x36, 0xcf, 0xbe, 0xf9, 0x5f, 0xaf, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x9b, 0x67, 0x9a, 0xbf, 0x02, 0x00, 0x00,
}