// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: io/casperlabs/casper/consensus/state/state.proto

package state

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// NOTE: Numeric values correspond to values of the domain
// AccessRights struct. DO NOT CHANGE.
type Key_URef_AccessRights int32

const (
	Key_URef_UNKNOWN        Key_URef_AccessRights = 0
	Key_URef_READ           Key_URef_AccessRights = 1
	Key_URef_WRITE          Key_URef_AccessRights = 2
	Key_URef_ADD            Key_URef_AccessRights = 4
	Key_URef_READ_ADD       Key_URef_AccessRights = 5
	Key_URef_READ_WRITE     Key_URef_AccessRights = 3
	Key_URef_ADD_WRITE      Key_URef_AccessRights = 6
	Key_URef_READ_ADD_WRITE Key_URef_AccessRights = 7
)

var Key_URef_AccessRights_name = map[int32]string{
	0: "UNKNOWN",
	1: "READ",
	2: "WRITE",
	4: "ADD",
	5: "READ_ADD",
	3: "READ_WRITE",
	6: "ADD_WRITE",
	7: "READ_ADD_WRITE",
}

var Key_URef_AccessRights_value = map[string]int32{
	"UNKNOWN":        0,
	"READ":           1,
	"WRITE":          2,
	"ADD":            4,
	"READ_ADD":       5,
	"READ_WRITE":     3,
	"ADD_WRITE":      6,
	"READ_ADD_WRITE": 7,
}

func (x Key_URef_AccessRights) String() string {
	return proto.EnumName(Key_URef_AccessRights_name, int32(x))
}

func (Key_URef_AccessRights) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{4, 2, 0}
}

// Value stored under a key in global state.
type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_IntValue
	//	*Value_BytesValue
	//	*Value_IntList
	//	*Value_StringValue
	//	*Value_Account
	//	*Value_Contract
	//	*Value_StringList
	//	*Value_NamedKey
	//	*Value_BigInt
	//	*Value_Key
	//	*Value_Unit
	//	*Value_LongValue
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{0}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_IntValue struct {
	IntValue int32 `protobuf:"varint,1,opt,name=int_value,json=intValue,proto3,oneof"`
}
type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,2,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}
type Value_IntList struct {
	IntList *IntList `protobuf:"bytes,3,opt,name=int_list,json=intList,proto3,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}
type Value_Account struct {
	Account *Account `protobuf:"bytes,5,opt,name=account,proto3,oneof"`
}
type Value_Contract struct {
	Contract *Contract `protobuf:"bytes,6,opt,name=contract,proto3,oneof"`
}
type Value_StringList struct {
	StringList *StringList `protobuf:"bytes,7,opt,name=string_list,json=stringList,proto3,oneof"`
}
type Value_NamedKey struct {
	NamedKey *NamedKey `protobuf:"bytes,8,opt,name=named_key,json=namedKey,proto3,oneof"`
}
type Value_BigInt struct {
	BigInt *BigInt `protobuf:"bytes,9,opt,name=big_int,json=bigInt,proto3,oneof"`
}
type Value_Key struct {
	Key *Key `protobuf:"bytes,10,opt,name=key,proto3,oneof"`
}
type Value_Unit struct {
	Unit *Unit `protobuf:"bytes,11,opt,name=unit,proto3,oneof"`
}
type Value_LongValue struct {
	LongValue uint64 `protobuf:"varint,12,opt,name=long_value,json=longValue,proto3,oneof"`
}

func (*Value_IntValue) isValue_Value()    {}
func (*Value_BytesValue) isValue_Value()  {}
func (*Value_IntList) isValue_Value()     {}
func (*Value_StringValue) isValue_Value() {}
func (*Value_Account) isValue_Value()     {}
func (*Value_Contract) isValue_Value()    {}
func (*Value_StringList) isValue_Value()  {}
func (*Value_NamedKey) isValue_Value()    {}
func (*Value_BigInt) isValue_Value()      {}
func (*Value_Key) isValue_Value()         {}
func (*Value_Unit) isValue_Value()        {}
func (*Value_LongValue) isValue_Value()   {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetIntValue() int32 {
	if x, ok := m.GetValue().(*Value_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *Value) GetBytesValue() []byte {
	if x, ok := m.GetValue().(*Value_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Value) GetIntList() *IntList {
	if x, ok := m.GetValue().(*Value_IntList); ok {
		return x.IntList
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValue().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetAccount() *Account {
	if x, ok := m.GetValue().(*Value_Account); ok {
		return x.Account
	}
	return nil
}

func (m *Value) GetContract() *Contract {
	if x, ok := m.GetValue().(*Value_Contract); ok {
		return x.Contract
	}
	return nil
}

func (m *Value) GetStringList() *StringList {
	if x, ok := m.GetValue().(*Value_StringList); ok {
		return x.StringList
	}
	return nil
}

func (m *Value) GetNamedKey() *NamedKey {
	if x, ok := m.GetValue().(*Value_NamedKey); ok {
		return x.NamedKey
	}
	return nil
}

func (m *Value) GetBigInt() *BigInt {
	if x, ok := m.GetValue().(*Value_BigInt); ok {
		return x.BigInt
	}
	return nil
}

func (m *Value) GetKey() *Key {
	if x, ok := m.GetValue().(*Value_Key); ok {
		return x.Key
	}
	return nil
}

func (m *Value) GetUnit() *Unit {
	if x, ok := m.GetValue().(*Value_Unit); ok {
		return x.Unit
	}
	return nil
}

func (m *Value) GetLongValue() uint64 {
	if x, ok := m.GetValue().(*Value_LongValue); ok {
		return x.LongValue
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_IntValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_IntList)(nil),
		(*Value_StringValue)(nil),
		(*Value_Account)(nil),
		(*Value_Contract)(nil),
		(*Value_StringList)(nil),
		(*Value_NamedKey)(nil),
		(*Value_BigInt)(nil),
		(*Value_Key)(nil),
		(*Value_Unit)(nil),
		(*Value_LongValue)(nil),
	}
}

type IntList struct {
	Values               []int32  `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntList) Reset()         { *m = IntList{} }
