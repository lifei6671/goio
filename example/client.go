package main

import (
	"github.com/lifei6671/goio"
	"fmt"
	"sync"
	"bytes"
)

func main()  {

	//buf := bytes.NewBufferString("abc")
	//
	//fmt.Println(buf.Len(),"=>",string(buf.Bytes()[0:1]))
	//fmt.Println(buf.Len(),"=>",string(buf.Bytes()[1:2]))
	//fmt.Println(buf.Len(),"=>",string(buf.Bytes()[2:3]))
	//
	//return
	c := goio.NewIMClient("192.168.3.104",8001)

	conn ,err := c.Dial()

	if err != nil {

	}

	wait := &sync.WaitGroup{}
	wait.Add(1)
	go func(conn *goio.IMConn,wait *sync.WaitGroup) {
		defer wait.Done()
		n,err := conn.Write([]byte(`Socket.IO 2.0.0发布了！

Socket.IO对于实时性的聊天应用开发非常方便，它基于Websocket协议开发，但可惜的是性能不足，内存占用也一直是个问题。但是令人非常期待的是，这次版本的发布主要带来了一些性能上的提升：

现在使用uws作为默认的websocket引擎。它将带来巨大的性能提升（特别是内存的消耗）(Engine.IO版本说明)

Engine.IO和Socket.IO的握手包被合并了，减少了一个连接的建立（#2833）。

现在可以为你的应用定义自定义的解析器（#2829）。查看示例可以获取更多的信息。

    需要注意的是，这个版本并不向下兼容，所以想升级到此版本的用户需要格外注意，因为：

engine.io-parser模块中有一个关于utf-8编码的突破性的变化。

一个使客户端的socket id匹配服务端的id的更新（#1058）

    相关更新包：

socket.io-redis 版本 5.x

socket.io-emitter 版本 3.x

    从CDN获取最新的客户端版本：`))
		if err != nil {
			fmt.Println(err,"=>",n)
		}else{
			fmt.Println(n)
		}
	}(conn,wait)

	wait.Add(1)
	go func(conn *goio.IMConn,wait *sync.WaitGroup) {
		defer wait.Done()
		n,err := conn.Write([]byte("各位同事，现在网管过来201室了，大家有关于电脑或网络的问题需要维修及维护的请及时提出！！！谢谢！@全体成员 "))

		if err != nil {
			fmt.Println(err,"=>",n)
		}else{
			fmt.Println(n)
		}
	}(conn,wait)

	wait.Wait()

	defer conn.Close()

	buf := bytes.NewBufferString("")
	buffer := make([]byte, 2048)
	for{
		n,err := conn.Read(buffer)
		if err == nil {
			buf.Write(buffer[:n])
		}else{
			break
		}
	}
}
