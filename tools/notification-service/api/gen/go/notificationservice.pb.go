// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: notificationservice.proto

package _go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The ErrCode enum defines errors for gRPC API functions. These are converted
// from the Go error types returned by gnoclient.
type ErrCode int32

const (
	// Undefined is the default value. It should never be set manually
	ErrCode_Undefined ErrCode = 0
	// TODO indicates that you plan to create an error later
	ErrCode_TODO ErrCode = 1
	// ErrNotImplemented indicates that a method is not implemented yet
	ErrCode_ErrNotImplemented ErrCode = 2
	// ErrInternal indicates an unknown error (without Code), i.e. in gRPC
	ErrCode_ErrInternal        ErrCode = 3
	ErrCode_ErrInvalidInput    ErrCode = 100
	ErrCode_ErrMissingInput    ErrCode = 101
	ErrCode_ErrSerialization   ErrCode = 102
	ErrCode_ErrDeserialization ErrCode = 103
	ErrCode_ErrInitService     ErrCode = 104
	ErrCode_ErrRunGRPCServer   ErrCode = 105
	ErrCode_ErrDB              ErrCode = 200
	ErrCode_ErrDBNotFound      ErrCode = 201
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:   "Undefined",
		1:   "TODO",
		2:   "ErrNotImplemented",
		3:   "ErrInternal",
		100: "ErrInvalidInput",
		101: "ErrMissingInput",
		102: "ErrSerialization",
		103: "ErrDeserialization",
		104: "ErrInitService",
		105: "ErrRunGRPCServer",
		200: "ErrDB",
		201: "ErrDBNotFound",
	}
	ErrCode_value = map[string]int32{
		"Undefined":          0,
		"TODO":               1,
		"ErrNotImplemented":  2,
		"ErrInternal":        3,
		"ErrInvalidInput":    100,
		"ErrMissingInput":    101,
		"ErrSerialization":   102,
		"ErrDeserialization": 103,
		"ErrInitService":     104,
		"ErrRunGRPCServer":   105,
		"ErrDB":              200,
		"ErrDBNotFound":      201,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_notificationservice_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_notificationservice_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{0}
}

type RegisterDeviceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterDeviceRequest) Reset() {
	*x = RegisterDeviceRequest{}
	mi := &file_notificationservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceRequest) ProtoMessage() {}

func (x *RegisterDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceRequest.ProtoReflect.Descriptor instead.
func (*RegisterDeviceRequest) Descriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterDeviceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterDeviceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegisterDeviceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterDeviceResponse) Reset() {
	*x = RegisterDeviceResponse{}
	mi := &file_notificationservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceResponse) ProtoMessage() {}

func (x *RegisterDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceResponse.ProtoReflect.Descriptor instead.
func (*RegisterDeviceResponse) Descriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{1}
}

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_notificationservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greeting      string                 `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_notificationservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{3}
}

func (x *HelloResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type HelloStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloStreamRequest) Reset() {
	*x = HelloStreamRequest{}
	mi := &file_notificationservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloStreamRequest) ProtoMessage() {}

func (x *HelloStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloStreamRequest.ProtoReflect.Descriptor instead.
func (*HelloStreamRequest) Descriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{4}
}

func (x *HelloStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greeting      string                 `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloStreamResponse) Reset() {
	*x = HelloStreamResponse{}
	mi := &file_notificationservice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloStreamResponse) ProtoMessage() {}

func (x *HelloStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloStreamResponse.ProtoReflect.Descriptor instead.
func (*HelloStreamResponse) Descriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{5}
}

