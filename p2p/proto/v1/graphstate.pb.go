// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: graphstate.proto

package qitmeer_p2p_v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GraphState struct {
	Total                uint32   `protobuf:"varint,100,opt,name=total,proto3" json:"total,omitempty"`
	Layer                uint32   `protobuf:"varint,101,opt,name=layer,proto3" json:"layer,omitempty"`
	MainHeight           uint32   `protobuf:"varint,102,opt,name=mainHeight,proto3" json:"mainHeight,omitempty"`
	MainOrder            uint32   `protobuf:"varint,103,opt,name=mainOrder,proto3" json:"mainOrder,omitempty"`
	Tips                 []*Hash  `protobuf:"bytes,104,rep,name=tips,proto3" json:"tips,omitempty" ssz-max:"100"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphState) Reset()         { *m = GraphState{} }
func (m *GraphState) String() string { return proto.CompactTextString(m) }
func (*GraphState) ProtoMessage()    {}
func (*GraphState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b0f1dfc60bedea3, []int{0}
}
func (m *GraphState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GraphState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GraphState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GraphState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphState.Merge(m, src)
}
func (m *GraphState) XXX_Size() int {
	return m.Size()
}
func (m *GraphState) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphState.DiscardUnknown(m)
}

var xxx_messageInfo_GraphState proto.InternalMessageInfo

func (m *GraphState) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GraphState) GetLayer() uint32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *GraphState) GetMainHeight() uint32 {
	if m != nil {
		return m.MainHeight
	}
	return 0
}

func (m *GraphState) GetMainOrder() uint32 {
	if m != nil {
		return m.MainOrder
	}
	return 0
}

func (m *GraphState) GetTips() []*Hash {
	if m != nil {
		return m.Tips
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphState)(nil), "qitmeer.p2p.v1.GraphState")
}

func init() { proto.RegisterFile("graphstate.proto", fileDescriptor_6b0f1dfc60bedea3) }

var fileDescriptor_6b0f1dfc60bedea3 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2f, 0x4a, 0x2c,
	0xc8, 0x28, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0xcc,
	0x2c, 0xc9, 0x4d, 0x4d, 0x2d, 0xd2, 0x2b, 0x30, 0x2a, 0xd0, 0x2b, 0x33, 0x94, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b,
	0x4b, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x5d, 0x8a, 0x37, 0x37, 0xb5, 0xb8,
	0x38, 0x31, 0x1d, 0x6a, 0x9a, 0xd2, 0x66, 0x46, 0x2e, 0x2e, 0x77, 0x90, 0x15, 0xc1, 0x20, 0x2b,
	0x84, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0x4b, 0x12, 0x73, 0x24, 0x52, 0x14, 0x18, 0x35, 0x78, 0x83,
	0x20, 0x1c, 0x90, 0x68, 0x4e, 0x62, 0x65, 0x6a, 0x91, 0x44, 0x2a, 0x44, 0x14, 0xcc, 0x11, 0x92,
	0xe3, 0xe2, 0xca, 0x4d, 0xcc, 0xcc, 0xf3, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x91, 0x48, 0x03, 0x4b,
	0x21, 0x89, 0x08, 0xc9, 0x70, 0x71, 0x82, 0x78, 0xfe, 0x45, 0x29, 0xa9, 0x45, 0x12, 0xe9, 0x60,
	0x69, 0x84, 0x80, 0x90, 0x35, 0x17, 0x4b, 0x49, 0x66, 0x41, 0xb1, 0x44, 0x86, 0x02, 0xb3, 0x06,
	0xb7, 0x91, 0x88, 0x1e, 0xaa, 0xaf, 0xf4, 0x3c, 0x12, 0x8b, 0x33, 0x9c, 0x04, 0x3f, 0xdd, 0x93,
	0xe7, 0x2d, 0x2e, 0xae, 0xd2, 0xcd, 0x4d, 0xac, 0xb0, 0x52, 0x32, 0x34, 0x30, 0x50, 0x0a, 0x02,
	0x6b, 0x72, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x7b, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6d,
	0x47, 0xb5, 0xbc, 0x30, 0x01, 0x00, 0x00,
}

func (m *GraphState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GraphState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GraphState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tips) > 0 {
		for iNdEx := len(m.Tips) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tips[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGraphstate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0xc2
		}
	}
	if m.MainOrder != 0 {
		i = encodeVarintGraphstate(dAtA, i, uint64(m.MainOrder))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xb8
	}
	if m.MainHeight != 0 {
		i = encodeVarintGraphstate(dAtA, i, uint64(m.MainHeight))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xb0
	}
	if m.Layer != 0 {
		i = encodeVarintGraphstate(dAtA, i, uint64(m.Layer))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa8
	}
	if m.Total != 0 {
		i = encodeVarintGraphstate(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa0
	}
	return len(dAtA) - i, nil
}

func encodeVarintGraphstate(dAtA []byte, offset int, v uint64) int {
	offset -= sovGraphstate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GraphState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 2 + sovGraphstate(uint64(m.Total))
	}
	if m.Layer != 0 {
		n += 2 + sovGraphstate(uint64(m.Layer))
	}
	if m.MainHeight != 0 {
		n += 2 + sovGraphstate(uint64(m.MainHeight))
	}
	if m.MainOrder != 0 {
		n += 2 + sovGraphstate(uint64(m.MainOrder))
	}
	if len(m.Tips) > 0 {
		for _, e := range m.Tips {
			l = e.Size()
			n += 2 + l + sovGraphstate(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGraphstate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGraphstate(x uint64) (n int) {
	return sovGraphstate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GraphState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGraphstate
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
			return fmt.Errorf("proto: GraphState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GraphState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 100:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraphstate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 101:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Layer", wireType)
			}
			m.Layer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraphstate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Layer |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 102:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainHeight", wireType)
			}
			m.MainHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraphstate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 103:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainOrder", wireType)
			}
			m.MainOrder = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraphstate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainOrder |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 104:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tips", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraphstate
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
				return ErrInvalidLengthGraphstate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGraphstate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tips = append(m.Tips, &Hash{})
			if err := m.Tips[len(m.Tips)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGraphstate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGraphstate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGraphstate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGraphstate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGraphstate
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
					return 0, ErrIntOverflowGraphstate
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
					return 0, ErrIntOverflowGraphstate
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
				return 0, ErrInvalidLengthGraphstate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGraphstate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGraphstate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGraphstate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGraphstate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGraphstate = fmt.Errorf("proto: unexpected end of group")
)