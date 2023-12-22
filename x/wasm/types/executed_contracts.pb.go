// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya/wasm/v1/executed_contracts.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// ExecutedContracts is a structure type that contains the list of executed contracts
// in a specific transaction.
type ExecutedContracts struct {
	ContractAddresses []string `protobuf:"bytes,1,rep,name=contract_addresses,json=contractAddresses,proto3" json:"contract_addresses,omitempty"`
}

func (m *ExecutedContracts) Reset()         { *m = ExecutedContracts{} }
func (m *ExecutedContracts) String() string { return proto.CompactTextString(m) }
func (*ExecutedContracts) ProtoMessage()    {}
func (*ExecutedContracts) Descriptor() ([]byte, []int) {
	return fileDescriptor_200f5f891d4ef8d1, []int{0}
}
func (m *ExecutedContracts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecutedContracts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecutedContracts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecutedContracts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutedContracts.Merge(m, src)
}
func (m *ExecutedContracts) XXX_Size() int {
	return m.Size()
}
func (m *ExecutedContracts) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutedContracts.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutedContracts proto.InternalMessageInfo

func (m *ExecutedContracts) GetContractAddresses() []string {
	if m != nil {
		return m.ContractAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecutedContracts)(nil), "furya.wasm.v1.ExecutedContracts")
}

func init() {
	proto.RegisterFile("furya/wasm/v1/executed_contracts.proto", fileDescriptor_200f5f891d4ef8d1)
}

var fileDescriptor_200f5f891d4ef8d1 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x49, 0x2d, 0x2a,
	0x4a, 0xd4, 0x2f, 0x4f, 0x2c, 0xce, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0xad, 0x48, 0x4d, 0x2e, 0x2d,
	0x49, 0x4d, 0x89, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x05, 0xab, 0xd3, 0x03, 0xa9, 0xd3, 0x2b, 0x33, 0x54, 0x72, 0xe2, 0x12,
	0x74, 0x85, 0x2a, 0x75, 0x86, 0xa9, 0x14, 0xd2, 0xe5, 0x12, 0x82, 0x69, 0x8b, 0x4f, 0x4c, 0x49,
	0x29, 0x4a, 0x2d, 0x2e, 0x4e, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x12, 0x84, 0xc9,
	0x38, 0xc2, 0x24, 0x9c, 0x5c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a,
	0x3b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x6c, 0xaf, 0x6e, 0x6e,
	0x7e, 0x5e, 0x6a, 0xa5, 0x7e, 0x72, 0x7e, 0x51, 0xaa, 0x7e, 0x99, 0x91, 0x7e, 0x05, 0xc4, 0xbd,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x07, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x66, 0x6b, 0xb9, 0x17, 0xca, 0x00, 0x00, 0x00,
}

func (m *ExecutedContracts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecutedContracts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExecutedContracts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddresses) > 0 {
		for iNdEx := len(m.ContractAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ContractAddresses[iNdEx])
			copy(dAtA[i:], m.ContractAddresses[iNdEx])
			i = encodeVarintExecutedContracts(dAtA, i, uint64(len(m.ContractAddresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintExecutedContracts(dAtA []byte, offset int, v uint64) int {
	offset -= sovExecutedContracts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExecutedContracts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ContractAddresses) > 0 {
		for _, s := range m.ContractAddresses {
			l = len(s)
			n += 1 + l + sovExecutedContracts(uint64(l))
		}
	}
	return n
}

func sovExecutedContracts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExecutedContracts(x uint64) (n int) {
	return sovExecutedContracts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExecutedContracts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutedContracts
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
			return fmt.Errorf("proto: ExecutedContracts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecutedContracts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutedContracts
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
				return ErrInvalidLengthExecutedContracts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecutedContracts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddresses = append(m.ContractAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecutedContracts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecutedContracts
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
func skipExecutedContracts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExecutedContracts
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
					return 0, ErrIntOverflowExecutedContracts
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
					return 0, ErrIntOverflowExecutedContracts
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
				return 0, ErrInvalidLengthExecutedContracts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExecutedContracts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExecutedContracts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExecutedContracts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExecutedContracts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExecutedContracts = fmt.Errorf("proto: unexpected end of group")
)