func (x *HelloStreamResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type ErrDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Codes         []ErrCode              `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=land.gno.dsocial.notificationservice.v1.ErrCode" json:"codes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrDetails) Reset() {
	*x = ErrDetails{}
	mi := &file_notificationservice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrDetails) ProtoMessage() {}

func (x *ErrDetails) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrDetails.ProtoReflect.Descriptor instead.
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return file_notificationservice_proto_rawDescGZIP(), []int{6}
}

func (x *ErrDetails) GetCodes() []ErrCode {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_notificationservice_proto protoreflect.FileDescriptor

const file_notificationservice_proto_rawDesc = "" +
	"\n" +
	"\x19notificationservice.proto\x12'land.gno.dsocial.notificationservice.v1\"G\n" +
	"\x15RegisterDeviceRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\"\x18\n" +
	"\x16RegisterDeviceResponse\"\"\n" +
	"\fHelloRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"+\n" +
	"\rHelloResponse\x12\x1a\n" +
	"\bgreeting\x18\x01 \x01(\tR\bgreeting\"(\n" +
	"\x12HelloStreamRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"1\n" +
	"\x13HelloStreamResponse\x12\x1a\n" +
	"\bgreeting\x18\x01 \x01(\tR\bgreeting\"T\n" +
	"\n" +
	"ErrDetails\x12F\n" +
	"\x05codes\x18\x01 \x03(\x0e20.land.gno.dsocial.notificationservice.v1.ErrCodeR\x05codes*\xec\x01\n" +
	"\aErrCode\x12\r\n" +
	"\tUndefined\x10\x00\x12\b\n" +
	"\x04TODO\x10\x01\x12\x15\n" +
	"\x11ErrNotImplemented\x10\x02\x12\x0f\n" +
	"\vErrInternal\x10\x03\x12\x13\n" +
	"\x0fErrInvalidInput\x10d\x12\x13\n" +
	"\x0fErrMissingInput\x10e\x12\x14\n" +
	"\x10ErrSerialization\x10f\x12\x16\n" +
	"\x12ErrDeserialization\x10g\x12\x12\n" +
	"\x0eErrInitService\x10h\x12\x14\n" +
	"\x10ErrRunGRPCServer\x10i\x12\n" +
	"\n" +
	"\x05ErrDB\x10\xc8\x01\x12\x12\n" +
	"\rErrDBNotFound\x10\xc9\x012\xae\x03\n" +
	"\x13NotificationService\x12\x91\x01\n" +
	"\x0eRegisterDevice\x12>.land.gno.dsocial.notificationservice.v1.RegisterDeviceRequest\x1a?.land.gno.dsocial.notificationservice.v1.RegisterDeviceResponse\x12v\n" +
	"\x05Hello\x125.land.gno.dsocial.notificationservice.v1.HelloRequest\x1a6.land.gno.dsocial.notificationservice.v1.HelloResponse\x12\x8a\x01\n" +
	"\vHelloStream\x12;.land.gno.dsocial.notificationservice.v1.HelloStreamRequest\x1a<.land.gno.dsocial.notificationservice.v1.HelloStreamResponse0\x01BIZAgithub.com/gnoverse/dsocial/tools/notification-service/api/gen/go\xa2\x02\x03RTGb\x06proto3"

var (
	file_notificationservice_proto_rawDescOnce sync.Once
	file_notificationservice_proto_rawDescData []byte
)

func file_notificationservice_proto_rawDescGZIP() []byte {
	file_notificationservice_proto_rawDescOnce.Do(func() {
		file_notificationservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notificationservice_proto_rawDesc), len(file_notificationservice_proto_rawDesc)))
	})
	return file_notificationservice_proto_rawDescData
}

var file_notificationservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notificationservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_notificationservice_proto_goTypes = []any{
	(ErrCode)(0),                   // 0: land.gno.dsocial.notificationservice.v1.ErrCode
	(*RegisterDeviceRequest)(nil),  // 1: land.gno.dsocial.notificationservice.v1.RegisterDeviceRequest
	(*RegisterDeviceResponse)(nil), // 2: land.gno.dsocial.notificationservice.v1.RegisterDeviceResponse
	(*HelloRequest)(nil),           // 3: land.gno.dsocial.notificationservice.v1.HelloRequest
	(*HelloResponse)(nil),          // 4: land.gno.dsocial.notificationservice.v1.HelloResponse
	(*HelloStreamRequest)(nil),     // 5: land.gno.dsocial.notificationservice.v1.HelloStreamRequest
	(*HelloStreamResponse)(nil),    // 6: land.gno.dsocial.notificationservice.v1.HelloStreamResponse
	(*ErrDetails)(nil),             // 7: land.gno.dsocial.notificationservice.v1.ErrDetails
}
var file_notificationservice_proto_depIdxs = []int32{
	0, // 0: land.gno.dsocial.notificationservice.v1.ErrDetails.codes:type_name -> land.gno.dsocial.notificationservice.v1.ErrCode
	1, // 1: land.gno.dsocial.notificationservice.v1.NotificationService.RegisterDevice:input_type -> land.gno.dsocial.notificationservice.v1.RegisterDeviceRequest
	3, // 2: land.gno.dsocial.notificationservice.v1.NotificationService.Hello:input_type -> land.gno.dsocial.notificationservice.v1.HelloRequest
	5, // 3: land.gno.dsocial.notificationservice.v1.NotificationService.HelloStream:input_type -> land.gno.dsocial.notificationservice.v1.HelloStreamRequest
	2, // 4: land.gno.dsocial.notificationservice.v1.NotificationService.RegisterDevice:output_type -> land.gno.dsocial.notificationservice.v1.RegisterDeviceResponse
	4, // 5: land.gno.dsocial.notificationservice.v1.NotificationService.Hello:output_type -> land.gno.dsocial.notificationservice.v1.HelloResponse
	6, // 6: land.gno.dsocial.notificationservice.v1.NotificationService.HelloStream:output_type -> land.gno.dsocial.notificationservice.v1.HelloStreamResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notificationservice_proto_init() }
func file_notificationservice_proto_init() {
	if File_notificationservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notificationservice_proto_rawDesc), len(file_notificationservice_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notificationservice_proto_goTypes,
		DependencyIndexes: file_notificationservice_proto_depIdxs,
		EnumInfos:         file_notificationservice_proto_enumTypes,
		MessageInfos:      file_notificationservice_proto_msgTypes,
	}.Build()
	File_notificationservice_proto = out.File
	file_notificationservice_proto_goTypes = nil
	file_notificationservice_proto_depIdxs = nil
}
