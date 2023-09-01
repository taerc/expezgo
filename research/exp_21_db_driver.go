package main

import (
	"context"
	"database/sql"
	"expezgo/modules/fakedb"
	"fmt"
	"log"
)

func init() {

	sql.Register("fakedb", &fakedb.Driver{})
}

func main() {

	ctx := context.Background()

	type User struct {
		Name    string `sql:"name"`
		Age     int
		Version string
	}
	var u User

	db, err := sql.Open("fakedb", "mydb://dalong@127.0.0.1/demoapp")
	if err != nil {
		fmt.Printf("some error %s", err.Error())
	}
	// _, err = db.Query("select name, age, version from demoapp", "1 ", "2")
	rows, err := db.QueryContext(ctx, "select * from demoapp", "1", "2")
	fmt.Println(rows)
	fmt.Println(rows.Columns())
	for rows.Next() {
		rows.Scan(&u.Name, &u.Age, &u.Version)
		fmt.Println(u.Name)
		fmt.Println(u.Age)
		fmt.Println(u.Version)
	}
	// _, err = db.ExecContext(ctx, "insert hello world", "1", "2")
	if err != nil {
		log.Fatal("some wrong for query", err.Error())
	}

}
