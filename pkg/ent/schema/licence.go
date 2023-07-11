package schema

import (
"entgo.io/ent"
"entgo.io/ent/dialect/entsql"
"entgo.io/ent/schema"
"entgo.io/ent/schema/field"
)

// Licence holds the schema definition for the Licence entity.
type Licence struct {
	ent.Schema
}

func (Licence) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entsql.Annotation{Table: "lic_sn"},
	}
}

// Fields of the Licence.
func (Licence) Fields() []ent.Field {
	// Fields of the Licence.
	return []ent.Field{
		field.Int64("id"),
		field.String("dev_uuid").MaxLen(64),
		field.String("lic_path").MaxLen(1024),
		field.Int("state").Default(0),
		field.String("task_id").Default("").MaxLen(64),
		field.Int64("create_time"),
	}
}

// Edges of the Licence.
func (Licence) Edges() []ent.Edge {
	return nil
}
