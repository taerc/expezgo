package main

import (
	"fmt"
	// "github.com/skip2/go-qrcode"
	"io/ioutil"
)

func loadFileAsString(name string) string {

	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(data)
}

type dataMode uint8

const (
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	dataModeNone dataMode = 1 << iota
	dataModeNumeric
	dataModeAlphanumeric
	dataModeByte
)
const (
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	TdataModeNone int = 10 + iota
	TdataModeNumeric
	TdataModeAlphanumeric
	TdataModeByte
)

func main() {
	//data := utils.GetRandomStr(1024)
	//data = "https://www.baidu.com?" + data

	//text_json := "/Users/rotaercw/Desktop/二维码收资.txt"
	txt_csv := "/Users/rotaercw/Desktop/1.csv"

	data := loadFileAsString(txt_csv)
	fmt.Println(data)
	// data = "https://www.baidu.com?" + data
	// fmt.Println(len(data))
	// err := qrcode.WriteFile(data, qrcode.Medium, 1080, "qr.png")
	// fmt.Println(err)

	fmt.Println(dataModeNone)
	fmt.Println(dataModeNumeric)
	fmt.Println(dataModeAlphanumeric)
	fmt.Println(dataModeByte)

}
