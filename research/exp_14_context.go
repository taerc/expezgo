package main

import (
	"context"
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
	fmt.Printf("*p %x \n", (*unsafe.Pointer)(unsafe.Pointer(&s)))

	ptr0 := uintptr((*[2]unsafe.Pointer)(unsafe.Pointer(&s))[0])
	ptr := uintptr((*[2]unsafe.Pointer)(unsafe.Pointer(&s))[1])
	//ptr3 := uintptr((*[2]unsafe.Pointer)(unsafe.Pointer(&s))[3])

	fmt.Println(reflect.TypeOf(ptr))
	fmt.Printf("ptr1 %x \n", ptr)
	fmt.Printf("ptr0 %x \n", ptr0)
	fmt.Printf("&s %x \n", &s)

}

func R1(ctx context.Context) {

	i := 0
	for {
		fmt.Println(fmt.Sprintf("R1 times %v ", i))
		i += 1

		select {

		case <-ctx.Done():
			fmt.Println("R1 Finished")
			return
		default:
			fmt.Println("R1 done")


		}
		time.Sleep(1 * time.Second)
	}

}

func R0(ctx context.Context) {

	//ctx.Deadline(10*time.Second, true)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	i := 0

	go R1(ctx)
	for {

		fmt.Println(fmt.Sprintf("times %v ", i))
		i += 1

		select {

		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			fmt.Println("done")


		}
		time.Sleep(1 * time.Second)
	}

}

func main() {

	ctx := context.Background()
	go R0(ctx)

	for {
		time.Sleep(10 * time.Second)

	}

}
