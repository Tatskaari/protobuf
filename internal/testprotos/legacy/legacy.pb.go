// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/legacy/legacy.proto

package legacy

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto2_v0_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v0.0.0-20160225-2fc053c5"
	proto2_v0_01 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v0.0.0-20160519-a4ab9ec5"
	proto2_v1_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.0.0-20180125-92554152"
	proto2_v1_1 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.1.0-20180430-b4deda09"
	proto2_v1_2 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.2.0-20180814-aa810b61"
	proto2_v1_21 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.2.1-20181126-8d0c54c1"
	proto3_v0_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v0.0.0-20160225-2fc053c5"
	proto3_v0_01 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v0.0.0-20160519-a4ab9ec5"
	proto3_v1_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.0.0-20180125-92554152"
	proto3_v1_1 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.1.0-20180430-b4deda09"
	proto3_v1_2 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.2.0-20180814-aa810b61"
	proto3_v1_21 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.2.1-20181126-8d0c54c1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Legacy struct {
	F1                   *proto2_v0_0.Message  `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
	F2                   *proto3_v0_0.Message  `protobuf:"bytes,2,opt,name=f2,proto3" json:"f2,omitempty"`
	F3                   *proto2_v0_01.Message `protobuf:"bytes,3,opt,name=f3,proto3" json:"f3,omitempty"`
	F4                   *proto3_v0_01.Message `protobuf:"bytes,4,opt,name=f4,proto3" json:"f4,omitempty"`
	F5                   *proto2_v1_0.Message  `protobuf:"bytes,5,opt,name=f5,proto3" json:"f5,omitempty"`
	F6                   *proto3_v1_0.Message  `protobuf:"bytes,6,opt,name=f6,proto3" json:"f6,omitempty"`
	F7                   *proto2_v1_1.Message  `protobuf:"bytes,7,opt,name=f7,proto3" json:"f7,omitempty"`
	F8                   *proto3_v1_1.Message  `protobuf:"bytes,8,opt,name=f8,proto3" json:"f8,omitempty"`
	F9                   *proto2_v1_2.Message  `protobuf:"bytes,9,opt,name=f9,proto3" json:"f9,omitempty"`
	F10                  *proto3_v1_2.Message  `protobuf:"bytes,10,opt,name=f10,proto3" json:"f10,omitempty"`
	F11                  *proto2_v1_21.Message `protobuf:"bytes,11,opt,name=f11,proto3" json:"f11,omitempty"`
	F12                  *proto3_v1_21.Message `protobuf:"bytes,12,opt,name=f12,proto3" json:"f12,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Legacy) Reset()         { *m = Legacy{} }
func (m *Legacy) String() string { return proto.CompactTextString(m) }
func (*Legacy) ProtoMessage()    {}
func (*Legacy) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fbf266527ec81d8, []int{0}
}

func (m *Legacy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Legacy.Unmarshal(m, b)
}
func (m *Legacy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Legacy.Marshal(b, m, deterministic)
}
func (m *Legacy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Legacy.Merge(m, src)
}
func (m *Legacy) XXX_Size() int {
	return xxx_messageInfo_Legacy.Size(m)
}
func (m *Legacy) XXX_DiscardUnknown() {
	xxx_messageInfo_Legacy.DiscardUnknown(m)
}

var xxx_messageInfo_Legacy proto.InternalMessageInfo

func (m *Legacy) GetF1() *proto2_v0_0.Message {
	if m != nil {
		return m.F1
	}
	return nil
}

func (m *Legacy) GetF2() *proto3_v0_0.Message {
	if m != nil {
		return m.F2
	}
	return nil
}

func (m *Legacy) GetF3() *proto2_v0_01.Message {
	if m != nil {
		return m.F3
	}
	return nil
}

func (m *Legacy) GetF4() *proto3_v0_01.Message {
	if m != nil {
		return m.F4
	}
	return nil
}

func (m *Legacy) GetF5() *proto2_v1_0.Message {
	if m != nil {
		return m.F5
	}
	return nil
}

func (m *Legacy) GetF6() *proto3_v1_0.Message {
	if m != nil {
		return m.F6
	}
	return nil
}

func (m *Legacy) GetF7() *proto2_v1_1.Message {
	if m != nil {
		return m.F7
	}
	return nil
}

func (m *Legacy) GetF8() *proto3_v1_1.Message {
	if m != nil {
		return m.F8
	}
	return nil
}

func (m *Legacy) GetF9() *proto2_v1_2.Message {
	if m != nil {
		return m.F9
	}
	return nil
}

func (m *Legacy) GetF10() *proto3_v1_2.Message {
	if m != nil {
		return m.F10
	}
	return nil
}

func (m *Legacy) GetF11() *proto2_v1_21.Message {
	if m != nil {
		return m.F11
	}
	return nil
}

func (m *Legacy) GetF12() *proto3_v1_21.Message {
	if m != nil {
		return m.F12
	}
	return nil
}

func init() {
	proto.RegisterType((*Legacy)(nil), "google.golang.org.Legacy")
}

func init() {
	proto.RegisterFile("internal/testprotos/legacy/legacy.proto", fileDescriptor_5fbf266527ec81d8)
}

