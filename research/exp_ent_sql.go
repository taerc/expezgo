package main

import (
	"context"
	"expezgo/pkg/ent"
	"expezgo/pkg/ent/licence"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func QueryLicence(ctx context.Context, client *ent.Client) (*ent.Licence, error) {
	u, e := client.Licence.Query().Where(licence.State(1)).Only(ctx)
	if e != nil {
		return nil, fmt.Errorf("failed queyring licence：%w", e)
	}
	return u, nil
}

func main() {

	client, err := ent.Open("mysql", "wp:wORd@2314@tcp(127.0.0.1:3306)/db_licence?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	//// Run the auto migration tool.
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	l, _ := QueryLicence(ctx, client)

	fmt.Println(l)
}
