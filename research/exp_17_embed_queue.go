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
	go cons()

	for {
		time.Sleep(5 * time.Second)

	}
}

func enqueue() {

	s := "this is a suffix"

	for {
		si := fmt.Sprintf("%s %d", s, i)
		queue.EnqueueString(si)
		fmt.Println(fmt.Sprintf("IN %s", si))
		i++
		time.Sleep(1 * time.Second)
	}
}

func dequeue() {

	// dequeue a string
	s, e := queue.Dequeue()
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
			continue
		}

		if e != nil {
			panic(e)
		}
		chanProcess <- s.ToString()
	}
}

func cons() {

	for s := range chanProcess {
		go processs(s)
	}
}

func processs(s string) {

	fmt.Println(fmt.Sprintf("OUT %s", s))
	time.Sleep(1 * time.Second)

}
