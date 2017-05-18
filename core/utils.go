package core

import (
	"encoding/binary"
)



//整形转换成字节
func IntToBytes(v uint32) []byte {

	b := make([]byte, 8)

	binary.LittleEndian.PutUint32(b, v)

	return b
}
