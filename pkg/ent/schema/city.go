package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

func (City) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "city"},
	}
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique(),
		field.String("name").MaxLen(256),
		field.Uint32("type").Default(0),
		field.Uint32("pid"),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provinces", Province.Type).Ref("cities"),
	}
}
