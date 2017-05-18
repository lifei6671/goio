package main

import "github.com/lifei6671/goio"

func main()  {
	s := goio.NewIMServer("",8001)

	s.Run()

}
