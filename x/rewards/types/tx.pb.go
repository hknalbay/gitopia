// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rewards/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCreateReward struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Recipient   string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	TotalAmount string `protobuf:"bytes,3,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
}

func (m *MsgCreateReward) Reset()         { *m = MsgCreateReward{} }
func (m *MsgCreateReward) String() string { return proto.CompactTextString(m) }
func (*MsgCreateReward) ProtoMessage()    {}
func (*MsgCreateReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0d543e6be8610a, []int{0}
}
func (m *MsgCreateReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateReward.Merge(m, src)
}
func (m *MsgCreateReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateReward proto.InternalMessageInfo

func (m *MsgCreateReward) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateReward) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *MsgCreateReward) GetTotalAmount() string {
	if m != nil {
		return m.TotalAmount
	}
	return ""
}

type MsgCreateRewardResponse struct {
}

func (m *MsgCreateRewardResponse) Reset()         { *m = MsgCreateRewardResponse{} }
func (m *MsgCreateRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRewardResponse) ProtoMessage()    {}
func (*MsgCreateRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0d543e6be8610a, []int{1}
}
func (m *MsgCreateRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRewardResponse.Merge(m, src)
}
func (m *MsgCreateRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRewardResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateReward)(nil), "gitopia.gitopia.rewards.MsgCreateReward")
	proto.RegisterType((*MsgCreateRewardResponse)(nil), "gitopia.gitopia.rewards.MsgCreateRewardResponse")
}

func init() { proto.RegisterFile("rewards/tx.proto", fileDescriptor_4c0d543e6be8610a) }

var fileDescriptor_4c0d543e6be8610a = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x2d, 0x4f,
	0x2c, 0x4a, 0x29, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0xcf,
	0x2c, 0xc9, 0x2f, 0xc8, 0x4c, 0xd4, 0x83, 0xd1, 0x50, 0x15, 0x52, 0xa2, 0x30, 0xa5, 0x50, 0x1a,
	0xa2, 0x5e, 0x29, 0x9b, 0x8b, 0xdf, 0xb7, 0x38, 0xdd, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0x08,
	0x2c, 0x23, 0x24, 0xc1, 0xc5, 0x9e, 0x0c, 0xe2, 0xe7, 0x17, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0xc1, 0xb8, 0x42, 0x32, 0x5c, 0x9c, 0x45, 0xa9, 0xc9, 0x99, 0x05, 0x99, 0xa9, 0x79, 0x25,
	0x12, 0x4c, 0x60, 0x39, 0x84, 0x80, 0x90, 0x02, 0x17, 0x77, 0x49, 0x7e, 0x49, 0x62, 0x8e, 0x63,
	0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x33, 0x58, 0x1e, 0x59, 0x48, 0x49, 0x92, 0x4b, 0x1c, 0xcd,
	0xb2, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3, 0x42, 0x2e, 0x66, 0xdf, 0xe2, 0x74,
	0xa1, 0x2c, 0x2e, 0x1e, 0x14, 0xb7, 0x68, 0xe8, 0xe1, 0xf0, 0x8f, 0x1e, 0x9a, 0x41, 0x52, 0x06,
	0xc4, 0xaa, 0x84, 0x59, 0xe9, 0xe4, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0x5a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0xd3, 0xe0,
	0x74, 0x85, 0x3e, 0x3c, 0xcc, 0x2b, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xe1, 0x68, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x61, 0xe5, 0x88, 0x66, 0x8b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateReward(ctx context.Context, in *MsgCreateReward, opts ...grpc.CallOption) (*MsgCreateRewardResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateReward(ctx context.Context, in *MsgCreateReward, opts ...grpc.CallOption) (*MsgCreateRewardResponse, error) {
	out := new(MsgCreateRewardResponse)
	err := c.cc.Invoke(ctx, "/gitopia.gitopia.rewards.Msg/CreateReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateReward(context.Context, *MsgCreateReward) (*MsgCreateRewardResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateReward(ctx context.Context, req *MsgCreateReward) (*MsgCreateRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReward not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitopia.gitopia.rewards.Msg/CreateReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateReward(ctx, req.(*MsgCreateReward))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitopia.gitopia.rewards.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReward",
			Handler:    _Msg_CreateReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rewards/tx.proto",
}

func (m *MsgCreateReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalAmount) > 0 {
		i -= len(m.TotalAmount)
		copy(dAtA[i:], m.TotalAmount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TotalAmount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TotalAmount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateRewardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