func (m *IntList) String() string { return proto.CompactTextString(m) }
func (*IntList) ProtoMessage()    {}
func (*IntList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{1}
}
func (m *IntList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntList.Unmarshal(m, b)
}
func (m *IntList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntList.Marshal(b, m, deterministic)
}
func (m *IntList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntList.Merge(m, src)
}
func (m *IntList) XXX_Size() int {
	return xxx_messageInfo_IntList.Size(m)
}
func (m *IntList) XXX_DiscardUnknown() {
	xxx_messageInfo_IntList.DiscardUnknown(m)
}

var xxx_messageInfo_IntList proto.InternalMessageInfo

func (m *IntList) GetValues() []int32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type StringList struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringList) Reset()         { *m = StringList{} }
func (m *StringList) String() string { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()    {}
func (*StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{2}
}
func (m *StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringList.Unmarshal(m, b)
}
func (m *StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringList.Marshal(b, m, deterministic)
}
func (m *StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringList.Merge(m, src)
}
func (m *StringList) XXX_Size() int {
	return xxx_messageInfo_StringList.Size(m)
}
func (m *StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_StringList proto.InternalMessageInfo

func (m *StringList) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type BigInt struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Number of bits: 128 | 256 | 512.
	BitWidth             uint32   `protobuf:"varint,2,opt,name=bit_width,json=bitWidth,proto3" json:"bit_width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigInt) Reset()         { *m = BigInt{} }
func (m *BigInt) String() string { return proto.CompactTextString(m) }
func (*BigInt) ProtoMessage()    {}
func (*BigInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{3}
}
func (m *BigInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigInt.Unmarshal(m, b)
}
func (m *BigInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigInt.Marshal(b, m, deterministic)
}
func (m *BigInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigInt.Merge(m, src)
}
func (m *BigInt) XXX_Size() int {
	return xxx_messageInfo_BigInt.Size(m)
}
func (m *BigInt) XXX_DiscardUnknown() {
	xxx_messageInfo_BigInt.DiscardUnknown(m)
}

var xxx_messageInfo_BigInt proto.InternalMessageInfo

func (m *BigInt) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *BigInt) GetBitWidth() uint32 {
	if m != nil {
		return m.BitWidth
	}
	return 0
}

type Key struct {
	// Types that are valid to be assigned to Value:
	//	*Key_Address_
	//	*Key_Hash_
	//	*Key_Uref
	//	*Key_Local_
	Value                isKey_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{4}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

type isKey_Value interface {
	isKey_Value()
}

type Key_Address_ struct {
	Address *Key_Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}
type Key_Hash_ struct {
	Hash *Key_Hash `protobuf:"bytes,2,opt,name=hash,proto3,oneof"`
}
type Key_Uref struct {
	Uref *Key_URef `protobuf:"bytes,3,opt,name=uref,proto3,oneof"`
}
type Key_Local_ struct {
	Local *Key_Local `protobuf:"bytes,4,opt,name=local,proto3,oneof"`
}

func (*Key_Address_) isKey_Value() {}
func (*Key_Hash_) isKey_Value()    {}
func (*Key_Uref) isKey_Value()     {}
func (*Key_Local_) isKey_Value()   {}

func (m *Key) GetValue() isKey_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Key) GetAddress() *Key_Address {
	if x, ok := m.GetValue().(*Key_Address_); ok {
		return x.Address
	}
	return nil
}

func (m *Key) GetHash() *Key_Hash {
	if x, ok := m.GetValue().(*Key_Hash_); ok {
		return x.Hash
	}
	return nil
}

func (m *Key) GetUref() *Key_URef {
	if x, ok := m.GetValue().(*Key_Uref); ok {
		return x.Uref
	}
	return nil
}

func (m *Key) GetLocal() *Key_Local {
	if x, ok := m.GetValue().(*Key_Local_); ok {
		return x.Local
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Key) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Key_Address_)(nil),
		(*Key_Hash_)(nil),
		(*Key_Uref)(nil),
		(*Key_Local_)(nil),
	}
}

type Key_Address struct {
	Account              []byte   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key_Address) Reset()         { *m = Key_Address{} }
func (m *Key_Address) String() string { return proto.CompactTextString(m) }
func (*Key_Address) ProtoMessage()    {}
func (*Key_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{4, 0}
}
func (m *Key_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key_Address.Unmarshal(m, b)
}
func (m *Key_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key_Address.Marshal(b, m, deterministic)
}
func (m *Key_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key_Address.Merge(m, src)
}
func (m *Key_Address) XXX_Size() int {
	return xxx_messageInfo_Key_Address.Size(m)
}
func (m *Key_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Key_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Key_Address proto.InternalMessageInfo

func (m *Key_Address) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

type Key_Hash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key_Hash) Reset()         { *m = Key_Hash{} }
func (m *Key_Hash) String() string { return proto.CompactTextString(m) }
func (*Key_Hash) ProtoMessage()    {}
func (*Key_Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{4, 1}
}
func (m *Key_Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key_Hash.Unmarshal(m, b)
}
func (m *Key_Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key_Hash.Marshal(b, m, deterministic)
}
func (m *Key_Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key_Hash.Merge(m, src)
}
func (m *Key_Hash) XXX_Size() int {
	return xxx_messageInfo_Key_Hash.Size(m)
}
func (m *Key_Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Key_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Key_Hash proto.InternalMessageInfo

func (m *Key_Hash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Key_URef struct {
	Uref                 []byte                `protobuf:"bytes,1,opt,name=uref,proto3" json:"uref,omitempty"`
	AccessRights         Key_URef_AccessRights `protobuf:"varint,2,opt,name=access_rights,json=accessRights,proto3,enum=io.casperlabs.casper.consensus.state.Key_URef_AccessRights" json:"access_rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Key_URef) Reset()         { *m = Key_URef{} }
func (m *Key_URef) String() string { return proto.CompactTextString(m) }
func (*Key_URef) ProtoMessage()    {}
func (*Key_URef) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{4, 2}
}
func (m *Key_URef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key_URef.Unmarshal(m, b)
}
func (m *Key_URef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key_URef.Marshal(b, m, deterministic)
}
func (m *Key_URef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key_URef.Merge(m, src)
}
func (m *Key_URef) XXX_Size() int {
	return xxx_messageInfo_Key_URef.Size(m)
}
func (m *Key_URef) XXX_DiscardUnknown() {
	xxx_messageInfo_Key_URef.DiscardUnknown(m)
}

var xxx_messageInfo_Key_URef proto.InternalMessageInfo

func (m *Key_URef) GetUref() []byte {
	if m != nil {
		return m.Uref
	}
	return nil
}

func (m *Key_URef) GetAccessRights() Key_URef_AccessRights {
	if m != nil {
		return m.AccessRights
	}
	return Key_URef_UNKNOWN
}

type Key_Local struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key_Local) Reset()         { *m = Key_Local{} }
func (m *Key_Local) String() string { return proto.CompactTextString(m) }
func (*Key_Local) ProtoMessage()    {}
func (*Key_Local) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{4, 3}
}
func (m *Key_Local) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key_Local.Unmarshal(m, b)
}
func (m *Key_Local) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key_Local.Marshal(b, m, deterministic)
}
func (m *Key_Local) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key_Local.Merge(m, src)
}
func (m *Key_Local) XXX_Size() int {
	return xxx_messageInfo_Key_Local.Size(m)
}
func (m *Key_Local) XXX_DiscardUnknown() {
	xxx_messageInfo_Key_Local.DiscardUnknown(m)
}

