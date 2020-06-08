package main

import (
	f "fmt"
	"net"
	"net/rpc"
)

type Calc int

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func main() {

	f.Println("rpc server running at 6000 port")

	rpc.Register(new(Calc))

	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()

		go rpc.ServeConn(conn)
	}

}

func (c *Calc) Sum(args Args, reply *Reply) error {

	reply.C = args.A + args.B // 두 값을 더하여 리턴값 구조체에 넣어줌
	return nil

}
