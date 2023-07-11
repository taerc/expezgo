package main

import (
	"fmt"
	"net/http"
	"path"
	"reflect"
	"time"
	"unsafe"
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

func test() {
	fmt.Println(path.Join("/api", "v1", "book"))
	fmt.Println(path.Join("/api", "v1", "/book"))
}

func test_pointer1(s any) {

	fmt.Printf("p %x \n", unsafe.Pointer(&s))
	fmt.Printf("*p %x \n", (*unsafe.Pointer)(unsafe.Pointer(&s)) )

	ptr0 := uintptr((*[2]unsafe.Pointer)(unsafe.Pointer(&s))[0])
	ptr := uintptr((*[2]unsafe.Pointer)(unsafe.Pointer(&s))[1])
	//ptr3 := uintptr((*[2]unsafe.Pointer)(unsafe.Pointer(&s))[3])

	fmt.Println(reflect.TypeOf(ptr))
	fmt.Printf("ptr1 %x \n", ptr)
	fmt.Printf("ptr0 %x \n", ptr0)
	fmt.Printf("&s %x \n", &s)

}


func main() {
	//http.HandleFunc("/", SayHello)
	//http.ListenAndServe(":8080", nil)
	//test()
	var age = 12
	//var age= "usuuuu"
	fmt.Printf("s %x \n", &age)
	test_pointer1(age)

	var name =  12
	//var name= "usuuuu"
	fmt.Printf("s %x \n", &name)
	test_pointer1(name)
	var addr uintptr = 0x85ebca6b

	fmt.Printf("data %x \n", (*[2]unsafe.Pointer)(unsafe.Pointer(&addr)))

}