var xxx_messageInfo_Key_Local proto.InternalMessageInfo

func (m *Key_Local) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type NamedKey struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key                  *Key     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedKey) Reset()         { *m = NamedKey{} }
func (m *NamedKey) String() string { return proto.CompactTextString(m) }
func (*NamedKey) ProtoMessage()    {}
func (*NamedKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{5}
}
func (m *NamedKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedKey.Unmarshal(m, b)
}
func (m *NamedKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedKey.Marshal(b, m, deterministic)
}
func (m *NamedKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedKey.Merge(m, src)
}
func (m *NamedKey) XXX_Size() int {
	return xxx_messageInfo_NamedKey.Size(m)
}
func (m *NamedKey) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedKey.DiscardUnknown(m)
}

var xxx_messageInfo_NamedKey proto.InternalMessageInfo

func (m *NamedKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedKey) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

type Contract struct {
	Body                 []byte           `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	NamedKeys            []*NamedKey      `protobuf:"bytes,2,rep,name=named_keys,json=namedKeys,proto3" json:"named_keys,omitempty"`
	ProtocolVersion      *ProtocolVersion `protobuf:"bytes,3,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{6}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contract.Unmarshal(m, b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return xxx_messageInfo_Contract.Size(m)
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Contract) GetNamedKeys() []*NamedKey {
	if m != nil {
		return m.NamedKeys
	}
	return nil
}

