package testdb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"code.google.com/p/gogoprotobuf/proto"
	"github.com/dmadrigalejos/GoWebService/server/message"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// read body
	body, _ := ioutil.ReadAll(r.Body)

	// convert body from proto to object
	data := new(message.Data)
	proto.Unmarshal(body, data)
	result := &message.Data{}

	switch strings.ToLower(*data.Method) {
	case "search":
		result = Search(data)
	case "insert":
		result = insert(data)
	case "update":
		result = update(data)
	default:
		key := "{Error: Method not supported.}"
		result.Data = &key
	}

	// convert result to proto
	msgBuf, _ := proto.Marshal(result)

	// send reply
	fmt.Fprintf(w, string(msgBuf))
}

func Search(data *message.Data) *message.Data {
	fmt.Printf("Method => %s\n", *data.Method)
	fmt.Printf("Data => %s\n", *data.Data)

	// query

	// create reply
	result := "{result:okay}"
	reply := &message.Data{Data: &result}

	return reply
}

func insert(data *message.Data) *message.Data {
	fmt.Printf("Method => %s\n", *data.Method)
	fmt.Printf("Data => %s\n", *data.Data)

	// query

	// create reply
	result := "{result:okay}"
	reply := &message.Data{Data: &result}

	return reply
}

func update(data *message.Data) *message.Data {
	fmt.Printf("Method => %s\n", *data.Method)
	fmt.Printf("Data => %s\n", *data.Data)

	// query

	// create reply
	result := "{result:okay}"
	reply := &message.Data{Data: &result}

	return reply
}
