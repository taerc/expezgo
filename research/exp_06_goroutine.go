package main

import (
	"fmt"
	"time"
)

const (
	MessageDispatchInit = 1 + iota
	MessageDispatchAddRoom
	MessageDispatchGetOff
)

const (
	ProcessStateUndo = 1 + iota
	ProcessStateDone
)

type DispatchMessage struct {
	Id    int
	Name  string
	Data  string
	State int
}

var messageChannel chan DispatchMessage
var dataChannel chan DispatchMessage
var syncSys chan DispatchMessage

func dispatch() {
	for {

		for i := 0; i < 10; i++ {
			dm := DispatchMessage{Id: i + 10,
				Name:  "wang",
				State: ProcessStateUndo,
			}
			fmt.Println("dispatch")
			dataChannel <- dm
			time.Sleep(1)

		}
		time.Sleep(10)
	}

}

func process() {
	for {
		data := <-dataChannel
		fmt.Println("process data ", data.Id, data.Name, data.Data, data.State)
		time.Sleep(1)
		data.State = ProcessStateDone
		data.Data = "HaveDone"
		sendMessage(data)
	}
}

func sendMessage(msg DispatchMessage) {
	//postProcess(msg)
	messageChannel <- msg
}
func postProcess() {

	for {
		data := <-messageChannel
		fmt.Println("post data ", data.Id, data.Name, data.Data, data.State)
	}
}

func messageDispatchMain() {
	messageChannel = make(chan DispatchMessage, 0)
	dataChannel = make(chan DispatchMessage, 0)
	syncSys = make(chan DispatchMessage, 0)

	go dispatch()
	go process()
	go postProcess()
	<-syncSys
}

// 监听多路消息通道

var case1Message chan DispatchMessage
var case2Message chan DispatchMessage

func listemMessage() {

	for {
		select {

		case m := <-case1Message:
			fmt.Println("message 1", m.Name, m.Id, m.Data)
		case m := <-case2Message:
			fmt.Println("message 2", m.Name, m.Id, m.Data)
		default:
			fmt.Println("default")
		}
	}
}

func sendCase1() {

	for {
		m := DispatchMessage{Name: "wangfm",
			Id:   100,
			Data: "test send case."}
		case1Message <- m
		time.Sleep(100)
	}

}

func sendCase2() {

	for {
		m := DispatchMessage{Name: "xidada",
			Id:   200,
			Data: "test send case2."}
		case1Message <- m
		time.Sleep(50)
	}
}

func channelTimeOut() {

}

func main() {

	case1Message = make(chan DispatchMessage, 0)
	case2Message = make(chan DispatchMessage, 0)
	go listemMessage()
	go sendCase1()
	//go sendCase2()
	syncSys = make(chan DispatchMessage, 0)
	<-syncSys

}
