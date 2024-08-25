package resp

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func reader(input string) string {
	reader := bufio.NewReader(strings.NewReader(input))
	dataType, _ := reader.ReadByte()
	if dataType != '$' {
		panic("Invalid Type: Accept only string")
	}
	size, _ := reader.ReadByte()
	strSize, _ := strconv.ParseInt(string(size), 10, 64)
	reader.ReadByte()
	reader.ReadByte()
	name := make([]byte, strSize)
	reader.Read(name)
	fmt.Println(string(name))
	return string(name)
}
