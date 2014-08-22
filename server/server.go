package main

import (
	"fmt"
	"net/http"

	"github.com/dmadrigalejos/GoWebService/server/testdb"
)

func main() {
	fmt.Println("Listening on http://localhost:8080")

	http.HandleFunc("/testdb", testdb.Handler)
	http.ListenAndServe(":8080", nil)
}
