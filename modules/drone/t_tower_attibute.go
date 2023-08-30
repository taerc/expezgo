// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TTowerAttibute struct {
	ent.Schema
}

func (TTowerAttibute) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("towerId").Comment("杆塔ID"), field.String("towerNo").Comment("塔号"), field.Int("lineId").Comment("线路ID"), field.Int8("attibuteValue").Comment("枚举：0：默认；1：起始塔；2：终端塔；3：支出塔；4：交跨塔；5：台区塔；"), field.String("attibuteName").Comment("属性名称"), field.Time("updateTime").Comment("更新时间"), field.Time("createTime").Comment("更新值")}
}
func (TTowerAttibute) Edges() []ent.Edge {
	return nil
}
func (TTowerAttibute) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "t_tower_attibute"}}
}
