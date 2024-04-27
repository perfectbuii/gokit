// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: hub.proto

package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "math/rand"
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

type GetHubByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HubID string `protobuf:"bytes,1,opt,name=hubID,proto3" json:"hubID,omitempty"`
}

func (x *GetHubByIDRequest) Reset() {
	*x = GetHubByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHubByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHubByIDRequest) ProtoMessage() {}

func (x *GetHubByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHubByIDRequest.ProtoReflect.Descriptor instead.
func (*GetHubByIDRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{0}
}

func (x *GetHubByIDRequest) GetHubID() string {
	if x != nil {
		return x.HubID
	}
	return ""
}

type GetHubByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Hub `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetHubByIDResponse) Reset() {
	*x = GetHubByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHubByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHubByIDResponse) ProtoMessage() {}

func (x *GetHubByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHubByIDResponse.ProtoReflect.Descriptor instead.
func (*GetHubByIDResponse) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{1}
}

func (x *GetHubByIDResponse) GetData() *Hub {
	if x != nil {
		return x.Data
	}
	return nil
}

type Hub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LocationId string `protobuf:"bytes,3,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *Hub) Reset() {
	*x = Hub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hub) ProtoMessage() {}

func (x *Hub) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hub.ProtoReflect.Descriptor instead.
func (*Hub) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{2}
}

func (x *Hub) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Hub) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hub) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

type CreateHubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hub *Hub `protobuf:"bytes,1,opt,name=hub,proto3" json:"hub,omitempty"`
}

func (x *CreateHubRequest) Reset() {
	*x = CreateHubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHubRequest) ProtoMessage() {}

func (x *CreateHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHubRequest.ProtoReflect.Descriptor instead.
func (*CreateHubRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{3}
}

func (x *CreateHubRequest) GetHub() *Hub {
	if x != nil {
		return x.Hub
	}
	return nil
}

type CreateHubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateHubResponse) Reset() {
	*x = CreateHubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHubResponse) ProtoMessage() {}

func (x *CreateHubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHubResponse.ProtoReflect.Descriptor instead.
func (*CreateHubResponse) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{4}
}

func (x *CreateHubResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetListHubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetListHubRequest) Reset() {
	*x = GetListHubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListHubRequest) ProtoMessage() {}

