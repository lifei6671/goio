package core

import (
	"errors"
	"bytes"
	"encoding/binary"
)

const MaxUint32 = 1<<32 - 1

type Packed struct {
	version uint32
	magic_num uint32
	command uint32
	length uint32
	data []byte
}

func NewPacked(v uint32,cmd uint32) *Packed {
	return &Packed{version:v, magic_num : 0, command: cmd,data : make([]byte,0)}
}

func (c *Packed) Enpack(b []byte) ([]byte,error) {

	if len(b) > MaxUint32 {
		return make([]byte,0),errors.New("pack too large")
	}
	c.data = b

	c.length = uint32(len(c.data))

	buf := bytes.NewBuffer(IntToBytes(c.version))
	buf.Write(IntToBytes(c.magic_num))
	buf.Write(IntToBytes(c.command))
	buf.Write(IntToBytes(c.length))
	buf.Write(c.data)
	return buf.Bytes(),nil
}

func (c *Packed) Depack(b []byte) (*Packed,error){
	length := len(b)

	if length < 16 {
		return nil,errors.New("pack error")
	}

	pack := &Packed{
		version : binary.LittleEndian.Uint32(b[0:3]),
		magic_num : binary.LittleEndian.Uint32(b[4:7]),
		command: binary.LittleEndian.Uint32(b[8:11]),
		length : binary.LittleEndian.Uint32(b[12:15]),
	}
	if length != int(16 + pack.length){
		return nil,errors.New("pack error")
	}
	pack.data = b[16: 16 + pack.length]

	return pack,nil
}

func (c *Packed) Bytes() []byte {
	return c.data
}