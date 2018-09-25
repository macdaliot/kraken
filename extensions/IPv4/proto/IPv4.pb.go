// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IPv4.proto

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

type IPv4 struct {
	Ip                   []byte   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Subnet               []byte   `protobuf:"bytes,3,opt,name=subnet,proto3" json:"subnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPv4) Reset()         { *m = IPv4{} }
func (m *IPv4) String() string { return proto.CompactTextString(m) }
func (*IPv4) ProtoMessage()    {}
func (*IPv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_IPv4_c24d1a3e220808d1, []int{0}
}
func (m *IPv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv4.Unmarshal(m, b)
}
func (m *IPv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv4.Marshal(b, m, deterministic)
}
func (dst *IPv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv4.Merge(dst, src)
}
func (m *IPv4) XXX_Size() int {
	return xxx_messageInfo_IPv4.Size(m)
}
func (m *IPv4) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv4.DiscardUnknown(m)
}

var xxx_messageInfo_IPv4 proto.InternalMessageInfo

func (m *IPv4) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *IPv4) GetSubnet() []byte {
	if m != nil {
		return m.Subnet
	}
	return nil
}

type Ethernet struct {
	Iface                string   `protobuf:"bytes,1,opt,name=iface,proto3" json:"iface,omitempty"`
	Mac                  []byte   `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`
	Mtu                  uint32   `protobuf:"varint,3,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Control              bool     `protobuf:"varint,4,opt,name=control,proto3" json:"control,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ethernet) Reset()         { *m = Ethernet{} }
func (m *Ethernet) String() string { return proto.CompactTextString(m) }
func (*Ethernet) ProtoMessage()    {}
func (*Ethernet) Descriptor() ([]byte, []int) {
	return fileDescriptor_IPv4_c24d1a3e220808d1, []int{1}
}
func (m *Ethernet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ethernet.Unmarshal(m, b)
}
func (m *Ethernet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ethernet.Marshal(b, m, deterministic)
}
func (dst *Ethernet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ethernet.Merge(dst, src)
}
func (m *Ethernet) XXX_Size() int {
	return xxx_messageInfo_Ethernet.Size(m)
}
func (m *Ethernet) XXX_DiscardUnknown() {
	xxx_messageInfo_Ethernet.DiscardUnknown(m)
}

var xxx_messageInfo_Ethernet proto.InternalMessageInfo

func (m *Ethernet) GetIface() string {
	if m != nil {
		return m.Iface
	}
	return ""
}

func (m *Ethernet) GetMac() []byte {
	if m != nil {
		return m.Mac
	}
	return nil
}

func (m *Ethernet) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Ethernet) GetControl() bool {
	if m != nil {
		return m.Control
	}
	return false
}

type IPv4OverEthernet struct {
	Ifaces               []*IPv4OverEthernet_ConfiguredInterface `protobuf:"bytes,1,rep,name=ifaces,proto3" json:"ifaces,omitempty"`
	Routes               []*IPv4                                 `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	DnsNameservers       []*IPv4                                 `protobuf:"bytes,3,rep,name=dns_nameservers,json=dnsNameservers,proto3" json:"dns_nameservers,omitempty"`
	DnsDomains           []*IPv4                                 `protobuf:"bytes,4,rep,name=dns_domains,json=dnsDomains,proto3" json:"dns_domains,omitempty"`
	Hostname             *DNSA                                   `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *IPv4OverEthernet) Reset()         { *m = IPv4OverEthernet{} }
func (m *IPv4OverEthernet) String() string { return proto.CompactTextString(m) }
func (*IPv4OverEthernet) ProtoMessage()    {}
func (*IPv4OverEthernet) Descriptor() ([]byte, []int) {
	return fileDescriptor_IPv4_c24d1a3e220808d1, []int{2}
}
func (m *IPv4OverEthernet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv4OverEthernet.Unmarshal(m, b)
}
func (m *IPv4OverEthernet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv4OverEthernet.Marshal(b, m, deterministic)
}
func (dst *IPv4OverEthernet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv4OverEthernet.Merge(dst, src)
}
func (m *IPv4OverEthernet) XXX_Size() int {
	return xxx_messageInfo_IPv4OverEthernet.Size(m)
}
func (m *IPv4OverEthernet) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv4OverEthernet.DiscardUnknown(m)
}

var xxx_messageInfo_IPv4OverEthernet proto.InternalMessageInfo

func (m *IPv4OverEthernet) GetIfaces() []*IPv4OverEthernet_ConfiguredInterface {
	if m != nil {
		return m.Ifaces
	}
	return nil
}

func (m *IPv4OverEthernet) GetRoutes() []*IPv4 {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *IPv4OverEthernet) GetDnsNameservers() []*IPv4 {
	if m != nil {
		return m.DnsNameservers
	}
	return nil
}

func (m *IPv4OverEthernet) GetDnsDomains() []*IPv4 {
	if m != nil {
		return m.DnsDomains
	}
	return nil
}

