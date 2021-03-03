// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: crypto.proto

package schema

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CryptoPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Hot    bool   `protobuf:"varint,2,opt,name=hot,proto3" json:"hot,omitempty"`
}

func (x *CryptoPayload) Reset() {
	*x = CryptoPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoPayload) ProtoMessage() {}

func (x *CryptoPayload) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoPayload.ProtoReflect.Descriptor instead.
func (*CryptoPayload) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoPayload) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *CryptoPayload) GetHot() bool {
	if x != nil {
		return x.Hot
	}
	return false
}

type CryptoPayloads struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payloads []*CryptoPayload `protobuf:"bytes,1,rep,name=payloads,proto3" json:"payloads,omitempty"`
}

func (x *CryptoPayloads) Reset() {
	*x = CryptoPayloads{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoPayloads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoPayloads) ProtoMessage() {}

func (x *CryptoPayloads) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoPayloads.ProtoReflect.Descriptor instead.
func (*CryptoPayloads) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoPayloads) GetPayloads() []*CryptoPayload {
	if x != nil {
		return x.Payloads
	}
	return nil
}

type CryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CryptoResponse) Reset() {
	*x = CryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoResponse) ProtoMessage() {}

func (x *CryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoResponse.ProtoReflect.Descriptor instead.
func (*CryptoResponse) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *CryptoResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CryptoResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_crypto_proto protoreflect.FileDescriptor

var file_crypto_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x6f,
	0x74, 0x22, 0x43, 0x0a, 0x0e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x40, 0x0a, 0x0e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x88, 0x02, 0x0a, 0x06, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x12, 0x15, 0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x15, 0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x12, 0x15,
	0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12,
	0x15, 0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x73, 0x12, 0x16, 0x2e, 0x63, 0x79, 0x72, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x1a, 0x16, 0x2e, 0x63, 0x79, 0x72, 0x70,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x67, 0x72, 0x69, 0x7a, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_crypto_proto_rawDescOnce sync.Once
	file_crypto_proto_rawDescData = file_crypto_proto_rawDesc
)

func file_crypto_proto_rawDescGZIP() []byte {
	file_crypto_proto_rawDescOnce.Do(func() {
		file_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_proto_rawDescData)
	})
	return file_crypto_proto_rawDescData
}

var file_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_crypto_proto_goTypes = []interface{}{
	(*CryptoPayload)(nil),  // 0: cyrpto.CryptoPayload
	(*CryptoPayloads)(nil), // 1: cyrpto.CryptoPayloads
	(*CryptoResponse)(nil), // 2: cyrpto.CryptoResponse
}
var file_crypto_proto_depIdxs = []int32{
	0, // 0: cyrpto.CryptoPayloads.payloads:type_name -> cyrpto.CryptoPayload
	0, // 1: cyrpto.Crypto.GetCrypto:input_type -> cyrpto.CryptoPayload
	0, // 2: cyrpto.Crypto.GetCryptos:input_type -> cyrpto.CryptoPayload
	0, // 3: cyrpto.Crypto.CreateCrypto:input_type -> cyrpto.CryptoPayload
	1, // 4: cyrpto.Crypto.CreateCryptos:input_type -> cyrpto.CryptoPayloads
	0, // 5: cyrpto.Crypto.GetCrypto:output_type -> cyrpto.CryptoPayload
	1, // 6: cyrpto.Crypto.GetCryptos:output_type -> cyrpto.CryptoPayloads
	2, // 7: cyrpto.Crypto.CreateCrypto:output_type -> cyrpto.CryptoResponse
	2, // 8: cyrpto.Crypto.CreateCryptos:output_type -> cyrpto.CryptoResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_crypto_proto_init() }
func file_crypto_proto_init() {
	if File_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoPayloads); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_proto_goTypes,
		DependencyIndexes: file_crypto_proto_depIdxs,
		MessageInfos:      file_crypto_proto_msgTypes,
	}.Build()
	File_crypto_proto = out.File
	file_crypto_proto_rawDesc = nil
	file_crypto_proto_goTypes = nil
	file_crypto_proto_depIdxs = nil
}
