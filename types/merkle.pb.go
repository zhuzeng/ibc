// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merkle.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	merkle.proto
	messages.proto
	queue.proto

It has these top-level messages:
	Op
	Data
	Branch
	MerkleProof
	IBCPacket
	QueueName
	StateKey
	StateValue
	MessageKey
	SendValue
	ReceiptValue
*/
package types

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// HashOp is the hashing algorithm we use at each level
type HashOp int32

const (
	HashOp_RIPEMD160 HashOp = 0
	HashOp_SHA224    HashOp = 1
	HashOp_SHA256    HashOp = 2
	HashOp_SHA384    HashOp = 3
	HashOp_SHA512    HashOp = 4
	HashOp_SHA3_224  HashOp = 5
	HashOp_SHA3_256  HashOp = 6
	HashOp_SHA3_384  HashOp = 7
	HashOp_SHA3_512  HashOp = 8
	HashOp_SHA256_X2 HashOp = 9
)

var HashOp_name = map[int32]string{
	0: "RIPEMD160",
	1: "SHA224",
	2: "SHA256",
	3: "SHA384",
	4: "SHA512",
	5: "SHA3_224",
	6: "SHA3_256",
	7: "SHA3_384",
	8: "SHA3_512",
	9: "SHA256_X2",
}
var HashOp_value = map[string]int32{
	"RIPEMD160": 0,
	"SHA224":    1,
	"SHA256":    2,
	"SHA384":    3,
	"SHA512":    4,
	"SHA3_224":  5,
	"SHA3_256":  6,
	"SHA3_384":  7,
	"SHA3_512":  8,
	"SHA256_X2": 9,
}

func (x HashOp) String() string {
	return proto.EnumName(HashOp_name, int32(x))
}
func (HashOp) EnumDescriptor() ([]byte, []int) { return fileDescriptorMerkle, []int{0} }

// If it is KeyValue, this is the data we want
// If it is SubTree, key is name of the tree,
//   value is root hash
type Data_DataType int32

const (
	Data_KeyValue Data_DataType = 0
	Data_SubTree  Data_DataType = 1
)

var Data_DataType_name = map[int32]string{
	0: "KeyValue",
	1: "SubTree",
}
var Data_DataType_value = map[string]int32{
	"KeyValue": 0,
	"SubTree":  1,
}

func (x Data_DataType) String() string {
	return proto.EnumName(Data_DataType_name, int32(x))
}
func (Data_DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMerkle, []int{1, 0} }

// Op represents one hash in a chain of hashes.
// An operation takes the output of the last level and returns
// a hash for the next level:
// Op(last) => Operation(prefix + last + sufix)
//
// A simple left/right hash would simply set prefix=left or
// suffix=right and leave the other blank. However, one could
// also represent the a Patricia trie proof by setting
// prefix to the rlp encoding of all nodes before the branch
// we select, and suffix to all those after the one we select.
type Op struct {
	Prefix []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix []byte `protobuf:"bytes,2,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Op     HashOp `protobuf:"varint,3,opt,name=op,proto3,enum=types.HashOp" json:"op,omitempty"`
}

func (m *Op) Reset()                    { *m = Op{} }
func (m *Op) String() string            { return proto.CompactTextString(m) }
func (*Op) ProtoMessage()               {}
func (*Op) Descriptor() ([]byte, []int) { return fileDescriptorMerkle, []int{0} }

func (m *Op) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Op) GetSuffix() []byte {
	if m != nil {
		return m.Suffix
	}
	return nil
}

func (m *Op) GetOp() HashOp {
	if m != nil {
		return m.Op
	}
	return HashOp_RIPEMD160
}

// Data is the end value stored,
// used to generate the initial hash store
type Data struct {
	// optional prefix allows second preimage resistance
	Prefix   []byte        `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Key      []byte        `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value    []byte        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Op       HashOp        `protobuf:"varint,4,opt,name=op,proto3,enum=types.HashOp" json:"op,omitempty"`
	DataType Data_DataType `protobuf:"varint,5,opt,name=dataType,proto3,enum=types.Data_DataType" json:"dataType,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptorMerkle, []int{1} }

func (m *Data) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Data) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Data) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Data) GetOp() HashOp {
	if m != nil {
		return m.Op
	}
	return HashOp_RIPEMD160
}