func (m *IPv4OverEthernet) GetHostname() *DNSA {
	if m != nil {
		return m.Hostname
	}
	return nil
}

type IPv4OverEthernet_ConfiguredInterface struct {
	Eth                  *Ethernet `protobuf:"bytes,1,opt,name=eth,proto3" json:"eth,omitempty"`
	Ip                   *IPv4     `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IPv4OverEthernet_ConfiguredInterface) Reset()         { *m = IPv4OverEthernet_ConfiguredInterface{} }
func (m *IPv4OverEthernet_ConfiguredInterface) String() string { return proto.CompactTextString(m) }
func (*IPv4OverEthernet_ConfiguredInterface) ProtoMessage()    {}
func (*IPv4OverEthernet_ConfiguredInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_IPv4_c24d1a3e220808d1, []int{2, 0}
}
func (m *IPv4OverEthernet_ConfiguredInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv4OverEthernet_ConfiguredInterface.Unmarshal(m, b)
}
func (m *IPv4OverEthernet_ConfiguredInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv4OverEthernet_ConfiguredInterface.Marshal(b, m, deterministic)
}
func (dst *IPv4OverEthernet_ConfiguredInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv4OverEthernet_ConfiguredInterface.Merge(dst, src)
}
func (m *IPv4OverEthernet_ConfiguredInterface) XXX_Size() int {
	return xxx_messageInfo_IPv4OverEthernet_ConfiguredInterface.Size(m)
}
func (m *IPv4OverEthernet_ConfiguredInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv4OverEthernet_ConfiguredInterface.DiscardUnknown(m)
}

var xxx_messageInfo_IPv4OverEthernet_ConfiguredInterface proto.InternalMessageInfo

func (m *IPv4OverEthernet_ConfiguredInterface) GetEth() *Ethernet {
	if m != nil {
		return m.Eth
	}
	return nil
}

func (m *IPv4OverEthernet_ConfiguredInterface) GetIp() *IPv4 {
	if m != nil {
		return m.Ip
	}
	return nil
}

type DNSA struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Domainname           string   `protobuf:"bytes,2,opt,name=domainname,proto3" json:"domainname,omitempty"`
	Ip                   *IPv4    `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNSA) Reset()         { *m = DNSA{} }
func (m *DNSA) String() string { return proto.CompactTextString(m) }
func (*DNSA) ProtoMessage()    {}
func (*DNSA) Descriptor() ([]byte, []int) {
	return fileDescriptor_IPv4_c24d1a3e220808d1, []int{3}
}
func (m *DNSA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSA.Unmarshal(m, b)
}
func (m *DNSA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSA.Marshal(b, m, deterministic)
}
func (dst *DNSA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSA.Merge(dst, src)
}
func (m *DNSA) XXX_Size() int {
	return xxx_messageInfo_DNSA.Size(m)
}
func (m *DNSA) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSA.DiscardUnknown(m)
}

var xxx_messageInfo_DNSA proto.InternalMessageInfo

func (m *DNSA) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *DNSA) GetDomainname() string {
	if m != nil {
		return m.Domainname
	}
	return ""
}

func (m *DNSA) GetIp() *IPv4 {
	if m != nil {
		return m.Ip
	}
	return nil
}

type DNSCNAME struct {
	Cname                string   `protobuf:"bytes,1,opt,name=cname,proto3" json:"cname,omitempty"`
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNSCNAME) Reset()         { *m = DNSCNAME{} }
func (m *DNSCNAME) String() string { return proto.CompactTextString(m) }
func (*DNSCNAME) ProtoMessage()    {}
func (*DNSCNAME) Descriptor() ([]byte, []int) {
	return fileDescriptor_IPv4_c24d1a3e220808d1, []int{4}
}
func (m *DNSCNAME) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSCNAME.Unmarshal(m, b)
}
func (m *DNSCNAME) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSCNAME.Marshal(b, m, deterministic)
}
func (dst *DNSCNAME) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSCNAME.Merge(dst, src)
}
func (m *DNSCNAME) XXX_Size() int {
	return xxx_messageInfo_DNSCNAME.Size(m)
}
func (m *DNSCNAME) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSCNAME.DiscardUnknown(m)
}

var xxx_messageInfo_DNSCNAME proto.InternalMessageInfo

func (m *DNSCNAME) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *DNSCNAME) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*IPv4)(nil), "proto.IPv4")
	proto.RegisterType((*Ethernet)(nil), "proto.Ethernet")
	proto.RegisterType((*IPv4OverEthernet)(nil), "proto.IPv4OverEthernet")
	proto.RegisterType((*IPv4OverEthernet_ConfiguredInterface)(nil), "proto.IPv4OverEthernet.ConfiguredInterface")
	proto.RegisterType((*DNSA)(nil), "proto.DNSA")
	proto.RegisterType((*DNSCNAME)(nil), "proto.DNSCNAME")
}

func init() { proto.RegisterFile("IPv4.proto", fileDescriptor_IPv4_c24d1a3e220808d1) }

