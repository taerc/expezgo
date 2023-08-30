// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TLineImage struct {
	ent.Schema
}

func (TLineImage) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("lineId").Comment("线路ID"), field.Int("towerId").Comment("杆塔ID"), field.Int("taskId").Comment("任务ID"), field.Int("imgId").Comment("图片ID")}
}
func (TLineImage) Edges() []ent.Edge {
	return nil
}
func (TLineImage) Annotations() []schema.Annotation {
	return nil
}
