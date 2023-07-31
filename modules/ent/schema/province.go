package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Province holds the schema definition for the Province entity.
type Province struct {
	ent.Schema
}

func (Province) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "province"},
	}
}

// Fields of the Province.
func (Province) Fields() []ent.Field {
	return []ent.Field{field.Uint32("id").Unique(),
		field.String("name").MaxLen(256),
		field.Uint32("type").Default(0)}
}

// Edges of the Province.
func (Province) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cities", City.Type),
	}
}
