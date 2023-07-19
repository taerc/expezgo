package main

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"expezgo/pkg/ent"
	"expezgo/pkg/ent/city"
	"expezgo/pkg/ent/county"
	"expezgo/pkg/ent/licence"
	"expezgo/pkg/ent/province"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/taerc/ezgo"
	"log"
)

func QueryLicence(ctx context.Context, client *ent.Client) (*ent.Licence, error) {
	u, e := client.Debug().Licence.Query().Where(licence.State(1)).Only(ctx)
	if e != nil {
		return nil, fmt.Errorf("failed queyring licence：%w", e)
	}
	return u, nil
}

func QueryProvince(ctx context.Context, client *ent.Client) ([]*ent.Province, error) {

	u, e := client.Debug().Province.Query().Where(province.ID(1)).WithCities(func(query *ent.CityQuery) {
		query.WithCounties()
	}).All(ctx)
	if e != nil {
		return nil, e
	}
	return u, nil
}

// QueryCityLeftJoin
// 测试两张表的 left join

func QueryCityLeftJoin(ctx context.Context, client *ent.Client) {

	var cities []struct {
		//ent.City
		ID           uint32 `sql:"id"`
		PID          uint32 `sql:"pid"`
		ProvinceName string `sql:"province_name"`
		CityName     string `sql:"city_name"`
	}

	client.Debug().City.Query().Select(city.FieldID).Modify(func(s *sql.Selector) {
		t := sql.Table(province.Table)
		s.LeftJoin(t).On(
			s.C(city.FieldPid),
			t.C(province.FieldID),
		).AppendSelect(
			sql.As(t.C(province.FieldName), "province_name"),
			sql.As(t.C(province.FieldID), "pid"),
			sql.As(s.C(city.FieldName), "city_name"),
		).Where(sql.EQ("pid", 1))
	}).ScanX(ctx, &cities)

	for _, d := range cities {
		fmt.Println(fmt.Sprintf("省 %s 市 %s ", d.ProvinceName, d.CityName))
	}
}

// Query3TbLeftJoin

func Query3TbLeftJoin(ctx context.Context, client *ent.Client) {
	var counties []struct {
		//ent.City
		ID uint32 `sql:"id"`
		//CID          uint32 `sql:"cid"`
		ProvinceName string `sql:"province_name"`
		CityName     string `sql:"city_name"`
		CountyName   string `sql:"county_name"`
	}

	client.Debug().City.Query().Select(city.FieldID).Modify(func(s *sql.Selector) {
		t := sql.Table(province.Table)
		c := sql.Table(county.Table)
		s.LeftJoin(t).On(
			s.C(city.FieldPid),
			t.C(province.FieldID),
		).AppendSelect(
			sql.As(t.C(province.FieldName), "province_name"),
			sql.As(s.C(city.FieldName), "city_name"),
		).Where(sql.EQ(t.C(province.FieldID), 1)).LeftJoin(c).On(
			s.C(city.FieldID),
			c.C(county.FieldPid)).AppendSelect(
			sql.As(c.C(county.FieldName), "county_name"),
		)
	}).ScanX(ctx, &counties)

	for _, d := range counties {
		fmt.Println(fmt.Sprintf("省 %s 市 %s 县 %s ", d.ProvinceName, d.CityName, d.CountyName))
	}

}

func entLog(ctx context.Context, args ...any) {
	ezgo.Info(nil, "ENT", args[0].(string))
}

func main() {

	client, err := ent.Open("mysql", "wp:wORd@2314@tcp(127.0.0.1:3306)/buckets?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	//client.Debug()
	ctx := context.Background()

	//// Run the auto migration tool.
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	//l, _ := QueryLicence(ctx, client)
	//l, _ := QueryProvince(ctx, client)
	//
	//fmt.Println()
	//fmt.Println("====")
	//if data, e := json.MarshalIndent(l, "", " "); e == nil {
	//	fmt.Println(string(data))
	//}

	//QueryCityLeftJoin(ctx, client)

	Query3TbLeftJoin(ctx, client)
}
