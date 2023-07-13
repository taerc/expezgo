package main

import "fmt"

func main() {

	citys := []string{"马鞍山市", "合肥市", "芜湖市", "铜陵市", "南京市", "苏州市", "扬州市", "无锡市", "南昌市", "九江市", "景德镇市", "新余市"}

	id := 1
	for i, c := range citys {

		for j := 0; j < 4; j += 1 {
			s := fmt.Sprintf("call add_county (%d, %d ,'%s',0);", id, i+1, c+fmt.Sprintf("-%03x", j))
			fmt.Println(s)
			id += 1
		}

	}

}
