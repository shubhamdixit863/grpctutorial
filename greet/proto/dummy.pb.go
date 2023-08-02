// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: dummy.proto

package proto

import (
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

type DummyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DummyResponse) Reset() {
	*x = DummyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyResponse) ProtoMessage() {}

func (x *DummyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyResponse.ProtoReflect.Descriptor instead.
func (*DummyResponse) Descriptor() ([]byte, []int) {
	return file_dummy_proto_rawDescGZIP(), []int{0}
}

func (x *DummyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DummyResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DummyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DummyRequest) Reset() {
	*x = DummyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyRequest) ProtoMessage() {}

func (x *DummyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyRequest.ProtoReflect.Descriptor instead.
func (*DummyRequest) Descriptor() ([]byte, []int) {
	return file_dummy_proto_rawDescGZIP(), []int{1}
}

func (x *DummyRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DummyRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_dummy_proto protoreflect.FileDescriptor

var file_dummy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x22, 0x43, 0x0a, 0x0d, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x0c, 0x44, 0x75, 0x6d,
	0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xaa, 0x02,
	0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36,
	0x0a, 0x09, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x18, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x43,
	0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x47, 0x0a, 0x18, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x50, 0x0a, 0x1f, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x75, 0x62, 0x68, 0x61, 0x6d,
	0x64, 0x69, 0x78, 0x69, 0x74, 0x38, 0x36, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x74, 0x75, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dummy_proto_rawDescOnce sync.Once
	file_dummy_proto_rawDescData = file_dummy_proto_rawDesc
)

func file_dummy_proto_rawDescGZIP() []byte {
	file_dummy_proto_rawDescOnce.Do(func() {
		file_dummy_proto_rawDescData = protoimpl.X.CompressGZIP(file_dummy_proto_rawDescData)
	})
	return file_dummy_proto_rawDescData
}

var file_dummy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dummy_proto_goTypes = []interface{}{
	(*DummyResponse)(nil), // 0: greet.DummyResponse
	(*DummyRequest)(nil),  // 1: greet.DummyRequest
}
var file_dummy_proto_depIdxs = []int32{
	1, // 0: greet.DummyService.DummyCall:input_type -> greet.DummyRequest
	1, // 1: greet.DummyService.DummyCallServerStreaming:input_type -> greet.DummyRequest
	1, // 2: greet.DummyService.DummyCallClientStreaming:input_type -> greet.DummyRequest
	1, // 3: greet.DummyService.DummyCallBiDirectionalStreaming:input_type -> greet.DummyRequest
	0, // 4: greet.DummyService.DummyCall:output_type -> greet.DummyResponse
	0, // 5: greet.DummyService.DummyCallServerStreaming:output_type -> greet.DummyResponse
	0, // 6: greet.DummyService.DummyCallClientStreaming:output_type -> greet.DummyResponse
	0, // 7: greet.DummyService.DummyCallBiDirectionalStreaming:output_type -> greet.DummyResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dummy_proto_init() }
func file_dummy_proto_init() {
	if File_dummy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dummy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyResponse); i {
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
		file_dummy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyRequest); i {
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
			RawDescriptor: file_dummy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dummy_proto_goTypes,
		DependencyIndexes: file_dummy_proto_depIdxs,
		MessageInfos:      file_dummy_proto_msgTypes,
	}.Build()
	File_dummy_proto = out.File
	file_dummy_proto_rawDesc = nil
	file_dummy_proto_goTypes = nil
	file_dummy_proto_depIdxs = nil
}
