package goio

import (
	"net"
	"fmt"
	"log"
)

type IMClient struct {
	ip string
	port uint
}

func NewIMClient(ip string,port uint) *IMClient {
	return &IMClient{ ip : ip, port : port}
}

func (c *IMClient) Dial() (*IMConn,error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d",c.ip,c.port))

	if err != nil {
		log.Println(err)
		return nil,err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &IMConn{ conn:conn},nil
}