func (m *Contract) GetProtocolVersion() *ProtocolVersion {
	if m != nil {
		return m.ProtocolVersion
	}
	return nil
}

type Account struct {
	PublicKey            []byte                    `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PurseId              *Key_URef                 `protobuf:"bytes,3,opt,name=purse_id,json=purseId,proto3" json:"purse_id,omitempty"`
	NamedKeys            []*NamedKey               `protobuf:"bytes,4,rep,name=named_keys,json=namedKeys,proto3" json:"named_keys,omitempty"`
	AssociatedKeys       []*Account_AssociatedKey  `protobuf:"bytes,5,rep,name=associated_keys,json=associatedKeys,proto3" json:"associated_keys,omitempty"`
	ActionThresholds     *Account_ActionThresholds `protobuf:"bytes,6,opt,name=action_thresholds,json=actionThresholds,proto3" json:"action_thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{7}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Account) GetPurseId() *Key_URef {
	if m != nil {
		return m.PurseId
	}
	return nil
}

func (m *Account) GetNamedKeys() []*NamedKey {
	if m != nil {
		return m.NamedKeys
	}
	return nil
}

func (m *Account) GetAssociatedKeys() []*Account_AssociatedKey {
	if m != nil {
		return m.AssociatedKeys
	}
	return nil
}

func (m *Account) GetActionThresholds() *Account_ActionThresholds {
	if m != nil {
		return m.ActionThresholds
	}
	return nil
}

type Account_AssociatedKey struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account_AssociatedKey) Reset()         { *m = Account_AssociatedKey{} }
func (m *Account_AssociatedKey) String() string { return proto.CompactTextString(m) }
func (*Account_AssociatedKey) ProtoMessage()    {}
func (*Account_AssociatedKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{7, 0}
}
func (m *Account_AssociatedKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account_AssociatedKey.Unmarshal(m, b)
}
func (m *Account_AssociatedKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account_AssociatedKey.Marshal(b, m, deterministic)
}
func (m *Account_AssociatedKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account_AssociatedKey.Merge(m, src)
}
func (m *Account_AssociatedKey) XXX_Size() int {
	return xxx_messageInfo_Account_AssociatedKey.Size(m)
}
func (m *Account_AssociatedKey) XXX_DiscardUnknown() {
	xxx_messageInfo_Account_AssociatedKey.DiscardUnknown(m)
}

var xxx_messageInfo_Account_AssociatedKey proto.InternalMessageInfo

