package goio

import (
	"net"
	"fmt"
	"log"
)

type SocketServer struct {
	ip string
	port uint
	protocol string
	isClose bool
	onConnection func()
	onMessage func([]byte)
	onError func(error)
	onClose func()
}

func NewSocketServer(ip string,port uint,protocol string) *SocketServer  {
	return &SocketServer{
		ip : ip,
		port : port,
		protocol : protocol,
		isClose : false,
		onConnection : func() {},
		onMessage : func([]byte) {},
		onError: func(error){},
		onClose : func() {},
	}
}


func (c *SocketServer) Run() error {

	netListen, err := net.Listen(c.protocol, fmt.Sprintf("%s:%d",c.ip,c.port))

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
		go func(conn net.Conn,c *SocketServer) {
			if c.onConnection != nil{
				go c.onConnection()
			}


		}(conn,c)

		//buf := bytes.NewBufferString("")
		//
		//buffer := make([]byte, 2048)
		//
		//for{
		//	n,err := conn.Read(buffer)
		//	if err != nil {
		//		log.Println(err)
		//		break
		//	}
		//	buf.Write(buffer[0:n])
		//
		//	pack := core.NewPacked(0,0)
		//
		//	pack,err = pack.Depack(buf)
		//	if(err == nil){
		//		log.Println(conn.RemoteAddr().String(), "receive data string => ", string(pack.Bytes()))
		//	}else{
		//		log.Println(err,n)
		//	}
		//}

	}

}

type AsyncConnectionCallback func()

func (c *SocketServer) AsyncConnection(handle AsyncConnectionCallback)  {
	c.onConnection = handle
}

type AsyncReceiveMessageCallback func()

func (c *SocketServer) AsyncMessage(handle func([]byte)) {
	c.onMessage = handle
}

func (c *SocketServer) AsyncError(handle func(error))  {
	c.onError = handle
}

func (c *SocketServer) AsyncClosed(handle func())  {
	c.onClose = handle
}

func (c *SocketServer) Close() {

}





