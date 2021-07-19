// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: event.proto

package _go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Greeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From                string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" parquet:"name=from, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	LastUpdated         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty" parquet:"name=last_updated, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	To                  *string                `protobuf:"bytes,3,opt,name=to,proto3,oneof" json:"to,omitempty" parquet:"name=to, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	PublisherCategories []int32                `protobuf:"zigzag32,4,rep,packed,name=publisher_categories,json=publisherCategories,proto3" json:"publisher_categories,omitempty" parquet:"name=publisher_categories, type=INT32 repetitiontype=REPEATED"`
	Address             *AddressPointer        `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Greeting) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *Greeting) GetTo() string {
	if x != nil && x.To != nil {
		return *x.To
	}
	return ""
}

func (x *Greeting) GetPublisherCategories() []int32 {
	if x != nil {
		return x.PublisherCategories
	}
	return nil
}

func (x *Greeting) GetAddress() *AddressPointer {
	if x != nil {
		return x.Address
	}
	return nil
}

type GreetingSimple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From                string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" parquet:"name=from,type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	LastUpdated         int64   `protobuf:"varint,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty" parquet:"name=last_updated, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	To                  *string `protobuf:"bytes,3,opt,name=to,proto3,oneof" json:"to,omitempty" parquet:"name=to,type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	PublisherCategories []int32 `protobuf:"zigzag32,4,rep,packed,name=publisher_categories,json=publisherCategories,proto3" json:"publisher_categories,omitempty" parquet:"name=publisher_categories, type=INT32, repetitiontype=REPEATED"`
}

func (x *GreetingSimple) Reset() {
	*x = GreetingSimple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingSimple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingSimple) ProtoMessage() {}

func (x *GreetingSimple) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingSimple.ProtoReflect.Descriptor instead.
func (*GreetingSimple) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingSimple) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GreetingSimple) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

func (x *GreetingSimple) GetTo() string {
	if x != nil && x.To != nil {
		return *x.To
	}
	return ""
}

func (x *GreetingSimple) GetPublisherCategories() []int32 {
	if x != nil {
		return x.PublisherCategories
	}
	return nil
}

type GreetingValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From                string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" parquet:"name=from,type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	LastUpdated         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty" parquet:"name=last_updated, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	To                  *string                `protobuf:"bytes,3,opt,name=to,proto3,oneof" json:"to,omitempty" parquet:"name=to,type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	PublisherCategories []int32                `protobuf:"zigzag32,4,rep,packed,name=publisher_categories,json=publisherCategories,proto3" json:"publisher_categories,omitempty" parquet:"name=publisher_categories, type=INT32, repetitiontype=REPEATED"`
	Address             *AddressValue          `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GreetingValue) Reset() {
	*x = GreetingValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingValue) ProtoMessage() {}

func (x *GreetingValue) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingValue.ProtoReflect.Descriptor instead.
func (*GreetingValue) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{2}
}

func (x *GreetingValue) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GreetingValue) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *GreetingValue) GetTo() string {
	if x != nil && x.To != nil {
		return *x.To
	}
	return ""
}

func (x *GreetingValue) GetPublisherCategories() []int32 {
	if x != nil {
		return x.PublisherCategories
	}
	return nil
}

func (x *GreetingValue) GetAddress() *AddressValue {
	if x != nil {
		return x.Address
	}
	return nil
}

type AddressPointer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street        *string        `protobuf:"bytes,1,opt,name=street,proto3,oneof" json:"street,omitempty"`
	PostCode      *uint32        `protobuf:"varint,2,opt,name=post_code,json=postCode,proto3,oneof" json:"post_code,omitempty"`
	Person        *PersonPointer `protobuf:"bytes,3,opt,name=person,proto3,oneof" json:"person,omitempty"`
	UseForBilling *bool          `protobuf:"varint,4,opt,name=use_for_billing,json=useForBilling,proto3,oneof" json:"use_for_billing,omitempty"`
}

func (x *AddressPointer) Reset() {
	*x = AddressPointer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressPointer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressPointer) ProtoMessage() {}

func (x *AddressPointer) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressPointer.ProtoReflect.Descriptor instead.
func (*AddressPointer) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{3}
}

