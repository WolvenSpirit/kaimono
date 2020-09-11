// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.10.0
// source: api.proto

package api

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    *string `protobuf:"bytes,1,req,name=Email" json:"Email,omitempty"`
	Password *string `protobuf:"bytes,2,req,name=Password" json:"Password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    *int32  `protobuf:"varint,1,req,name=Code" json:"Code,omitempty"`
	Message *string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Token   *string `protobuf:"bytes,3,opt,name=Token" json:"Token,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *AuthResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *AuthResponse) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    *string `protobuf:"bytes,1,req,name=Email" json:"Email,omitempty"`
	Password *string `protobuf:"bytes,2,req,name=Password" json:"Password,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *RegistrationRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *RegistrationRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table *string `protobuf:"bytes,1,req,name=Table" json:"Table,omitempty"`
	On    *string `protobuf:"bytes,2,req,name=On" json:"On,omitempty"`
}

func (x *Join) Reset() {
	*x = Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join) ProtoMessage() {}

func (x *Join) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join.ProtoReflect.Descriptor instead.
func (*Join) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Join) GetTable() string {
	if x != nil && x.Table != nil {
		return *x.Table
	}
	return ""
}

func (x *Join) GetOn() string {
	if x != nil && x.On != nil {
		return *x.On
	}
	return ""
}

type Insert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cols   []string `protobuf:"bytes,1,rep,name=Cols" json:"Cols,omitempty"`
	Values [][]byte `protobuf:"bytes,2,rep,name=Values" json:"Values,omitempty"`
}

func (x *Insert) Reset() {
	*x = Insert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insert) ProtoMessage() {}

func (x *Insert) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insert.ProtoReflect.Descriptor instead.
func (*Insert) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Insert) GetCols() []string {
	if x != nil {
		return x.Cols
	}
	return nil
}

func (x *Insert) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type RequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation        *string  `protobuf:"bytes,1,req,name=Operation" json:"Operation,omitempty"`
	Table            []string `protobuf:"bytes,2,rep,name=Table" json:"Table,omitempty"`
	Want             []string `protobuf:"bytes,3,rep,name=Want" json:"Want,omitempty"`
	Where            []string `protobuf:"bytes,4,rep,name=Where" json:"Where,omitempty"`
	Join             []*Join  `protobuf:"bytes,5,rep,name=Join" json:"Join,omitempty"`
	Define           []string `protobuf:"bytes,6,rep,name=Define" json:"Define,omitempty"`
	Insert           *Insert  `protobuf:"bytes,7,opt,name=Insert" json:"Insert,omitempty"`
	UpdateAssignment *string  `protobuf:"bytes,8,opt,name=UpdateAssignment" json:"UpdateAssignment,omitempty"`
	Token            *string  `protobuf:"bytes,9,opt,name=token" json:"token,omitempty"`
}

func (x *RequestBody) Reset() {
	*x = RequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBody) ProtoMessage() {}

func (x *RequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBody.ProtoReflect.Descriptor instead.
func (*RequestBody) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *RequestBody) GetOperation() string {
	if x != nil && x.Operation != nil {
		return *x.Operation
	}
	return ""
}

func (x *RequestBody) GetTable() []string {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *RequestBody) GetWant() []string {
	if x != nil {
		return x.Want
	}
	return nil
}

func (x *RequestBody) GetWhere() []string {
	if x != nil {
		return x.Where
	}
	return nil
}

func (x *RequestBody) GetJoin() []*Join {
	if x != nil {
		return x.Join
	}
	return nil
}

func (x *RequestBody) GetDefine() []string {
	if x != nil {
		return x.Define
	}
	return nil
}

func (x *RequestBody) GetInsert() *Insert {
	if x != nil {
		return x.Insert
	}
	return nil
}

func (x *RequestBody) GetUpdateAssignment() string {
	if x != nil && x.UpdateAssignment != nil {
		return *x.UpdateAssignment
	}
	return ""
}

func (x *RequestBody) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type ResponseBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=StatusCode" json:"StatusCode,omitempty"`
	Payload    []byte  `protobuf:"bytes,2,opt,name=Payload" json:"Payload,omitempty"`
	Message    *string `protobuf:"bytes,3,opt,name=Message" json:"Message,omitempty"`
}

func (x *ResponseBody) Reset() {
	*x = ResponseBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBody) ProtoMessage() {}

func (x *ResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBody.ProtoReflect.Descriptor instead.
func (*ResponseBody) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseBody) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *ResponseBody) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ResponseBody) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x52, 0x0a,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x47, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2c, 0x0a, 0x04, 0x6a, 0x6f,
	0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6e, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x6e, 0x22, 0x34, 0x0a, 0x06, 0x69, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x43, 0x6f, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x81,
	0x02, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x57, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x04,
	0x4a, 0x6f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x6a, 0x6f, 0x69,
	0x6e, 0x52, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12,
	0x1f, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x62, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa7, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0c,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x0d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x6f, 0x64, 0x79, 0x1a, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x0d, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64,
	0x79, 0x1a, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),        // 0: LoginRequest
	(*AuthResponse)(nil),        // 1: AuthResponse
	(*RegistrationRequest)(nil), // 2: RegistrationRequest
	(*Join)(nil),                // 3: join
	(*Insert)(nil),              // 4: insert
	(*RequestBody)(nil),         // 5: RequestBody
	(*ResponseBody)(nil),        // 6: ResponseBody
}
var file_api_proto_depIdxs = []int32{
	3, // 0: RequestBody.Join:type_name -> join
	4, // 1: RequestBody.Insert:type_name -> insert
	0, // 2: ApiService.Login:input_type -> LoginRequest
	2, // 3: ApiService.Register:input_type -> RegistrationRequest
	5, // 4: ApiService.Select:input_type -> RequestBody
	5, // 5: ApiService.Insert:input_type -> RequestBody
	5, // 6: ApiService.Update:input_type -> RequestBody
	5, // 7: ApiService.Delete:input_type -> RequestBody
	5, // 8: ApiService.Create:input_type -> RequestBody
	1, // 9: ApiService.Login:output_type -> AuthResponse
	1, // 10: ApiService.Register:output_type -> AuthResponse
	6, // 11: ApiService.Select:output_type -> ResponseBody
	6, // 12: ApiService.Insert:output_type -> ResponseBody
	6, // 13: ApiService.Update:output_type -> ResponseBody
	6, // 14: ApiService.Delete:output_type -> ResponseBody
	6, // 15: ApiService.Create:output_type -> ResponseBody
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insert); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBody); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBody); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}