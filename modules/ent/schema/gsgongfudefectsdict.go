package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GSGongFuDefectsDict holds the schema definition for the GSGongFuDefectsDict entity.
type GSGongFuDefectsDict struct {
	ent.Schema
}

func (GSGongFuDefectsDict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "gs_gongfu_defects_dictory"},
	}
}

// Fields of the GSGongFuDefectsDict.
func (GSGongFuDefectsDict) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("obj_id"),
		field.String("sblxbm"),
		field.String("sblxmc"),
		field.String("zsblxbm"),
		field.String("zsblxmc"),
		field.String("bjsblxbm"),
		field.String("bjsblxmc"),
		field.String("bwbm"),
		field.String("bwmc"),
		field.String("qxmsbm"),
		field.String("qxms"),
		field.String("flyjbm"),
		field.String("flyj"),
		field.String("qxxzbm"),
		field.String("qxxzmc"),
		field.Int64("created_at"),
		field.Int64("updated_at"),
		field.Int64("deleted_at"),
	}
}

// Edges of the GSGongFuDefectsDict.
func (GSGongFuDefectsDict) Edges() []ent.Edge {
	return nil
}
