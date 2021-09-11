// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/service.proto

package mygrpc

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

type ResStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	ErrCode    string `protobuf:"bytes,2,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMessage string `protobuf:"bytes,3,opt,name=errMessage,proto3" json:"errMessage,omitempty"`
}

func (x *ResStatus) Reset() {
	*x = ResStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResStatus) ProtoMessage() {}

func (x *ResStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResStatus.ProtoReflect.Descriptor instead.
func (*ResStatus) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResStatus) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ResStatus) GetErrCode() string {
	if x != nil {
		return x.ErrCode
	}
	return ""
}

func (x *ResStatus) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

type ResEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *ResStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResEmpty) Reset() {
	*x = ResEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResEmpty) ProtoMessage() {}

func (x *ResEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResEmpty.ProtoReflect.Descriptor instead.
func (*ResEmpty) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResEmpty) GetStatus() *ResStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type RequestPaging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *int32 `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit  *int32 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *RequestPaging) Reset() {
	*x = RequestPaging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPaging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPaging) ProtoMessage() {}

func (x *RequestPaging) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPaging.ProtoReflect.Descriptor instead.
func (*RequestPaging) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *RequestPaging) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *RequestPaging) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type ReqEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging *RequestPaging `protobuf:"bytes,3,opt,name=paging,proto3,oneof" json:"paging,omitempty"`
	Offset *int32         `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"` // deprecated in favor of paging
	Limit  *int32         `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`   // deprecated in favor of paging
}

func (x *ReqEmpty) Reset() {
	*x = ReqEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqEmpty) ProtoMessage() {}

func (x *ReqEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqEmpty.ProtoReflect.Descriptor instead.
func (*ReqEmpty) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReqEmpty) GetPaging() *RequestPaging {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *ReqEmpty) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *ReqEmpty) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type ReqMLResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string  `protobuf:"bytes,1,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"` // deprecated in favor of Result
	Result      string  `protobuf:"bytes,5,opt,name=Result,proto3" json:"Result,omitempty"`
	CamName     string  `protobuf:"bytes,4,opt,name=CamName,proto3" json:"CamName,omitempty"`
	ImageURL    string  `protobuf:"bytes,2,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Score       float32 `protobuf:"fixed32,3,opt,name=Score,proto3" json:"Score,omitempty"`
}

func (x *ReqMLResult) Reset() {
	*x = ReqMLResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMLResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMLResult) ProtoMessage() {}

func (x *ReqMLResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMLResult.ProtoReflect.Descriptor instead.
func (*ReqMLResult) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReqMLResult) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *ReqMLResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ReqMLResult) GetCamName() string {
	if x != nil {
		return x.CamName
	}
	return ""
}

func (x *ReqMLResult) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *ReqMLResult) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type ReqPullMLResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging *RequestPaging `protobuf:"bytes,1,opt,name=paging,proto3,oneof" json:"paging,omitempty"`
	// int32 gateID = 2;
	LaneID int32 `protobuf:"varint,3,opt,name=laneID,proto3" json:"laneID,omitempty"`
}

func (x *ReqPullMLResult) Reset() {
	*x = ReqPullMLResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPullMLResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPullMLResult) ProtoMessage() {}

func (x *ReqPullMLResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPullMLResult.ProtoReflect.Descriptor instead.
func (*ReqPullMLResult) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *ReqPullMLResult) GetPaging() *RequestPaging {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *ReqPullMLResult) GetLaneID() int32 {
	if x != nil {
		return x.LaneID
	}
	return 0
}

type ReqListContainerOCRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging *RequestPaging `protobuf:"bytes,1,opt,name=paging,proto3,oneof" json:"paging,omitempty"`
	// int32 gateID = 2;
	LaneID int32 `protobuf:"varint,3,opt,name=laneID,proto3" json:"laneID,omitempty"`
}

func (x *ReqListContainerOCRs) Reset() {
	*x = ReqListContainerOCRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqListContainerOCRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqListContainerOCRs) ProtoMessage() {}

func (x *ReqListContainerOCRs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqListContainerOCRs.ProtoReflect.Descriptor instead.
func (*ReqListContainerOCRs) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *ReqListContainerOCRs) GetPaging() *RequestPaging {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *ReqListContainerOCRs) GetLaneID() int32 {
	if x != nil {
		return x.LaneID
	}
	return 0
}

type ResMLResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      *ResStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SuggestID   int32      `protobuf:"varint,5,opt,name=SuggestID,proto3" json:"SuggestID,omitempty"`
	ContainerID string     `protobuf:"bytes,2,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	ImageURL    string     `protobuf:"bytes,3,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Score       float32    `protobuf:"fixed32,4,opt,name=Score,proto3" json:"Score,omitempty"`
}