var fileDescriptor_5fbf266527ec81d8 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd3, 0xcf, 0xce, 0xd2, 0x40,
	0x14, 0x05, 0xf0, 0x00, 0x82, 0x5a, 0xdc, 0xc8, 0x6a, 0xe2, 0xca, 0xb8, 0xd1, 0x98, 0x74, 0x3a,
	0xf7, 0x76, 0xa6, 0x4c, 0x89, 0x2b, 0x13, 0x77, 0xba, 0x71, 0xe9, 0xc6, 0x4c, 0xcb, 0x74, 0x24,
	0xa9, 0xd4, 0xd0, 0x42, 0xe2, 0xcb, 0xfa, 0x2c, 0x86, 0x5b, 0x9a, 0x14, 0xf4, 0xe3, 0x1b, 0xfe,
	0xac, 0x08, 0x4d, 0x7e, 0xe7, 0x9c, 0xc5, 0xdc, 0xe0, 0xed, 0x6a, 0xdd, 0xd8, 0xcd, 0xda, 0x94,
	0x51, 0x63, 0xeb, 0xe6, 0xd7, 0xa6, 0x6a, 0xaa, 0x3a, 0x2a, 0xad, 0x33, 0xf9, 0xef, 0xc3, 0x0f,
	0xa7, 0x8f, 0xb3, 0x97, 0xae, 0xaa, 0x5c, 0x69, 0xb9, 0xab, 0x4a, 0xb3, 0x76, 0xbc, 0xda, 0xb8,
	0x57, 0x9f, 0xce, 0x58, 0xfa, 0x87, 0x7c, 0x27, 0xb8, 0xe0, 0x22, 0x44, 0x01, 0x89, 0x40, 0x54,
	0x21, 0x16, 0xb9, 0x50, 0x71, 0xae, 0x48, 0xb4, 0xc9, 0x8f, 0xc7, 0xc4, 0xf7, 0x89, 0x39, 0x5e,
	0xa3, 0x20, 0x0d, 0x8d, 0x34, 0x59, 0x6a, 0x6f, 0x5a, 0x73, 0x75, 0x0c, 0xf2, 0x1d, 0x74, 0x31,
	0x5a, 0x00, 0xaa, 0x30, 0x45, 0xa5, 0x24, 0x28, 0xbc, 0x74, 0xcd, 0x3d, 0x62, 0x68, 0x0d, 0x74,
	0x31, 0x32, 0x16, 0x61, 0x26, 0x97, 0x76, 0x69, 0x44, 0x7a, 0xf9, 0x9a, 0xdb, 0x63, 0x68, 0x0d,
	0x76, 0x31, 0x1a, 0x64, 0x68, 0x8c, 0x06, 0x91, 0x25, 0x70, 0xf9, 0x9a, 0xdb, 0x63, 0x0e, 0x6b,
	0x80, 0x62, 0x00, 0x30, 0x09, 0xf5, 0x52, 0xe4, 0x4a, 0xe6, 0x57, 0xad, 0x39, 0x1f, 0xf3, 0xe6,
	0xcf, 0x38, 0x98, 0x7c, 0x26, 0x35, 0x5b, 0x04, 0xc3, 0x02, 0xd8, 0xe0, 0xf5, 0xe0, 0xdd, 0x14,
	0xdf, 0xf3, 0x7f, 0xce, 0xaf, 0x05, 0xf8, 0xbd, 0x3b, 0x0a, 0xfe, 0xc5, 0xd6, 0xb5, 0x71, 0xf6,
	0xeb, 0xb0, 0x00, 0xb2, 0xc8, 0x86, 0xe7, 0x6d, 0xfc, 0x3f, 0x8b, 0x64, 0x63, 0x36, 0xf2, 0xea,
	0x55, 0x90, 0xf6, 0x6c, 0x4c, 0x56, 0xb2, 0x27, 0x5e, 0xbd, 0xc7, 0x56, 0x92, 0x55, 0x6c, 0xec,
	0xd1, 0x4b, 0x0f, 0xbd, 0x67, 0x15, 0xd9, 0x84, 0x4d, 0x3c, 0x7a, 0x4f, 0x6d, 0x42, 0x76, 0xce,
	0x9e, 0x7a, 0xf5, 0xca, 0x58, 0xf4, 0xec, 0x9c, 0xac, 0x66, 0xcf, 0xbc, 0x7a, 0x8f, 0xad, 0x26,
	0x9b, 0xb2, 0xe7, 0x5e, 0xbd, 0x1a, 0x64, 0xcf, 0xa6, 0xb3, 0x0f, 0xc1, 0xa8, 0x00, 0xc1, 0x02,
	0xaf, 0xe2, 0x3e, 0xde, 0xb3, 0x56, 0x03, 0x9b, 0xfa, 0x54, 0xef, 0x5f, 0x6a, 0x5f, 0x43, 0xab,
	0x91, 0xbd, 0xf0, 0xe9, 0x3e, 0xd5, 0xf8, 0x71, 0xf1, 0x4d, 0xbb, 0x55, 0xf3, 0x63, 0x9b, 0xf1,
	0xbc, 0xfa, 0x19, 0xb5, 0xaa, 0x3d, 0x90, 0x6c, 0x5b, 0x44, 0x3b, 0x8c, 0x1e, 0xbe, 0xa3, 0x6c,
	0xd2, 0x26, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x78, 0x39, 0xb9, 0x4a, 0xb5, 0x06, 0x00, 0x00,
}
