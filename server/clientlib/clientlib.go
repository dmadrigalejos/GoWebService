package clientlib

import (
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

func (c *Client) Request(method string, args ...interface{}) *message.Data {
	// create request message
	request := &message.Data{Value: &method}

	for _, arg := range args {
		value, _ := arg.(string)
		param := &message.Data{Value: &value}
		request.Data = append(request.Data, param)
	}

	// convert to proto
	msgBuf, _ := proto.Marshal(request)

	// send post request
	r, _ := http.Post(c.Url, "text/json", bytes.NewBuffer(msgBuf))
	response, _ := ioutil.ReadAll(r.Body)

	// convert response
	data := new(message.Data)
	proto.Unmarshal(response, data)

	return data
}
