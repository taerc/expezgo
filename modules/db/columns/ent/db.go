package ent

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/taerc/ezgo"
)


var db *Client = nil

func DB() *Client {

	if db != nil {
		return db
	}
	drv, e := ezgo.EntDBDriver(ezgo.Default)
	if e != nil {
		return nil
	}
	db = NewClient(Driver(dialect.DebugWithContext(drv, func(ctx context.Context, a ...any) {
	})))
	return db

}
