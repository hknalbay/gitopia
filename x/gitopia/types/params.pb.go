// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DistributionProportion struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Proportion int64  `protobuf:"varint,2,opt,name=proportion,proto3" json:"proportion,omitempty" yaml:"proportion"`
}

func (m *DistributionProportion) Reset()         { *m = DistributionProportion{} }
func (m *DistributionProportion) String() string { return proto.CompactTextString(m) }
func (*DistributionProportion) ProtoMessage()    {}
func (*DistributionProportion) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdae11692a018c3a, []int{0}
}
func (m *DistributionProportion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportion.Merge(m, src)
}
func (m *DistributionProportion) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportion) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportion.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportion proto.InternalMessageInfo

func (m *DistributionProportion) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DistributionProportion) GetProportion() int64 {
	if m != nil {
		return m.Proportion
	}
	return 0
}

type DistributionProportions struct {
	EcosystemProportion DistributionProportion `protobuf:"bytes,1,opt,name=ecosystem_proportion,json=ecosystemProportion,proto3" json:"ecosystem_proportion" yaml:"ecosystem_proportion"`
	TeamProportion      DistributionProportion `protobuf:"bytes,2,opt,name=team_proportion,json=teamProportion,proto3" json:"team_proportion" yaml:"team_proportion"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdae11692a018c3a, []int{1}
}
func (m *DistributionProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportions.Merge(m, src)
}
func (m *DistributionProportions) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportions proto.InternalMessageInfo

func (m *DistributionProportions) GetEcosystemProportion() DistributionProportion {
	if m != nil {
		return m.EcosystemProportion
	}
	return DistributionProportion{}
}

func (m *DistributionProportions) GetTeamProportion() DistributionProportion {
	if m != nil {
		return m.TeamProportion
	}
	return DistributionProportion{}
}

// Params defines the parameters for the module.
type Params struct {
	NextInflationTime       time.Time               `protobuf:"bytes,1,opt,name=next_inflation_time,json=nextInflationTime,proto3,stdtime" json:"next_inflation_time" yaml:"next_inflation_time"`
	DistributionProportions DistributionProportions `protobuf:"bytes,2,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions" yaml:"distribution_proportions"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdae11692a018c3a, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetNextInflationTime() time.Time {
	if m != nil {
		return m.NextInflationTime
	}
	return time.Time{}
}

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func init() {
	proto.RegisterType((*DistributionProportion)(nil), "gitopia.gitopia.gitopia.DistributionProportion")
	proto.RegisterType((*DistributionProportions)(nil), "gitopia.gitopia.gitopia.DistributionProportions")
	proto.RegisterType((*Params)(nil), "gitopia.gitopia.gitopia.Params")
}

func init() { proto.RegisterFile("gitopia/params.proto", fileDescriptor_cdae11692a018c3a) }

var fileDescriptor_cdae11692a018c3a = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xce, 0x44, 0x59, 0x71, 0x84, 0x95, 0xcd, 0xd6, 0x6d, 0x89, 0x90, 0x59, 0x46, 0xd0, 0x45,
	0x64, 0x22, 0x8a, 0x97, 0x3d, 0x86, 0x5e, 0xbc, 0x95, 0xe0, 0xc9, 0x4b, 0x99, 0x34, 0xd3, 0x38,
	0xd0, 0x74, 0x42, 0x66, 0x0a, 0xed, 0xc1, 0xb3, 0xd7, 0x1e, 0x3c, 0xe8, 0x41, 0xf0, 0xe7, 0xf4,
	0xd8, 0xa3, 0xa7, 0x28, 0xed, 0x3f, 0xe8, 0x2f, 0x90, 0xc9, 0x24, 0x6d, 0x2c, 0xe9, 0xc1, 0x3d,
	0xcd, 0xeb, 0x7b, 0xdf, 0xfb, 0xde, 0xd7, 0xef, 0x23, 0xb0, 0x93, 0x70, 0x25, 0x32, 0x4e, 0xfd,
	0x8c, 0xe6, 0x34, 0x95, 0x24, 0xcb, 0x85, 0x12, 0x4e, 0xb7, 0xea, 0x92, 0xa3, 0xd7, 0xed, 0x24,
	0x22, 0x11, 0x25, 0xc6, 0xd7, 0x95, 0x81, 0xbb, 0x28, 0x11, 0x22, 0x99, 0x30, 0xbf, 0xfc, 0x15,
	0xcd, 0xc6, 0xbe, 0xe2, 0x29, 0x93, 0x8a, 0xa6, 0x99, 0x01, 0xe0, 0xcf, 0xf0, 0xaa, 0xcf, 0xa5,
	0xca, 0x79, 0x34, 0x53, 0x5c, 0x4c, 0x07, 0xb9, 0xc8, 0x44, 0xae, 0x2b, 0xe7, 0x15, 0x7c, 0x40,
	0xe3, 0x38, 0x67, 0x52, 0xf6, 0xc0, 0x35, 0xb8, 0x79, 0x18, 0x38, 0xbb, 0x02, 0x9d, 0x2f, 0x68,
	0x3a, 0xb9, 0xc5, 0xd5, 0x00, 0x87, 0x35, 0xc4, 0x79, 0x07, 0x61, 0xb6, 0xdf, 0xed, 0xd9, 0xd7,
	0xe0, 0xe6, 0x5e, 0xf0, 0x64, 0x57, 0xa0, 0x0b, 0xb3, 0x70, 0x98, 0xe1, 0xb0, 0x01, 0xc4, 0x3f,
	0x6c, 0xd8, 0x6d, 0xbf, 0x2f, 0x9d, 0x2f, 0x00, 0x76, 0xd8, 0x48, 0xc8, 0x85, 0x54, 0x2c, 0x1d,
	0x36, 0xd8, 0xb5, 0x9c, 0x47, 0x6f, 0x7c, 0x72, 0xc2, 0x0a, 0xd2, 0x4e, 0x18, 0x3c, 0x5b, 0x15,
	0xc8, 0xda, 0x15, 0xe8, 0xa9, 0x91, 0xd4, 0x46, 0x8d, 0xc3, 0xcb, 0x7d, 0xbb, 0x61, 0xc5, 0x1c,
	0x3e, 0x56, 0x8c, 0xfe, 0xa3, 0xc1, 0xbe, 0x9b, 0x06, 0xaf, 0xd2, 0x70, 0x65, 0x34, 0x1c, 0xb1,
	0xe2, 0xf0, 0x5c, 0x77, 0x0e, 0x78, 0xfc, 0xdd, 0x86, 0x67, 0x83, 0x32, 0x7f, 0x27, 0x87, 0x97,
	0x53, 0x36, 0x57, 0x43, 0x3e, 0x1d, 0x4f, 0xa8, 0x1e, 0x0e, 0x75, 0x96, 0x95, 0x19, 0x2e, 0x31,
	0x41, 0x93, 0x3a, 0x68, 0xf2, 0xa1, 0x0e, 0x3a, 0x78, 0x5e, 0xdd, 0x74, 0xcd, 0xcd, 0x16, 0x12,
	0xbc, 0xfc, 0x8d, 0x40, 0x78, 0xa1, 0x27, 0xef, 0xeb, 0x81, 0xde, 0x77, 0xbe, 0x02, 0xd8, 0x8b,
	0x1b, 0xff, 0xa4, 0xa1, 0x55, 0x56, 0x16, 0xbc, 0xfe, 0x4f, 0x0b, 0x64, 0xf0, 0xa2, 0xd2, 0x83,
	0x8c, 0x9e, 0x53, 0xfc, 0x38, 0xec, 0xc6, 0xed, 0x0c, 0xb7, 0xf7, 0xbf, 0xfd, 0x44, 0x56, 0xd0,
	0x5f, 0x6d, 0x3c, 0xb0, 0xde, 0x78, 0xe0, 0xcf, 0xc6, 0x03, 0xcb, 0xad, 0x67, 0xad, 0xb7, 0x9e,
	0xf5, 0x6b, 0xeb, 0x59, 0x1f, 0x5f, 0x26, 0x5c, 0x7d, 0x9a, 0x45, 0x64, 0x24, 0x52, 0xbf, 0xfe,
	0x8a, 0xea, 0x77, 0xbe, 0xaf, 0xd4, 0x22, 0x63, 0x32, 0x3a, 0x2b, 0x1d, 0x7b, 0xfb, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x73, 0x00, 0xbe, 0x04, 0x6f, 0x03, 0x00, 0x00,
}

func (m *DistributionProportion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proportion != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Proportion))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DistributionProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TeamProportion.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.EcosystemProportion.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.NextInflationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.NextInflationTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintParams(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DistributionProportion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Proportion != 0 {
		n += 1 + sovParams(uint64(m.Proportion))
	}
	return n
}

func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EcosystemProportion.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.TeamProportion.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.NextInflationTime)
	n += 1 + l + sovParams(uint64(l))
	l = m.DistributionProportions.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DistributionProportion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DistributionProportion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proportion", wireType)
			}
			m.Proportion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Proportion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DistributionProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EcosystemProportion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EcosystemProportion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamProportion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TeamProportion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextInflationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.NextInflationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
