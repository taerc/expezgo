package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/beeker1121/goque"
)

var queue *goque.Queue = nil

var i = 0

var chanProcess = make(chan string, 5)

func main() {
	var e error
	queue, e = goque.OpenQueue("data_dir")

	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer queue.Close()
	defer close(chanProcess)

	go enqueue()
	go chanDequeue()

	for {
		time.Sleep(5 * time.Second)

	}
}

func enqueue() {

	s := "this is a suffix"

	for {
		queue.EnqueueString(fmt.Sprintf("%s %d", s, i))
		fmt.Println("inqueue ")
		i++
		time.Sleep(1 * time.Second)
	}

}

func dequeue() {

	// dequeue a string
	s, e := queue.Dequeue()
	fmt.Println(e.Error())
	if e != nil && errors.Is(e, goque.ErrEmpty) {
		fmt.Println("queue is empty")
		return
	}

	if e != nil {
		panic(e)
	}
	fmt.Println(s.ToString())
}

func chanDequeue() {
	// dequeue a string

	for {
		s, e := queue.Dequeue()
		if e != nil && errors.Is(e, goque.ErrEmpty) {
			fmt.Println("queue is empty")
			time.Sleep(3 * time.Second)
			return
		}

		if e != nil {
			panic(e)
		}

		chanProcess <- s.ToString()
		go cons(chanProcess)

	}
}

func cons(ta chan string) {
	fmt.Println(<-ta)
	time.Sleep(5 * time.Second)
}
