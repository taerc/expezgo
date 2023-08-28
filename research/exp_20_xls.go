package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	defectsDictPath := "/Users/rotaercw/Downloads/巡视标准与缺陷标准对照表.xlsx"
	file, xlsxerr := xlsx.OpenFile(defectsDictPath) //+xlsxfilename)

	if xlsxerr != nil {
		fmt.Println(xlsxerr)
	}
	sheet := file.Sheets[2]
	fmt.Println(sheet.Name)

	for i := 0; i < sheet.MaxRow; i++ {

		row := sheet.Row(i)
		if row == nil {
			continue
		}
		for j := 0; j < sheet.MaxCol; j++ {
			fmt.Println(fmt.Sprintf("%d:%s", j, row.Cells[j].Value))
		}

		break
		// for j:=0; j < sheet.MaxCol; j++ { }

	}

}
