package util

import (
	"math/big"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
)

func TestEncodeToHexString(t *testing.T) {
	res := EncodeToHexString([]byte{147, 37, 61, 17})
	assert.Equal(t, res, "93253d11", "they should be equal")
}

func TestDecodeHexString(t *testing.T) {
	res := DecodeHexString("1d526a")
	assert.Equal(t, res, []byte{29, 82, 106}, "they should be equal")
}

func TestBlake2b256_0(t *testing.T) {
	res := Blake2b256([]byte{0})
	assert.Equal(t, EncodeToHexString(res), "03170a2e7597b7b7e3d84c05391d139a62b157e78786d8c082f29dcf4c111314", "they should be equal")
}

func TestBlake2b256_LEN_8(t *testing.T) {
	res := Blake2b256([]byte{147, 68, 51, 37, 29, 6, 244, 90})
	assert.Equal(t, EncodeToHexString(res), "f0f84944e0ccfa9e67383e6a448291787d208c8e46adc849f714078663d1dd36", "they should be equal")
}

func TestAbiUint32ToBytes(t *testing.T) {
	res := AbiUint32ToBytes(uint32(10))
	assert.Equal(t, res, []byte{10, 0, 0, 0})
}

func TestAbiUint64ToBytes(t *testing.T) {
	res := AbiUint64ToBytes(uint64(67305985))
	assert.Equal(t, res, []byte{1, 2, 3, 4, 0, 0, 0, 0})
}

func TestAbiBytesToBytes(t *testing.T) {
	res := AbiBytesToBytes([]byte{147, 68, 51, 37, 29, 6, 244, 90})
	assert.Equal(t, res, []byte{8, 0, 0, 0, 147, 68, 51, 37, 29, 6, 244, 90})
}

func TestAbiStringToBytes(t *testing.T) {
	res := AbiStringToBytes("안녕하세요")
	assert.Equal(t, res, []byte{15, 0, 0, 0, 236, 149, 136, 235, 133, 149, 237, 149, 152, 236, 132, 184, 236, 154, 148})
}

func TestAbiBigIntTobytes(t *testing.T) {
	src := new(big.Int)
	src, err := src.SetString("256", 10)
	if !err {
		panic(err)
	}
	res := AbiBigIntTobytes(src)
	assert.Equal(t, res, []byte{2, 0, 1})
}

func TestAbiMakeArgs(t *testing.T) {

	res := AbiMakeArgs([][]byte{
		[]byte{10, 0, 0, 0},
		[]byte{1, 2, 3, 4, 0, 0, 0, 0},
		[]byte{8, 0, 0, 0, 147, 68, 51, 37, 29, 6, 244, 90},
	})
	assert.Equal(t, res, []byte{3, 0, 0, 0, 4, 0, 0, 0, 10, 0, 0, 0, 8, 0, 0, 0, 1, 2, 3, 4, 0, 0, 0, 0, 12, 0, 0, 0, 8, 0, 0, 0, 147, 68, 51, 37, 29, 6, 244, 90})
}

func TestAbiOptionToBytes(t *testing.T) {

	res := AbiOptionToBytes([]byte{})
	assert.Equal(t, []byte{0}, res)

	res = AbiOptionToBytes(AbiUint64ToBytes(uint64(10)))
	assert.Equal(t, []byte{1, 10, 0, 0, 0, 0, 0, 0, 0}, res)
}

func TestJsonStringToDeployArgs(t *testing.T) {
	inputStr := `[{"name":"amount","value":{"int_value":123456}},{"name":"fee","value":{"int_value":54321}}]`
	args, err := JsonStringToDeployArgs(inputStr)
	assert.NoError(t, err)

	assert.Equal(t, "amount", args[0].GetName())
	assert.Equal(t, int32(123456), args[0].GetValue().GetIntValue())
	assert.Equal(t, "fee", args[1].GetName())
	assert.Equal(t, int32(54321), args[1].GetValue().GetIntValue())

	m := &jsonpb.Marshaler{}
	str := "["
	for idx, arg := range args {
		if idx != 0 {
			str += ","
		}
		s, _ := m.MarshalToString(arg)
		str += s
	}
	str += "]"

	args2, err := JsonStringToDeployArgs(str)
	assert.NoError(t, err)
	assert.Equal(t, args[0].GetName(), args2[0].GetName())
	assert.Equal(t, args[0].GetValue().GetIntValue(), args2[0].GetValue().GetIntValue())
	assert.Equal(t, args[1].GetName(), args2[1].GetName())
	assert.Equal(t, args[1].GetValue().GetIntValue(), args2[1].GetValue().GetIntValue())
}
