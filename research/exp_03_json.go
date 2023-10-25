package main

import (
	"encoding/json"
	"fmt"
)

func order(key string, columns ...string) {

	fmt.Println(key)

	for _, v := range columns {
		fmt.Println(v)
	}

}

type AUser struct {
	Name string `json:"name"`
	Age  int64  `json:"age,string"`
}

func testJsonInt64String() {
	u := AUser{
		Name: "wangfm",
		Age:  12,
	}

	if b, e := json.Marshal(&u); e == nil {
		fmt.Println(e)
		fmt.Println(string(b))

		m := AUser{}

		json.Unmarshal(b, &m)

		fmt.Println(m.Age)
		fmt.Println(m.Name)
	}
}

func PrintSize(a []int) {
	fmt.Println(len(a))
}

func testPointer() {
	a := []int{2, 4, 5, 7, 8, 10, 1}
	fmt.Println(len(a))
	//PrintSize(&a[3])

}

func testSlice() []int {
	n := make([]int, 10)

	n[0] = 1
	n[3] = 6

	for _, v := range n {
		fmt.Println(v)
	}
	fmt.Println(len(n))
	mp := make(map[string]int, 5)
	mp["data"] = 10
	mp["name1"] = 20
	mp["data2"] = 35
	mp["name"] = 55
	mp["root"] = 55
	mp["glass"] = 55
	fmt.Println(len(mp))

	for k, v := range mp {
		fmt.Println(k, v)

	}
	return n
}

func constJson() {
	bts, e := json.Marshal(1)

	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println(string(bts))

}

type mediaInfo struct {
	Id   int `json:"id"`
	Data map[string]int
}

func mapJson() {

	mi := mediaInfo{}
	mi.Id = 13
	mi.Data = make(map[string]int)
	mi.Data["wang"] = 12
	mi.Data["chen"] = 24
	bts, e := json.Marshal(mi)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(string(bts))
}

func main() {
	mapJson()
}
