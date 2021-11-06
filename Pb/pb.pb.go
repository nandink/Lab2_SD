// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pb.proto

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
		mi := &file_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensaje) ProtoMessage() {}

func (x *Mensaje) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[0]
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
	return file_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Mensaje) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Mensajito2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Mensajito2) Reset() {
	*x = Mensajito2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensajito2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensajito2) ProtoMessage() {}

func (x *Mensajito2) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mensajito2.ProtoReflect.Descriptor instead.
func (*Mensajito2) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{1}
}

func (x *Mensajito2) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Jugada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jugador []int32 `protobuf:"varint,1,rep,packed,name=jugador,proto3" json:"jugador,omitempty"`
	Ronda   int32   `protobuf:"varint,2,opt,name=ronda,proto3" json:"ronda,omitempty"`
	Muertos []int32 `protobuf:"varint,3,rep,packed,name=muertos,proto3" json:"muertos,omitempty"`
}

func (x *Jugada) Reset() {
	*x = Jugada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jugada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jugada) ProtoMessage() {}

func (x *Jugada) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jugada.ProtoReflect.Descriptor instead.
func (*Jugada) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Jugada) GetJugador() []int32 {
	if x != nil {
		return x.Jugador
	}
	return nil
}

func (x *Jugada) GetRonda() int32 {
	if x != nil {
		return x.Ronda
	}
	return 0
}

func (x *Jugada) GetMuertos() []int32 {
	if x != nil {
		return x.Muertos
	}
	return nil
}

type Monto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monto int64 `protobuf:"varint,1,opt,name=monto,proto3" json:"monto,omitempty"`
}

func (x *Monto) Reset() {
	*x = Monto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monto) ProtoMessage() {}

func (x *Monto) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monto.ProtoReflect.Descriptor instead.
func (*Monto) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{3}
}

func (x *Monto) GetMonto() int64 {
	if x != nil {
		return x.Monto
	}
	return 0
}

var File_pb_proto protoreflect.FileDescriptor

var file_pb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x1d,
	0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x1c, 0x0a,
	0x0a, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x69, 0x74, 0x6f, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x06, 0x4a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6e, 0x64, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x72, 0x6f, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x75, 0x65, 0x72, 0x74, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x75, 0x65, 0x72, 0x74, 0x6f, 0x73, 0x22,
	0x1d, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x32, 0x92,
	0x01, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x53, 0x76, 0x12, 0x26, 0x0a, 0x08,
	0x44, 0x69, 0x6d, 0x65, 0x48, 0x6f, 0x6c, 0x61, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x4a, 0x75,
	0x67, 0x61, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x69, 0x74, 0x6f, 0x32, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x69, 0x74, 0x6f, 0x32, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0c, 0x4d, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64,
	0x61, 0x22, 0x00, 0x32, 0x38, 0x0a, 0x07, 0x4c, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x4e, 0x12, 0x2d,
	0x0a, 0x0d, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73, 0x12,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x1a, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x69, 0x74, 0x6f, 0x32, 0x22, 0x00, 0x32, 0x36, 0x0a,
	0x09, 0x4c, 0x69, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x7a, 0x6f, 0x12, 0x29, 0x0a, 0x0a, 0x50, 0x65,
	0x64, 0x69, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x6e, 0x73, 0x61, 0x6a, 0x69, 0x74, 0x6f, 0x32, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f,
	0x6e, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x4c, 0x61, 0x62, 0x32, 0x5f, 0x53, 0x44,
	0x2f, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_rawDescOnce sync.Once
	file_pb_proto_rawDescData = file_pb_proto_rawDesc
)

func file_pb_proto_rawDescGZIP() []byte {
	file_pb_proto_rawDescOnce.Do(func() {
		file_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_rawDescData)
	})
	return file_pb_proto_rawDescData
}

var file_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_proto_goTypes = []interface{}{
	(*Mensaje)(nil),    // 0: pb.Mensaje
	(*Mensajito2)(nil), // 1: pb.Mensajito2
	(*Jugada)(nil),     // 2: pb.Jugada
	(*Monto)(nil),      // 3: pb.Monto
}
var file_pb_proto_depIdxs = []int32{
	0, // 0: pb.ClienteSv.DimeHola:input_type -> pb.Mensaje
	1, // 1: pb.ClienteSv.MandarJugadores:input_type -> pb.Mensajito2
	2, // 2: pb.ClienteSv.MandarJugada:input_type -> pb.Jugada
	2, // 3: pb.LiderNN.EnviarJugadas:input_type -> pb.Jugada
	1, // 4: pb.LiderPozo.PedirMonto:input_type -> pb.Mensajito2
	0, // 5: pb.ClienteSv.DimeHola:output_type -> pb.Mensaje
	1, // 6: pb.ClienteSv.MandarJugadores:output_type -> pb.Mensajito2
	2, // 7: pb.ClienteSv.MandarJugada:output_type -> pb.Jugada
	1, // 8: pb.LiderNN.EnviarJugadas:output_type -> pb.Mensajito2
	3, // 9: pb.LiderPozo.PedirMonto:output_type -> pb.Monto
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_proto_init() }
func file_pb_proto_init() {
	if File_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mensajito2); i {
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
		file_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jugada); i {
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
		file_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monto); i {
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
			RawDescriptor: file_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_pb_proto_goTypes,
		DependencyIndexes: file_pb_proto_depIdxs,
		MessageInfos:      file_pb_proto_msgTypes,
	}.Build()
	File_pb_proto = out.File
	file_pb_proto_rawDesc = nil
	file_pb_proto_goTypes = nil
	file_pb_proto_depIdxs = nil
}
