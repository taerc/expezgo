package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (es *echoServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("open ")

	return
}

func (es *echoServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println("close ", c.RemoteAddr().String())
	return
}
func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Echo synchronously.
	fmt.Println(string(frame))
	out = frame
	return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

// CeilToPowerOfTwo returns n if it is a power-of-two, otherwise the next-highest power-of-two.
func CeilToPowerOfTwo(n int) int {
	fmt.Println(n)
	n--

	fmt.Println(n)
	n |= n >> 1
	fmt.Println(n)
	n |= n >> 2
	fmt.Println(n)
	n |= n >> 4
	fmt.Println(n)
	n |= n >> 8
	fmt.Println(n)
	n |= n >> 16
	fmt.Println(n)
	n++
	fmt.Println(n)

	return n
}

func CeilToPowerOfx(n int) int {

	fmt.Println("n")
	fmt.Println(n)

	n = ((n + 15) >> 4) << 4
	fmt.Println(n)

	return n
}

func main() {
	var port int
	var multicore bool

	CeilToPowerOfTwo(100)
	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()
	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}