func (m *Data) GetDataType() Data_DataType {
	if m != nil {
		return m.DataType
	}
	return Data_KeyValue
}

// Branch will hash data and then pass it through operations
// from first to last in order to calculate the root node.
//
// Visualize Branch as representing the data closest to
// root as the first item, and the leaf as the last item.
type Branch struct {
	// if either are non-empty, enforce this prefix on all
	// leaf/inner nodes to provide second preimage resistence
	PrefixLeaf  []byte `protobuf:"bytes,1,opt,name=prefixLeaf,proto3" json:"prefixLeaf,omitempty"`
	PrefixInner []byte `protobuf:"bytes,2,opt,name=prefixInner,proto3" json:"prefixInner,omitempty"`
	// this is the data to get the original hash,
	// and a set of operations to calculate the root hash
	Data       *Data `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	Operations []*Op `protobuf:"bytes,4,rep,name=operations" json:"operations,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptorMerkle, []int{2} }

func (m *Branch) GetPrefixLeaf() []byte {
	if m != nil {
		return m.PrefixLeaf
	}
	return nil
}

func (m *Branch) GetPrefixInner() []byte {
	if m != nil {
		return m.PrefixInner
	}
	return nil
}

func (m *Branch) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Branch) GetOperations() []*Op {
	if m != nil {
		return m.Operations
	}
	return nil
}

type MerkleProof struct {
	// branches start from the value, and then may
	// include multiple subtree branches to embed it
	//
	// The first branch must have dataType KeyValue
	// Following branches must have dataType SubTree
	Branches []*Branch `protobuf:"bytes,1,rep,name=branches" json:"branches,omitempty"`
}

func (m *MerkleProof) Reset()                    { *m = MerkleProof{} }
func (m *MerkleProof) String() string            { return proto.CompactTextString(m) }
func (*MerkleProof) ProtoMessage()               {}
func (*MerkleProof) Descriptor() ([]byte, []int) { return fileDescriptorMerkle, []int{3} }

func (m *MerkleProof) GetBranches() []*Branch {
	if m != nil {
		return m.Branches
	}
	return nil
}

func init() {
	proto.RegisterType((*Op)(nil), "types.Op")
	proto.RegisterType((*Data)(nil), "types.Data")
	proto.RegisterType((*Branch)(nil), "types.Branch")
	proto.RegisterType((*MerkleProof)(nil), "types.MerkleProof")
	proto.RegisterEnum("types.HashOp", HashOp_name, HashOp_value)
	proto.RegisterEnum("types.Data_DataType", Data_DataType_name, Data_DataType_value)
}

func init() { proto.RegisterFile("merkle.proto", fileDescriptorMerkle) }