var fileDescriptor_IPv4_c24d1a3e220808d1 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x5f, 0xef, 0x93, 0x30,
	0x14, 0x0d, 0x85, 0x21, 0x5e, 0xf4, 0xb7, 0xa5, 0x1a, 0xd3, 0xcc, 0xc4, 0x20, 0x3e, 0x48, 0xa2,
	0xe1, 0x61, 0xee, 0xd1, 0x97, 0x65, 0xdb, 0xc3, 0x1e, 0x44, 0xd3, 0xc5, 0x57, 0x17, 0x06, 0x9d,
	0x23, 0x71, 0x2d, 0x69, 0xcb, 0x3e, 0xaf, 0x1f, 0xc5, 0xb4, 0x30, 0x64, 0x7f, 0x9e, 0xe8, 0xb9,
	0xe7, 0x9e, 0x73, 0xee, 0x6d, 0x01, 0xd8, 0xfc, 0x38, 0xcf, 0xd3, 0x5a, 0x0a, 0x2d, 0xf0, 0xc8,
	0x7e, 0xe2, 0x14, 0x3c, 0x53, 0xc4, 0x4f, 0x80, 0xaa, 0x9a, 0xa0, 0xc8, 0x49, 0x5e, 0x50, 0x54,
	0xd5, 0xf8, 0x0d, 0xf8, 0xaa, 0xd9, 0x73, 0xa6, 0x89, 0x6b, 0x6b, 0x1d, 0x8a, 0x7f, 0x41, 0xb0,
	0xd6, 0x47, 0x26, 0x39, 0xd3, 0xf8, 0x35, 0x8c, 0xaa, 0x43, 0x5e, 0x30, 0xe2, 0x44, 0x4e, 0xf2,
	0x9c, 0xb6, 0x00, 0x4f, 0xc0, 0x3d, 0xe5, 0x45, 0x67, 0x65, 0x8e, 0xb6, 0xa2, 0x1b, 0x6b, 0xf4,
	0x92, 0x9a, 0x23, 0x26, 0xf0, 0xac, 0x10, 0x5c, 0x4b, 0xf1, 0x87, 0x78, 0x91, 0x93, 0x04, 0xf4,
	0x02, 0xe3, 0xbf, 0x08, 0x26, 0x66, 0xa0, 0xef, 0x67, 0x26, 0xfb, 0xa0, 0x25, 0xf8, 0xd6, 0x5b,
	0x11, 0x27, 0x72, 0x93, 0x70, 0xf6, 0xa9, 0xdd, 0x21, 0xbd, 0x6d, 0x4c, 0x97, 0x82, 0x1f, 0xaa,
	0xdf, 0x8d, 0x64, 0xe5, 0x86, 0x6b, 0x26, 0x8d, 0x86, 0x76, 0x52, 0xfc, 0x01, 0x7c, 0x29, 0x1a,
	0xcd, 0x14, 0x41, 0xd6, 0x24, 0x1c, 0x98, 0xd0, 0x8e, 0xc2, 0x73, 0x18, 0x97, 0x5c, 0xed, 0x78,
	0x7e, 0x62, 0x8a, 0xc9, 0x33, 0x93, 0x8a, 0xb8, 0xf7, 0xdd, 0x4f, 0x25, 0x57, 0xd9, 0xff, 0x16,
	0xfc, 0x19, 0x42, 0xa3, 0x2a, 0xc5, 0x29, 0xaf, 0xb8, 0x22, 0xde, 0xbd, 0x02, 0x4a, 0xae, 0x56,
	0x2d, 0x8d, 0x3f, 0x42, 0x70, 0x14, 0x4a, 0x9b, 0x0c, 0x32, 0x8a, 0x9c, 0x41, 0xeb, 0x2a, 0xdb,
	0x2e, 0x68, 0x4f, 0x4e, 0x7f, 0xc2, 0xab, 0x07, 0x0b, 0xe1, 0xf7, 0xe0, 0x32, 0x7d, 0xb4, 0x97,
	0x1e, 0xce, 0xc6, 0x9d, 0xf4, 0x72, 0x05, 0xd4, 0x70, 0xf8, 0x6d, 0xff, 0x9a, 0x37, 0x73, 0xa0,
	0xaa, 0x8e, 0x77, 0xe0, 0x99, 0x20, 0x3c, 0x1d, 0xcc, 0xd1, 0xbe, 0x60, 0x8f, 0xf1, 0x3b, 0x80,
	0x76, 0x1b, 0xcb, 0x22, 0xcb, 0x0e, 0x2a, 0x5d, 0x80, 0xfb, 0x38, 0xe0, 0x2b, 0x04, 0xab, 0x6c,
	0xbb, 0xcc, 0x16, 0xdf, 0xd6, 0xe6, 0x1f, 0x29, 0x06, 0x09, 0x2d, 0xb8, 0x8a, 0x46, 0xd7, 0xd1,
	0x7b, 0xdf, 0xba, 0x7d, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x3e, 0x99, 0x8a, 0xad, 0x02,
	0x00, 0x00,
}
