// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TTaskImage struct {
	ent.Schema
}

func (TTaskImage) Fields() []ent.Field {
	return []ent.Field{field.Int("id"),
 field.Int("taskId"),
 field.Int("towerId"),
 field.Int("imgId")}
}
func (TTaskImage) Edges() []ent.Edge {
	return nil
}
func (TTaskImage) Annotations() []schema.Annotation {
	return nil
}
