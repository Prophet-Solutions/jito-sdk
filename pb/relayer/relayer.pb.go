// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.24.4
// source: relayer.proto

package relayer

import (
	packet "github.com/Prophet-Solutions/jito-sdk/pb/packet"
	shared "github.com/Prophet-Solutions/jito-sdk/pb/shared"
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

type GetTpuConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTpuConfigsRequest) Reset() {
	*x = GetTpuConfigsRequest{}
	mi := &file_relayer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTpuConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTpuConfigsRequest) ProtoMessage() {}

func (x *GetTpuConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relayer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTpuConfigsRequest.ProtoReflect.Descriptor instead.
func (*GetTpuConfigsRequest) Descriptor() ([]byte, []int) {
	return file_relayer_proto_rawDescGZIP(), []int{0}
}

type GetTpuConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tpu        *shared.Socket `protobuf:"bytes,1,opt,name=tpu,proto3" json:"tpu,omitempty"`
	TpuForward *shared.Socket `protobuf:"bytes,2,opt,name=tpu_forward,json=tpuForward,proto3" json:"tpu_forward,omitempty"`
}

func (x *GetTpuConfigsResponse) Reset() {
	*x = GetTpuConfigsResponse{}
	mi := &file_relayer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTpuConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTpuConfigsResponse) ProtoMessage() {}

func (x *GetTpuConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relayer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTpuConfigsResponse.ProtoReflect.Descriptor instead.
func (*GetTpuConfigsResponse) Descriptor() ([]byte, []int) {
	return file_relayer_proto_rawDescGZIP(), []int{1}
}

func (x *GetTpuConfigsResponse) GetTpu() *shared.Socket {
	if x != nil {
		return x.Tpu
	}
	return nil
}

func (x *GetTpuConfigsResponse) GetTpuForward() *shared.Socket {
	if x != nil {
		return x.TpuForward
	}
	return nil
}

type SubscribePacketsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribePacketsRequest) Reset() {
	*x = SubscribePacketsRequest{}
	mi := &file_relayer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribePacketsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribePacketsRequest) ProtoMessage() {}

func (x *SubscribePacketsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relayer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribePacketsRequest.ProtoReflect.Descriptor instead.
func (*SubscribePacketsRequest) Descriptor() ([]byte, []int) {
	return file_relayer_proto_rawDescGZIP(), []int{2}
}

type SubscribePacketsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Types that are assignable to Msg:
	//
	//	*SubscribePacketsResponse_Heartbeat
	//	*SubscribePacketsResponse_Batch
	Msg isSubscribePacketsResponse_Msg `protobuf_oneof:"msg"`
}

func (x *SubscribePacketsResponse) Reset() {
	*x = SubscribePacketsResponse{}
	mi := &file_relayer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribePacketsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribePacketsResponse) ProtoMessage() {}

func (x *SubscribePacketsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relayer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribePacketsResponse.ProtoReflect.Descriptor instead.
func (*SubscribePacketsResponse) Descriptor() ([]byte, []int) {
	return file_relayer_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribePacketsResponse) GetHeader() *shared.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (m *SubscribePacketsResponse) GetMsg() isSubscribePacketsResponse_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *SubscribePacketsResponse) GetHeartbeat() *shared.Heartbeat {
	if x, ok := x.GetMsg().(*SubscribePacketsResponse_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (x *SubscribePacketsResponse) GetBatch() *packet.PacketBatch {
	if x, ok := x.GetMsg().(*SubscribePacketsResponse_Batch); ok {
		return x.Batch
	}
	return nil
}

type isSubscribePacketsResponse_Msg interface {
	isSubscribePacketsResponse_Msg()
}

type SubscribePacketsResponse_Heartbeat struct {
	Heartbeat *shared.Heartbeat `protobuf:"bytes,2,opt,name=heartbeat,proto3,oneof"`
}

type SubscribePacketsResponse_Batch struct {
	Batch *packet.PacketBatch `protobuf:"bytes,3,opt,name=batch,proto3,oneof"`
}

func (*SubscribePacketsResponse_Heartbeat) isSubscribePacketsResponse_Msg() {}

func (*SubscribePacketsResponse_Batch) isSubscribePacketsResponse_Msg() {}

var File_relayer_proto protoreflect.FileDescriptor

var file_relayer_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x74, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x03, 0x74, 0x70, 0x75, 0x12, 0x2f, 0x0a, 0x0b, 0x74, 0x70, 0x75, 0x5f, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x74, 0x70,
	0x75, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x00,
	0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48,
	0x00, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x32,
	0xb8, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x54, 0x70, 0x75, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x70, 0x75, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a,
	0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x41,
	0x49, 0x4f, 0x2d, 0x44, 0x65, 0x76, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x6a, 0x69, 0x74, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relayer_proto_rawDescOnce sync.Once
	file_relayer_proto_rawDescData = file_relayer_proto_rawDesc
)

func file_relayer_proto_rawDescGZIP() []byte {
	file_relayer_proto_rawDescOnce.Do(func() {
		file_relayer_proto_rawDescData = protoimpl.X.CompressGZIP(file_relayer_proto_rawDescData)
	})
	return file_relayer_proto_rawDescData
}

var file_relayer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_relayer_proto_goTypes = []any{
	(*GetTpuConfigsRequest)(nil),     // 0: relayer.GetTpuConfigsRequest
	(*GetTpuConfigsResponse)(nil),    // 1: relayer.GetTpuConfigsResponse
	(*SubscribePacketsRequest)(nil),  // 2: relayer.SubscribePacketsRequest
	(*SubscribePacketsResponse)(nil), // 3: relayer.SubscribePacketsResponse
	(*shared.Socket)(nil),            // 4: shared.Socket
	(*shared.Header)(nil),            // 5: shared.Header
	(*shared.Heartbeat)(nil),         // 6: shared.Heartbeat
	(*packet.PacketBatch)(nil),       // 7: packet.PacketBatch
}
var file_relayer_proto_depIdxs = []int32{
	4, // 0: relayer.GetTpuConfigsResponse.tpu:type_name -> shared.Socket
	4, // 1: relayer.GetTpuConfigsResponse.tpu_forward:type_name -> shared.Socket
	5, // 2: relayer.SubscribePacketsResponse.header:type_name -> shared.Header
	6, // 3: relayer.SubscribePacketsResponse.heartbeat:type_name -> shared.Heartbeat
	7, // 4: relayer.SubscribePacketsResponse.batch:type_name -> packet.PacketBatch
	0, // 5: relayer.Relayer.GetTpuConfigs:input_type -> relayer.GetTpuConfigsRequest
	2, // 6: relayer.Relayer.SubscribePackets:input_type -> relayer.SubscribePacketsRequest
	1, // 7: relayer.Relayer.GetTpuConfigs:output_type -> relayer.GetTpuConfigsResponse
	3, // 8: relayer.Relayer.SubscribePackets:output_type -> relayer.SubscribePacketsResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_relayer_proto_init() }
func file_relayer_proto_init() {
	if File_relayer_proto != nil {
		return
	}
	file_relayer_proto_msgTypes[3].OneofWrappers = []any{
		(*SubscribePacketsResponse_Heartbeat)(nil),
		(*SubscribePacketsResponse_Batch)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relayer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relayer_proto_goTypes,
		DependencyIndexes: file_relayer_proto_depIdxs,
		MessageInfos:      file_relayer_proto_msgTypes,
	}.Build()
	File_relayer_proto = out.File
	file_relayer_proto_rawDesc = nil
	file_relayer_proto_goTypes = nil
	file_relayer_proto_depIdxs = nil
}
