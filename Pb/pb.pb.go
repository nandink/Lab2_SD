// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: Pb/pb.proto

package Pb

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

type Mensaje struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Mensaje) Reset() {
	*x = Mensaje{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pb_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensaje) ProtoMessage() {}

func (x *Mensaje) ProtoReflect() protoreflect.Message {
	mi := &file_Pb_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mensaje.ProtoReflect.Descriptor instead.
func (*Mensaje) Descriptor() ([]byte, []int) {
	return file_Pb_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Mensaje) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_Pb_pb_proto protoreflect.FileDescriptor

var file_Pb_pb_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x50, 0x62, 0x2f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x32, 0x33, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x53, 0x76, 0x12, 0x26, 0x0a,
	0x08, 0x44, 0x69, 0x6d, 0x65, 0x48, 0x6f, 0x6c, 0x61, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x6c, 0x61, 0x62, 0x32, 0x2f, 0x50, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Pb_pb_proto_rawDescOnce sync.Once
	file_Pb_pb_proto_rawDescData = file_Pb_pb_proto_rawDesc
)

func file_Pb_pb_proto_rawDescGZIP() []byte {
	file_Pb_pb_proto_rawDescOnce.Do(func() {
		file_Pb_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_Pb_pb_proto_rawDescData)
	})
	return file_Pb_pb_proto_rawDescData
}

var file_Pb_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Pb_pb_proto_goTypes = []interface{}{
	(*Mensaje)(nil), // 0: pb.Mensaje
}
var file_Pb_pb_proto_depIdxs = []int32{
	0, // 0: pb.ClienteSv.DimeHola:input_type -> pb.Mensaje
	0, // 1: pb.ClienteSv.DimeHola:output_type -> pb.Mensaje
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Pb_pb_proto_init() }
func file_Pb_pb_proto_init() {
	if File_Pb_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Pb_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mensaje); i {
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
			RawDescriptor: file_Pb_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Pb_pb_proto_goTypes,
		DependencyIndexes: file_Pb_pb_proto_depIdxs,
		MessageInfos:      file_Pb_pb_proto_msgTypes,
	}.Build()
	File_Pb_pb_proto = out.File
	file_Pb_pb_proto_rawDesc = nil
	file_Pb_pb_proto_goTypes = nil
	file_Pb_pb_proto_depIdxs = nil
}
