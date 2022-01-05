// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: account/account.proto

package account

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

type PagingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo   uint32 `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize uint32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PagingReq) Reset() {
	*x = PagingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingReq) ProtoMessage() {}

func (x *PagingReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingReq.ProtoReflect.Descriptor instead.
func (*PagingReq) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{0}
}

func (x *PagingReq) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *PagingReq) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Gender   string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Role     uint32 `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *AccountInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AccountInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *AccountInfo) GetRole() uint32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type AccounttRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AccounttRes) Reset() {
	*x = AccounttRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccounttRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccounttRes) ProtoMessage() {}

func (x *AccounttRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccounttRes.ProtoReflect.Descriptor instead.
func (*AccounttRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccounttRes) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

type AccountListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       int32          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	AccountList []*AccountInfo `protobuf:"bytes,2,rep,name=accountList,proto3" json:"accountList,omitempty"`
}

func (x *AccountListRes) Reset() {
	*x = AccountListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListRes) ProtoMessage() {}

func (x *AccountListRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListRes.ProtoReflect.Descriptor instead.
func (*AccountListRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountListRes) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AccountListRes) GetAccountList() []*AccountInfo {
	if x != nil {
		return x.AccountList
	}
	return nil
}

type MobileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *MobileReq) Reset() {
	*x = MobileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileReq) ProtoMessage() {}

func (x *MobileReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileReq.ProtoReflect.Descriptor instead.
func (*MobileReq) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{4}
}

func (x *MobileReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type MobileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *MobileRes) Reset() {
	*x = MobileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileRes) ProtoMessage() {}

func (x *MobileRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileRes.ProtoReflect.Descriptor instead.
func (*MobileRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{5}
}

func (x *MobileRes) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

type IDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDReq) Reset() {
	*x = IDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReq) ProtoMessage() {}

func (x *IDReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReq.ProtoReflect.Descriptor instead.
func (*IDReq) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{6}
}

func (x *IDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *IDRes) Reset() {
	*x = IDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDRes) ProtoMessage() {}

func (x *IDRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDRes.ProtoReflect.Descriptor instead.
func (*IDRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{7}
}

func (x *IDRes) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

type AddAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AddAccountReq) Reset() {
	*x = AddAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAccountReq) ProtoMessage() {}

func (x *AddAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAccountReq.ProtoReflect.Descriptor instead.
func (*AddAccountReq) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{8}
}

func (x *AddAccountReq) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

type AddAccountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *AddAccountRes) Reset() {
	*x = AddAccountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAccountRes) ProtoMessage() {}

func (x *AddAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAccountRes.ProtoReflect.Descriptor instead.
func (*AddAccountRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{9}
}

func (x *AddAccountRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type UpdateAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *UpdateAccountReq) Reset() {
	*x = UpdateAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountReq) ProtoMessage() {}

func (x *UpdateAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountReq.ProtoReflect.Descriptor instead.
func (*UpdateAccountReq) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateAccountReq) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

type UpdateAccountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UpdateAccountRes) Reset() {
	*x = UpdateAccountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRes) ProtoMessage() {}