func (x *ResMLResult) Reset() {
	*x = ResMLResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMLResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMLResult) ProtoMessage() {}

func (x *ResMLResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMLResult.ProtoReflect.Descriptor instead.
func (*ResMLResult) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{7}
}

func (x *ResMLResult) GetStatus() *ResStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ResMLResult) GetSuggestID() int32 {
	if x != nil {
		return x.SuggestID
	}
	return 0
}

func (x *ResMLResult) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *ResMLResult) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *ResMLResult) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type ContainerOCR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Score       float32 `protobuf:"fixed32,5,opt,name=Score,proto3" json:"Score,omitempty"`
	ContainerID string  `protobuf:"bytes,2,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	ImageURL    string  `protobuf:"bytes,3,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	BIC         string  `protobuf:"bytes,6,opt,name=BIC,proto3" json:"BIC,omitempty"`
	Serial      string  `protobuf:"bytes,7,opt,name=Serial,proto3" json:"Serial,omitempty"`
	Checksum    string  `protobuf:"bytes,8,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
	CreatedAt   int32   `protobuf:"varint,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *ContainerOCR) Reset() {
	*x = ContainerOCR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerOCR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerOCR) ProtoMessage() {}

func (x *ContainerOCR) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerOCR.ProtoReflect.Descriptor instead.
func (*ContainerOCR) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{8}
}

func (x *ContainerOCR) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ContainerOCR) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ContainerOCR) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *ContainerOCR) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *ContainerOCR) GetBIC() string {
	if x != nil {
		return x.BIC
	}
	return ""
}

func (x *ContainerOCR) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *ContainerOCR) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

func (x *ContainerOCR) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ContainerTracking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OCRID       int32   `protobuf:"varint,9,opt,name=OCRID,proto3" json:"OCRID,omitempty"`
	Score       float32 `protobuf:"fixed32,5,opt,name=Score,proto3" json:"Score,omitempty"`
	ContainerID string  `protobuf:"bytes,2,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	ImageURL    string  `protobuf:"bytes,3,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	BIC         string  `protobuf:"bytes,6,opt,name=BIC,proto3" json:"BIC,omitempty"`
	Serial      string  `protobuf:"bytes,7,opt,name=Serial,proto3" json:"Serial,omitempty"`
	Checksum    string  `protobuf:"bytes,8,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
	CreatedAt   int32   `protobuf:"varint,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *ContainerTracking) Reset() {
	*x = ContainerTracking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerTracking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerTracking) ProtoMessage() {}

func (x *ContainerTracking) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerTracking.ProtoReflect.Descriptor instead.
func (*ContainerTracking) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{9}
}

func (x *ContainerTracking) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ContainerTracking) GetOCRID() int32 {
	if x != nil {
		return x.OCRID
	}
	return 0
}

func (x *ContainerTracking) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ContainerTracking) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *ContainerTracking) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *ContainerTracking) GetBIC() string {
	if x != nil {
		return x.BIC
	}
	return ""
}

func (x *ContainerTracking) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *ContainerTracking) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

