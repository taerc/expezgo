// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TImage struct {
	ent.Schema
}

func (TImage) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.String("imgName"), field.Int32("imgWidth"), field.Int32("imgHeight"), field.Float("longtitude"), field.Float("latitude"), field.Float32("altitude"), field.Time("photoTime"), field.Int8("imgType"), field.String("thumbnailId"), field.String("mediaId"), field.String("comprId"), field.Int8("is_ai_check"), field.Int8("is_ok"), field.Int8("is_check"), field.Time("checkTime"), field.String("xmlPath"), field.Time("createTime")}
}
func (TImage) Edges() []ent.Edge {
	return nil
}
func (TImage) Annotations() []schema.Annotation {
	return nil
}