func (m *Account_AssociatedKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Account_AssociatedKey) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type Account_ActionThresholds struct {
	DeploymentThreshold    uint32   `protobuf:"varint,1,opt,name=deployment_threshold,json=deploymentThreshold,proto3" json:"deployment_threshold,omitempty"`
	KeyManagementThreshold uint32   `protobuf:"varint,2,opt,name=key_management_threshold,json=keyManagementThreshold,proto3" json:"key_management_threshold,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Account_ActionThresholds) Reset()         { *m = Account_ActionThresholds{} }
func (m *Account_ActionThresholds) String() string { return proto.CompactTextString(m) }
func (*Account_ActionThresholds) ProtoMessage()    {}
func (*Account_ActionThresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{7, 1}
}
func (m *Account_ActionThresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account_ActionThresholds.Unmarshal(m, b)
}
func (m *Account_ActionThresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account_ActionThresholds.Marshal(b, m, deterministic)
}
func (m *Account_ActionThresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account_ActionThresholds.Merge(m, src)
}
func (m *Account_ActionThresholds) XXX_Size() int {
	return xxx_messageInfo_Account_ActionThresholds.Size(m)
}
func (m *Account_ActionThresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_Account_ActionThresholds.DiscardUnknown(m)
}

var xxx_messageInfo_Account_ActionThresholds proto.InternalMessageInfo

func (m *Account_ActionThresholds) GetDeploymentThreshold() uint32 {
	if m != nil {
		return m.DeploymentThreshold
	}
	return 0
}

func (m *Account_ActionThresholds) GetKeyManagementThreshold() uint32 {
	if m != nil {
		return m.KeyManagementThreshold
	}
	return 0
}

type Unit struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unit) Reset()         { *m = Unit{} }
func (m *Unit) String() string { return proto.CompactTextString(m) }
func (*Unit) ProtoMessage()    {}
func (*Unit) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{8}
}
func (m *Unit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unit.Unmarshal(m, b)
}
func (m *Unit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unit.Marshal(b, m, deterministic)
}
func (m *Unit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unit.Merge(m, src)
}
func (m *Unit) XXX_Size() int {
	return xxx_messageInfo_Unit.Size(m)
}
func (m *Unit) XXX_DiscardUnknown() {
	xxx_messageInfo_Unit.DiscardUnknown(m)
}

var xxx_messageInfo_Unit proto.InternalMessageInfo

type ProtocolVersion struct {
	Major                uint32   `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                uint32   `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                uint32   `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtocolVersion) Reset()         { *m = ProtocolVersion{} }
func (m *ProtocolVersion) String() string { return proto.CompactTextString(m) }
func (*ProtocolVersion) ProtoMessage()    {}
func (*ProtocolVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5c211b7c38b10a, []int{9}
}
func (m *ProtocolVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolVersion.Unmarshal(m, b)
}
func (m *ProtocolVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolVersion.Marshal(b, m, deterministic)
}
func (m *ProtocolVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolVersion.Merge(m, src)
}
func (m *ProtocolVersion) XXX_Size() int {
	return xxx_messageInfo_ProtocolVersion.Size(m)
}
func (m *ProtocolVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolVersion proto.InternalMessageInfo

func (m *ProtocolVersion) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *ProtocolVersion) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *ProtocolVersion) GetPatch() uint32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func init() {
	proto.RegisterEnum("io.casperlabs.casper.consensus.state.Key_URef_AccessRights", Key_URef_AccessRights_name, Key_URef_AccessRights_value)
	proto.RegisterType((*Value)(nil), "io.casperlabs.casper.consensus.state.Value")
	proto.RegisterType((*IntList)(nil), "io.casperlabs.casper.consensus.state.IntList")
	proto.RegisterType((*StringList)(nil), "io.casperlabs.casper.consensus.state.StringList")
	proto.RegisterType((*BigInt)(nil), "io.casperlabs.casper.consensus.state.BigInt")
	proto.RegisterType((*Key)(nil), "io.casperlabs.casper.consensus.state.Key")
	proto.RegisterType((*Key_Address)(nil), "io.casperlabs.casper.consensus.state.Key.Address")
	proto.RegisterType((*Key_Hash)(nil), "io.casperlabs.casper.consensus.state.Key.Hash")
	proto.RegisterType((*Key_URef)(nil), "io.casperlabs.casper.consensus.state.Key.URef")
	proto.RegisterType((*Key_Local)(nil), "io.casperlabs.casper.consensus.state.Key.Local")
	proto.RegisterType((*NamedKey)(nil), "io.casperlabs.casper.consensus.state.NamedKey")
	proto.RegisterType((*Contract)(nil), "io.casperlabs.casper.consensus.state.Contract")
	proto.RegisterType((*Account)(nil), "io.casperlabs.casper.consensus.state.Account")
	proto.RegisterType((*Account_AssociatedKey)(nil), "io.casperlabs.casper.consensus.state.Account.AssociatedKey")
	proto.RegisterType((*Account_ActionThresholds)(nil), "io.casperlabs.casper.consensus.state.Account.ActionThresholds")
	proto.RegisterType((*Unit)(nil), "io.casperlabs.casper.consensus.state.Unit")
	proto.RegisterType((*ProtocolVersion)(nil), "io.casperlabs.casper.consensus.state.ProtocolVersion")
}

func init() {
	proto.RegisterFile("io/casperlabs/casper/consensus/state/state.proto", fileDescriptor_bc5c211b7c38b10a)
}

