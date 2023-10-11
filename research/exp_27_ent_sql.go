package main

import (
	"context"
	"expezgo/modules/ent"
	"expezgo/modules/ent/city"
	"expezgo/modules/ent/county"
	"expezgo/modules/ent/licence"
	"expezgo/modules/ent/province"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/taerc/ezgo"
	"github.com/taerc/ezgo/conf"
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

func RawSqlExample(ctx context.Context, client *ent.Client) {

	rows, e := client.QueryContext(ctx, "select `id`, `name`, `type` from `province`")

	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer rows.Close()

	type Pros struct {
		ID   uint32 `sql:"id"`
		Name string `sql:"name"`
		Type uint32 `sql:"type"`
	}

	provinces := make([]Pros, 0)

	for rows.Next() {

		var p Pros
		if se := rows.Scan(&p.ID, &p.Name, &p.Type); se == nil {
			provinces = append(provinces, p)
		} else {
			fmt.Println(se.Error())
			continue
		}
	}

	fmt.Println("Result:")
	fmt.Println(fmt.Sprintf("Len [%d]", len(provinces)))

	for _, p := range provinces {
		fmt.Println(p.ID, p.Name)
	}

}

func entLog(ctx context.Context, args ...any) {
	ezgo.Info(nil, "ENT", args[0].(string))
}

func EdgeQuery(ctx context.Context, client *ent.Client) {

	cites, e := client.City.Query().Where(city.Pid(2)).WithProvinces().All(ctx)

	if e != nil {
		fmt.Println(e)
	}

	for c, v := range cites {
		fmt.Println(c)
		fmt.Println(v.Edges.Provinces.Name)
		fmt.Println(v.Name)
	}

	fmt.Println("===")
	provices, e := client.Province.Query().Where(province.ID(2)).WithCities().All(ctx)

	if e != nil {
		fmt.Println(e)
	}

	for c, p := range provices {
		fmt.Println(c)
		fmt.Println(p.Name)
		for i, c := range p.Edges.Cities {
			fmt.Println(i, c.Name)
		}
	}

}

func DyTableName(ctx context.Context, client *ent.Client) {

	// original
	cities, e := client.City.Query().All(ctx)
	if e != nil {
		fmt.Println(e)
	}
	for i, c := range cities {
		fmt.Println(i, c.Name)
	}

	// dyname
	table := "city_010"
	cities, e = client.City.Query().Select(city.FieldID).Where(func(s *sql.Selector) {
		table := sql.Table(table)
		s.From(table)
	}).All(ctx)
	if e != nil {
		fmt.Println(e)
	}
	for i, c := range cities {
		fmt.Println(i, c.Name)
	}

}

func DBD() *ent.Client {

	drv, e := ezgo.EntDBDriver("mysql")

	if e != nil {
		fmt.Println(e)
		return nil
	}

	c := ent.NewClient(ent.Driver(dialect.DebugWithContext(drv, func(ctx context.Context, a ...any) {
	})))

	return c.Debug()
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

	//client, err := ent.Open("mysql", "wp:wORd@2314@tcp(127.0.0.1:3306)/buckets?parseTime=True")
	//if err != nil {
	//	log.Fatalf("failed opening connection to mysql: %v", err)
	//}
	//defer client.Close()
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

	//Query3TbLeftJoin(ctx, client)
	//time.Sleep(3 * time.Second)

	// RawSqlExample(ctx, DBD())

	// edges

	// EdgeQuery(ctx, DBD())

	DyTableName(ctx, DBD())

}
