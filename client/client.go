package main

import (
	"fmt"
	"github.com/dmadrigalejos/GoWebService/server/clientlib"
)

func main() {
	cli := clientlib.New("http://localhost:8080/testdb")
	reply := cli.Request("Search", "{url:www.google.com}")
	fmt.Println(reply)
}
