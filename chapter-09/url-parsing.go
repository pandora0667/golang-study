package main

import (
	f "fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	f.Println(u.Scheme)
	f.Println(u.User)
	f.Println(u.User.Username())

	password, _ := u.User.Password()
	f.Println(password)

	f.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	f.Println(host)
	f.Println(port)

	f.Println(u.Path)
	f.Println(u.Fragment)

	f.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	f.Println(m)
	f.Println(m["k"][0])

}
