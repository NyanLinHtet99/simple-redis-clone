package main

import (
	"fmt"
	"net"
	"strings"

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
		value, err := res.Read()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if value.Typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}
		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected command length > 0")
			continue
		}
		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]
		writer := resp.NewWriter(conn)
		handler, ok := resp.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Typ: "string", Str: ""})
			continue
		}
		result := handler(args)
		writer.Write(result)
	}
}
