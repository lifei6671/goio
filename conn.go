package goio

import (
	"net"
	"time"
	"github.com/lifei6671/goio/core"
	"bytes"
)

type IMConn struct {
	conn net.Conn
}

func (c *IMConn) Write(b []byte) (int,error){
	pack := core.NewPacked(1,1)

	b ,err := pack.Enpack(b)
	if err != nil {
		return 0,err
	}

	return c.conn.Write(b)
}

func (c *IMConn) Read(b []byte)  (int,error) {
	pack := core.NewPacked(0,0)
	var bb []byte

	_,err := c.conn.Read(bb)

	if err != nil {
		return 0,err
	}
	buf := bytes.NewBuffer(bb)
	pack,err = pack.Depack(buf)

	if err != nil {
		return 0,err
	}
	b = pack.Bytes()
	return len(b),nil
}

func (c *IMConn) Close()error {
	return c.conn.Close()
}

func (c *IMConn) LocalAddr() net.Addr  {
	return c.conn.LocalAddr()
}

func (c *IMConn)RemoteAddr() net.Addr  {
	return c.conn.RemoteAddr()
}

func (c *IMConn) SetDeadline(t time.Time) error  {
	return c.conn.SetDeadline(t)
}

func (c *IMConn)SetReadDeadline(t time.Time) error  {
	return c.conn.SetReadDeadline(t)
}

func (c *IMConn)SetWriteDeadline(t time.Time) error  {
	return c.conn.SetWriteDeadline(t)
}





















