package fakedb

import (
	"context"
	"database/sql/driver"
	"errors"
	"log"
)

// MyStmt for sql statement
type MyStmt struct {
}

// Close  implement for stmt
func (stmt *MyStmt) Close() error {
	return nil
}

// Query  implement for Query
func (stmt *MyStmt) Query(args []driver.Value) (driver.Rows, error) {
	log.Println("do query", args)
	myrows := MyRowS{
		Size: 3,
	}
	return &myrows, nil
}

// NumInput row numbers
func (stmt *MyStmt) NumInput() int {
	// don't know how many row numbers
	return -1
}

// Exec exec  implement
func (stmt *MyStmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.New("some wrong")
}

func (stmt *MyStmt) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	log.Println("do query context", args)
	log.Println("do query context", query)
	return nil, nil
}
