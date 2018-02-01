// Code generated by protoc-gen-go. DO NOT EDIT.
// source: folderstat.proto

/*
Package folderstat is a generated protocol buffer package.

It is generated from these files:
	folderstat.proto
	folderstat_service.proto

It has these top-level messages:
	GetFolderStatNowRequest
	GetFolderStatNowResponse
	GetFolderStatByDateRequest
	GetFolderStatByDateResponse
	FolderStatistics
	ExecuteResponse
*/
package folderstat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// 获取当前统计请求参数
type GetFolderStatNowRequest struct {
	// *
	// 项目id
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId" json:"folder_id,omitempty"`
}

func (m *GetFolderStatNowRequest) Reset()                    { *m = GetFolderStatNowRequest{} }
func (m *GetFolderStatNowRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFolderStatNowRequest) ProtoMessage()               {}
func (*GetFolderStatNowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetFolderStatNowRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

// *
// 获取当前统计返回值
type GetFolderStatNowResponse struct {
	// *
	// 执行结果
	Result *ExecuteResponse `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	// *
	// 项目统计实体
	Folderstat *FolderStatistics `protobuf:"bytes,2,opt,name=folderstat" json:"folderstat,omitempty"`
}

func (m *GetFolderStatNowResponse) Reset()                    { *m = GetFolderStatNowResponse{} }
func (m *GetFolderStatNowResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFolderStatNowResponse) ProtoMessage()               {}
func (*GetFolderStatNowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetFolderStatNowResponse) GetResult() *ExecuteResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetFolderStatNowResponse) GetFolderstat() *FolderStatistics {
	if m != nil {
		return m.Folderstat
	}
	return nil
}

// *
// 获取时间段的项目统计请求参数
type GetFolderStatByDateRequest struct {
	// *
	// 项目id
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId" json:"folder_id,omitempty"`
	// *
	// 开始时间
	StartDate int64 `protobuf:"varint,2,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	// *
	// 截止时间
	EndDate int64 `protobuf:"varint,3,opt,name=end_date,json=endDate" json:"end_date,omitempty"`
}

func (m *GetFolderStatByDateRequest) Reset()                    { *m = GetFolderStatByDateRequest{} }
func (m *GetFolderStatByDateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFolderStatByDateRequest) ProtoMessage()               {}
func (*GetFolderStatByDateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetFolderStatByDateRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *GetFolderStatByDateRequest) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *GetFolderStatByDateRequest) GetEndDate() int64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

// *
// 获取时间段的项目统计返回值
type GetFolderStatByDateResponse struct {
	// *
	// 执行结果
	Result *ExecuteResponse `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	// *
	// 项目统计实体数组
	Folderstats []*FolderStatistics `protobuf:"bytes,2,rep,name=folderstats" json:"folderstats,omitempty"`
}

func (m *GetFolderStatByDateResponse) Reset()                    { *m = GetFolderStatByDateResponse{} }
func (m *GetFolderStatByDateResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFolderStatByDateResponse) ProtoMessage()               {}
func (*GetFolderStatByDateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetFolderStatByDateResponse) GetResult() *ExecuteResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetFolderStatByDateResponse) GetFolderstats() []*FolderStatistics {
	if m != nil {
		return m.Folderstats
	}
	return nil
}

