// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: targets.proto

package grpc_gen

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

type TargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TargetRequest) Reset() {
	*x = TargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_targets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetRequest) ProtoMessage() {}

func (x *TargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_targets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetRequest.ProtoReflect.Descriptor instead.
func (*TargetRequest) Descriptor() ([]byte, []int) {
	return file_targets_proto_rawDescGZIP(), []int{0}
}

type TargetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ticker             string  `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	User               *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	ValuationRatio     string  `protobuf:"bytes,4,opt,name=ValuationRatio,proto3" json:"ValuationRatio,omitempty"`
	Value              float32 `protobuf:"fixed32,5,opt,name=Value,proto3" json:"Value,omitempty"`
	FinancialReport    string  `protobuf:"bytes,6,opt,name=FinancialReport,proto3" json:"FinancialReport,omitempty"`
	Achieved           bool    `protobuf:"varint,7,opt,name=Achieved,proto3" json:"Achieved,omitempty"`
	NotificationMethod string  `protobuf:"bytes,8,opt,name=NotificationMethod,proto3" json:"NotificationMethod,omitempty"`
}

func (x *TargetItem) Reset() {
	*x = TargetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_targets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetItem) ProtoMessage() {}

func (x *TargetItem) ProtoReflect() protoreflect.Message {
	mi := &file_targets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetItem.ProtoReflect.Descriptor instead.
func (*TargetItem) Descriptor() ([]byte, []int) {
	return file_targets_proto_rawDescGZIP(), []int{1}
}

func (x *TargetItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TargetItem) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *TargetItem) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TargetItem) GetValuationRatio() string {
	if x != nil {
		return x.ValuationRatio
	}
	return ""
}

func (x *TargetItem) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TargetItem) GetFinancialReport() string {
	if x != nil {
		return x.FinancialReport
	}
	return ""
}

func (x *TargetItem) GetAchieved() bool {
	if x != nil {
		return x.Achieved
	}
	return false
}

func (x *TargetItem) GetNotificationMethod() string {
	if x != nil {
		return x.NotificationMethod
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Telegram string `protobuf:"bytes,4,opt,name=telegram,proto3" json:"telegram,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_targets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_targets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_targets_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetTelegram() string {
	if x != nil {
		return x.Telegram
	}
	return ""
}

type TargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*TargetItem `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *TargetResponse) Reset() {
	*x = TargetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_targets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetResponse) ProtoMessage() {}

func (x *TargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_targets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetResponse.ProtoReflect.Descriptor instead.
func (*TargetResponse) Descriptor() ([]byte, []int) {
	return file_targets_proto_rawDescGZIP(), []int{3}
}

func (x *TargetResponse) GetTargets() []*TargetItem {
	if x != nil {
		return x.Targets
	}
	return nil
}

var File_targets_proto protoreflect.FileDescriptor

var file_targets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x0f, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x83, 0x02, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x56, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x46, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x41, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x5c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x22, 0x37, 0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x32, 0x3f, 0x0a,
	0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x0e, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f,
	0x5a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_targets_proto_rawDescOnce sync.Once
	file_targets_proto_rawDescData = file_targets_proto_rawDesc
)

func file_targets_proto_rawDescGZIP() []byte {
	file_targets_proto_rawDescOnce.Do(func() {
		file_targets_proto_rawDescData = protoimpl.X.CompressGZIP(file_targets_proto_rawDescData)
	})
	return file_targets_proto_rawDescData
}

var file_targets_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_targets_proto_goTypes = []any{
	(*TargetRequest)(nil),  // 0: TargetRequest
	(*TargetItem)(nil),     // 1: TargetItem
	(*User)(nil),           // 2: User
	(*TargetResponse)(nil), // 3: TargetResponse
}
var file_targets_proto_depIdxs = []int32{
	2, // 0: TargetItem.user:type_name -> User
	1, // 1: TargetResponse.targets:type_name -> TargetItem
	0, // 2: TargetsService.GetTargets:input_type -> TargetRequest
	3, // 3: TargetsService.GetTargets:output_type -> TargetResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_targets_proto_init() }
func file_targets_proto_init() {
	if File_targets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_targets_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TargetRequest); i {
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
		file_targets_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TargetItem); i {
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
		file_targets_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_targets_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TargetResponse); i {
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
			RawDescriptor: file_targets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_targets_proto_goTypes,
		DependencyIndexes: file_targets_proto_depIdxs,
		MessageInfos:      file_targets_proto_msgTypes,
	}.Build()
	File_targets_proto = out.File
	file_targets_proto_rawDesc = nil
	file_targets_proto_goTypes = nil
	file_targets_proto_depIdxs = nil
}