func (x *GetListHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListHubRequest.ProtoReflect.Descriptor instead.
func (*GetListHubRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{5}
}

func (x *GetListHubRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListHubRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetListHubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Hub `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetListHubResponse) Reset() {
	*x = GetListHubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListHubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListHubResponse) ProtoMessage() {}

func (x *GetListHubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListHubResponse.ProtoReflect.Descriptor instead.
func (*GetListHubResponse) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{6}
}

func (x *GetListHubResponse) GetData() []*Hub {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetListHubResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateHubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hub *Hub `protobuf:"bytes,1,opt,name=hub,proto3" json:"hub,omitempty"`
}

func (x *UpdateHubRequest) Reset() {
	*x = UpdateHubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHubRequest) ProtoMessage() {}

func (x *UpdateHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHubRequest.ProtoReflect.Descriptor instead.
func (*UpdateHubRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateHubRequest) GetHub() *Hub {
	if x != nil {
		return x.Hub
	}
	return nil
}

type UpdateHubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateHubResponse) Reset() {
	*x = UpdateHubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHubResponse) ProtoMessage() {}

func (x *UpdateHubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHubResponse.ProtoReflect.Descriptor instead.
func (*UpdateHubResponse) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateHubResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_hub_proto protoreflect.FileDescriptor

var file_hub_proto_rawDesc = []byte{
	0x0a, 0x09, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x68, 0x75, 0x62,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x75, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x68, 0x75, 0x62, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x68, 0x75, 0x62, 0x2e, 0x48, 0x75, 0x62, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x53, 0x0a,
	0x03, 0x48, 0x75, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3a, 0x07, 0xe2, 0xbe, 0x02, 0x03, 0x48,
	0x75, 0x62, 0x22, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x68, 0x75, 0x62, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x48, 0x75, 0x62, 0x52, 0x03, 0x68,
	0x75, 0x62, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x48,
	0x75, 0x62, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2e,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x68, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x48, 0x75, 0x62, 0x52, 0x03, 0x68, 0x75, 0x62, 0x22, 0x2d,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xd5, 0x02,
	0x0a, 0x0a, 0x48, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x68, 0x75, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x62, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x75, 0x62, 0x73, 0x2f, 0x7b, 0x68,
	0x75, 0x62, 0x49, 0x44, 0x7d, 0x12, 0x4f, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x75, 0x62, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x75, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x75, 0x62, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x75, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x75, 0x62, 0x73, 0x12, 0x4f, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75,
	0x62, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x1a, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x75,
	0x62, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x62, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hub_proto_rawDescOnce sync.Once
	file_hub_proto_rawDescData = file_hub_proto_rawDesc
)

func file_hub_proto_rawDescGZIP() []byte {
	file_hub_proto_rawDescOnce.Do(func() {
		file_hub_proto_rawDescData = protoimpl.X.CompressGZIP(file_hub_proto_rawDescData)
	})
	return file_hub_proto_rawDescData
}

var file_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hub_proto_goTypes = []interface{}{
	(*GetHubByIDRequest)(nil),  // 0: hub.GetHubByIDRequest
	(*GetHubByIDResponse)(nil), // 1: hub.GetHubByIDResponse
	(*Hub)(nil),                // 2: hub.Hub
	(*CreateHubRequest)(nil),   // 3: hub.CreateHubRequest
	(*CreateHubResponse)(nil),  // 4: hub.CreateHubResponse
	(*GetListHubRequest)(nil),  // 5: hub.GetListHubRequest
	(*GetListHubResponse)(nil), // 6: hub.GetListHubResponse
	(*UpdateHubRequest)(nil),   // 7: hub.UpdateHubRequest
	(*UpdateHubResponse)(nil),  // 8: hub.UpdateHubResponse
}
var file_hub_proto_depIdxs = []int32{
	2, // 0: hub.GetHubByIDResponse.data:type_name -> hub.Hub
	2, // 1: hub.CreateHubRequest.hub:type_name -> hub.Hub
	2, // 2: hub.GetListHubResponse.data:type_name -> hub.Hub
	2, // 3: hub.UpdateHubRequest.hub:type_name -> hub.Hub
	0, // 4: hub.HubService.GetHubByID:input_type -> hub.GetHubByIDRequest
	3, // 5: hub.HubService.CreateHub:input_type -> hub.CreateHubRequest
	5, // 6: hub.HubService.GetList:input_type -> hub.GetListHubRequest
	7, // 7: hub.HubService.UpdateHub:input_type -> hub.UpdateHubRequest
	1, // 8: hub.HubService.GetHubByID:output_type -> hub.GetHubByIDResponse
	4, // 9: hub.HubService.CreateHub:output_type -> hub.CreateHubResponse
	6, // 10: hub.HubService.GetList:output_type -> hub.GetListHubResponse
	8, // 11: hub.HubService.UpdateHub:output_type -> hub.UpdateHubResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_hub_proto_init() }
func file_hub_proto_init() {
	if File_hub_proto != nil {
		return
	}
	file_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHubByIDRequest); i {
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
		file_hub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHubByIDResponse); i {
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
		file_hub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hub); i {
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
		file_hub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHubRequest); i {
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
		file_hub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHubResponse); i {
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
		file_hub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListHubRequest); i {
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
		file_hub_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListHubResponse); i {
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
		file_hub_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHubRequest); i {
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
		file_hub_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHubResponse); i {
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
			RawDescriptor: file_hub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hub_proto_goTypes,
		DependencyIndexes: file_hub_proto_depIdxs,
		MessageInfos:      file_hub_proto_msgTypes,
	}.Build()
	File_hub_proto = out.File
	file_hub_proto_rawDesc = nil
	file_hub_proto_goTypes = nil
	file_hub_proto_depIdxs = nil
}
