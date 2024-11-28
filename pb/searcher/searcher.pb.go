// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.24.4
// source: searcher.proto

package searcher

import (
	bundle "github.com/Prophet-Solutions/jito-sdk/pb/bundle"
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

type SlotList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slots []uint64 `protobuf:"varint,1,rep,packed,name=slots,proto3" json:"slots,omitempty"`
}

func (x *SlotList) Reset() {
	*x = SlotList{}
	mi := &file_searcher_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlotList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotList) ProtoMessage() {}

func (x *SlotList) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotList.ProtoReflect.Descriptor instead.
func (*SlotList) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{0}
}

func (x *SlotList) GetSlots() []uint64 {
	if x != nil {
		return x.Slots
	}
	return nil
}

type ConnectedLeadersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mapping of validator pubkey to leader slots for the current epoch.
	ConnectedValidators map[string]*SlotList `protobuf:"bytes,1,rep,name=connected_validators,json=connectedValidators,proto3" json:"connected_validators,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectedLeadersResponse) Reset() {
	*x = ConnectedLeadersResponse{}
	mi := &file_searcher_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectedLeadersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedLeadersResponse) ProtoMessage() {}

func (x *ConnectedLeadersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedLeadersResponse.ProtoReflect.Descriptor instead.
func (*ConnectedLeadersResponse) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectedLeadersResponse) GetConnectedValidators() map[string]*SlotList {
	if x != nil {
		return x.ConnectedValidators
	}
	return nil
}

type SendBundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundle *bundle.Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *SendBundleRequest) Reset() {
	*x = SendBundleRequest{}
	mi := &file_searcher_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBundleRequest) ProtoMessage() {}

func (x *SendBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBundleRequest.ProtoReflect.Descriptor instead.
func (*SendBundleRequest) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{2}
}

func (x *SendBundleRequest) GetBundle() *bundle.Bundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

type SendBundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// server uuid for the bundle
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *SendBundleResponse) Reset() {
	*x = SendBundleResponse{}
	mi := &file_searcher_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBundleResponse) ProtoMessage() {}

func (x *SendBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBundleResponse.ProtoReflect.Descriptor instead.
func (*SendBundleResponse) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{3}
}

func (x *SendBundleResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type NextScheduledLeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defaults to the currently connected region if no region provided.
	Regions []string `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *NextScheduledLeaderRequest) Reset() {
	*x = NextScheduledLeaderRequest{}
	mi := &file_searcher_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NextScheduledLeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextScheduledLeaderRequest) ProtoMessage() {}

func (x *NextScheduledLeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextScheduledLeaderRequest.ProtoReflect.Descriptor instead.
func (*NextScheduledLeaderRequest) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{4}
}

func (x *NextScheduledLeaderRequest) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

type NextScheduledLeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the current slot the backend is on
	CurrentSlot uint64 `protobuf:"varint,1,opt,name=current_slot,json=currentSlot,proto3" json:"current_slot,omitempty"`
	// the slot of the next leader
	NextLeaderSlot uint64 `protobuf:"varint,2,opt,name=next_leader_slot,json=nextLeaderSlot,proto3" json:"next_leader_slot,omitempty"`
	// the identity pubkey (base58) of the next leader
	NextLeaderIdentity string `protobuf:"bytes,3,opt,name=next_leader_identity,json=nextLeaderIdentity,proto3" json:"next_leader_identity,omitempty"`
	// the block engine region of the next leader
	NextLeaderRegion string `protobuf:"bytes,4,opt,name=next_leader_region,json=nextLeaderRegion,proto3" json:"next_leader_region,omitempty"`
}

func (x *NextScheduledLeaderResponse) Reset() {
	*x = NextScheduledLeaderResponse{}
	mi := &file_searcher_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NextScheduledLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextScheduledLeaderResponse) ProtoMessage() {}

func (x *NextScheduledLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextScheduledLeaderResponse.ProtoReflect.Descriptor instead.
func (*NextScheduledLeaderResponse) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{5}
}

func (x *NextScheduledLeaderResponse) GetCurrentSlot() uint64 {
	if x != nil {
		return x.CurrentSlot
	}
	return 0
}

func (x *NextScheduledLeaderResponse) GetNextLeaderSlot() uint64 {
	if x != nil {
		return x.NextLeaderSlot
	}
	return 0
}

func (x *NextScheduledLeaderResponse) GetNextLeaderIdentity() string {
	if x != nil {
		return x.NextLeaderIdentity
	}
	return ""
}

func (x *NextScheduledLeaderResponse) GetNextLeaderRegion() string {
	if x != nil {
		return x.NextLeaderRegion
	}
	return ""
}

type ConnectedLeadersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectedLeadersRequest) Reset() {
	*x = ConnectedLeadersRequest{}
	mi := &file_searcher_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectedLeadersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedLeadersRequest) ProtoMessage() {}

func (x *ConnectedLeadersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedLeadersRequest.ProtoReflect.Descriptor instead.
func (*ConnectedLeadersRequest) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{6}
}

type ConnectedLeadersRegionedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defaults to the currently connected region if no region provided.
	Regions []string `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *ConnectedLeadersRegionedRequest) Reset() {
	*x = ConnectedLeadersRegionedRequest{}
	mi := &file_searcher_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectedLeadersRegionedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedLeadersRegionedRequest) ProtoMessage() {}

func (x *ConnectedLeadersRegionedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedLeadersRegionedRequest.ProtoReflect.Descriptor instead.
func (*ConnectedLeadersRegionedRequest) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{7}
}

func (x *ConnectedLeadersRegionedRequest) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

type ConnectedLeadersRegionedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectedValidators map[string]*ConnectedLeadersResponse `protobuf:"bytes,1,rep,name=connected_validators,json=connectedValidators,proto3" json:"connected_validators,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectedLeadersRegionedResponse) Reset() {
	*x = ConnectedLeadersRegionedResponse{}
	mi := &file_searcher_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectedLeadersRegionedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedLeadersRegionedResponse) ProtoMessage() {}

func (x *ConnectedLeadersRegionedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedLeadersRegionedResponse.ProtoReflect.Descriptor instead.
func (*ConnectedLeadersRegionedResponse) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{8}
}

func (x *ConnectedLeadersRegionedResponse) GetConnectedValidators() map[string]*ConnectedLeadersResponse {
	if x != nil {
		return x.ConnectedValidators
	}
	return nil
}

type GetTipAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTipAccountsRequest) Reset() {
	*x = GetTipAccountsRequest{}
	mi := &file_searcher_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTipAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTipAccountsRequest) ProtoMessage() {}

func (x *GetTipAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTipAccountsRequest.ProtoReflect.Descriptor instead.
func (*GetTipAccountsRequest) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{9}
}

type GetTipAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *GetTipAccountsResponse) Reset() {
	*x = GetTipAccountsResponse{}
	mi := &file_searcher_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTipAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTipAccountsResponse) ProtoMessage() {}

func (x *GetTipAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTipAccountsResponse.ProtoReflect.Descriptor instead.
func (*GetTipAccountsResponse) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{10}
}

func (x *GetTipAccountsResponse) GetAccounts() []string {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type SubscribeBundleResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeBundleResultsRequest) Reset() {
	*x = SubscribeBundleResultsRequest{}
	mi := &file_searcher_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeBundleResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeBundleResultsRequest) ProtoMessage() {}

func (x *SubscribeBundleResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeBundleResultsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeBundleResultsRequest) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{11}
}

type GetRegionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRegionsRequest) Reset() {
	*x = GetRegionsRequest{}
	mi := &file_searcher_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRegionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionsRequest) ProtoMessage() {}

func (x *GetRegionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionsRequest.ProtoReflect.Descriptor instead.
func (*GetRegionsRequest) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{12}
}

type GetRegionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The region the client is currently connected to
	CurrentRegion string `protobuf:"bytes,1,opt,name=current_region,json=currentRegion,proto3" json:"current_region,omitempty"`
	// Regions that are online and ready for connections
	// All regions: https://jito-labs.gitbook.io/mev/systems/connecting/mainnet
	AvailableRegions []string `protobuf:"bytes,2,rep,name=available_regions,json=availableRegions,proto3" json:"available_regions,omitempty"`
}

func (x *GetRegionsResponse) Reset() {
	*x = GetRegionsResponse{}
	mi := &file_searcher_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRegionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionsResponse) ProtoMessage() {}

func (x *GetRegionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searcher_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionsResponse.ProtoReflect.Descriptor instead.
func (*GetRegionsResponse) Descriptor() ([]byte, []int) {
	return file_searcher_proto_rawDescGZIP(), []int{13}
}

func (x *GetRegionsResponse) GetCurrentRegion() string {
	if x != nil {
		return x.CurrentRegion
	}
	return ""
}

func (x *GetRegionsResponse) GetAvailableRegions() []string {
	if x != nil {
		return x.AvailableRegions
	}
	return nil
}

var File_searcher_proto protoreflect.FileDescriptor

var file_searcher_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x1a, 0x0c, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x08, 0x53, 0x6c, 0x6f, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x18, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x5a, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x22, 0x28, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x1a, 0x4e, 0x65,
	0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x1b, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x6e, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12,
	0x30, 0x0a, 0x14, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e,
	0x65, 0x78, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e,
	0x65, 0x78, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22,
	0x19, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x1f, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x20, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x14,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x1a, 0x6a, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x69, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x1f, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x68, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x32,
	0x9c, 0x05, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x27, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x49, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1b,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x65, 0x64, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x69, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1f,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x70,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x69,
	0x70, 0x65, 0x72, 0x41, 0x49, 0x4f, 0x2d, 0x44, 0x65, 0x76, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x6a,
	0x69, 0x74, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_searcher_proto_rawDescOnce sync.Once
	file_searcher_proto_rawDescData = file_searcher_proto_rawDesc
)

