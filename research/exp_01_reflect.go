package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name string
	Id   string
	Age  int
}

func display(value interface{}) {
	fmt.Println("display")
	fmt.Println(reflect.ValueOf(value)) // 获取到的是对象的指针形式
	fmt.Println(reflect.ValueOf(value).Elem())

	fmt.Println("name")
	//vType := reflect.TypeOf(value)
	vValue := reflect.ValueOf(value).Elem() // 获取到对象的值的形式
	fmt.Println(reflect.ValueOf(vValue.Field(0).Interface()))
}

type SdkItem struct {
	Name       string `gorm:"column:name" json:"name"`
	TargetArch string `gorm:"column:target_arch" json:"target_arch"`
	XpuType    string `gorm:"column:xpu_type" json:"xpu_type"`
	UsrID      int    `gorm:"column:usr_id" json:"usr_id"`
}
type SdkListItem struct {
	SdkItem
	Id         int64  `gorm:"column:id"`
	Version    string `gorm:"column:version" json:"version"`
	Desc       string `gorm:"column:desc" json:"desc"`
	CreateTime string `gorm:"column:create_time" json:"create_time"`
	UsrName    string `gorm:"column:usr_name" json:"usr_name"`
}

type ItemPage struct {
	TotalNum    int
	CurrentPage int
	PageLimit   int
	Order       string
}
type ReqSdkItemCondition struct {
	SdkItem
	ItemPage
	Begin string
	End   string
	sql   string
}

type SQLCondition struct {
	Op    string
	Value interface{}
}
type SQLBuilder struct {
	where map[string]SQLCondition
	and   map[string]SQLCondition
}

func (sb *SQLBuilder) BuildSqlCount(table string) string {
	sql := "select count(1) from `" + table + "` " + sb.BuildWhere() + sb.BuildAnd()
	return sql
}

func (sb *SQLBuilder) Where(key string, op string, value interface{}) {
	sb.where[key] = SQLCondition{
		Op:    op,
		Value: value,
	}
}

func (sb *SQLBuilder) And(key string, op string, value interface{}) {
	sb.and[key] = SQLCondition{
		Op:    op,
		Value: value,
	}
}

func (sb *SQLBuilder) BuildAnd() string {

	cond := ""
	for k, v := range sb.and {
		cond = sb.append(cond, k, v.Op, v.Value)
	}
	if len(cond) == 0 {
		return ""
	}
	return " and " + cond

}

func (sb *SQLBuilder) BuildWhere() string {

	cond := ""
	for k, v := range sb.where {
		cond = sb.append(cond, k, v.Op, v.Value)
	}
	if len(cond) == 0 {
		return ""
	}
	return " where " + cond
}

func (sb *SQLBuilder) append(sql string, key string, op string, value interface{}) string {
	if reflect.ValueOf(value).Kind() == reflect.String && reflect.ValueOf(value).String() != "" {
		if sql != "" {
			sql += " and "
		}
		sql += "`" + key + "`"
		sql += op
		sql += "\"" + reflect.ValueOf(value).String() + "\""
		reflect.ValueOf(value)
	} else if reflect.ValueOf(value).Kind() == reflect.Int && reflect.ValueOf(value).Int() != 0 {
		if sql != "" {
			sql += " and "
		}
		sql += "`" + key + "`"
		sql += op
		sql += strconv.FormatInt(reflect.ValueOf(value).Int(), 10)
	}
	return sql
}

func DefaultSqlBuilder() *SQLBuilder {

	return &SQLBuilder{
		where: make(map[string]SQLCondition, 0),
		and:   make(map[string]SQLCondition, 0),
	}

}

func testSql() {

	//u := &User{
	//	Name: "wangfangming",
	//	Id:   "123456",
	//	Age:  17,
	//}
	//display(u)

	req := &ReqSdkItemCondition{
		SdkItem: SdkItem{
			Name:       "wang",
			TargetArch: "x86",
			XpuType:    "altas200",
			UsrID:      1,
		},
		ItemPage: ItemPage{
			TotalNum:  100,
			PageLimit: 15,
		},
	}

	sb := DefaultSqlBuilder()

	sb.Where("name", "=", req.Name)
	sb.Where("name", "=", req.Name)
	sb.Where("usr_id", "=", req.UsrID)
	sb.Where("usr_page", "=", req.UsrID)
	sb.Where("create_time", "<", 1000)
	sb.And("create_time", ">", 500)

	sql := sb.BuildSqlCount("device_info")
	fmt.Println(sql)

}

type MInt64 int64

type ReqStruct struct {
	ID   json.Number `json:"id,string,omitempty"`
	Name string      `json:"name,omitempty"`
}

func (u *MInt64) UnmarshalJSON(bs []byte) error {

	fmt.Println("call me")
	d := string(bs)
	if d == "" {
		*u = 0
		return nil
	}
	if unquoted, err := strconv.Unquote(d); err == nil {
		d = unquoted
	}
	x, e := strconv.ParseInt(d, 10, 64)
	if e != nil {
		return e
	}
	*u = MInt64(x)
	return nil
}

func (u *MInt64) Unmarshal(bs []byte) error {

	fmt.Println("call me")
	d := string(bs)
	if d == "" {
		*u = 0
		return nil
	}
	if unquoted, err := strconv.Unquote(d); err == nil {
		d = unquoted
	}
	x, e := strconv.ParseInt(d, 10, 64)
	if e != nil {
		return e
	}
	*u = MInt64(x)
	return nil
}

func testJson() {
	js := []byte(`{"id": "", "name": ""}`)

	req := ReqStruct{}
	if e := json.Unmarshal(js, &req); e != nil {
		fmt.Println(e)
	}
	fmt.Println(req.ID, req.Name)

	data, e := json.Marshal(req)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(data))

}
func main() {
	testJson()
}
