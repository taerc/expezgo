package main

import (
	"bytes"
	"fmt"
	"github.com/taerc/ezgo"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func test_01() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

type SimpleNotice struct {
	Title    string
	ImageUrl string
	UrlName  string
	Url      string
	Project  string
	Tag      string
	Author   string
	Items    []string
}

func test_02() {

	tplText := `
**项目** : {{.Project}}
{{if .ImageUrl}}![image]({{.ImageUrl}}) {{end}}
**标题**: {{.Title}}
**标签**: {{.Tag}}
**作者**: {{.Author}}
**详情**:
{{if .Url}}![链接]({{.Url}}) {{end}}
{{- range $i, $e := .Items }}
* {{$e}}
{{- end }}
`
	sn := SimpleNotice{Project: "齐鲁石化", Title: "项目更新", Tag: "v222222", Author: "wfm",
		Url:      "http://www.baidu.com/test/iage",
		ImageUrl: "http:///www.123123.com",
		Items:    []string{"wwww wwww", "eee eee", "fffff"}}

	tpl, err := template.New("hello").Parse(tplText)
	if err != nil {
		fmt.Printf("failed parse tpltext,err:%s\n", err.Error())
		return
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, sn)
	if err != nil {
		fmt.Printf("failed execute tpltext,err:%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", buf.String())

}

//func StringSplits(str string, seps []string) []string {
//
//	strs := []string{str}
//	strsBuf := make([]string, 0)
//	for _, sp := range seps {
//		for _, s := range  strs {
//			sa := strings.Split(s, sp)
//			strsBuf = append(strsBuf, sa...)
//		}
//		strs = strsBuf
//		strsBuf = append([]string{}) // clear slice
//	}
//	return strs
//}
func test_split() {

	s := "this is a book; 中国共产党；你好吗;ss,ss; DDDD ; adkfafla；iskskdidi"

	sa := ezgo.StringSplits(s, []string{";", "；", ","})

	for _, v := range sa {
		fmt.Println(v)
	}

}

func test_pointer() {

	array := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}

	mask := []byte{0xff, 0x3f, 0x2f, 0x1f}
	for i := 0; i < 8; i += 4 {
		array[i] ^= mask[0]
		array[i+1] ^= mask[1]
		array[i+2] ^= mask[2]
		array[i+3] ^= mask[3]
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("%c ", array[i])
	}
	for i := 0; i < 8; i += 4 {
		array[i] ^= mask[0]
		array[i+1] ^= mask[1]
		array[i+2] ^= mask[2]
		array[i+3] ^= mask[3]
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("%c ", array[i])
	}

}

func main() {
	//pngreader()
	//test_02()
	//test_split()
	test_pointer()

}