func (x *AddressPointer) GetStreet() string {
	if x != nil && x.Street != nil {
		return *x.Street
	}
	return ""
}

func (x *AddressPointer) GetPostCode() uint32 {
	if x != nil && x.PostCode != nil {
		return *x.PostCode
	}
	return 0
}

func (x *AddressPointer) GetPerson() *PersonPointer {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *AddressPointer) GetUseForBilling() bool {
	if x != nil && x.UseForBilling != nil {
		return *x.UseForBilling
	}
	return false
}

type PersonPointer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName *string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3,oneof" json:"first_name,omitempty"`
	LastName  *string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3,oneof" json:"last_name,omitempty"`
}

func (x *PersonPointer) Reset() {
	*x = PersonPointer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonPointer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonPointer) ProtoMessage() {}

func (x *PersonPointer) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonPointer.ProtoReflect.Descriptor instead.
func (*PersonPointer) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{4}
}

func (x *PersonPointer) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *PersonPointer) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

type AddressValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street        string       `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	PostCode      uint32       `protobuf:"varint,2,opt,name=post_code,json=postCode,proto3" json:"post_code,omitempty"`
	Person        *PersonValue `protobuf:"bytes,3,opt,name=person,proto3" json:"person,omitempty"`
	UseForBilling bool         `protobuf:"varint,4,opt,name=use_for_billing,json=useForBilling,proto3" json:"use_for_billing,omitempty"`
}

func (x *AddressValue) Reset() {
	*x = AddressValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressValue) ProtoMessage() {}

func (x *AddressValue) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressValue.ProtoReflect.Descriptor instead.
func (*AddressValue) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5}
}

func (x *AddressValue) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *AddressValue) GetPostCode() uint32 {
	if x != nil {
		return x.PostCode
	}
	return 0
}

func (x *AddressValue) GetPerson() *PersonValue {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *AddressValue) GetUseForBilling() bool {
	if x != nil {
		return x.UseForBilling
	}
	return false
}

type PersonValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *PersonValue) Reset() {
	*x = PersonValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonValue) ProtoMessage() {}

func (x *PersonValue) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonValue.ProtoReflect.Descriptor instead.
func (*PersonValue) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{6}
}

func (x *PersonValue) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PersonValue) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x84, 0x04, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x58, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x44,
	0x9a, 0x84, 0x9e, 0x03, 0x3f, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x66, 0x72, 0x6f, 0x6d, 0x2c, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x42, 0x59,
	0x54, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x2c, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x3d, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x41, 0x52, 0x59, 0x22, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x4c, 0x9a,
	0x84, 0x9e, 0x03, 0x47, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x2c, 0x20,
	0x74, 0x79, 0x70, 0x65, 0x3d, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x2c, 0x20, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54,
	0x41, 0x4d, 0x50, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x22, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x57, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x42, 0x9a, 0x84, 0x9e, 0x03, 0x3d, 0x70, 0x61, 0x72, 0x71, 0x75,
	0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74, 0x6f, 0x2c, 0x20, 0x74, 0x79, 0x70,
	0x65, 0x3d, 0x42, 0x59, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x2c, 0x20, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3d, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x44, 0x49, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x41, 0x52, 0x59, 0x22, 0x48, 0x00, 0x52, 0x02, 0x74, 0x6f, 0x88, 0x01,
	0x01, 0x12, 0x7f, 0x0a, 0x14, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x11, 0x42,
	0x4c, 0x9a, 0x84, 0x9e, 0x03, 0x47, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2c, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x49,
	0x4e, 0x54, 0x33, 0x32, 0x20, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x74,
	0x79, 0x70, 0x65, 0x3d, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x45, 0x44, 0x22, 0x52, 0x13, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x74, 0x6f, 0x22, 0xe4, 0x03, 0x0a, 0x0e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x6b, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x57, 0x9a, 0x84, 0x9e,
	0x03, 0x52, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x66, 0x72, 0x6f, 0x6d, 0x2c, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x42, 0x59, 0x54, 0x45, 0x5f, 0x41,
	0x52, 0x52, 0x41, 0x59, 0x2c, 0x20, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x74,
	0x79, 0x70, 0x65, 0x3d, 0x55, 0x54, 0x46, 0x38, 0x2c, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x3d, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x41, 0x52, 0x59, 0x22, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x6f, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x4c, 0x9a, 0x84, 0x9e, 0x03, 0x47, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x2c, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x2c, 0x20, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x22, 0x52, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x6a, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x55, 0x9a, 0x84, 0x9e, 0x03, 0x50, 0x70, 0x61,
	0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74, 0x6f, 0x2c, 0x74,
	0x79, 0x70, 0x65, 0x3d, 0x42, 0x59, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x2c, 0x20,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x55, 0x54,
	0x46, 0x38, 0x2c, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3d, 0x50, 0x4c, 0x41,
	0x49, 0x4e, 0x5f, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x52, 0x59, 0x22, 0x48, 0x00,
	0x52, 0x02, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x80, 0x01, 0x0a, 0x14, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x11, 0x42, 0x4d, 0x9a, 0x84, 0x9e, 0x03, 0x48, 0x70, 0x61, 0x72,
	0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2c,
	0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x2c, 0x20, 0x72, 0x65, 0x70,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x52, 0x45, 0x50, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x22, 0x52, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x74,
	0x6f, 0x22, 0xaf, 0x04, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x6b, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x57, 0x9a, 0x84, 0x9e, 0x03, 0x52, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x66, 0x72, 0x6f, 0x6d, 0x2c, 0x74, 0x79, 0x70, 0x65, 0x3d,
	0x42, 0x59, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x2c, 0x20, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x55, 0x54, 0x46, 0x38, 0x2c, 0x20,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3d, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x44,
	0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x52, 0x59, 0x22, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x4c, 0x9a, 0x84, 0x9e, 0x03, 0x47, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65,
	0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x2c, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x49, 0x4e, 0x54, 0x36, 0x34,
	0x2c, 0x20, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x3d,
	0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53,
	0x22, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x6a,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x55, 0x9a, 0x84, 0x9e, 0x03,
	0x50, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74,
	0x6f, 0x2c, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x42, 0x59, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41,
	0x59, 0x2c, 0x20, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65,
	0x3d, 0x55, 0x54, 0x46, 0x38, 0x2c, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3d,
	0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x52, 0x59,
	0x22, 0x48, 0x00, 0x52, 0x02, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x80, 0x01, 0x0a, 0x14, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x11, 0x42, 0x4d, 0x9a, 0x84, 0x9e, 0x03, 0x48,
	0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x2c, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x2c, 0x20,
	0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x52,
	0x45, 0x50, 0x45, 0x41, 0x54, 0x45, 0x44, 0x22, 0x52, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x02, 0x52, 0x06, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x5f,
	0x66, 0x6f, 0x72, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x03, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x75, 0x73,
	0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x72, 0x0a,
	0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73,
	0x65, 0x46, 0x6f, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x49, 0x0a, 0x0b, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x76, 0x69, 0x64, 0x61, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_event_proto_goTypes = []interface{}{
	(*Greeting)(nil),              // 0: event.Greeting
	(*GreetingSimple)(nil),        // 1: event.GreetingSimple
	(*GreetingValue)(nil),         // 2: event.GreetingValue
	(*AddressPointer)(nil),        // 3: event.AddressPointer
	(*PersonPointer)(nil),         // 4: event.PersonPointer
	(*AddressValue)(nil),          // 5: event.AddressValue
	(*PersonValue)(nil),           // 6: event.PersonValue
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_event_proto_depIdxs = []int32{
	7, // 0: event.Greeting.last_updated:type_name -> google.protobuf.Timestamp
	3, // 1: event.Greeting.address:type_name -> event.AddressPointer
	7, // 2: event.GreetingValue.last_updated:type_name -> google.protobuf.Timestamp
	5, // 3: event.GreetingValue.address:type_name -> event.AddressValue
	4, // 4: event.AddressPointer.person:type_name -> event.PersonPointer
	6, // 5: event.AddressValue.person:type_name -> event.PersonValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	file_tagger_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeting); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingSimple); i {
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
		file_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingValue); i {
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
		file_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressPointer); i {
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
		file_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonPointer); i {
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
		file_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressValue); i {
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
		file_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonValue); i {
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
	file_event_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_event_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_event_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_event_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_event_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
