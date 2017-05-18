package main

import "github.com/lifei6671/goio"

func main()  {
	c := goio.NewIMClient("192.168.3.104",8001)

	conn ,err := c.Dial()

	if err != nil {

	}
	conn.Write([]byte("aaaaaaaaaaaa"))
	conn.Write([]byte("bbbbbbbbbbbbbb"))

	defer conn.Close()

}
