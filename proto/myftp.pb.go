// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myftp.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UploadResponse_Status int32

const (
	UploadResponse_UNKNOWN UploadResponse_Status = 0
	UploadResponse_OK      UploadResponse_Status = 1
	UploadResponse_FAILED  UploadResponse_Status = 2
)

var UploadResponse_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "FAILED",
}

var UploadResponse_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"OK":      1,
	"FAILED":  2,
}

func (x UploadResponse_Status) String() string {
	return proto.EnumName(UploadResponse_Status_name, int32(x))
}

func (UploadResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{1, 0}
}

type UploadRequest struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{0}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *UploadRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type UploadResponse struct {
	Message              string                `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               UploadResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=proto.UploadResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{1}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func (m *UploadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadResponse) GetStatus() UploadResponse_Status {
	if m != nil {
		return m.Status
	}
	return UploadResponse_UNKNOWN
}

type DownloadRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{2}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DownloadResponse struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadResponse) Reset()         { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()    {}
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{3}
}

func (m *DownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResponse.Unmarshal(m, b)
}
func (m *DownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResponse.Marshal(b, m, deterministic)
}
func (m *DownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResponse.Merge(m, src)
}
func (m *DownloadResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadResponse.Size(m)
}
func (m *DownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResponse proto.InternalMessageInfo

func (m *DownloadResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Files                []*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{5}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetFiles() []*FileInfo {
	if m != nil {
		return m.Files
	}
	return nil
}

type FileInfo struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64                `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Mode                 uint32               `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d33c66d3c1a633, []int{6}
}

func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileInfo) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *FileInfo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.UploadResponse_Status", UploadResponse_Status_name, UploadResponse_Status_value)
	proto.RegisterType((*UploadRequest)(nil), "proto.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "proto.UploadResponse")
	proto.RegisterType((*DownloadRequest)(nil), "proto.DownloadRequest")
	proto.RegisterType((*DownloadResponse)(nil), "proto.DownloadResponse")
	proto.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto.RegisterType((*FileInfo)(nil), "proto.FileInfo")
}

func init() { proto.RegisterFile("myftp.proto", fileDescriptor_07d33c66d3c1a633) }

var fileDescriptor_07d33c66d3c1a633 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x75, 0x36, 0x6d, 0xba, 0xb9, 0xe9, 0xb6, 0x61, 0xfc, 0x0a, 0x41, 0x30, 0x0c, 0x14, 0x22,
	0x48, 0x2a, 0xb1, 0xa2, 0x2f, 0x3e, 0x14, 0xda, 0x85, 0x65, 0x97, 0x2c, 0x8c, 0x2e, 0x3e, 0x67,
	0xcd, 0x64, 0x09, 0x24, 0x99, 0xe8, 0x4c, 0x10, 0x7d, 0xf2, 0xdd, 0x9f, 0xe3, 0x1f, 0x94, 0xcc,
	0x4c, 0xdc, 0x8f, 0xee, 0x53, 0xee, 0x3d, 0x39, 0x73, 0xb8, 0xe7, 0x1c, 0x70, 0xeb, 0x9f, 0x85,
	0x6c, 0xe3, 0xf6, 0x3b, 0x97, 0x1c, 0x9f, 0xaa, 0x4f, 0xf0, 0x72, 0xc3, 0xf9, 0xa6, 0x62, 0xd7,
	0x6a, 0x5b, 0x77, 0xc5, 0xb5, 0x2c, 0x6b, 0x26, 0x64, 0x56, 0x1b, 0x1e, 0xb9, 0x87, 0xc9, 0xaa,
	0xad, 0x78, 0x96, 0x53, 0xf6, 0xad, 0x63, 0x42, 0x62, 0x1f, 0xce, 0xbe, 0xf2, 0x46, 0xb2, 0x46,
	0xfa, 0x28, 0x44, 0xd1, 0x39, 0x1d, 0x56, 0x1c, 0xc0, 0xb8, 0x28, 0x2b, 0x96, 0x66, 0x35, 0xf3,
	0x47, 0x21, 0x8a, 0x1c, 0xfa, 0x7f, 0x27, 0x7f, 0x10, 0x5c, 0x0c, 0x3a, 0xa2, 0xe5, 0x8d, 0x60,
	0xbd, 0x50, 0xcd, 0x84, 0xc8, 0x36, 0x4c, 0x09, 0x39, 0x74, 0x58, 0xf1, 0x0d, 0xd8, 0x42, 0x66,
	0xb2, 0x13, 0x4a, 0xe6, 0x22, 0x79, 0xa1, 0x6f, 0x89, 0xf7, 0x05, 0xe2, 0x4f, 0x8a, 0x43, 0x0d,
	0x97, 0xbc, 0x02, 0x5b, 0x23, 0xd8, 0x85, 0xb3, 0x55, 0x3a, 0x4f, 0x97, 0x5f, 0x52, 0xef, 0x11,
	0xb6, 0x61, 0xb4, 0x9c, 0x7b, 0x08, 0x03, 0xd8, 0xd3, 0xdb, 0xd9, 0xe2, 0xfe, 0xce, 0x1b, 0x91,
	0x2b, 0xb8, 0xbc, 0xe3, 0x3f, 0x9a, 0x5d, 0x5b, 0x18, 0x4e, 0x9a, 0xfe, 0x70, 0x7d, 0x8a, 0x9a,
	0xc9, 0x6b, 0xf0, 0xb6, 0xb4, 0xed, 0xd5, 0xc7, 0xed, 0x93, 0x09, 0xb8, 0x8b, 0x52, 0x48, 0x23,
	0x48, 0xde, 0xc1, 0xb9, 0x5e, 0xcd, 0xc3, 0x2b, 0x38, 0xed, 0xd3, 0x10, 0x3e, 0x0a, 0xad, 0xc8,
	0x4d, 0x2e, 0x8d, 0xa7, 0x69, 0x59, 0xb1, 0x59, 0x53, 0x70, 0xaa, 0xff, 0x92, 0xdf, 0x08, 0xc6,
	0x03, 0x76, 0xec, 0xa8, 0x1e, 0x13, 0xe5, 0x2f, 0x9d, 0xb0, 0x45, 0xd5, 0xdc, 0x63, 0x35, 0xcf,
	0x99, 0x6f, 0x85, 0x28, 0x9a, 0x50, 0x35, 0xe3, 0x0f, 0xe0, 0x74, 0x6d, 0x9e, 0x49, 0x96, 0xdf,
	0x4a, 0xff, 0x24, 0x44, 0x91, 0x9b, 0x04, 0xb1, 0x6e, 0x3b, 0x1e, 0xda, 0x8e, 0x3f, 0x0f, 0x6d,
	0xd3, 0x2d, 0x39, 0xf9, 0x8b, 0xc0, 0x9a, 0xca, 0x16, 0xbf, 0x07, 0x5b, 0x27, 0x8e, 0x9f, 0x1c,
	0x14, 0xa0, 0x1c, 0x06, 0x4f, 0x8f, 0xd6, 0x12, 0x21, 0x7c, 0x03, 0x4e, 0x6f, 0xbd, 0xb7, 0x21,
	0x30, 0x36, 0xac, 0x9d, 0x6c, 0x82, 0xc7, 0x7b, 0x98, 0x09, 0xe8, 0x23, 0x8c, 0x87, 0xb4, 0xf1,
	0x33, 0x43, 0x38, 0x68, 0x29, 0x78, 0xfe, 0x00, 0xd7, 0x8f, 0xdf, 0xa0, 0xb5, 0xad, 0xfe, 0xbc,
	0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xbd, 0x9c, 0xa2, 0xe6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FtpClient is the client API for Ftp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FtpClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Ftp_UploadClient, error)
	ListFiles(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Ftp_DownloadClient, error)
}

type ftpClient struct {
	cc *grpc.ClientConn
}

func NewFtpClient(cc *grpc.ClientConn) FtpClient {
	return &ftpClient{cc}
}

func (c *ftpClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Ftp_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ftp_serviceDesc.Streams[0], "/proto.Ftp/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &ftpUploadClient{stream}
	return x, nil
}

type Ftp_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type ftpUploadClient struct {
	grpc.ClientStream
}

func (x *ftpUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ftpUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ftpClient) ListFiles(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Ftp/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ftpClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Ftp_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ftp_serviceDesc.Streams[1], "/proto.Ftp/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &ftpDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ftp_DownloadClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type ftpDownloadClient struct {
	grpc.ClientStream
}

func (x *ftpDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FtpServer is the server API for Ftp service.
type FtpServer interface {
	Upload(Ftp_UploadServer) error
	ListFiles(context.Context, *ListRequest) (*ListResponse, error)
	Download(*DownloadRequest, Ftp_DownloadServer) error
}

func RegisterFtpServer(s *grpc.Server, srv FtpServer) {
	s.RegisterService(&_Ftp_serviceDesc, srv)
}

func _Ftp_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FtpServer).Upload(&ftpUploadServer{stream})
}

type Ftp_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type ftpUploadServer struct {
	grpc.ServerStream
}

func (x *ftpUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ftpUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ftp_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FtpServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ftp/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FtpServer).ListFiles(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ftp_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FtpServer).Download(m, &ftpDownloadServer{stream})
}

type Ftp_DownloadServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type ftpDownloadServer struct {
	grpc.ServerStream
}

func (x *ftpDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Ftp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Ftp",
	HandlerType: (*FtpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _Ftp_ListFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Ftp_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _Ftp_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "myftp.proto",
}
