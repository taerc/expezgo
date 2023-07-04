package main

import "fmt"
import "net/url"

var srcData []string = []string{"A", "B", "C", "D"}
var maskData []int = []int{0, 0, 0, 0}
var dataDomain int = 4
var dataTarget int = 2
var dstData []string = []string{"", ""}

func BFS(n, m int) {

	if n == 2 {
		for i := 0; i < 2; i++ {
			fmt.Printf("%s ", dstData[i])
		}
		fmt.Println("")
		return
	}

	for i := n; i < 4; i++ {
		if maskData[i] == 0 && i >= m {
			maskData[i] = 1
			dstData[n] = srcData[i]
			BFS(n+1, i)
			maskData[i] = 0

		}
	}

}
func combination_0x01() {
	// c2/4
	BFS(0, 0)
}

func combination_0x02()  {


	s := "?token=7025551298334294016LdgL1HE"

	u, e := url.Parse(s)
	if e != nil {
		fmt.Println(e)
	}

	token := u.Query().Get("token")
	sec := u.Query().Get("sec")
	fmt.Println(token)
	fmt.Println(sec)

}
func main() {
	//combination_0x01()
	combination_0x02()
}
