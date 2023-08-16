package main

import (
	"fmt"

	"github.com/mozillazg/go-pinyin"
)

func main() {

	hans := "枚举出所有的运维班组"
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]

	fmt.Println(pinyin.LazyConvert(hans, nil))
	// [zhong guo ren]

}
