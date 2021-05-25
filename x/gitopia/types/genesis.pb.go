// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/genesis.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the capability module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	CommentList    []*Comment    `protobuf:"bytes,5,rep,name=commentList,proto3" json:"commentList,omitempty"`
	IssueList      []*Issue      `protobuf:"bytes,4,rep,name=issueList,proto3" json:"issueList,omitempty"`
	RepositoryList []*Repository `protobuf:"bytes,3,rep,name=repositoryList,proto3" json:"repositoryList,omitempty"`
	UserList       []*User       `protobuf:"bytes,2,rep,name=userList,proto3" json:"userList,omitempty"`
	WhoisList      []*Whois      `protobuf:"bytes,1,rep,name=whoisList,proto3" json:"whoisList,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe28ed7a80acf9ab, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCommentList() []*Comment {
	if m != nil {
		return m.CommentList
	}
	return nil
}

func (m *GenesisState) GetIssueList() []*Issue {
	if m != nil {
		return m.IssueList
	}
	return nil
}

func (m *GenesisState) GetRepositoryList() []*Repository {
	if m != nil {
		return m.RepositoryList
	}
	return nil
}

func (m *GenesisState) GetUserList() []*User {
	if m != nil {
		return m.UserList
	}
	return nil
}

func (m *GenesisState) GetWhoisList() []*Whois {
	if m != nil {
		return m.WhoisList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "gitopia.gitopia.gitopia.GenesisState")
}

func init() { proto.RegisterFile("gitopia/genesis.proto", fileDescriptor_fe28ed7a80acf9ab) }

var fileDescriptor_fe28ed7a80acf9ab = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xcf, 0x2c, 0xc9,
	0x2f, 0xc8, 0x4c, 0xd4, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x87, 0x0a, 0xeb, 0xa1, 0xd1, 0x52, 0x70, 0xf5, 0xc9, 0xf9, 0xb9, 0xb9, 0xa9,
	0x79, 0x25, 0x10, 0xf5, 0x52, 0xc2, 0x30, 0xe1, 0xcc, 0xe2, 0xe2, 0xd2, 0x54, 0xa8, 0xa0, 0x04,
	0x4c, 0xb0, 0x28, 0xb5, 0x20, 0xbf, 0x38, 0xb3, 0x24, 0xbf, 0xa8, 0x12, 0x2a, 0x23, 0x04, 0x93,
	0x29, 0x2d, 0x4e, 0x2d, 0x42, 0x37, 0xa2, 0x3c, 0x23, 0x1f, 0xe6, 0x0e, 0xa5, 0x6b, 0x4c, 0x5c,
	0x3c, 0xee, 0x10, 0x97, 0x05, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x39, 0x71, 0x71, 0x43, 0x6d, 0xf6,
	0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x55, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd0, 0xc3, 0xe1, 0x5c, 0x3d,
	0x67, 0x88, 0xda, 0x20, 0x64, 0x4d, 0x42, 0x36, 0x5c, 0x9c, 0x60, 0x67, 0x82, 0x4d, 0x60, 0x01,
	0x9b, 0x20, 0x87, 0xd3, 0x04, 0x4f, 0x90, 0xca, 0x20, 0x84, 0x06, 0x21, 0x6f, 0x2e, 0x3e, 0x84,
	0x7f, 0xc0, 0x46, 0x30, 0x83, 0x8d, 0x50, 0xc6, 0x69, 0x44, 0x10, 0x5c, 0x79, 0x10, 0x9a, 0x56,
	0x21, 0x4b, 0x2e, 0x0e, 0x50, 0x10, 0x80, 0x8d, 0x61, 0x02, 0x1b, 0x23, 0x8b, 0xd3, 0x98, 0xd0,
	0xe2, 0xd4, 0xa2, 0x20, 0xb8, 0x72, 0x90, 0x2f, 0xc0, 0x21, 0x05, 0xd6, 0xcb, 0x48, 0xc0, 0x17,
	0xe1, 0x20, 0x95, 0x41, 0x08, 0x0d, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x4f,
	0x1c, 0x50, 0xba, 0x02, 0xce, 0x2a, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xc7, 0x92, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x64, 0x01, 0x74, 0x46, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CommentList) > 0 {
		for iNdEx := len(m.CommentList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CommentList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.IssueList) > 0 {
		for iNdEx := len(m.IssueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IssueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RepositoryList) > 0 {
		for iNdEx := len(m.RepositoryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RepositoryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.UserList) > 0 {
		for iNdEx := len(m.UserList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.WhoisList) > 0 {
		for iNdEx := len(m.WhoisList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhoisList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WhoisList) > 0 {
		for _, e := range m.WhoisList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserList) > 0 {
		for _, e := range m.UserList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RepositoryList) > 0 {
		for _, e := range m.RepositoryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IssueList) > 0 {
		for _, e := range m.IssueList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CommentList) > 0 {
		for _, e := range m.CommentList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhoisList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhoisList = append(m.WhoisList, &Whois{})
			if err := m.WhoisList[len(m.WhoisList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserList = append(m.UserList, &User{})
			if err := m.UserList[len(m.UserList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoryList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RepositoryList = append(m.RepositoryList, &Repository{})
			if err := m.RepositoryList[len(m.RepositoryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssueList = append(m.IssueList, &Issue{})
			if err := m.IssueList[len(m.IssueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommentList = append(m.CommentList, &Comment{})
			if err := m.CommentList[len(m.CommentList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
