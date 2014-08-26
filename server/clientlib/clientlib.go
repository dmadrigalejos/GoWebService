package clientlib

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"

	"code.google.com/p/gogoprotobuf/proto"
	"github.com/dmadrigalejos/GoWebService/server/message"
)

type Client struct {
	Url string
}

func New(url string) *Client {
	return &Client{url}
}

func (c *Client) Request(method, data string) *message.Data {
	// create request message
	request := &message.Data{Method: &method, Data: &data}

	fmt.Println(request)

	// convert to proto
	msgBuf, _ := proto.Marshal(request)

	// send post request
	r, _ := http.Post(c.Url, "application/zip", bytes.NewBuffer(msgBuf))
	response, _ := ioutil.ReadAll(r.Body)

	// convert response
	reply := new(message.Data)
	proto.Unmarshal(response, reply)

	return reply
}
