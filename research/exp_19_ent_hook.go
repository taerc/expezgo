package main

import (
	"context"
	"expezgo/modules/ent"
	"expezgo/modules/ent/hook"
	"expezgo/modules/ent/user"
	"fmt"

	"entgo.io/ent/dialect"
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

	// c.Use(func(next ent.Mutator) ent.Mutator {
	// 	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	// 		fmt.Println("mutate", m.Type())
	// 		if e := m.SetField("created_at", ezgo.GetUnixTimeStamp()); e != nil {
	// 			fmt.Println("set time error ", e.Error())
	// 		}
	// 		return next.Mutate(ctx, m)
	// 	})
	// })

	c.Intercept(ent.InterceptFunc(func(q ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			fmt.Println("query ..... ....")
			return q.Query(ctx, query)
		})
	}))
	c.Use(hook.On(updateTimestamp(), ent.OpUpdate|ent.OpUpdateOne))

	// ent.InterceptFunc(func(next ent.Querier) ent.Querier {
	// 	return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
	// 		// Do something before the query execution.
	// 		fmt.Println("this is BEFER ..")
	// 		value, err := next.Query(ctx, query)
	// 		// Do something after the query execution.
	// 		fmt.Println("this is END ...")
	// 		return value, err
	// 	})
	// })
	return c
}

// func updateTimestamp(next ent.Mutator) ent.Mutator {
// 	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
// 		if e := m.SetField("created_at", ezgo.GetUnixTimeStamp()); e != nil {
// 			fmt.Println("set time error ", e.Error())
// 		}
// 		return next.Mutate(ctx, m)
// 	})
// }

func updateTimestamp() func(next ent.Mutator) ent.Mutator {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if e := m.SetField("updated_at", ezgo.GetUnixTimeStamp()); e != nil {
				fmt.Println("set time error ", e.Error())
			}
			return next.Mutate(ctx, m)
		})
	}
}

func AddUser(ctx context.Context, c *ent.Client) {

	_, e := c.Debug().User.Query().All(ctx)
	if e != nil {
		fmt.Println(e.Error())
	}
}

func UpdateUser(ctx context.Context, c *ent.Client) {

	if n, e := c.User.Update().Where(user.Username("wangfangming")).SetUsername("fangming").Save(ctx); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(n)
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

	AddUser(ctx, DBD())

	// UpdateUser(ctx, DBD())

}
