// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: v1/orders.proto

package v1

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

type GetOrderHistoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetOrderHistoryReq) Reset() {
	*x = GetOrderHistoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryReq) ProtoMessage() {}

func (x *GetOrderHistoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryReq.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryReq) Descriptor() ([]byte, []int) {
	return file_v1_orders_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrderHistoryReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetOrderHistoryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrderHistoryRes) Reset() {
	*x = GetOrderHistoryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryRes) ProtoMessage() {}

func (x *GetOrderHistoryRes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryRes.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryRes) Descriptor() ([]byte, []int) {
	return file_v1_orders_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderHistoryRes) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Net         string  `protobuf:"bytes,2,opt,name=net,proto3" json:"net,omitempty"`
	CoinAddrUrl string  `protobuf:"bytes,3,opt,name=coin_addr_url,json=coinAddrUrl,proto3" json:"coin_addr_url,omitempty"`
	Status      string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt   int64   `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ConfirmedAt int64   `protobuf:"varint,6,opt,name=confirmed_at,json=confirmedAt,proto3" json:"confirmed_at,omitempty"`
	Am          float64 `protobuf:"fixed64,7,opt,name=am,proto3" json:"am,omitempty"`
	ToWallet    string  `protobuf:"bytes,8,opt,name=to_wallet,json=toWallet,proto3" json:"to_wallet,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_v1_orders_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetNet() string {
	if x != nil {
		return x.Net
	}
	return ""
}

func (x *Order) GetCoinAddrUrl() string {
	if x != nil {
		return x.CoinAddrUrl
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Order) GetConfirmedAt() int64 {
	if x != nil {
		return x.ConfirmedAt
	}
	return 0
}

func (x *Order) GetAm() float64 {
	if x != nil {
		return x.Am
	}
	return 0
}

func (x *Order) GetToWallet() string {
	if x != nil {
		return x.ToWallet
	}
	return ""
}

type CreateOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Net    string `protobuf:"bytes,2,opt,name=net,proto3" json:"net,omitempty"`
	Am     int64  `protobuf:"varint,3,opt,name=am,proto3" json:"am,omitempty"`
}

func (x *CreateOrderReq) Reset() {
	*x = CreateOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderReq) ProtoMessage() {}

func (x *CreateOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderReq.ProtoReflect.Descriptor instead.
func (*CreateOrderReq) Descriptor() ([]byte, []int) {
	return file_v1_orders_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderReq) GetNet() string {
	if x != nil {
		return x.Net
	}
	return ""
}

func (x *CreateOrderReq) GetAm() int64 {
	if x != nil {
		return x.Am
	}
	return 0
}

type CreateOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CoinAddrUrl string  `protobuf:"bytes,2,opt,name=coin_addr_url,json=coinAddrUrl,proto3" json:"coin_addr_url,omitempty"`
	Am          float64 `protobuf:"fixed64,3,opt,name=am,proto3" json:"am,omitempty"`
	ToWallet    string  `protobuf:"bytes,4,opt,name=to_wallet,json=toWallet,proto3" json:"to_wallet,omitempty"`
}

func (x *CreateOrderResp) Reset() {
	*x = CreateOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResp) ProtoMessage() {}

func (x *CreateOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResp.ProtoReflect.Descriptor instead.
func (*CreateOrderResp) Descriptor() ([]byte, []int) {
	return file_v1_orders_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderResp) GetCoinAddrUrl() string {
	if x != nil {
		return x.CoinAddrUrl
	}
	return ""
}

func (x *CreateOrderResp) GetAm() float64 {
	if x != nil {
		return x.Am
	}
	return 0
}

func (x *CreateOrderResp) GetToWallet() string {
	if x != nil {
		return x.ToWallet
	}
	return ""
}

type CheckOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckOrderReq) Reset() {
	*x = CheckOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOrderReq) ProtoMessage() {}

func (x *CheckOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOrderReq.ProtoReflect.Descriptor instead.
func (*CheckOrderReq) Descriptor() ([]byte, []int) {
	return file_v1_orders_proto_rawDescGZIP(), []int{5}
}

func (x *CheckOrderReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CheckOrderResp) Reset() {
	*x = CheckOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOrderResp) ProtoMessage() {}

func (x *CheckOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOrderResp.ProtoReflect.Descriptor instead.
func (*CheckOrderResp) Descriptor() ([]byte, []int) {
	return file_v1_orders_proto_rawDescGZIP(), []int{6}
}

func (x *CheckOrderResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_v1_orders_proto protoreflect.FileDescriptor

var file_v1_orders_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x0e, 0x76, 0x31,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6e, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x02, 0x61, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x22, 0x4b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x65, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x6d, 0x22,
	0x72, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x69, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x02, 0x61, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xef,
	0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x1d, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x63, 0x72, 0x79, 0x70, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_orders_proto_rawDescOnce sync.Once
	file_v1_orders_proto_rawDescData = file_v1_orders_proto_rawDesc
)

func file_v1_orders_proto_rawDescGZIP() []byte {
	file_v1_orders_proto_rawDescOnce.Do(func() {
		file_v1_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_orders_proto_rawDescData)
	})
	return file_v1_orders_proto_rawDescData
}

var file_v1_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_orders_proto_goTypes = []interface{}{
	(*GetOrderHistoryReq)(nil), // 0: crypay.v1.GetOrderHistoryReq
	(*GetOrderHistoryRes)(nil), // 1: crypay.v1.GetOrderHistoryRes
	(*Order)(nil),              // 2: crypay.v1.Order
	(*CreateOrderReq)(nil),     // 3: crypay.v1.CreateOrderReq
	(*CreateOrderResp)(nil),    // 4: crypay.v1.CreateOrderResp
	(*CheckOrderReq)(nil),      // 5: crypay.v1.CheckOrderReq
	(*CheckOrderResp)(nil),     // 6: crypay.v1.CheckOrderResp
}
var file_v1_orders_proto_depIdxs = []int32{
	2, // 0: crypay.v1.GetOrderHistoryRes.orders:type_name -> crypay.v1.Order
	3, // 1: crypay.v1.OrdersService.CreateOrder:input_type -> crypay.v1.CreateOrderReq
	5, // 2: crypay.v1.OrdersService.CheckOrder:input_type -> crypay.v1.CheckOrderReq
	0, // 3: crypay.v1.OrdersService.GetOrderHistory:input_type -> crypay.v1.GetOrderHistoryReq
	4, // 4: crypay.v1.OrdersService.CreateOrder:output_type -> crypay.v1.CreateOrderResp
	6, // 5: crypay.v1.OrdersService.CheckOrder:output_type -> crypay.v1.CheckOrderResp
	1, // 6: crypay.v1.OrdersService.GetOrderHistory:output_type -> crypay.v1.GetOrderHistoryRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_orders_proto_init() }
func file_v1_orders_proto_init() {
	if File_v1_orders_proto != nil {
		return
	}
	file_v1_empty_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryReq); i {
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
		file_v1_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryRes); i {
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
		file_v1_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_v1_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderReq); i {
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
		file_v1_orders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResp); i {
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
		file_v1_orders_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOrderReq); i {
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
		file_v1_orders_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOrderResp); i {
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
			RawDescriptor: file_v1_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_orders_proto_goTypes,
		DependencyIndexes: file_v1_orders_proto_depIdxs,
		MessageInfos:      file_v1_orders_proto_msgTypes,
	}.Build()
	File_v1_orders_proto = out.File
	file_v1_orders_proto_rawDesc = nil
	file_v1_orders_proto_goTypes = nil
	file_v1_orders_proto_depIdxs = nil
}