func (x *UpdateAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRes.ProtoReflect.Descriptor instead.
func (*UpdateAccountRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateAccountRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type CheckAccountByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CheckAccountByIDReq) Reset() {
	*x = CheckAccountByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccountByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccountByIDReq) ProtoMessage() {}

func (x *CheckAccountByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccountByIDReq.ProtoReflect.Descriptor instead.
func (*CheckAccountByIDReq) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{12}
}

func (x *CheckAccountByIDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CheckAccountByIDReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CheckAccountByIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CheckAccountByIDRes) Reset() {
	*x = CheckAccountByIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccountByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccountByIDRes) ProtoMessage() {}

func (x *CheckAccountByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccountByIDRes.ProtoReflect.Descriptor instead.
func (*CheckAccountByIDRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{13}
}

func (x *CheckAccountByIDRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DeleteAccountReq) Reset() {
	*x = DeleteAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountReq) ProtoMessage() {}

func (x *DeleteAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountReq.ProtoReflect.Descriptor instead.
func (*DeleteAccountReq) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteAccountReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteAccountReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteAccountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteAccountRes) Reset() {
	*x = DeleteAccountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountRes) ProtoMessage() {}

func (x *DeleteAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountRes.ProtoReflect.Descriptor instead.
func (*DeleteAccountRes) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{15}
}

func (x *DeleteAccountRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_account_account_proto protoreflect.FileDescriptor

var file_account_account_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x3f, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x3d, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x09,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0x3b, 0x0a, 0x09, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x17,
	0x0a, 0x05, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x05, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x3f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x22, 0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x41, 0x0a, 0x13, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a,
	0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x22, 0x3e, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xdb, 0x03, 0x0a, 0x0e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x4e, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12,
	0x45, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_account_proto_rawDescOnce sync.Once
	file_account_account_proto_rawDescData = file_account_account_proto_rawDesc
)

func file_account_account_proto_rawDescGZIP() []byte {
	file_account_account_proto_rawDescOnce.Do(func() {
		file_account_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_account_proto_rawDescData)
	})
	return file_account_account_proto_rawDescData
}

var file_account_account_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_account_account_proto_goTypes = []interface{}{
	(*PagingReq)(nil),           // 0: account.PagingReq
	(*AccountInfo)(nil),         // 1: account.AccountInfo
	(*AccounttRes)(nil),         // 2: account.AccounttRes
	(*AccountListRes)(nil),      // 3: account.AccountListRes
	(*MobileReq)(nil),           // 4: account.MobileReq
	(*MobileRes)(nil),           // 5: account.MobileRes
	(*IDReq)(nil),               // 6: account.IDReq
	(*IDRes)(nil),               // 7: account.IDRes
	(*AddAccountReq)(nil),       // 8: account.AddAccountReq
	(*AddAccountRes)(nil),       // 9: account.AddAccountRes
	(*UpdateAccountReq)(nil),    // 10: account.UpdateAccountReq
	(*UpdateAccountRes)(nil),    // 11: account.UpdateAccountRes
	(*CheckAccountByIDReq)(nil), // 12: account.CheckAccountByIDReq
	(*CheckAccountByIDRes)(nil), // 13: account.CheckAccountByIDRes
	(*DeleteAccountReq)(nil),    // 14: account.DeleteAccountReq
	(*DeleteAccountRes)(nil),    // 15: account.DeleteAccountRes
}
var file_account_account_proto_depIdxs = []int32{
	1,  // 0: account.AccounttRes.account:type_name -> account.AccountInfo
	1,  // 1: account.AccountListRes.accountList:type_name -> account.AccountInfo
	1,  // 2: account.MobileRes.account:type_name -> account.AccountInfo
	1,  // 3: account.IDRes.account:type_name -> account.AccountInfo
	1,  // 4: account.AddAccountReq.account:type_name -> account.AccountInfo
	1,  // 5: account.UpdateAccountReq.account:type_name -> account.AccountInfo
	0,  // 6: account.AccountService.GetAccountList:input_type -> account.PagingReq
	4,  // 7: account.AccountService.GetAccountByMobile:input_type -> account.MobileReq
	6,  // 8: account.AccountService.GetAccountByID:input_type -> account.IDReq
	8,  // 9: account.AccountService.AddAccount:input_type -> account.AddAccountReq
	10, // 10: account.AccountService.UpdateAccount:input_type -> account.UpdateAccountReq
	12, // 11: account.AccountService.CheckAccountByID:input_type -> account.CheckAccountByIDReq
	14, // 12: account.AccountService.DeleceAccount:input_type -> account.DeleteAccountReq
	3,  // 13: account.AccountService.GetAccountList:output_type -> account.AccountListRes
	5,  // 14: account.AccountService.GetAccountByMobile:output_type -> account.MobileRes
	7,  // 15: account.AccountService.GetAccountByID:output_type -> account.IDRes
	9,  // 16: account.AccountService.AddAccount:output_type -> account.AddAccountRes
	11, // 17: account.AccountService.UpdateAccount:output_type -> account.UpdateAccountRes
	13, // 18: account.AccountService.CheckAccountByID:output_type -> account.CheckAccountByIDRes
	15, // 19: account.AccountService.DeleceAccount:output_type -> account.DeleteAccountRes
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_account_account_proto_init() }
func file_account_account_proto_init() {
	if File_account_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingReq); i {
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
		file_account_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
		file_account_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccounttRes); i {
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
		file_account_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListRes); i {
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
		file_account_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileReq); i {
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
		file_account_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileRes); i {
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
		file_account_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReq); i {
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
		file_account_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDRes); i {
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
		file_account_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAccountReq); i {
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
		file_account_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAccountRes); i {
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
		file_account_account_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountReq); i {
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
		file_account_account_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRes); i {
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
		file_account_account_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccountByIDReq); i {
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
		file_account_account_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccountByIDRes); i {
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
		file_account_account_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountReq); i {
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
		file_account_account_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountRes); i {
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
			RawDescriptor: file_account_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_account_proto_goTypes,
		DependencyIndexes: file_account_account_proto_depIdxs,
		MessageInfos:      file_account_account_proto_msgTypes,
	}.Build()
	File_account_account_proto = out.File
	file_account_account_proto_rawDesc = nil
	file_account_account_proto_goTypes = nil
	file_account_account_proto_depIdxs = nil
}
