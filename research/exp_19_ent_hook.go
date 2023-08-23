package main

import (
	"context"
	"expezgo/modules/ent"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/shopspring/decimal"
	"github.com/taerc/ezgo"
	"github.com/taerc/ezgo/conf"
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

func AddUser(ctx context.Context, c *ent.Client) {
	_, e := c.Debug().User.Create().SetUsername("wangfangming").SetPassword("123456").
		SetRealname("wfm").
		SetLoginLat(decimal.Decimal{}).Save(ctx)
	if e != nil {
		fmt.Println(e.Error())
	}
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
	}

	//ezgo.WithComponentMySQL("mysql", &c)

	ezgo.LoadComponent(ezgo.WithComponentMySQL("mysql", &c))

	//client, err := ent.Open("mysql", "wp:wORd@2314@tcp(127.0.0.1:3306)/buckets?parseTime=True")
	//if err != nil {
	//	log.Fatalf("failed opening connection to mysql: %v", err)
	//}
	//defer client.Close()
	//client.Debug()
	ctx := context.Background()

	AddUser(ctx, DBD())

}
