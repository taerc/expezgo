package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
)

func calHmacSha256(data, key string) string {
	k := []byte(key)
	h := hmac.New(sha256.New, k)
	h.Write([]byte(data))
	//sha := hex.EncodeToString(h.Sum(nil))
	//fmt.Println("sss :", sha)
	//return  base64.StdEncoding.EncodeToString([]byte(sha))
	return  base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func calMd5Content(data []byte) string {
	h := md5.New()
	h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func fixHeadString(method, accept, md5, contenType, date string) string {
	head := method + "\n" +
		accept + "\n" +
		md5 + "\n" +
		contenType + "\n"

	return head
}

func main()  {

	host := "http://172.10.50.239"
	url := "/artemis/api/resource/v1/cameras"

	//appKey := "23885619"
	//appSecret := "jccEvp7mPo4ZVXMSInqd"
	appKey := "22918394"
	appSecret := "FiPnSrHq5nUVh63XwaG9"
	//contentType := "application/json;charset=UTF-8"
	//client := &http.Client{}
	data := []byte(`{"pageNo": 1,"pageSize": 20,"treeCode": "0"}`)

	req, _ := http.NewRequest("POST", host + url, bytes.NewBuffer(data))
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("x-ca-key", appKey)
	req.Header.Set("x-ca-nonce", "a397a255-53be-41fc-996b-521413e9e22d")
	req.Header.Set("x-ca-signature", "fRS00drn0opGb0gErY6S7MwtVaAHY0YhKndTPlnoNmU=")
	req.Header.Set("x-ca-signature-headers", "x-ca-key,x-ca-nonce,x-ca-timestamp")
	req.Header.Set("x-ca-timestamp", "1673080126260154")

	h1 := fixHeadString("POST", "application/json", "F3SVYJEJaJHT1m2cfPkbdQ==", "application/json;charset=UTF-8", "")
	// test dict ordering
	h11 := "x-ca-key" + ":" + appKey + "\n" +
	"x-ca-nonce" + ":" + "a397a255-53be-41fc-996b-521413e9e22d" + "\n" +
	"x-ca-timestamp" + ":" + "1673080126260154" + "\n"

	// 测试结果与海康sdk保持一致
	fmt.Println(calHmacSha256(h1 + h11 + url, appSecret))

	// md5 content
	fmt.Println(calMd5Content(data))
	// F3SVYJEJaJHT1m2cfPkbdQ==

}
