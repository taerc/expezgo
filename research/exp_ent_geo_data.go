package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func generateGEOData() {

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

const (
	_             = iota
	CodeParamBase = 10000 * iota
	CodeAuthBase
	CodePermBase
	CodeResourceBase
	CodeSystemBase
	CodeDataBase
	CodeServiceBase
)

const (
	CodeServiceAdminBase = CodeServiceBase + iota*500
	CodeServiceLineBase
	CodeServiceTowerBase
)

func constData() {

	fmt.Println(CodeDataBase)
	fmt.Println(CodeSystemBase)
	fmt.Println(CodeServiceTowerBase, CodeServiceLineBase, CodeServiceAdminBase)

}

func GetMessageByCode(c int) string {
	return "!!!!!!"
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d:%s : %s", e.Code, GetMessageByCode(e.Code), e.Message)
}

func NewError(c int, m string) error {
	return &Error{
		Code:    c,
		Message: m,
	}
}
func NewErrorFromError(c int, e error) error {
	return &Error{
		Code:    c,
		Message: e.Error(),
	}
}

const CodeErrorBase = 80000
func GetErrorCode(e error) int {
	es := e.Error()
	if idx := strings.Index(es, ":"); idx != -1 {
		if c, e := strconv.ParseInt(es[0:idx], 10, 64); e == nil {
			return int(c)
		}
	}
	return CodeErrorBase
}

func genError() error {
	return NewErrorFromError(200, errors.New("This is a error testing!"))
}

func testError() {
	e := genError()
	fmt.Println(GetErrorCode(e))
	fmt.Println(e.Error())

	fmt.Println(GetErrorCode(errors.New("This is a invalid error")))
}

func main() {

	testError()

}
