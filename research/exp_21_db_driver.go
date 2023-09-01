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

	db, err := sql.Open("fakedb", "mydb://dalong@127.0.0.1/demoapp")
	if err != nil {
		fmt.Printf("some error %s", err.Error())
	}
	// _, err = db.Query("select name, age, version from demoapp", "1 ", "2")
	_, err = db.QueryContext(ctx, "select * from demoapp", "1", "2")
	if err != nil {
		log.Fatal("some wrong for query", err.Error())
	}

}