// *
// 项目统计实体
type FolderStatistics struct {
	// *
	// 项目id
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId" json:"folder_id,omitempty"`
	// *
	// 创建时间
	CreateTime int64 `protobuf:"varint,2,opt,name=createTime" json:"createTime,omitempty"`
	// *
	// 统计的当天日期
	Date int64 `protobuf:"varint,3,opt,name=date" json:"date,omitempty"`
	// *
	// 任务总数
	Amount int32 `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
	// *
	// 进行中任务总数
	Underway int32 `protobuf:"varint,5,opt,name=underway" json:"underway,omitempty"`
	// *
	// 正常状态进行中任务数
	UnderwayN int32 `protobuf:"varint,6,opt,name=underway_n,json=underwayN" json:"underway_n,omitempty"`
	// *
	// 逾期进行中任务数
	UnderwayA int32 `protobuf:"varint,7,opt,name=underway_a,json=underwayA" json:"underway_a,omitempty"`
	// *
	// 未设定截止日期的进行中任务数
	UnderwayU int32 `protobuf:"varint,8,opt,name=underway_u,json=underwayU" json:"underway_u,omitempty"`
	// *
	// 总逾期时间 单位：小时
	Timespan int64 `protobuf:"varint,9,opt,name=timespan" json:"timespan,omitempty"`
	// *
	// 进行中任务总逾期时间 单位：小时
	TimespanUnd int64 `protobuf:"varint,10,opt,name=timespan_und,json=timespanUnd" json:"timespan_und,omitempty"`
	// *
	// 已完成任务总逾期时间 单位：小时
	TimespanCom int64 `protobuf:"varint,11,opt,name=timespan_com,json=timespanCom" json:"timespan_com,omitempty"`
	// *
	// 已完成任务总数
	Completed int32 `protobuf:"varint,12,opt,name=completed" json:"completed,omitempty"`
	// *
	// 正常已完成任务数
	CompletedN int32 `protobuf:"varint,13,opt,name=completed_n,json=completedN" json:"completed_n,omitempty"`
	// *
	// 逾期已完成任务总数
	CompletedA int32 `protobuf:"varint,14,opt,name=completed_a,json=completedA" json:"completed_a,omitempty"`
	// *
	// 未设定截止日期已完成任务数
	CompletedU int32 `protobuf:"varint,15,opt,name=completed_u,json=completedU" json:"completed_u,omitempty"`
	// *
	// 参与成员数
	MemberAmount int32 `protobuf:"varint,16,opt,name=member_amount,json=memberAmount" json:"member_amount,omitempty"`
	// *
	// 任务负责人数
	ChargeAmount int32 `protobuf:"varint,17,opt,name=charge_amount,json=chargeAmount" json:"charge_amount,omitempty"`
	// *
	// 当天新增任务数
	NewTasks int32 `protobuf:"varint,18,opt,name=new_tasks,json=newTasks" json:"new_tasks,omitempty"`
	// *
	// 当天完成任务数
	CompletedTasks int32 `protobuf:"varint,19,opt,name=completed_tasks,json=completedTasks" json:"completed_tasks,omitempty"`
	// *
	// 进行中燃尽时间
	TimespanBurnDown int64 `protobuf:"varint,20,opt,name=timespan_burnDown,json=timespanBurnDown" json:"timespan_burnDown,omitempty"`
}

func (m *FolderStatistics) Reset()                    { *m = FolderStatistics{} }
func (m *FolderStatistics) String() string            { return proto.CompactTextString(m) }
func (*FolderStatistics) ProtoMessage()               {}
func (*FolderStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FolderStatistics) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *FolderStatistics) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FolderStatistics) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *FolderStatistics) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *FolderStatistics) GetUnderway() int32 {
	if m != nil {
		return m.Underway
	}
	return 0
}

func (m *FolderStatistics) GetUnderwayN() int32 {
	if m != nil {
		return m.UnderwayN
	}
	return 0
}

func (m *FolderStatistics) GetUnderwayA() int32 {
	if m != nil {
		return m.UnderwayA
	}
	return 0
}

func (m *FolderStatistics) GetUnderwayU() int32 {
	if m != nil {
		return m.UnderwayU
	}
	return 0
}

func (m *FolderStatistics) GetTimespan() int64 {
	if m != nil {
		return m.Timespan
	}
	return 0
}

func (m *FolderStatistics) GetTimespanUnd() int64 {
	if m != nil {
		return m.TimespanUnd
	}
	return 0
}

func (m *FolderStatistics) GetTimespanCom() int64 {
	if m != nil {
		return m.TimespanCom
	}
	return 0
}

func (m *FolderStatistics) GetCompleted() int32 {
	if m != nil {
		return m.Completed
	}
	return 0
}

func (m *FolderStatistics) GetCompletedN() int32 {
	if m != nil {
		return m.CompletedN
	}
	return 0
}

func (m *FolderStatistics) GetCompletedA() int32 {
	if m != nil {
		return m.CompletedA
	}
	return 0
}

func (m *FolderStatistics) GetCompletedU() int32 {
	if m != nil {
		return m.CompletedU
	}
	return 0
}

func (m *FolderStatistics) GetMemberAmount() int32 {
	if m != nil {
		return m.MemberAmount
	}
	return 0
}

func (m *FolderStatistics) GetChargeAmount() int32 {
	if m != nil {
		return m.ChargeAmount
	}
	return 0
}

func (m *FolderStatistics) GetNewTasks() int32 {
	if m != nil {
		return m.NewTasks
	}
	return 0
}

func (m *FolderStatistics) GetCompletedTasks() int32 {
	if m != nil {
		return m.CompletedTasks
	}
	return 0
}

func (m *FolderStatistics) GetTimespanBurnDown() int64 {
	if m != nil {
		return m.TimespanBurnDown
	}
	return 0
}

// *
// 执行结果
type ExecuteResponse struct {
	// *
	// 执行状态
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// *
	// 信息
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
}

func (m *ExecuteResponse) Reset()                    { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()               {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ExecuteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ExecuteResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFolderStatNowRequest)(nil), "folderstat.GetFolderStatNowRequest")
	proto.RegisterType((*GetFolderStatNowResponse)(nil), "folderstat.GetFolderStatNowResponse")
	proto.RegisterType((*GetFolderStatByDateRequest)(nil), "folderstat.GetFolderStatByDateRequest")
	proto.RegisterType((*GetFolderStatByDateResponse)(nil), "folderstat.GetFolderStatByDateResponse")
	proto.RegisterType((*FolderStatistics)(nil), "folderstat.FolderStatistics")
	proto.RegisterType((*ExecuteResponse)(nil), "folderstat.ExecuteResponse")
}

func init() { proto.RegisterFile("folderstat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x95, 0xa6, 0xf9, 0xe3, 0x49, 0xd2, 0xa4, 0x0b, 0xa2, 0x4b, 0x53, 0x20, 0xa4, 0x07,
	0x2a, 0x21, 0xf5, 0x40, 0x25, 0x4e, 0x08, 0x29, 0x25, 0x80, 0x38, 0x90, 0x83, 0x69, 0xce, 0xd6,
	0xc6, 0x3b, 0x84, 0x88, 0x78, 0x37, 0xec, 0x1f, 0x85, 0x7e, 0x08, 0x2e, 0x88, 0x0f, 0x8c, 0xbc,
	0x6b, 0x3b, 0x76, 0x40, 0xf4, 0xc0, 0x6d, 0xe7, 0xbd, 0xdf, 0x64, 0x66, 0x9f, 0x57, 0x81, 0xc1,
	0x67, 0xb9, 0xe6, 0xa8, 0xb4, 0x61, 0xe6, 0x72, 0xa3, 0xa4, 0x91, 0x04, 0x76, 0xca, 0xf8, 0x25,
	0x9c, 0xbc, 0x47, 0xf3, 0xce, 0x09, 0x9f, 0x0c, 0x33, 0x33, 0xb9, 0x0d, 0xf1, 0x9b, 0x45, 0x6d,
	0xc8, 0x10, 0x02, 0x0f, 0x46, 0x2b, 0x4e, 0x6b, 0xa3, 0xda, 0x45, 0x10, 0xb6, 0xbd, 0xf0, 0x81,
	0x8f, 0x7f, 0xd4, 0x80, 0xfe, 0xd9, 0xa8, 0x37, 0x52, 0x68, 0x24, 0x57, 0xd0, 0x54, 0xa8, 0xed,
	0xda, 0xb8, 0xb6, 0xce, 0x8b, 0xe1, 0x65, 0x69, 0x87, 0xb7, 0xdf, 0x31, 0xb6, 0x06, 0x73, 0x38,
	0xcc, 0x50, 0xf2, 0x0a, 0x4a, 0x7b, 0xd1, 0x03, 0xd7, 0x78, 0x56, 0x6e, 0xdc, 0xcd, 0x5a, 0x69,
	0xb3, 0x8a, 0x75, 0x58, 0xbe, 0x87, 0x86, 0xd3, 0xca, 0x3a, 0xd7, 0xb7, 0x53, 0x96, 0x0e, 0xb9,
	0xfb, 0x2a, 0xe4, 0x11, 0x80, 0x36, 0x4c, 0x99, 0x88, 0x33, 0x83, 0x6e, 0x70, 0x3d, 0x0c, 0x9c,
	0x92, 0xfe, 0x04, 0x79, 0x08, 0x6d, 0x14, 0xdc, 0x9b, 0x75, 0x67, 0xb6, 0x50, 0xf0, 0xd4, 0x1a,
	0xff, 0xac, 0xc1, 0xf0, 0xaf, 0x53, 0xff, 0x27, 0x87, 0xd7, 0xd0, 0xd9, 0x51, 0x9a, 0x1e, 0x8c,
	0xea, 0x77, 0x06, 0x51, 0x6e, 0x18, 0xff, 0x6a, 0xc0, 0x60, 0x9f, 0xf8, 0x77, 0x00, 0x8f, 0x01,
	0x62, 0x85, 0xcc, 0xe0, 0xcd, 0x2a, 0xc9, 0x03, 0x28, 0x29, 0x84, 0xc0, 0x61, 0xe9, 0xf6, 0xee,
	0x4c, 0x1e, 0x40, 0x93, 0x25, 0xd2, 0x0a, 0x43, 0x0f, 0x47, 0xb5, 0x8b, 0x46, 0x98, 0x55, 0xe4,
	0x14, 0xda, 0x56, 0x70, 0x54, 0x5b, 0x76, 0x4b, 0x1b, 0xce, 0x29, 0xea, 0x34, 0xe8, 0xfc, 0x1c,
	0x09, 0xda, 0x74, 0x6e, 0x90, 0x2b, 0xb3, 0x8a, 0xcd, 0x68, 0xab, 0x6a, 0x4f, 0x2a, 0xb6, 0xa5,
	0xed, 0xaa, 0x3d, 0x4f, 0x07, 0x9b, 0x55, 0x82, 0x7a, 0xc3, 0x04, 0x0d, 0xdc, 0xa2, 0x45, 0x4d,
	0x9e, 0x42, 0x37, 0x3f, 0x47, 0x56, 0x70, 0x0a, 0xce, 0xef, 0xe4, 0xda, 0x5c, 0xf0, 0x0a, 0x12,
	0xcb, 0x84, 0x76, 0xaa, 0xc8, 0x1b, 0x99, 0x90, 0x33, 0x08, 0x62, 0x99, 0x6c, 0xd6, 0x68, 0x90,
	0xd3, 0xae, 0x9f, 0x5f, 0x08, 0xe4, 0x09, 0x74, 0x8a, 0x22, 0x12, 0xb4, 0xe7, 0x7c, 0x28, 0xa4,
	0x59, 0x15, 0x60, 0xf4, 0x68, 0x0f, 0x98, 0x54, 0x01, 0x4b, 0xfb, 0x7b, 0xc0, 0x9c, 0x9c, 0x43,
	0x2f, 0xc1, 0x64, 0x81, 0x2a, 0xca, 0xa2, 0x1f, 0x38, 0xa4, 0xeb, 0xc5, 0x89, 0xff, 0x00, 0xe7,
	0xd0, 0x8b, 0xbf, 0x30, 0xb5, 0xc4, 0x1c, 0x3a, 0xf6, 0x90, 0x17, 0x33, 0x68, 0x08, 0x81, 0xc0,
	0x6d, 0x64, 0x98, 0xfe, 0xaa, 0x29, 0xf1, 0x9f, 0x49, 0xe0, 0xf6, 0x26, 0xad, 0xc9, 0x33, 0xe8,
	0xef, 0xf6, 0xf0, 0xc8, 0x3d, 0x87, 0x1c, 0x15, 0xb2, 0x07, 0x9f, 0xc3, 0x71, 0x91, 0xd9, 0xc2,
	0x2a, 0x31, 0x95, 0x5b, 0x41, 0xef, 0xbb, 0xe0, 0x06, 0xb9, 0x71, 0x9d, 0xe9, 0xe3, 0x29, 0xf4,
	0xf7, 0x5e, 0x3c, 0xa1, 0xd0, 0xd2, 0x36, 0x8e, 0x51, 0x6b, 0xf7, 0x24, 0xdb, 0x61, 0x5e, 0x92,
	0x13, 0x68, 0xa1, 0x52, 0x51, 0xa2, 0x97, 0xee, 0x39, 0x06, 0x61, 0x13, 0x95, 0xfa, 0xa8, 0x97,
	0x8b, 0xa6, 0xfb, 0x07, 0xbb, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x84, 0x7a, 0x29, 0xd5,
	0x04, 0x00, 0x00,
}
