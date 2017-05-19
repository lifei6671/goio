package core

import (
	"errors"
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)


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

	if len(b) > math.MaxUint32 {
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

func (c *Packed) Depack(buf *bytes.Buffer) (*Packed,error){

	if buf.Len() < 16 {
		return nil,errors.New(fmt.Sprintf("pack shot 16 bit => %d",buf.Len()))
	}

	pack := &Packed{}

	b := make([]byte,16)

	n,err := buf.Read(b)

	//fmt.Println("buf => ",b);
	if n != 16 {
		return nil,errors.New(fmt.Sprintf("pack length error => %d",n))
	}
	if err != nil {
		return nil,err
	}
	pack.version =  binary.LittleEndian.Uint32(b[:4])
	pack.magic_num =  binary.LittleEndian.Uint32(b[4:8])
	pack.command = binary.LittleEndian.Uint32(b[8:12])
	pack.length = binary.LittleEndian.Uint32(b[12:16])


	if buf.Len() < int(pack.length){
		return nil,errors.New(fmt.Sprintf("pack error => %d ", pack.length))
	}

	b = make([]byte,pack.length)

	n,err = buf.Read(b)

	if n != int(pack.length) {
		return nil,errors.New(fmt.Sprintf("pack length error => %d",n))
	}
	if err != nil {
		return nil,err
	}

	pack.data = b

	return pack,nil
}

func (c *Packed) Bytes() []byte {
	return c.data
}

func (c *Packed) DataLength() int {
	return len(c.data)
}