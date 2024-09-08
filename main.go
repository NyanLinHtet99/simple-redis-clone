package main

import (
	"net"

	resp "github.com/NyanLinHtet99/simple-redis-clone/RESP"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		res := resp.NewResp(conn)
		_, err := res.Read()
		if err != nil {
			panic(err)
		}
		writer := resp.NewWriter(conn)
		writer.Write(resp.Value{Typ: "string", Str: "OK"})
	}
}
