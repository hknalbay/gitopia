// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/comment.proto

package v013

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Comment_Type int32

const (
	Comment_ISSUE       Comment_Type = 0
	Comment_PULLREQUEST Comment_Type = 1
)

var Comment_Type_name = map[int32]string{
	0: "ISSUE",
	1: "PULLREQUEST",
}

var Comment_Type_value = map[string]int32{
	"ISSUE":       0,
	"PULLREQUEST": 1,
}

func (x Comment_Type) String() string {
	return proto.EnumName(Comment_Type_name, int32(x))
}

func (Comment_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61a8a10ae7d09fb4, []int{0, 0}
}

type Comment struct {
	Creator           string       `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id                uint64       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	ParentId          uint64       `protobuf:"varint,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
	CommentIid        uint64       `protobuf:"varint,4,opt,name=commentIid,proto3" json:"commentIid,omitempty"`
	Body              string       `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Attachments       []string     `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
	DiffHunk          string       `protobuf:"bytes,7,opt,name=diffHunk,proto3" json:"diffHunk,omitempty"`
	Path              string       `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	System            bool         `protobuf:"varint,9,opt,name=system,proto3" json:"system,omitempty"`
	AuthorAssociation string       `protobuf:"bytes,10,opt,name=authorAssociation,proto3" json:"authorAssociation,omitempty"`
	CreatedAt         int64        `protobuf:"varint,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt         int64        `protobuf:"varint,12,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	CommentType       Comment_Type `protobuf:"varint,13,opt,name=commentType,proto3,enum=gitopia.gitopia.gitopia.Comment_Type" json:"commentType,omitempty"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_61a8a10ae7d09fb4, []int{0}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return m.Size()
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Comment) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Comment) GetCommentIid() uint64 {
	if m != nil {
		return m.CommentIid
	}
	return 0
}

func (m *Comment) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Comment) GetAttachments() []string {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Comment) GetDiffHunk() string {
	if m != nil {
		return m.DiffHunk
	}
	return ""
}

func (m *Comment) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Comment) GetSystem() bool {
	if m != nil {
		return m.System
	}
	return false
}

func (m *Comment) GetAuthorAssociation() string {
	if m != nil {
		return m.AuthorAssociation
	}
	return ""
}

func (m *Comment) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Comment) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Comment) GetCommentType() Comment_Type {
	if m != nil {
		return m.CommentType
	}
	return Comment_ISSUE
}

func init() {
	// proto.RegisterEnum("gitopia.gitopia.gitopia.Comment_Type", Comment_Type_name, Comment_Type_value)
	// proto.RegisterType((*Comment)(nil), "gitopia.gitopia.gitopia.Comment")
}

func init() { proto.RegisterFile("gitopia/comment.proto", fileDescriptor_61a8a10ae7d09fb4) }

var fileDescriptor_61a8a10ae7d09fb4 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xdd, 0xea, 0xd3, 0x30,
	0x14, 0x6f, 0xd6, 0xfe, 0xb7, 0xf5, 0x54, 0xe7, 0x0c, 0x7e, 0x84, 0x21, 0xa5, 0x0c, 0x84, 0x22,
	0xd2, 0x81, 0x3e, 0xc1, 0xd4, 0xa1, 0x83, 0x5d, 0x68, 0xb7, 0xdd, 0x78, 0x97, 0x35, 0x5d, 0x1b,
	0xa4, 0x4d, 0x69, 0x53, 0xb0, 0x6f, 0xe1, 0x43, 0x79, 0xe1, 0xe5, 0x2e, 0xbd, 0x94, 0xed, 0x45,
	0x24, 0x69, 0x57, 0x87, 0xe2, 0xd5, 0xf9, 0x7d, 0x9d, 0x93, 0x36, 0x39, 0xf0, 0x38, 0xe1, 0x52,
	0x14, 0x9c, 0x2e, 0x22, 0x91, 0x65, 0x71, 0x2e, 0x83, 0xa2, 0x14, 0x52, 0xe0, 0xa7, 0x9d, 0x1c,
	0xfc, 0x55, 0x67, 0x8f, 0x12, 0x91, 0x08, 0x9d, 0x59, 0x28, 0xd4, 0xc6, 0xe7, 0xdf, 0x4d, 0x18,
	0xbd, 0x6d, 0x07, 0x60, 0x02, 0xa3, 0xa8, 0x8c, 0xa9, 0x14, 0x25, 0x41, 0x1e, 0xf2, 0xed, 0xf0,
	0x4a, 0xf1, 0x04, 0x06, 0x9c, 0x91, 0x81, 0x87, 0x7c, 0x2b, 0x1c, 0x70, 0x86, 0x67, 0x30, 0x2e,
	0x68, 0x19, 0xe7, 0x72, 0xcd, 0x88, 0xa9, 0xd5, 0x9e, 0x63, 0x17, 0xa0, 0xfb, 0xa2, 0x35, 0x67,
	0xc4, 0xd2, 0xee, 0x8d, 0x82, 0x31, 0x58, 0x07, 0xc1, 0x1a, 0x72, 0xa7, 0x8f, 0xd0, 0x18, 0x7b,
	0xe0, 0x50, 0x29, 0x69, 0x94, 0xaa, 0x50, 0x45, 0x86, 0x9e, 0xe9, 0xdb, 0xe1, 0xad, 0xa4, 0x4e,
	0x64, 0xfc, 0x78, 0xfc, 0x50, 0xe7, 0x5f, 0xc8, 0x48, 0x77, 0xf6, 0x5c, 0x4d, 0x2c, 0xa8, 0x4c,
	0xc9, 0xb8, 0x9d, 0xa8, 0x30, 0x7e, 0x02, 0xc3, 0xaa, 0xa9, 0x64, 0x9c, 0x11, 0xdb, 0x43, 0xfe,
	0x38, 0xec, 0x18, 0x7e, 0x09, 0x0f, 0x69, 0x2d, 0x53, 0x51, 0x2e, 0xab, 0x4a, 0x44, 0x9c, 0x4a,
	0x2e, 0x72, 0x02, 0xba, 0xf1, 0x5f, 0x03, 0x3f, 0x03, 0x5b, 0x5f, 0x41, 0xcc, 0x96, 0x92, 0x38,
	0x1e, 0xf2, 0xcd, 0xf0, 0x8f, 0xa0, 0xdc, 0xba, 0x60, 0x9d, 0x7b, 0xaf, 0x75, 0x7b, 0x01, 0xbf,
	0x07, 0xa7, 0xfb, 0xeb, 0x5d, 0x53, 0xc4, 0xe4, 0xbe, 0x87, 0xfc, 0xc9, 0xab, 0xe7, 0xc1, 0x7f,
	0x9e, 0x27, 0xe8, 0x1e, 0x21, 0x50, 0xe1, 0xf0, 0xb6, 0x73, 0x3e, 0x07, 0x4b, 0x55, 0x6c, 0xc3,
	0xdd, 0x7a, 0xbb, 0xdd, 0xaf, 0xa6, 0x06, 0x7e, 0x00, 0xce, 0xc7, 0xfd, 0x66, 0x13, 0xae, 0x3e,
	0xed, 0x57, 0xdb, 0xdd, 0x14, 0xbd, 0x79, 0xf7, 0xe3, 0xec, 0xa2, 0xd3, 0xd9, 0x45, 0xbf, 0xce,
	0x2e, 0xfa, 0x76, 0x71, 0x8d, 0xd3, 0xc5, 0x35, 0x7e, 0x5e, 0x5c, 0xe3, 0xf3, 0x8b, 0x84, 0xcb,
	0xb4, 0x3e, 0x04, 0x91, 0xc8, 0x16, 0xd7, 0x8d, 0xb9, 0xd6, 0xaf, 0x3d, 0x92, 0x4d, 0x11, 0x57,
	0x87, 0xa1, 0xde, 0x89, 0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x07, 0x17, 0x6f, 0xbd, 0x5b,
	0x02, 0x00, 0x00,
}

func (m *Comment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Comment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Comment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommentType != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.CommentType))
		i--
		dAtA[i] = 0x68
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x60
	}
	if m.CreatedAt != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x58
	}
	if len(m.AuthorAssociation) > 0 {
		i -= len(m.AuthorAssociation)
		copy(dAtA[i:], m.AuthorAssociation)
		i = encodeVarintComment(dAtA, i, uint64(len(m.AuthorAssociation)))
		i--
		dAtA[i] = 0x52
	}
	if m.System {
		i--
		if m.System {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintComment(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DiffHunk) > 0 {
		i -= len(m.DiffHunk)
		copy(dAtA[i:], m.DiffHunk)
		i = encodeVarintComment(dAtA, i, uint64(len(m.DiffHunk)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Attachments) > 0 {
		for iNdEx := len(m.Attachments) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Attachments[iNdEx])
			copy(dAtA[i:], m.Attachments[iNdEx])
			i = encodeVarintComment(dAtA, i, uint64(len(m.Attachments[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintComment(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CommentIid != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.CommentIid))
		i--
		dAtA[i] = 0x20
	}
	if m.ParentId != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.ParentId))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintComment(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintComment(dAtA []byte, offset int, v uint64) int {
	offset -= sovComment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Comment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovComment(uint64(m.Id))
	}
	if m.ParentId != 0 {
		n += 1 + sovComment(uint64(m.ParentId))
	}
	if m.CommentIid != 0 {
		n += 1 + sovComment(uint64(m.CommentIid))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if len(m.Attachments) > 0 {
		for _, s := range m.Attachments {
			l = len(s)
			n += 1 + l + sovComment(uint64(l))
		}
	}
	l = len(m.DiffHunk)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.System {
		n += 2
	}
	l = len(m.AuthorAssociation)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovComment(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovComment(uint64(m.UpdatedAt))
	}
	if m.CommentType != 0 {
		n += 1 + sovComment(uint64(m.CommentType))
	}
	return n
}

func sovComment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComment(x uint64) (n int) {
	return sovComment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Comment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
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
			return fmt.Errorf("proto: Comment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Comment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentId", wireType)
			}
			m.ParentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParentId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentIid", wireType)
			}
			m.CommentIid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommentIid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attachments", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attachments = append(m.Attachments, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiffHunk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DiffHunk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field System", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.System = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorAssociation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorAssociation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentType", wireType)
			}
			m.CommentType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommentType |= Comment_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
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
func skipComment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComment
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
					return 0, ErrIntOverflowComment
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
					return 0, ErrIntOverflowComment
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
				return 0, ErrInvalidLengthComment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComment = fmt.Errorf("proto: unexpected end of group")
)
