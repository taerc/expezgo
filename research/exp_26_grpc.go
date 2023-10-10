package main

import (
	"fmt"
	"net"
)

type GSCodec struct {
}

func (gs *GSCodec) Marshal(v interface{}) ([]byte, error) {

	return nil, nil
}

func (gs *GSCodec) Unmarshal(data []byte, v interface{}) error {

	return nil
}

func (gs *GSCodec) Name() string {
	return "GSCodec"
}

func main() {

	lis, e := net.Listen("tcp", "9999")

	if e != nil {
		fmt.Println("listen ", e.Error())
	}

}