var fileDescriptor_bc5c211b7c38b10a = []byte{
	// 1035 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x6d, 0x6f, 0x1a, 0xc7,
	0x13, 0x07, 0x73, 0xc7, 0xdd, 0x0d, 0x60, 0xf3, 0xdf, 0x7f, 0x64, 0x21, 0xa2, 0xa8, 0x84, 0xe4,
	0x05, 0xad, 0x0a, 0x24, 0xae, 0x2a, 0x55, 0xb2, 0x5a, 0x15, 0xd7, 0x69, 0xc0, 0xb1, 0xdd, 0xea,
	0x1c, 0xc7, 0x52, 0xa3, 0xea, 0xb2, 0x77, 0xb7, 0x86, 0xad, 0xcf, 0xb7, 0xe8, 0x76, 0x49, 0x44,
	0xdf, 0xf4, 0x03, 0xf4, 0xc3, 0xf5, 0x7b, 0xb4, 0x5f, 0xa0, 0x2f, 0xab, 0x7d, 0xe0, 0xb0, 0x51,
	0xa5, 0x1e, 0xed, 0x1b, 0xb4, 0xf3, 0xf0, 0xfb, 0xcd, 0xec, 0xdc, 0xcc, 0x2c, 0xf0, 0x8c, 0xb2,
	0x61, 0x84, 0xf9, 0x9c, 0x64, 0x09, 0x0e, 0xb9, 0x39, 0x0e, 0x23, 0x96, 0x72, 0x92, 0xf2, 0x05,
	0x1f, 0x72, 0x81, 0x05, 0xd1, 0xbf, 0x83, 0x79, 0xc6, 0x04, 0x43, 0x4f, 0x29, 0x1b, 0xac, 0x11,
	0xe6, 0x38, 0xc8, 0x11, 0x03, 0xe5, 0xdb, 0xfd, 0xc3, 0x06, 0xfb, 0x0d, 0x4e, 0x16, 0x04, 0x3d,
	0x02, 0x8f, 0xa6, 0x22, 0x78, 0x2f, 0x85, 0x56, 0xb9, 0x53, 0xee, 0xd9, 0xe3, 0x92, 0xef, 0xd2,
	0x54, 0x68, 0xf3, 0x63, 0xa8, 0x85, 0x4b, 0x41, 0xb8, 0x71, 0xd8, 0xe9, 0x94, 0x7b, 0xf5, 0x71,
	0xc9, 0x07, 0xa5, 0xd4, 0x2e, 0x27, 0x20, 0xdd, 0x83, 0x84, 0x72, 0xd1, 0xaa, 0x74, 0xca, 0xbd,
	0xda, 0x41, 0x7f, 0x50, 0x24, 0x89, 0xc1, 0x24, 0x15, 0xa7, 0x94, 0x8b, 0x71, 0xc9, 0x77, 0xa8,
	0x3e, 0xa2, 0x27, 0x50, 0xe7, 0x22, 0xa3, 0xe9, 0xd4, 0xc4, 0xb3, 0x3a, 0xe5, 0x9e, 0x37, 0x2e,
	0xf9, 0x35, 0xad, 0xd5, 0x01, 0x27, 0xe0, 0xe0, 0x28, 0x62, 0x8b, 0x54, 0xb4, 0xec, 0x6d, 0xe2,
	0x8d, 0x34, 0x48, 0xc6, 0x33, 0x78, 0x74, 0x0a, 0x6e, 0xc4, 0x52, 0x91, 0xe1, 0x48, 0xb4, 0xaa,
	0x8a, 0x6b, 0x50, 0x8c, 0xeb, 0x1b, 0x83, 0x92, 0xc5, 0x5a, 0x31, 0xa0, 0x0b, 0x30, 0x79, 0xea,
	0x62, 0x38, 0x8a, 0xf0, 0x59, 0x31, 0xc2, 0x0b, 0x05, 0x34, 0xf5, 0x00, 0x9e, 0x4b, 0xe8, 0x0c,
	0xbc, 0x14, 0xdf, 0x92, 0x38, 0xb8, 0x21, 0xcb, 0x96, 0xbb, 0x4d, 0x8e, 0xe7, 0x12, 0xf6, 0x8a,
	0x2c, 0x65, 0x8e, 0xa9, 0x39, 0xa3, 0x97, 0xe0, 0x84, 0x74, 0x1a, 0xd0, 0x54, 0xb4, 0x3c, 0x45,
	0xf6, 0x69, 0x31, 0xb2, 0x23, 0x3a, 0x9d, 0xa8, 0xda, 0x55, 0x43, 0x75, 0x42, 0x5f, 0x42, 0x45,
	0x66, 0x04, 0x8a, 0xe4, 0xe3, 0x62, 0x24, 0x3a, 0x19, 0x89, 0x43, 0x5f, 0x83, 0xb5, 0x48, 0xa9,
	0x68, 0xd5, 0x14, 0xfe, 0x93, 0x62, 0xf8, 0xcb, 0x94, 0xca, 0x14, 0x14, 0x12, 0x7d, 0x04, 0x90,
	0xb0, 0xbc, 0x53, 0xea, 0x9d, 0x72, 0xcf, 0x1a, 0x97, 0x7c, 0x4f, 0xea, 0x54, 0x9f, 0x1c, 0x39,
	0x60, 0x2b, 0x5b, 0xf7, 0x31, 0x38, 0xa6, 0xd7, 0xd0, 0x3e, 0x54, 0x95, 0x8e, 0xb7, 0xca, 0x9d,
	0x4a, 0xcf, 0xf6, 0x8d, 0xd4, 0x7d, 0x0a, 0xb0, 0xfe, 0x02, 0x1b, 0x5e, 0x5e, 0xee, 0x75, 0x08,
	0x55, 0x5d, 0x07, 0xf4, 0xc0, 0x70, 0xab, 0x91, 0xf1, 0x7c, 0x2d, 0xa0, 0x87, 0xe0, 0x85, 0x54,
	0x04, 0x1f, 0x68, 0x2c, 0x66, 0x6a, 0x56, 0x1a, 0xbe, 0x1b, 0x52, 0x71, 0x25, 0xe5, 0xee, 0xaf,
	0x36, 0x54, 0xe4, 0x17, 0x38, 0x03, 0x07, 0xc7, 0x71, 0x46, 0x38, 0x57, 0xe0, 0xda, 0xc1, 0xf3,
	0xc2, 0xc5, 0x1b, 0x8c, 0x34, 0x50, 0xb5, 0xb0, 0x3e, 0xa2, 0x63, 0xb0, 0x66, 0x98, 0xeb, 0x70,
	0x85, 0x5b, 0x43, 0x72, 0x8d, 0x31, 0x9f, 0xc9, 0x62, 0x4a, 0xb4, 0x64, 0x59, 0x64, 0xe4, 0xda,
	0x0c, 0xf0, 0x16, 0x2c, 0x97, 0x3e, 0xb9, 0x56, 0x9f, 0x24, 0x23, 0xd7, 0xe8, 0x25, 0xd8, 0x09,
	0x8b, 0x70, 0xa2, 0xe6, 0xb6, 0x76, 0x30, 0x2c, 0x4e, 0x73, 0x2a, 0x61, 0xe3, 0x92, 0xaf, 0xf1,
	0xed, 0x27, 0xe0, 0x98, 0xab, 0xa2, 0xd6, 0x7a, 0xda, 0x65, 0xb9, 0xea, 0xf9, 0xf0, 0xb6, 0xdb,
	0x60, 0xc9, 0x3b, 0x20, 0x64, 0x2a, 0xa0, 0xcd, 0xea, 0xdc, 0xfe, 0xb3, 0x0c, 0x96, 0x4c, 0x4d,
	0x1a, 0xd5, 0xc5, 0x8c, 0x51, 0xa5, 0xf9, 0x0e, 0x1a, 0x38, 0x8a, 0x08, 0xe7, 0x41, 0x46, 0xa7,
	0x33, 0xc1, 0x55, 0xed, 0x76, 0x0f, 0x0e, 0xb7, 0xbb, 0xb5, 0xdc, 0x27, 0x84, 0x73, 0x5f, 0x51,
	0xf8, 0x75, 0x7c, 0x47, 0xea, 0xfe, 0x0c, 0xf5, 0xbb, 0x56, 0x54, 0x03, 0xe7, 0xf2, 0xfc, 0xd5,
	0xf9, 0x77, 0x57, 0xe7, 0xcd, 0x12, 0x72, 0xc1, 0xf2, 0x5f, 0x8c, 0x8e, 0x9b, 0x65, 0xe4, 0x81,
	0x7d, 0xe5, 0x4f, 0x5e, 0xbf, 0x68, 0xee, 0x20, 0x07, 0x2a, 0xa3, 0xe3, 0xe3, 0xa6, 0x85, 0xea,
	0xe0, 0x4a, 0x6b, 0x20, 0x25, 0x1b, 0xed, 0x02, 0x28, 0x49, 0xbb, 0x55, 0x50, 0x03, 0xbc, 0xd1,
	0xf1, 0x4a, 0xac, 0x22, 0x04, 0xbb, 0x2b, 0x67, 0xa3, 0x73, 0xda, 0x0f, 0xc1, 0x56, 0xd5, 0xfc,
	0xbb, 0xba, 0xac, 0x67, 0xe2, 0x2d, 0xb8, 0xab, 0xfd, 0x20, 0x1d, 0xe5, 0x7e, 0x30, 0xbd, 0xac,
	0xce, 0xe8, 0x50, 0x8f, 0xf7, 0xce, 0x96, 0xe3, 0xad, 0x86, 0xbb, 0xfb, 0x5b, 0x19, 0xdc, 0xd5,
	0x86, 0x94, 0xec, 0x21, 0x8b, 0x97, 0xab, 0x34, 0xe4, 0x19, 0x9d, 0x01, 0xe4, 0x4b, 0x4d, 0x96,
	0xbf, 0xb2, 0xfd, 0x56, 0xf3, 0xbd, 0xd5, 0x4e, 0xe3, 0xe8, 0x1d, 0x34, 0xd5, 0xeb, 0x17, 0xb1,
	0x24, 0x78, 0x4f, 0x32, 0x4e, 0x59, 0x6a, 0x3a, 0xf9, 0xf3, 0x62, 0xa4, 0xdf, 0x1b, 0xf4, 0x1b,
	0x0d, 0xf6, 0xf7, 0xe6, 0xf7, 0x15, 0xdd, 0xdf, 0x2d, 0x70, 0xcc, 0xfb, 0x81, 0x1e, 0x01, 0xcc,
	0x17, 0x61, 0x42, 0x23, 0xb5, 0x92, 0xf5, 0xb5, 0x3c, 0xad, 0x91, 0xd5, 0x9c, 0x80, 0x3b, 0x5f,
	0x64, 0x9c, 0x04, 0x34, 0xfe, 0x77, 0xe3, 0xe4, 0x3b, 0x0a, 0x3f, 0x89, 0x37, 0xca, 0x64, 0xfd,
	0xd7, 0x32, 0xc5, 0xb0, 0x87, 0x39, 0x67, 0x11, 0xc5, 0x62, 0xc5, 0x69, 0x2b, 0xce, 0xc3, 0xad,
	0x1e, 0xd0, 0xc1, 0x28, 0x27, 0x91, 0x01, 0x76, 0xf1, 0x5d, 0x91, 0xa3, 0x1b, 0xf8, 0x1f, 0x8e,
	0x04, 0x65, 0x69, 0x20, 0x66, 0x19, 0xe1, 0x33, 0x96, 0xc4, 0xdc, 0x3c, 0xae, 0x5f, 0x6d, 0x19,
	0x47, 0xd1, 0xbc, 0xce, 0x59, 0xfc, 0x26, 0xde, 0xd0, 0xb4, 0xbf, 0x85, 0xc6, 0xbd, 0x6c, 0xfe,
	0xe9, 0xe3, 0xec, 0x43, 0xf5, 0x03, 0x91, 0x33, 0x69, 0xd6, 0xb3, 0x91, 0xda, 0xbf, 0x40, 0x73,
	0x33, 0x1a, 0x7a, 0x0e, 0x0f, 0x62, 0x32, 0x4f, 0xd8, 0xf2, 0x96, 0xa4, 0x62, 0x7d, 0x19, 0x45,
	0xda, 0xf0, 0xff, 0xbf, 0xb6, 0xe5, 0x18, 0xf4, 0x05, 0xb4, 0x6e, 0xc8, 0x32, 0xb8, 0xc5, 0x29,
	0x9e, 0x92, 0x0d, 0x98, 0x0e, 0xb8, 0x7f, 0x43, 0x96, 0x67, 0xb9, 0x39, 0x47, 0x9e, 0x58, 0xee,
	0x4e, 0xb3, 0x72, 0x62, 0xb9, 0x4e, 0xd3, 0xed, 0x56, 0xc1, 0x92, 0x2f, 0x5d, 0xf7, 0x02, 0xf6,
	0x36, 0x1a, 0x53, 0xbe, 0x3b, 0xb7, 0xf8, 0x27, 0x96, 0x99, 0x24, 0xb4, 0xa0, 0xb4, 0x34, 0x65,
	0x99, 0x89, 0xa1, 0x05, 0xa9, 0x9d, 0x63, 0x11, 0xcd, 0x54, 0x17, 0x36, 0x7c, 0x2d, 0x1c, 0xfd,
	0xf8, 0xc3, 0xdb, 0x29, 0x15, 0xb3, 0x45, 0x38, 0x88, 0xd8, 0xed, 0x70, 0x16, 0xe3, 0xa8, 0x7f,
	0xef, 0x4f, 0x66, 0x9f, 0x90, 0xfe, 0x34, 0x9b, 0x47, 0xfd, 0x29, 0xeb, 0x2f, 0x04, 0x4d, 0x86,
	0x6a, 0x14, 0xc2, 0xc5, 0xf5, 0xb0, 0xc8, 0x7f, 0xd1, 0xb0, 0xaa, 0xdc, 0x3f, 0xfb, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xdf, 0xd6, 0x52, 0x55, 0xba, 0x0a, 0x00, 0x00,
}
