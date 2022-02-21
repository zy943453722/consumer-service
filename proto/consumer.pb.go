// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.14.0
// source: proto/consumer.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetConsumerListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source     string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company    string `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Group      string `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	PageNumber int64  `protobuf:"varint,5,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	PageSize   int64  `protobuf:"varint,6,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	All        int64  `protobuf:"varint,7,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *GetConsumerListReq) Reset() {
	*x = GetConsumerListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerListReq) ProtoMessage() {}

func (x *GetConsumerListReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerListReq.ProtoReflect.Descriptor instead.
func (*GetConsumerListReq) Descriptor() ([]byte, []int) {
	return file_proto_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *GetConsumerListReq) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetConsumerListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetConsumerListReq) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *GetConsumerListReq) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GetConsumerListReq) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetConsumerListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetConsumerListReq) GetAll() int64 {
	if x != nil {
		return x.All
	}
	return 0
}

type ConsumerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId    string `protobuf:"bytes,1,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	Source        string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	SourceName    string `protobuf:"bytes,3,opt,name=sourceName,proto3" json:"sourceName,omitempty"`
	ConsumerName  string `protobuf:"bytes,4,opt,name=consumerName,proto3" json:"consumerName,omitempty"`
	CompanyName   string `protobuf:"bytes,5,opt,name=companyName,proto3" json:"companyName,omitempty"`
	UserGroup     string `protobuf:"bytes,6,opt,name=userGroup,proto3" json:"userGroup,omitempty"`
	UserGroupName string `protobuf:"bytes,7,opt,name=userGroupName,proto3" json:"userGroupName,omitempty"`
	UpdateTime    string `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	UpdateErp     string `protobuf:"bytes,9,opt,name=updateErp,proto3" json:"updateErp,omitempty"`
}

func (x *ConsumerInfo) Reset() {
	*x = ConsumerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerInfo) ProtoMessage() {}

func (x *ConsumerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerInfo.ProtoReflect.Descriptor instead.
func (*ConsumerInfo) Descriptor() ([]byte, []int) {
	return file_proto_consumer_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumerInfo) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *ConsumerInfo) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ConsumerInfo) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *ConsumerInfo) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *ConsumerInfo) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *ConsumerInfo) GetUserGroup() string {
	if x != nil {
		return x.UserGroup
	}
	return ""
}

func (x *ConsumerInfo) GetUserGroupName() string {
	if x != nil {
		return x.UserGroupName
	}
	return ""
}

func (x *ConsumerInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *ConsumerInfo) GetUpdateErp() string {
	if x != nil {
		return x.UpdateErp
	}
	return ""
}

type GetConsumerListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*ConsumerInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	PageNumber int64           `protobuf:"varint,2,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	PageSize   int64           `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	TotalCount int64           `protobuf:"varint,4,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *GetConsumerListResp) Reset() {
	*x = GetConsumerListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerListResp) ProtoMessage() {}

func (x *GetConsumerListResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerListResp.ProtoReflect.Descriptor instead.
func (*GetConsumerListResp) Descriptor() ([]byte, []int) {
	return file_proto_consumer_proto_rawDescGZIP(), []int{2}
}

func (x *GetConsumerListResp) GetData() []*ConsumerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetConsumerListResp) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetConsumerListResp) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetConsumerListResp) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_proto_consumer_proto protoreflect.FileDescriptor

var file_proto_consumer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x22, 0xae,
	0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x75,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x72, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x72, 0x70, 0x22,
	0x9a, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x5d, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_consumer_proto_rawDescOnce sync.Once
	file_proto_consumer_proto_rawDescData = file_proto_consumer_proto_rawDesc
)

func file_proto_consumer_proto_rawDescGZIP() []byte {
	file_proto_consumer_proto_rawDescOnce.Do(func() {
		file_proto_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_consumer_proto_rawDescData)
	})
	return file_proto_consumer_proto_rawDescData
}

var file_proto_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_consumer_proto_goTypes = []interface{}{
	(*GetConsumerListReq)(nil),  // 0: proto.GetConsumerListReq
	(*ConsumerInfo)(nil),        // 1: proto.ConsumerInfo
	(*GetConsumerListResp)(nil), // 2: proto.GetConsumerListResp
}
var file_proto_consumer_proto_depIdxs = []int32{
	1, // 0: proto.GetConsumerListResp.data:type_name -> proto.ConsumerInfo
	0, // 1: proto.ConsumerService.GetConsumerList:input_type -> proto.GetConsumerListReq
	2, // 2: proto.ConsumerService.GetConsumerList:output_type -> proto.GetConsumerListResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_consumer_proto_init() }
func file_proto_consumer_proto_init() {
	if File_proto_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerListReq); i {
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
		file_proto_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerInfo); i {
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
		file_proto_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerListResp); i {
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
			RawDescriptor: file_proto_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_consumer_proto_goTypes,
		DependencyIndexes: file_proto_consumer_proto_depIdxs,
		MessageInfos:      file_proto_consumer_proto_msgTypes,
	}.Build()
	File_proto_consumer_proto = out.File
	file_proto_consumer_proto_rawDesc = nil
	file_proto_consumer_proto_goTypes = nil
	file_proto_consumer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	GetConsumerList(ctx context.Context, in *GetConsumerListReq, opts ...grpc.CallOption) (*GetConsumerListResp, error)
}

type consumerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerServiceClient(cc grpc.ClientConnInterface) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) GetConsumerList(ctx context.Context, in *GetConsumerListReq, opts ...grpc.CallOption) (*GetConsumerListResp, error) {
	out := new(GetConsumerListResp)
	err := c.cc.Invoke(ctx, "/proto.ConsumerService/GetConsumerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
type ConsumerServiceServer interface {
	GetConsumerList(context.Context, *GetConsumerListReq) (*GetConsumerListResp, error)
}

// UnimplementedConsumerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsumerServiceServer struct {
}

func (*UnimplementedConsumerServiceServer) GetConsumerList(context.Context, *GetConsumerListReq) (*GetConsumerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerList not implemented")
}

func RegisterConsumerServiceServer(s *grpc.Server, srv ConsumerServiceServer) {
	s.RegisterService(&_ConsumerService_serviceDesc, srv)
}

func _ConsumerService_GetConsumerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsumerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetConsumerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsumerService/GetConsumerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetConsumerList(ctx, req.(*GetConsumerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsumerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConsumerList",
			Handler:    _ConsumerService_GetConsumerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consumer.proto",
}