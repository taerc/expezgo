package main

import (
	"context"
	"expezgo/modules/ent"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"github.com/taerc/ezgo"
	"github.com/taerc/ezgo/conf"
	"github.com/tealeg/xlsx"
)

func DBD() *ent.Client {

	drv, e := ezgo.EntDBDriver("mysql")

	if e != nil {
		fmt.Println(e)
		return nil
	}

	c := ent.NewClient(ent.Driver(dialect.DebugWithContext(drv, func(ctx context.Context, a ...any) {
	})))

	return c
}

func main() {

	c := conf.MySQLConf{
		MySQLHostname: "127.0.0.1",
		MySQLPort:     "3306",
		MySQLUserName: "wp",
		MySQLPass:     "wORd@2314",
		MySQLDBName:   "buckets",
		Charset:       "utf8mb4",
		ParseTime:     "true",
		Loc:           "Local",
		MaxLifeTime:   "1h",
	}

	//ezgo.WithComponentMySQL("mysql", &c)

	ezgo.LoadComponent(ezgo.WithComponentMySQL("mysql", &c))
	ctx := context.Background()
	cli := DBD()

	defectsDictPath := "/Users/rotaercw/Downloads/巡视标准与缺陷标准对照表.xlsx"
	file, xlsxerr := xlsx.OpenFile(defectsDictPath) //+xlsxfilename)

	if xlsxerr != nil {
		fmt.Println(xlsxerr)
	}
	sheet := file.Sheets[2]
	fmt.Println(sheet.Name)

	for i := 1; i < sheet.MaxRow; i++ {

		row := sheet.Row(i)
		if row == nil {
			continue
		}
		// for j := 0; j < sheet.MaxCol; j++ {
		// 	fmt.Println(fmt.Sprintf("%d:%s", j, row.Cells[j].Value))

		// }
		n := time.Now().Unix()
		if _, e := cli.Debug().GSGongFuDefectsDict.Create().SetObjID(row.Cells[0].String()).
			SetSblxbm(row.Cells[1].String()).
			SetSblxmc(row.Cells[2].String()).
			SetZsblxbm(row.Cells[3].String()).
			SetZsblxmc(row.Cells[4].String()).
			SetBjsblxbm(row.Cells[5].String()).
			SetBjsblxmc(row.Cells[6].String()).
			SetBwbm(row.Cells[7].String()).
			SetBwmc(row.Cells[8].String()).
			SetQxmsbm(row.Cells[9].String()).
			SetQxms(row.Cells[10].String()).
			SetFlyjbm(row.Cells[11].String()).
			SetFlyj(row.Cells[12].String()).
			SetQxxzbm(row.Cells[13].String()).
			SetQxxzmc(row.Cells[14].String()).SetCreatedAt(n).SetUpdatedAt(n).SetDeletedAt(0).Save(ctx); e != nil {
			fmt.Println(e)
		}
		// for j:=0; j < sheet.MaxCol; j++ { }

	}

}
