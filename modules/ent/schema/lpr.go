package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// LPR holds the schema definition for the LPR entity.
type LPR struct {
	ent.Schema
}

func (LPR) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "lpr"},
	}

}

// Fields of the LPR.
func (LPR) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique(),
		//field.Float32("one_year").Default(0),
		//field.Float32("five_year").Default(0),
		field.Other("one_year", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,5)",
		}).Default(decimal.Zero),

		field.Other("five_year", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,5)",
		}).Default(decimal.Zero),
		field.String("date").MaxLen(32).Default("2023-07-31"),
		field.Int64("create_at").Immutable().Default(0),
		field.Int64("update_at").Default(0),
	}
}

// Edges of the LPR.
func (LPR) Edges() []ent.Edge {
	return nil
}