func (x *ContainerTracking) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ResListContainerOCRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *ResStatus      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total  int32           `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Ocrs   []*ContainerOCR `protobuf:"bytes,2,rep,name=ocrs,proto3" json:"ocrs,omitempty"`
}

func (x *ResListContainerOCRs) Reset() {
	*x = ResListContainerOCRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResListContainerOCRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResListContainerOCRs) ProtoMessage() {}

func (x *ResListContainerOCRs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResListContainerOCRs.ProtoReflect.Descriptor instead.
func (*ResListContainerOCRs) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{10}
}

func (x *ResListContainerOCRs) GetStatus() *ResStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ResListContainerOCRs) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ResListContainerOCRs) GetOcrs() []*ContainerOCR {
	if x != nil {
		return x.Ocrs
	}
	return nil
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x65, 0x0a, 0x09, 0x52,
	0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x33, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x02, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x93, 0x01, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x4d, 0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x66, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x4c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x06, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x6e, 0x65, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x44, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x6b, 0x0a, 0x14, 0x52, 0x65,
	0x71, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4f, 0x43,
	0x52, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x4d,
	0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4f, 0x43, 0x52, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x49, 0x43, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x42, 0x49, 0x43, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf1, 0x01, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x43, 0x52, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4f, 0x43, 0x52, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x42,
	0x49, 0x43, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x42, 0x49, 0x43, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x7d, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x4f, 0x43, 0x52, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x6f, 0x63, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4f, 0x43, 0x52, 0x52, 0x04, 0x6f, 0x63, 0x72, 0x73, 0x32, 0xc3,
	0x01, 0x0a, 0x06, 0x4d, 0x79, 0x47, 0x52, 0x50, 0x43, 0x12, 0x30, 0x0a, 0x0b, 0x6e, 0x65, 0x77,
	0x4d, 0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x52, 0x65, 0x71, 0x4d, 0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0c, 0x70,
	0x75, 0x6c, 0x6c, 0x4d, 0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x4c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x4c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4f, 0x43, 0x52, 0x73, 0x12, 0x1a, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x4f, 0x43, 0x52, 0x73, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4f, 0x43, 0x52, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x6d, 0x79, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_service_proto_goTypes = []interface{}{
	(*ResStatus)(nil),            // 0: main.ResStatus
	(*ResEmpty)(nil),             // 1: main.ResEmpty
	(*RequestPaging)(nil),        // 2: main.RequestPaging
	(*ReqEmpty)(nil),             // 3: main.ReqEmpty
	(*ReqMLResult)(nil),          // 4: main.ReqMLResult
	(*ReqPullMLResult)(nil),      // 5: main.ReqPullMLResult
	(*ReqListContainerOCRs)(nil), // 6: main.ReqListContainerOCRs
	(*ResMLResult)(nil),          // 7: main.ResMLResult
	(*ContainerOCR)(nil),         // 8: main.ContainerOCR
	(*ContainerTracking)(nil),    // 9: main.ContainerTracking
	(*ResListContainerOCRs)(nil), // 10: main.ResListContainerOCRs
}
var file_proto_service_proto_depIdxs = []int32{
	0,  // 0: main.ResEmpty.status:type_name -> main.ResStatus
	2,  // 1: main.ReqEmpty.paging:type_name -> main.RequestPaging
	2,  // 2: main.ReqPullMLResult.paging:type_name -> main.RequestPaging
	2,  // 3: main.ReqListContainerOCRs.paging:type_name -> main.RequestPaging
	0,  // 4: main.ResMLResult.status:type_name -> main.ResStatus
	0,  // 5: main.ResListContainerOCRs.status:type_name -> main.ResStatus
	8,  // 6: main.ResListContainerOCRs.ocrs:type_name -> main.ContainerOCR
	4,  // 7: main.MyGRPC.newMLResult:input_type -> main.ReqMLResult
	5,  // 8: main.MyGRPC.pullMLResult:input_type -> main.ReqPullMLResult
	6,  // 9: main.MyGRPC.listContainerOCRs:input_type -> main.ReqListContainerOCRs
	1,  // 10: main.MyGRPC.newMLResult:output_type -> main.ResEmpty
	7,  // 11: main.MyGRPC.pullMLResult:output_type -> main.ResMLResult
	10, // 12: main.MyGRPC.listContainerOCRs:output_type -> main.ResListContainerOCRs
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResStatus); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResEmpty); i {
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
		file_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPaging); i {
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
		file_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqEmpty); i {
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
		file_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMLResult); i {
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
		file_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPullMLResult); i {
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
		file_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqListContainerOCRs); i {
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
		file_proto_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResMLResult); i {
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
		file_proto_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerOCR); i {
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
		file_proto_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerTracking); i {
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
		file_proto_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResListContainerOCRs); i {
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
	file_proto_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_proto_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_proto_service_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_proto_service_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
