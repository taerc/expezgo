package main

import (
	"fmt"
	"net/http"
	"time"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("sayHello")
	fmt.Println(&request)

	go func() {
		for range time.Tick(time.Second) {

			select {
			case <-request.Context().Done():
				fmt.Println("request is outgoing")
				return
			default:
				fmt.Println("Current request is in progress!")
			}
		}
	}()
	time.Sleep(2 * time.Second)
	writer.Write([]byte("Hi, New Request Comes"))
}

func main() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8080", nil)

}
