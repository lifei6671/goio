package goio

import (
	"net"
	"fmt"
	"log"
	"github.com/lifei6671/goio/core"
)

type IMServer struct {
	ip string
	port uint
}

type Handler interface {

}
func NewIMServer(ip string,port uint) *IMServer  {
	return &IMServer{
		ip : ip,
		port : port,
	}
}

func (c *IMServer) Run() error {

	netListen, err := net.Listen("tcp", fmt.Sprintf("%s:%d",c.ip,c.port))

	if err != nil {
		return err
	}
	defer netListen.Close()
	for{
		conn,err := netListen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		buffer := make([]byte, 2048)

		for{
			_,err := conn.Read(buffer)
			if err != nil {
				log.Println(err)
				break
			}
			pack := core.NewPacked(0,0)
			pack,err = pack.Depack(buffer)

			log.Println(conn.RemoteAddr().String(), "receive data string:\n", string(pack.Bytes()))
		}

	}

}
