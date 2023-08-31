package main

import (
	"database/sql"
	"expezgo/modules/fakedb"
	"fmt"
	"log"
)

func init() {

	sql.Register("fakedb", &fakedb.Driver{})
}

func main() {

	db, err := sql.Open("fakedb", "mydb://dalong@127.0.0.1/demoapp")
	if err != nil {
		fmt.Printf("some error %s", err.Error())
	}
	_, err = db.Query("select name, age, version from demoapp", "1 ", "2")
	if err != nil {
		log.Fatal("some wrong for query", err.Error())
	}
	// for rows.Next() {
	// 	var user fakedb.MyUser
	// 	if err := rows.Scan(&user.Name, &user.Age, &user.Version); err != nil {
	// 		log.Println("scan value erro", err.Error())
	// 	} else {
	// 		log.Println(user)
	// 	}
	// }

}
