package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
	
}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}package main

import (
	f "fmt"
	"net"
)

func main()  {

	f.Println("server running 8888 port")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for  {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn)  {

	data := make([]byte, 4096)

	for  {
		n, err := conn.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}

}