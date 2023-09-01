package fakedb

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
)

// Conn for db open
type Conn struct {
}

// Prepare statement for prepare exec
func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return &MyStmt{}, nil
}

// Close close db connection
func (c *Conn) Close() error {
	return errors.New("can't close connection")
}

// Begin begin
func (c *Conn) Begin() (driver.Tx, error) {
	return nil, errors.New("not support tx")
}

func (c *Conn) Query(query string, args []driver.Value) (driver.Result, error) {

	fmt.Println("query", query)
	fmt.Println("args", args)

	return nil, nil
}

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	fmt.Println("context")
	fmt.Println("query", query)
	fmt.Println("args", args)

	myrows := MyRowS{
		Size: 3,
	}
	return &myrows, nil
}

func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {

	fmt.Println("exec Context")
	fmt.Println("query", query)
	fmt.Println("args", args)

	return nil, nil
}
