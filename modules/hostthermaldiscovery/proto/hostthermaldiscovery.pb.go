// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hostthermaldiscovery.proto

package proto

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

type HostDiscoveryConfig struct {
	PollingInterval      string   `protobuf:"bytes,1,opt,name=polling_interval,json=pollingInterval,proto3" json:"polling_interval,omitempty"`
	TempSensorPath       string   `protobuf:"bytes,2,opt,name=temp_sensor_path,json=tempSensorPath,proto3" json:"temp_sensor_path,omitempty"`
	FreqSensorUrl        string   `protobuf:"bytes,3,opt,name=freq_sensor_url,json=freqSensorUrl,proto3" json:"freq_sensor_url,omitempty"`
	LogThermalData       bool     `protobuf:"varint,4,opt,name=log_thermal_data,json=logThermalData,proto3" json:"log_thermal_data,omitempty"`
	LogHere              string   `protobuf:"bytes,5,opt,name=log_here,json=logHere,proto3" json:"log_here,omitempty"`
	LowerNormal          int32    `protobuf:"varint,6,opt,name=lower_normal,json=lowerNormal,proto3" json:"lower_normal,omitempty"`
	UpperNormal          int32    `protobuf:"varint,7,opt,name=upper_normal,json=upperNormal,proto3" json:"upper_normal,omitempty"`
	LowerHigh            int32    `protobuf:"varint,8,opt,name=lower_high,json=lowerHigh,proto3" json:"lower_high,omitempty"`
	UpperHigh            int32    `protobuf:"varint,9,opt,name=upper_high,json=upperHigh,proto3" json:"upper_high,omitempty"`
	LowerCritical        int32    `protobuf:"varint,10,opt,name=lower_critical,json=lowerCritical,proto3" json:"lower_critical,omitempty"`
	UpperCritical        int32    `protobuf:"varint,11,opt,name=upper_critical,json=upperCritical,proto3" json:"upper_critical,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostDiscoveryConfig) Reset()         { *m = HostDiscoveryConfig{} }
func (m *HostDiscoveryConfig) String() string { return proto.CompactTextString(m) }
func (*HostDiscoveryConfig) ProtoMessage()    {}
func (*HostDiscoveryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_hostthermaldiscovery_01a159d510988b58, []int{0}
}
func (m *HostDiscoveryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostDiscoveryConfig.Unmarshal(m, b)
}
func (m *HostDiscoveryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostDiscoveryConfig.Marshal(b, m, deterministic)
}
func (dst *HostDiscoveryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostDiscoveryConfig.Merge(dst, src)
}
func (m *HostDiscoveryConfig) XXX_Size() int {
	return xxx_messageInfo_HostDiscoveryConfig.Size(m)
}
func (m *HostDiscoveryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HostDiscoveryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HostDiscoveryConfig proto.InternalMessageInfo

func (m *HostDiscoveryConfig) GetPollingInterval() string {
	if m != nil {
		return m.PollingInterval
	}
	return ""
}

func (m *HostDiscoveryConfig) GetTempSensorPath() string {
	if m != nil {
		return m.TempSensorPath
	}
	return ""
}

func (m *HostDiscoveryConfig) GetFreqSensorUrl() string {
	if m != nil {
		return m.FreqSensorUrl
	}
	return ""
}

func (m *HostDiscoveryConfig) GetLogThermalData() bool {
	if m != nil {
		return m.LogThermalData
	}
	return false
}

func (m *HostDiscoveryConfig) GetLogHere() string {
	if m != nil {
		return m.LogHere
	}
	return ""
}

func (m *HostDiscoveryConfig) GetLowerNormal() int32 {
	if m != nil {
		return m.LowerNormal
	}
	return 0
}

func (m *HostDiscoveryConfig) GetUpperNormal() int32 {
	if m != nil {
		return m.UpperNormal
	}
	return 0
}

func (m *HostDiscoveryConfig) GetLowerHigh() int32 {
	if m != nil {
		return m.LowerHigh
	}
	return 0
}

func (m *HostDiscoveryConfig) GetUpperHigh() int32 {
	if m != nil {
		return m.UpperHigh
	}
	return 0
}

func (m *HostDiscoveryConfig) GetLowerCritical() int32 {
	if m != nil {
		return m.LowerCritical
	}
	return 0
}

func (m *HostDiscoveryConfig) GetUpperCritical() int32 {
	if m != nil {
		return m.UpperCritical
	}
	return 0
}

func init() {
	proto.RegisterType((*HostDiscoveryConfig)(nil), "proto.HostDiscoveryConfig")
}

func init() {
	proto.RegisterFile("hostthermaldiscovery.proto", fileDescriptor_hostthermaldiscovery_01a159d510988b58)
}

var fileDescriptor_hostthermaldiscovery_01a159d510988b58 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd1, 0xcb, 0x4a, 0x03, 0x31,
	0x14, 0xc6, 0x71, 0xc6, 0xda, 0x5b, 0x6a, 0x2f, 0xc4, 0x4d, 0x14, 0x84, 0x2a, 0x28, 0xe3, 0xc6,
	0x8d, 0x8f, 0xd0, 0x2e, 0xea, 0x46, 0xa4, 0xea, 0x3a, 0xc4, 0x36, 0x4d, 0x02, 0xe9, 0x9c, 0x78,
	0x26, 0xad, 0xf8, 0x48, 0xbe, 0xa5, 0xe4, 0x64, 0x46, 0x57, 0x03, 0xff, 0xf3, 0x9b, 0x8f, 0x81,
	0x61, 0x97, 0x16, 0xea, 0x18, 0xad, 0xc6, 0xbd, 0xf2, 0x5b, 0x57, 0x6f, 0xe0, 0xa8, 0xf1, 0xfb,
	0x21, 0x20, 0x44, 0xe0, 0x5d, 0x7a, 0xdc, 0xfc, 0x74, 0xd8, 0xf9, 0x0a, 0xea, 0xb8, 0x6c, 0xcf,
	0x0b, 0xa8, 0x76, 0xce, 0xf0, 0x7b, 0x36, 0x0b, 0xe0, 0xbd, 0xab, 0x8c, 0x74, 0x55, 0xd4, 0x78,
	0x54, 0x5e, 0x14, 0xf3, 0xa2, 0x1c, 0xae, 0xa7, 0x4d, 0x7f, 0x6a, 0x32, 0x2f, 0xd9, 0x2c, 0xea,
	0x7d, 0x90, 0xb5, 0xae, 0x6a, 0x40, 0x19, 0x54, 0xb4, 0xe2, 0x84, 0xe8, 0x24, 0xf5, 0x57, 0xca,
	0x2f, 0x2a, 0x5a, 0x7e, 0xc7, 0xa6, 0x3b, 0xd4, 0x9f, 0xad, 0x3c, 0xa0, 0x17, 0x1d, 0x82, 0xe3,
	0x94, 0x33, 0x7c, 0x47, 0x5a, 0xf4, 0x60, 0x64, 0xf3, 0xe5, 0x72, 0xab, 0xa2, 0x12, 0xa7, 0xf3,
	0xa2, 0x1c, 0xac, 0x27, 0x1e, 0xcc, 0x5b, 0xce, 0x4b, 0x15, 0x15, 0xbf, 0x60, 0x83, 0x24, 0xad,
	0x46, 0x2d, 0xba, 0x34, 0xd5, 0xf7, 0x60, 0x56, 0x1a, 0x35, 0xbf, 0x66, 0x67, 0x1e, 0xbe, 0x34,
	0xca, 0x0a, 0x12, 0x17, 0xbd, 0x79, 0x51, 0x76, 0xd7, 0x23, 0x6a, 0xcf, 0x94, 0x12, 0x39, 0x84,
	0xf0, 0x4f, 0xfa, 0x99, 0x50, 0x6b, 0xc8, 0x15, 0x63, 0x79, 0xc5, 0x3a, 0x63, 0xc5, 0x80, 0xc0,
	0x90, 0xca, 0xca, 0x19, 0x9b, 0xce, 0x79, 0x81, 0xce, 0xc3, 0x7c, 0xa6, 0x42, 0xe7, 0x5b, 0x36,
	0xc9, 0x6f, 0x6f, 0xd0, 0x45, 0xb7, 0x51, 0x5e, 0x30, 0x22, 0x63, 0xaa, 0x8b, 0x26, 0x26, 0x96,
	0x57, 0xfe, 0xd8, 0x28, 0x33, 0xaa, 0x2d, 0xfb, 0xe8, 0xd1, 0x2f, 0x7b, 0xfc, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x43, 0xeb, 0x7e, 0x04, 0xd7, 0x01, 0x00, 0x00,
}
