// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TMedium struct {
	ent.Schema
}

func (TMedium) Fields() []ent.Field {
	return []ent.Field{field.Int("id"),
 field.String("bucket").Comment("Bucket NO"),
 field.String("mediaId").Comment("指纹(MD5)"),
 field.String("fileName").Comment("文件名称"),
 field.Int("size").Comment("文件大小（单位：b）"),
 field.String("format").Comment("图片Format"),
 field.Time("createTime").Comment("创建时间")}
}
func (TMedium) Edges() []ent.Edge {
	return nil
}
func (TMedium) Annotations() []schema.Annotation {
	return nil
}
