package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("code").Default(""),
		field.String("username").Default(""),
		field.String("password").Default(""),
		field.String("salt").Default(""),
		field.String("realname").Default(""),
		field.String("avatar").Default(""),
		field.Int("flight_cert_type").Default(0),
		field.Int64("org_id").Optional().Default(0),
		field.Int64("deadline").Default(0),
		field.String("remark").Default(""),
		field.Int64("login_at").Default(0),
		field.Other("login_lng", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,8)",
		}).Default(decimal.Zero),
		field.Other("login_lat", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(12,8)",
		}).Default(decimal.Zero),
		field.String("login_model").Default(""),
		field.String("login_device").Default(""),
		field.String("login_token").Default(""),
		field.Int64("created_at").Immutable().Default(0),
		field.Int64("updated_at").Default(0),
		field.Int64("deleted_at").Default(0),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
