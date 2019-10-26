// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vconf.proto

package vconf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type GetFileResponse struct {
	FileContent          []byte   `protobuf:"bytes,1,opt,name=fileContent,proto3" json:"fileContent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileResponse) Reset()         { *m = GetFileResponse{} }
func (m *GetFileResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileResponse) ProtoMessage()    {}
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bda77b776bdcc01c, []int{0}
}

func (m *GetFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileResponse.Unmarshal(m, b)
}
func (m *GetFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileResponse.Marshal(b, m, deterministic)
}
func (m *GetFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileResponse.Merge(m, src)
}
func (m *GetFileResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileResponse.Size(m)
}
func (m *GetFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileResponse proto.InternalMessageInfo

func (m *GetFileResponse) GetFileContent() []byte {
	if m != nil {
		return m.FileContent
	}
	return nil
}

type GetFileRequest struct {
	AppName              string   `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	AppVersion           string   `protobuf:"bytes,2,opt,name=appVersion,proto3" json:"appVersion,omitempty"`
	FilePath             string   `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileRequest) Reset()         { *m = GetFileRequest{} }
func (m *GetFileRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileRequest) ProtoMessage()    {}
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bda77b776bdcc01c, []int{1}
}

func (m *GetFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileRequest.Unmarshal(m, b)
}
func (m *GetFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileRequest.Marshal(b, m, deterministic)
}
func (m *GetFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileRequest.Merge(m, src)
}
func (m *GetFileRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileRequest.Size(m)
}
func (m *GetFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileRequest proto.InternalMessageInfo

func (m *GetFileRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *GetFileRequest) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *GetFileRequest) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFileResponse)(nil), "GetFileResponse")
	proto.RegisterType((*GetFileRequest)(nil), "GetFileRequest")
}

func init() { proto.RegisterFile("vconf.proto", fileDescriptor_bda77b776bdcc01c) }

var fileDescriptor_bda77b776bdcc01c = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4b, 0xce, 0xcf,
	0x4b, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe6, 0xe2, 0x77, 0x4f, 0x2d, 0x71, 0xcb,
	0xcc, 0x49, 0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x52, 0xe0, 0xe2, 0x4e, 0xcb,
	0xcc, 0x49, 0x75, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09,
	0x42, 0x16, 0x52, 0x4a, 0xe3, 0xe2, 0x83, 0x6b, 0x2a, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0,
	0x62, 0x4f, 0x2c, 0x28, 0xf0, 0x4b, 0xcc, 0x4d, 0x05, 0xab, 0xe7, 0x0c, 0x82, 0x71, 0x85, 0xe4,
	0xb8, 0xb8, 0x12, 0x0b, 0x0a, 0xc2, 0x52, 0x8b, 0x8a, 0x33, 0xf3, 0xf3, 0x24, 0x98, 0xc0, 0x92,
	0x48, 0x22, 0x42, 0x52, 0x5c, 0x1c, 0x20, 0xa3, 0x03, 0x12, 0x4b, 0x32, 0x24, 0x98, 0xc1, 0xb2,
	0x70, 0xbe, 0x91, 0x3d, 0x17, 0xaf, 0x73, 0x7e, 0x5e, 0x5a, 0x66, 0x7a, 0x69, 0x51, 0x62, 0x09,
	0x48, 0xb1, 0x1e, 0x17, 0x3b, 0xd4, 0x62, 0x21, 0x7e, 0x3d, 0x54, 0x27, 0x48, 0x09, 0xe8, 0xa1,
	0x79, 0x44, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x49, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce,
	0x01, 0x58, 0xd7, 0xf3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, "/Configuration/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
}

// UnimplementedConfigurationServer can be embedded to have forward compatible implementations.
type UnimplementedConfigurationServer struct {
}

func (*UnimplementedConfigurationServer) GetFile(ctx context.Context, req *GetFileRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Configuration/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFile",
			Handler:    _Configuration_GetFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vconf.proto",
}
