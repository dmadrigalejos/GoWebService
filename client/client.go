package main

import (
	"fmt"
	"github.com/dmadrigalejos/GoWebService/server/clientlib"
)

func main() {
	cli := clientlib.New("http://localhost:8080/testdb")
	reply := cli.Request("save", "www.google.com", "123", "abc", "xmy", "qqq")
	fmt.Println(reply)
}
