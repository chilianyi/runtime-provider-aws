// Code generated by protoc-gen-go. DO NOT EDIT.
// source: category.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Category struct {
	CategoryId           *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Locale               *wrappers.StringValue `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	OwnerPath            *wrappers.StringValue `protobuf:"bytes,4,opt,name=owner_path,json=ownerPath,proto3" json:"owner_path,omitempty"`
	CreateTime           *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Icon                 *wrappers.StringValue `protobuf:"bytes,8,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func (m *Category) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Category) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *Category) GetOwnerPath() *wrappers.StringValue {
	if m != nil {
		return m.OwnerPath
	}
	return nil
}

func (m *Category) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Category) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Category) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Category) GetIcon() *wrappers.StringValue {
	if m != nil {
		return m.Icon
	}
	return nil
}

type DescribeCategoriesRequest struct {
	SearchWord *wrappers.StringValue `protobuf:"bytes,1,opt,name=search_word,json=searchWord,proto3" json:"search_word,omitempty"`
	SortKey    *wrappers.StringValue `protobuf:"bytes,2,opt,name=sort_key,json=sortKey,proto3" json:"sort_key,omitempty"`
	Reverse    *wrappers.BoolValue   `protobuf:"bytes,3,opt,name=reverse,proto3" json:"reverse,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// default is 0
	Offset               uint32   `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	DisplayColumns       []string `protobuf:"bytes,6,rep,name=display_columns,json=displayColumns,proto3" json:"display_columns,omitempty"`
	CategoryId           []string `protobuf:"bytes,7,rep,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name                 []string `protobuf:"bytes,8,rep,name=name,proto3" json:"name,omitempty"`
	Owner                []string `protobuf:"bytes,9,rep,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeCategoriesRequest) Reset()         { *m = DescribeCategoriesRequest{} }
func (m *DescribeCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeCategoriesRequest) ProtoMessage()    {}
func (*DescribeCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{1}
}

func (m *DescribeCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeCategoriesRequest.Unmarshal(m, b)
}
func (m *DescribeCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeCategoriesRequest.Marshal(b, m, deterministic)
}
func (m *DescribeCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeCategoriesRequest.Merge(m, src)
}
func (m *DescribeCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeCategoriesRequest.Size(m)
}
func (m *DescribeCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeCategoriesRequest proto.InternalMessageInfo

func (m *DescribeCategoriesRequest) GetSearchWord() *wrappers.StringValue {
	if m != nil {
		return m.SearchWord
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetSortKey() *wrappers.StringValue {
	if m != nil {
		return m.SortKey
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetReverse() *wrappers.BoolValue {
	if m != nil {
		return m.Reverse
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeCategoriesRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeCategoriesRequest) GetDisplayColumns() []string {
	if m != nil {
		return m.DisplayColumns
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetCategoryId() []string {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DescribeCategoriesRequest) GetOwner() []string {
	if m != nil {
		return m.Owner
	}
	return nil
}

type DescribeCategoriesResponse struct {
	TotalCount           uint32      `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	CategorySet          []*Category `protobuf:"bytes,2,rep,name=category_set,json=categorySet,proto3" json:"category_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DescribeCategoriesResponse) Reset()         { *m = DescribeCategoriesResponse{} }
func (m *DescribeCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeCategoriesResponse) ProtoMessage()    {}
func (*DescribeCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{2}
}

func (m *DescribeCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeCategoriesResponse.Unmarshal(m, b)
}
func (m *DescribeCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *DescribeCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeCategoriesResponse.Merge(m, src)
}
func (m *DescribeCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeCategoriesResponse.Size(m)
}
func (m *DescribeCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeCategoriesResponse proto.InternalMessageInfo

func (m *DescribeCategoriesResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeCategoriesResponse) GetCategorySet() []*Category {
	if m != nil {
		return m.CategorySet
	}
	return nil
}

type CreateCategoryRequest struct {
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// the i18n of this category, json format, sample: {"zh_cn": "数据库", "en": "database"}
	Locale               *wrappers.StringValue `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Icon                 *wrappers.BytesValue  `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateCategoryRequest) Reset()         { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()    {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{3}
}

func (m *CreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryRequest.Unmarshal(m, b)
}
func (m *CreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *CreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryRequest.Merge(m, src)
}
func (m *CreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryRequest.Size(m)
}
func (m *CreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryRequest proto.InternalMessageInfo

func (m *CreateCategoryRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateCategoryRequest) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *CreateCategoryRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CreateCategoryRequest) GetIcon() *wrappers.BytesValue {
	if m != nil {
		return m.Icon
	}
	return nil
}

type CreateCategoryResponse struct {
	CategoryId           *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateCategoryResponse) Reset()         { *m = CreateCategoryResponse{} }
func (m *CreateCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryResponse) ProtoMessage()    {}
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{4}
}

func (m *CreateCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryResponse.Unmarshal(m, b)
}
func (m *CreateCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryResponse.Marshal(b, m, deterministic)
}
func (m *CreateCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryResponse.Merge(m, src)
}
func (m *CreateCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryResponse.Size(m)
}
func (m *CreateCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryResponse proto.InternalMessageInfo

func (m *CreateCategoryResponse) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

type ModifyCategoryRequest struct {
	CategoryId *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name       *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// the i18n of this category, json format, sample: {"zh_cn": "数据库", "en": "database"}
	Locale               *wrappers.StringValue `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Icon                 *wrappers.BytesValue  `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ModifyCategoryRequest) Reset()         { *m = ModifyCategoryRequest{} }
func (m *ModifyCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyCategoryRequest) ProtoMessage()    {}
func (*ModifyCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{5}
}

func (m *ModifyCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyCategoryRequest.Unmarshal(m, b)
}
func (m *ModifyCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyCategoryRequest.Marshal(b, m, deterministic)
}
func (m *ModifyCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyCategoryRequest.Merge(m, src)
}
func (m *ModifyCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyCategoryRequest.Size(m)
}
func (m *ModifyCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyCategoryRequest proto.InternalMessageInfo

func (m *ModifyCategoryRequest) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func (m *ModifyCategoryRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ModifyCategoryRequest) GetLocale() *wrappers.StringValue {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *ModifyCategoryRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ModifyCategoryRequest) GetIcon() *wrappers.BytesValue {
	if m != nil {
		return m.Icon
	}
	return nil
}

type ModifyCategoryResponse struct {
	CategoryId           *wrappers.StringValue `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ModifyCategoryResponse) Reset()         { *m = ModifyCategoryResponse{} }
func (m *ModifyCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyCategoryResponse) ProtoMessage()    {}
func (*ModifyCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{6}
}

func (m *ModifyCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyCategoryResponse.Unmarshal(m, b)
}
func (m *ModifyCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyCategoryResponse.Marshal(b, m, deterministic)
}
func (m *ModifyCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyCategoryResponse.Merge(m, src)
}
func (m *ModifyCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyCategoryResponse.Size(m)
}
func (m *ModifyCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyCategoryResponse proto.InternalMessageInfo

func (m *ModifyCategoryResponse) GetCategoryId() *wrappers.StringValue {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

type DeleteCategoriesRequest struct {
	CategoryId           []string `protobuf:"bytes,1,rep,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCategoriesRequest) Reset()         { *m = DeleteCategoriesRequest{} }
func (m *DeleteCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCategoriesRequest) ProtoMessage()    {}
func (*DeleteCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{7}
}

func (m *DeleteCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCategoriesRequest.Unmarshal(m, b)
}
func (m *DeleteCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCategoriesRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCategoriesRequest.Merge(m, src)
}
func (m *DeleteCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCategoriesRequest.Size(m)
}
func (m *DeleteCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCategoriesRequest proto.InternalMessageInfo

func (m *DeleteCategoriesRequest) GetCategoryId() []string {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

type DeleteCategoriesResponse struct {
	CategoryId           []string `protobuf:"bytes,1,rep,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCategoriesResponse) Reset()         { *m = DeleteCategoriesResponse{} }
func (m *DeleteCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCategoriesResponse) ProtoMessage()    {}
func (*DeleteCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6ef5ed29d8d1a1, []int{8}
}

func (m *DeleteCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCategoriesResponse.Unmarshal(m, b)
}
func (m *DeleteCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCategoriesResponse.Merge(m, src)
}
func (m *DeleteCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCategoriesResponse.Size(m)
}
func (m *DeleteCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCategoriesResponse proto.InternalMessageInfo

func (m *DeleteCategoriesResponse) GetCategoryId() []string {
	if m != nil {
		return m.CategoryId
	}
	return nil
}

func init() {
	proto.RegisterType((*Category)(nil), "openpitrix.Category")
	proto.RegisterType((*DescribeCategoriesRequest)(nil), "openpitrix.DescribeCategoriesRequest")
	proto.RegisterType((*DescribeCategoriesResponse)(nil), "openpitrix.DescribeCategoriesResponse")
	proto.RegisterType((*CreateCategoryRequest)(nil), "openpitrix.CreateCategoryRequest")
	proto.RegisterType((*CreateCategoryResponse)(nil), "openpitrix.CreateCategoryResponse")
	proto.RegisterType((*ModifyCategoryRequest)(nil), "openpitrix.ModifyCategoryRequest")
	proto.RegisterType((*ModifyCategoryResponse)(nil), "openpitrix.ModifyCategoryResponse")
	proto.RegisterType((*DeleteCategoriesRequest)(nil), "openpitrix.DeleteCategoriesRequest")
	proto.RegisterType((*DeleteCategoriesResponse)(nil), "openpitrix.DeleteCategoriesResponse")
}

func init() { proto.RegisterFile("category.proto", fileDescriptor_1c6ef5ed29d8d1a1) }

var fileDescriptor_1c6ef5ed29d8d1a1 = []byte{
	// 798 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x55, 0x3e, 0x9a, 0xa4, 0x93, 0xd7, 0xb4, 0x9d, 0x7e, 0x3c, 0xbf, 0xbc, 0xea, 0x25, 0xf5,
	0xfb, 0xaa, 0x2a, 0x1a, 0x43, 0xa8, 0x54, 0xa9, 0x15, 0x48, 0x6d, 0xba, 0x41, 0xa8, 0x12, 0x72,
	0x11, 0x95, 0xd8, 0x44, 0x13, 0xfb, 0xc6, 0x19, 0xe1, 0x78, 0xcc, 0x78, 0xd2, 0x90, 0x2d, 0x0b,
	0x96, 0x2c, 0xca, 0x0f, 0xe0, 0x27, 0xf0, 0x63, 0x90, 0x58, 0xb0, 0x66, 0xc7, 0x86, 0x9f, 0x80,
	0x3c, 0x1e, 0xa7, 0x89, 0x13, 0xda, 0x20, 0x10, 0x88, 0x55, 0xe4, 0xb9, 0xe7, 0xd8, 0x67, 0xee,
	0x39, 0xf7, 0x06, 0x95, 0x2c, 0x22, 0xc0, 0x61, 0x7c, 0x50, 0xf3, 0x39, 0x13, 0x0c, 0x23, 0xe6,
	0x83, 0xe7, 0x53, 0xc1, 0xe9, 0xb3, 0xf2, 0x5f, 0x0e, 0x63, 0x8e, 0x0b, 0x86, 0xac, 0xb4, 0x7a,
	0x6d, 0xa3, 0xcf, 0x89, 0xef, 0x03, 0x0f, 0x22, 0x6c, 0xb9, 0x92, 0xac, 0x0b, 0xda, 0x85, 0x40,
	0x90, 0xae, 0xaf, 0x00, 0x1b, 0x0a, 0x40, 0x7c, 0x6a, 0x10, 0xcf, 0x63, 0x82, 0x08, 0xca, 0xbc,
	0x98, 0x7e, 0x43, 0xfe, 0x58, 0x3b, 0x0e, 0x78, 0x3b, 0x41, 0x9f, 0x38, 0x0e, 0x70, 0x83, 0xf9,
	0x12, 0x31, 0x89, 0xd6, 0xdf, 0x67, 0x50, 0xa1, 0xa1, 0xb4, 0xe2, 0x3b, 0xa8, 0x18, 0xeb, 0x6e,
	0x52, 0x5b, 0x4b, 0x55, 0x53, 0x5b, 0xc5, 0xfa, 0x46, 0x2d, 0xfa, 0x5c, 0x2d, 0xd6, 0x53, 0x3b,
	0x15, 0x9c, 0x7a, 0xce, 0x23, 0xe2, 0xf6, 0xc0, 0x44, 0x31, 0xe1, 0x9e, 0x8d, 0x6f, 0xa2, 0xac,
	0x47, 0xba, 0xa0, 0xa5, 0x67, 0xe0, 0x49, 0x24, 0xde, 0x45, 0x39, 0x97, 0x59, 0xc4, 0x05, 0x2d,
	0x33, 0x03, 0x47, 0x61, 0xf1, 0x01, 0x42, 0xac, 0xef, 0x01, 0x6f, 0xfa, 0x44, 0x74, 0xb4, 0xec,
	0x0c, 0xcc, 0x79, 0x89, 0x7f, 0x40, 0x44, 0x07, 0x1f, 0xa0, 0xa2, 0xc5, 0x81, 0x08, 0x68, 0x86,
	0x6d, 0xd5, 0xe6, 0x24, 0xbb, 0x3c, 0xc1, 0x7e, 0x18, 0xf7, 0xdc, 0x44, 0x11, 0x3c, 0x3c, 0x08,
	0xc9, 0x3d, 0xdf, 0x1e, 0x92, 0x73, 0xd7, 0x93, 0x23, 0xb8, 0x24, 0xdf, 0x45, 0x45, 0x1b, 0x02,
	0x8b, 0x53, 0x69, 0x86, 0x96, 0x9f, 0x41, 0xf7, 0x28, 0x21, 0x6c, 0x2f, 0xb5, 0x98, 0xa7, 0x15,
	0x66, 0x69, 0x6f, 0x88, 0xd4, 0x3f, 0xa6, 0xd1, 0x1f, 0xc7, 0xf2, 0x0d, 0x2d, 0x50, 0x26, 0x53,
	0x08, 0x4c, 0x78, 0xda, 0x83, 0x40, 0x84, 0x6e, 0x07, 0x40, 0xb8, 0xd5, 0x69, 0xf6, 0x19, 0x9f,
	0xd1, 0xed, 0x88, 0x70, 0xc6, 0xb8, 0x8d, 0xf7, 0x50, 0x21, 0x60, 0x5c, 0x34, 0x9f, 0xc0, 0x60,
	0x26, 0xc7, 0xf3, 0x21, 0xfa, 0x3e, 0x0c, 0xf0, 0x2e, 0xca, 0x73, 0x38, 0x07, 0x1e, 0xc4, 0xae,
	0x4f, 0x36, 0xf0, 0x88, 0x31, 0x57, 0xb1, 0x14, 0x14, 0xaf, 0xa2, 0x39, 0x97, 0x76, 0xa9, 0x90,
	0x7e, 0x2f, 0x98, 0xd1, 0x03, 0x5e, 0x47, 0x39, 0xd6, 0x6e, 0x07, 0x20, 0xa4, 0x91, 0x0b, 0xa6,
	0x7a, 0xc2, 0xff, 0xa3, 0x45, 0x9b, 0x06, 0xbe, 0x4b, 0x06, 0x4d, 0x8b, 0xb9, 0xbd, 0xae, 0x17,
	0x68, 0xb9, 0x6a, 0x66, 0x6b, 0xde, 0x2c, 0xa9, 0xe3, 0x46, 0x74, 0x8a, 0x2b, 0xe3, 0x91, 0xcf,
	0x4b, 0xd0, 0x68, 0xa8, 0xb1, 0x0a, 0x75, 0x41, 0x56, 0xa2, 0xd8, 0xae, 0xa2, 0x39, 0x19, 0x28,
	0x6d, 0x5e, 0x1e, 0x46, 0x0f, 0xfa, 0x39, 0x2a, 0x4f, 0x6b, 0x76, 0xe0, 0x33, 0x2f, 0x80, 0xf0,
	0x43, 0x82, 0x09, 0xe2, 0x36, 0x2d, 0xd6, 0xf3, 0x84, 0xec, 0xf6, 0x82, 0x89, 0xe4, 0x51, 0x23,
	0x3c, 0xc1, 0x7b, 0xe8, 0xb7, 0xa1, 0x92, 0xf0, 0x42, 0xe9, 0x6a, 0x66, 0xab, 0x58, 0x5f, 0xad,
	0x5d, 0x6e, 0x8e, 0x5a, 0x3c, 0xa8, 0xe6, 0x50, 0xf3, 0x29, 0x08, 0xfd, 0x53, 0x0a, 0xad, 0x35,
	0x64, 0x46, 0x87, 0x75, 0xe5, 0xf0, 0x8f, 0x1a, 0xc8, 0x44, 0xb2, 0xb3, 0x5f, 0x9b, 0x6c, 0x43,
	0x25, 0x3b, 0x1a, 0xc6, 0x3f, 0x27, 0xe3, 0x30, 0x10, 0x10, 0x8c, 0x06, 0xfb, 0x0c, 0xad, 0x27,
	0x6f, 0xac, 0xda, 0xfc, 0x6d, 0x2b, 0x4c, 0x7f, 0x93, 0x46, 0x6b, 0x27, 0xcc, 0xa6, 0xed, 0x41,
	0xb2, 0x97, 0xbf, 0xc8, 0x6e, 0xfc, 0x19, 0x56, 0x24, 0x1b, 0xf6, 0x7d, 0xac, 0xd8, 0x47, 0xbf,
	0x1f, 0x83, 0x0b, 0x62, 0xca, 0xe6, 0xaa, 0x24, 0xdf, 0x9c, 0x18, 0x5a, 0xfd, 0x00, 0x69, 0x93,
	0xdc, 0xcb, 0x41, 0xbc, 0x92, 0x5c, 0x7f, 0x97, 0x45, 0x8b, 0xf1, 0x65, 0x4e, 0x88, 0x47, 0x1c,
	0xe0, 0xf8, 0x75, 0x0a, 0xe1, 0xc9, 0xe1, 0xc6, 0xff, 0x8e, 0x4e, 0xe7, 0x17, 0x37, 0x6d, 0xf9,
	0xbf, 0xeb, 0x60, 0x91, 0x34, 0x7d, 0xff, 0xe2, 0x70, 0x13, 0x57, 0x6c, 0x05, 0xa8, 0x5a, 0x43,
	0x44, 0xb5, 0x4f, 0x45, 0xa7, 0xda, 0xa6, 0xae, 0x00, 0xfe, 0xfc, 0xed, 0x87, 0x57, 0xe9, 0x25,
	0x5c, 0x32, 0xce, 0x6f, 0x19, 0x97, 0x08, 0xfc, 0x22, 0x85, 0x4a, 0xe3, 0x33, 0x81, 0x37, 0xc7,
	0x76, 0xc7, 0xb4, 0x0d, 0x51, 0xd6, 0xaf, 0x82, 0x28, 0x55, 0x3b, 0x17, 0x87, 0xcb, 0x78, 0x31,
	0xfa, 0x17, 0x8c, 0x35, 0x0d, 0xa4, 0x8a, 0x15, 0x3d, 0xa1, 0x62, 0x3f, 0xb5, 0x2d, 0x85, 0x8c,
	0x27, 0x62, 0x5c, 0xc8, 0xd4, 0xf1, 0x1a, 0x17, 0x32, 0x3d, 0x50, 0x4a, 0x48, 0x57, 0x16, 0x13,
	0x42, 0xea, 0x53, 0x84, 0xbc, 0x4c, 0xa1, 0xa5, 0x64, 0x0a, 0xf0, 0xdf, 0xe3, 0x56, 0x4c, 0xcd,
	0x57, 0xf9, 0x9f, 0xab, 0x41, 0x4a, 0x8e, 0x71, 0x71, 0xb8, 0x82, 0x97, 0x6d, 0x59, 0x1e, 0xf1,
	0x2a, 0x12, 0xb4, 0x3d, 0x29, 0xe8, 0x28, 0xfb, 0x38, 0xed, 0xb7, 0x5a, 0x39, 0x99, 0xfc, 0xdb,
	0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x71, 0x19, 0xba, 0xec, 0x23, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CategoryManagerClient is the client API for CategoryManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CategoryManagerClient interface {
	DescribeCategories(ctx context.Context, in *DescribeCategoriesRequest, opts ...grpc.CallOption) (*DescribeCategoriesResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	ModifyCategory(ctx context.Context, in *ModifyCategoryRequest, opts ...grpc.CallOption) (*ModifyCategoryResponse, error)
	DeleteCategories(ctx context.Context, in *DeleteCategoriesRequest, opts ...grpc.CallOption) (*DeleteCategoriesResponse, error)
}

type categoryManagerClient struct {
	cc *grpc.ClientConn
}

func NewCategoryManagerClient(cc *grpc.ClientConn) CategoryManagerClient {
	return &categoryManagerClient{cc}
}

func (c *categoryManagerClient) DescribeCategories(ctx context.Context, in *DescribeCategoriesRequest, opts ...grpc.CallOption) (*DescribeCategoriesResponse, error) {
	out := new(DescribeCategoriesResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/DescribeCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryManagerClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryManagerClient) ModifyCategory(ctx context.Context, in *ModifyCategoryRequest, opts ...grpc.CallOption) (*ModifyCategoryResponse, error) {
	out := new(ModifyCategoryResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/ModifyCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryManagerClient) DeleteCategories(ctx context.Context, in *DeleteCategoriesRequest, opts ...grpc.CallOption) (*DeleteCategoriesResponse, error) {
	out := new(DeleteCategoriesResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.CategoryManager/DeleteCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryManagerServer is the server API for CategoryManager service.
type CategoryManagerServer interface {
	DescribeCategories(context.Context, *DescribeCategoriesRequest) (*DescribeCategoriesResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	ModifyCategory(context.Context, *ModifyCategoryRequest) (*ModifyCategoryResponse, error)
	DeleteCategories(context.Context, *DeleteCategoriesRequest) (*DeleteCategoriesResponse, error)
}

func RegisterCategoryManagerServer(s *grpc.Server, srv CategoryManagerServer) {
	s.RegisterService(&_CategoryManager_serviceDesc, srv)
}

func _CategoryManager_DescribeCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).DescribeCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/DescribeCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).DescribeCategories(ctx, req.(*DescribeCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryManager_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryManager_ModifyCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).ModifyCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/ModifyCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).ModifyCategory(ctx, req.(*ModifyCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryManager_DeleteCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryManagerServer).DeleteCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.CategoryManager/DeleteCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryManagerServer).DeleteCategories(ctx, req.(*DeleteCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CategoryManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.CategoryManager",
	HandlerType: (*CategoryManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeCategories",
			Handler:    _CategoryManager_DescribeCategories_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryManager_CreateCategory_Handler,
		},
		{
			MethodName: "ModifyCategory",
			Handler:    _CategoryManager_ModifyCategory_Handler,
		},
		{
			MethodName: "DeleteCategories",
			Handler:    _CategoryManager_DeleteCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}