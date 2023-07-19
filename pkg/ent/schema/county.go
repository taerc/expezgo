package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// County holds the schema definition for the County entity.
type County struct {
	ent.Schema
}

func (County) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "county"},
	}
}

// Fields of the County.
func (County) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique(),
		field.String("name").MaxLen(256),
		field.Uint32("type").Default(0),
		field.Uint32("pid").Optional(),
	}
}

// Edges of the County.
func (County) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("city", City.Type).Field("pid").Ref("counties").Unique(),
	}
}
