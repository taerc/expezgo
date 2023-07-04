package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"airiacloud/modules/guandan/types"
	"github.com/gorilla/websocket"
)

//var addr = flag.String("addr", "localhost:8080", "http service address")

// SendMsg @description: WebSocket发送数据通用方法
// @parameter c
// @parameter msg
func sendMsg(client *websocket.Conn, msg types.AiWeWsReq) {
	//sendMsgObject.Client.Lock.RLock()
	//defer sendMsgObject.Client.Lock.RUnlock()
	msgByte, err := json.Marshal(msg)
	if err != nil {
		log.Println("send msg [%v] marsha1 err:%v", string(msgByte), err)
		return
	}
	err = client.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Println("send msg SetWriteDeadline [%v] err:%v", string(msgByte), err)
		return
	}
	w, err := client.NextWriter(websocket.TextMessage)
	if err != nil {
		err = client.Close()
		if err != nil {
			log.Println("close client err: %v", err)
		}
	}
	_, err = w.Write(msgByte)
	if err != nil {
		log.Println("Write msg [%v] err: %v", string(msgByte), err)
	}
	if err := w.Close(); err != nil {
		err = client.Close()
		if err != nil {
			log.Println("close err: %v", err)
		}
	}
}

func main() {
	//flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "172.10.50.238:5021", Path: "/airiacloud/api/ai_whipped_egg/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	//done := make(chan struct{})

	ws := types.AiWeWsReq{}
	ws.Id = types.JoinRoomReqId
	usr := types.UserInfo{
		UserId:   1234,
		UserName: "A1",
		UserRole: 1,
		Location: 0,
	}
	ws.Data = usr
	sendMsg(c, ws)

	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", message)

	for {
		select {
		case <-time.After(3 * time.Second):
			log.Println("Information")
		}
	}
}
