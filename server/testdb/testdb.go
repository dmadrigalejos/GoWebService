package testdb

import (
	"fmt"
	"io/ioutil"
	"net/http"

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

	switch *data.Value {
	case "exist":
		result = exist(data)
	case "save":
		result = save(data)
	case "update":
		result = update(data)
	default:
		key := "Error"
		result = &message.Data{Key: &key}
	}

	// convert result to proto
	msgBuf, _ := proto.Marshal(result)

	// send reply
	fmt.Fprintf(w, string(msgBuf))
}

func exist(data *message.Data) *message.Data {
	fmt.Printf("Method => %s\n", *data.Value)
	for i := range data.Data {
		fmt.Printf("Param%d => %s\n", i, *data.Data[i].Value)
	}

	// query

	// create reply
	resultKey := "result"
	resultValue := "okay"
	result := &message.Data{Key: &resultKey, Value: &resultValue}

	return result
}

func save(data *message.Data) *message.Data {
	fmt.Printf("Method => %s\n", *data.Value)
	for i := range data.Data {
		fmt.Printf("Param%d => %s\n", i, *data.Data[i].Value)
	}

	// query

	// create reply
	resultKey := "result"
	resultValue := "okay"
	result := &message.Data{Key: &resultKey, Value: &resultValue}

	return result
}

func update(data *message.Data) *message.Data {
	fmt.Printf("Method => %s\n", *data.Value)
	for i := range data.Data {
		fmt.Printf("Param%d => %s\n", i, *data.Data[i].Value)
	}

	// query

	// create reply
	resultKey := "result"
	resultValue := "okay"
	result := &message.Data{Key: &resultKey, Value: &resultValue}

	return result
}
