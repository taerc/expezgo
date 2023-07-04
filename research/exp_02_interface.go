package main

import "fmt"

var OK int = 200
var FAIL int = 20

type GinController interface {
	Process(name string) int

	BeginAction(name string) int
	Action(name string) int
	EndAction(name string) int
}

type GinControllerTemplate struct {
}

func (gc *GinControllerTemplate) BeginAction(name string) int {
	fmt.Println("template begin")
	if name == "wangfm" {
		return OK
	}

	return FAIL
}

func (gc *GinControllerTemplate) Action(name string) int {
	fmt.Println("template action")
	return OK
}

func (gc *GinControllerTemplate) EndAction(name string) int {
	fmt.Println("template endaction")
	return OK
}

// template
// 运行时多态的问题
func (gc *GinControllerTemplate) Process(name string) int {

	if gc.BeginAction(name) == OK {
		gc.Action(name)
		return gc.EndAction(name)
	}
	return FAIL
}

type GinControllerRegister struct {
	GinControllerTemplate
}
type GinControllerQueryDevice struct {
	GinControllerTemplate
}

// 注册 接口
func (gcr *GinControllerRegister) Action(name string) int {
	fmt.Println("GinControllerRegister")
	return OK
}
func (gc *GinControllerRegister) Process(name string) int {



	if gc.BeginAction(name) == OK {
		gc.Action(name)
		return gc.EndAction(name)
	}
	return FAIL
}

func Process(gc interface{}) {

	if a, ok := gc.(GinController); ok {
		if c :=a.BeginAction("wangfm"); c == OK {
			c = a.Action("name")
			c = a.EndAction("name")
			fmt.Println(c)
		}
	} else {
		fmt.Println(ok)
	}
}

func do2(gc interface{}) {

	return
}

func main() {

	gcr := &GinControllerRegister{}
	Process(gcr)
	//gcr.Action("name")
	//do(gcr)

}