var fileDescriptorMerkle = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xcd, 0xea, 0x2b, 0xf2, 0x48, 0x29, 0xcb, 0x12, 0x8a, 0x2e, 0x6d, 0x85, 0xa0, 0xe0, 0xf4,
	0x60, 0x12, 0x25, 0x36, 0xb9, 0xb6, 0xa4, 0x90, 0xd0, 0x06, 0x87, 0x75, 0x28, 0xbd, 0x99, 0x75,
	0xbb, 0xc2, 0xc6, 0xae, 0x76, 0x59, 0xc9, 0xa5, 0xfa, 0x13, 0x3d, 0xf6, 0xe7, 0xf4, 0xb7, 0x95,
	0xfd, 0xb0, 0xaa, 0x8b, 0x2f, 0x62, 0xde, 0x9b, 0x99, 0xf7, 0x66, 0x46, 0x0b, 0xe9, 0x0f, 0xae,
	0xb6, 0x3b, 0x3e, 0x91, 0x4a, 0xb4, 0x82, 0x84, 0x6d, 0x27, 0x79, 0x53, 0x2c, 0xc0, 0x9b, 0x4b,
	0xf2, 0x12, 0x22, 0xa9, 0x78, 0xb5, 0xf9, 0x95, 0xa1, 0x1c, 0x8d, 0x53, 0xea, 0x90, 0xe6, 0x9b,
	0x7d, 0xa5, 0x79, 0xcf, 0xf2, 0x16, 0x91, 0x57, 0xe0, 0x09, 0x99, 0xf9, 0x39, 0x1a, 0xbf, 0x28,
	0xcf, 0x26, 0x46, 0x69, 0x72, 0xcf, 0x9a, 0xf5, 0x5c, 0x52, 0x4f, 0xc8, 0xe2, 0x2f, 0x82, 0xe0,
	0x8e, 0xb5, 0xec, 0xa8, 0x2e, 0x06, 0x7f, 0xcb, 0x3b, 0x27, 0xaa, 0x43, 0x72, 0x0e, 0xe1, 0x4f,
	0xb6, 0xdb, 0x73, 0x23, 0x9a, 0x52, 0x0b, 0x9c, 0x4f, 0x70, 0xc4, 0x87, 0x5c, 0x42, 0xfc, 0x9d,
	0xb5, 0xec, 0xb9, 0x93, 0x3c, 0x0b, 0x4d, 0xd1, 0xb9, 0x2b, 0xd2, 0xee, 0xe6, 0xa3, 0x73, 0xb4,
	0xaf, 0x2a, 0xde, 0x42, 0x7c, 0x60, 0x49, 0x0a, 0xf1, 0x27, 0xde, 0x7d, 0xd1, 0x46, 0xf8, 0x84,
	0x24, 0x70, 0xba, 0xd8, 0xaf, 0x9e, 0x15, 0xe7, 0x18, 0x15, 0x7f, 0x10, 0x44, 0x1f, 0x14, 0xab,
	0xbf, 0xad, 0xc9, 0x6b, 0x00, 0x3b, 0xf4, 0x67, 0xce, 0x2a, 0xb7, 0xc6, 0x80, 0x21, 0x39, 0x24,
	0x16, 0x3d, 0xd4, 0x35, 0x57, 0x6e, 0xa5, 0x21, 0x45, 0xde, 0x40, 0xa0, 0xfd, 0xcd, 0x66, 0x49,
	0x99, 0x0c, 0x26, 0xa4, 0x26, 0x41, 0x2e, 0x00, 0x84, 0xe4, 0x8a, 0xb5, 0x1b, 0x51, 0x37, 0x59,
	0x90, 0xfb, 0xe3, 0xa4, 0x1c, 0xb9, 0xb2, 0xb9, 0xa4, 0x83, 0x64, 0x71, 0x0b, 0xc9, 0xa3, 0xf9,
	0x8b, 0x4f, 0x4a, 0x88, 0x8a, 0x5c, 0x40, 0xbc, 0x32, 0x63, 0xf2, 0x26, 0x43, 0xa6, 0xef, 0x70,
	0x25, 0x3b, 0x3d, 0xed, 0xd3, 0xef, 0x7e, 0x23, 0x88, 0xec, 0xe9, 0xc8, 0x19, 0x8c, 0xe8, 0xc3,
	0xd3, 0xc7, 0xc7, 0xbb, 0xab, 0xd9, 0x25, 0x3e, 0x21, 0x00, 0xd1, 0xe2, 0xfe, 0x7d, 0x59, 0xde,
	0x60, 0x74, 0x88, 0xa7, 0x33, 0xec, 0xb9, 0xf8, 0xfa, 0xf6, 0x06, 0xfb, 0x2e, 0x9e, 0x5e, 0x95,
	0x38, 0xd0, 0x77, 0xd3, 0xfc, 0x52, 0x77, 0x84, 0xff, 0xd1, 0x74, 0x86, 0xa3, 0x1e, 0xe9, 0xae,
	0xd3, 0x1e, 0xe9, 0xbe, 0x58, 0xdb, 0x5a, 0xed, 0xe5, 0xd7, 0x12, 0x8f, 0x56, 0x91, 0x79, 0x87,
	0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x52, 0xc9, 0xdd, 0x97, 0x02, 0x00, 0x00,
}