func file_searcher_proto_rawDescGZIP() []byte {
	file_searcher_proto_rawDescOnce.Do(func() {
		file_searcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_searcher_proto_rawDescData)
	})
	return file_searcher_proto_rawDescData
}

var file_searcher_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_searcher_proto_goTypes = []any{
	(*SlotList)(nil),                         // 0: searcher.SlotList
	(*ConnectedLeadersResponse)(nil),         // 1: searcher.ConnectedLeadersResponse
	(*SendBundleRequest)(nil),                // 2: searcher.SendBundleRequest
	(*SendBundleResponse)(nil),               // 3: searcher.SendBundleResponse
	(*NextScheduledLeaderRequest)(nil),       // 4: searcher.NextScheduledLeaderRequest
	(*NextScheduledLeaderResponse)(nil),      // 5: searcher.NextScheduledLeaderResponse
	(*ConnectedLeadersRequest)(nil),          // 6: searcher.ConnectedLeadersRequest
	(*ConnectedLeadersRegionedRequest)(nil),  // 7: searcher.ConnectedLeadersRegionedRequest
	(*ConnectedLeadersRegionedResponse)(nil), // 8: searcher.ConnectedLeadersRegionedResponse
	(*GetTipAccountsRequest)(nil),            // 9: searcher.GetTipAccountsRequest
	(*GetTipAccountsResponse)(nil),           // 10: searcher.GetTipAccountsResponse
	(*SubscribeBundleResultsRequest)(nil),    // 11: searcher.SubscribeBundleResultsRequest
	(*GetRegionsRequest)(nil),                // 12: searcher.GetRegionsRequest
	(*GetRegionsResponse)(nil),               // 13: searcher.GetRegionsResponse
	nil,                                      // 14: searcher.ConnectedLeadersResponse.ConnectedValidatorsEntry
	nil,                                      // 15: searcher.ConnectedLeadersRegionedResponse.ConnectedValidatorsEntry
	(*bundle.Bundle)(nil),                    // 16: bundle.Bundle
	(*bundle.BundleResult)(nil),              // 17: bundle.BundleResult
}
var file_searcher_proto_depIdxs = []int32{
	14, // 0: searcher.ConnectedLeadersResponse.connected_validators:type_name -> searcher.ConnectedLeadersResponse.ConnectedValidatorsEntry
	16, // 1: searcher.SendBundleRequest.bundle:type_name -> bundle.Bundle
	15, // 2: searcher.ConnectedLeadersRegionedResponse.connected_validators:type_name -> searcher.ConnectedLeadersRegionedResponse.ConnectedValidatorsEntry
	0,  // 3: searcher.ConnectedLeadersResponse.ConnectedValidatorsEntry.value:type_name -> searcher.SlotList
	1,  // 4: searcher.ConnectedLeadersRegionedResponse.ConnectedValidatorsEntry.value:type_name -> searcher.ConnectedLeadersResponse
	11, // 5: searcher.SearcherService.SubscribeBundleResults:input_type -> searcher.SubscribeBundleResultsRequest
	2,  // 6: searcher.SearcherService.SendBundle:input_type -> searcher.SendBundleRequest
	4,  // 7: searcher.SearcherService.GetNextScheduledLeader:input_type -> searcher.NextScheduledLeaderRequest
	6,  // 8: searcher.SearcherService.GetConnectedLeaders:input_type -> searcher.ConnectedLeadersRequest
	7,  // 9: searcher.SearcherService.GetConnectedLeadersRegioned:input_type -> searcher.ConnectedLeadersRegionedRequest
	9,  // 10: searcher.SearcherService.GetTipAccounts:input_type -> searcher.GetTipAccountsRequest
	12, // 11: searcher.SearcherService.GetRegions:input_type -> searcher.GetRegionsRequest
	17, // 12: searcher.SearcherService.SubscribeBundleResults:output_type -> bundle.BundleResult
	3,  // 13: searcher.SearcherService.SendBundle:output_type -> searcher.SendBundleResponse
	5,  // 14: searcher.SearcherService.GetNextScheduledLeader:output_type -> searcher.NextScheduledLeaderResponse
	1,  // 15: searcher.SearcherService.GetConnectedLeaders:output_type -> searcher.ConnectedLeadersResponse
	8,  // 16: searcher.SearcherService.GetConnectedLeadersRegioned:output_type -> searcher.ConnectedLeadersRegionedResponse
	10, // 17: searcher.SearcherService.GetTipAccounts:output_type -> searcher.GetTipAccountsResponse
	13, // 18: searcher.SearcherService.GetRegions:output_type -> searcher.GetRegionsResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_searcher_proto_init() }
func file_searcher_proto_init() {
	if File_searcher_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_searcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_searcher_proto_goTypes,
		DependencyIndexes: file_searcher_proto_depIdxs,
		MessageInfos:      file_searcher_proto_msgTypes,
	}.Build()
	File_searcher_proto = out.File
	file_searcher_proto_rawDesc = nil
	file_searcher_proto_goTypes = nil
	file_searcher_proto_depIdxs = nil
}
