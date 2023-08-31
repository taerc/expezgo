// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TPlatformAccount struct {
	ent.Schema
}

func (TPlatformAccount) Fields() []ent.Field {
	return []ent.Field{field.Int("id"),
 field.Int("orgId").Comment("所属组织单位ID"),
 field.String("orgName").Comment("所属组织单位名称"),
 field.Int("operaId").Comment("运维班组ID"),
 field.String("operaName").Comment("运维班组名称"),
 field.String("account").Comment("账号名称（限制20个字符）"),
 field.String("userName").Comment("用户名"),
 field.String("passwd").Comment("密码（明文限制10个字符）"),
 field.String("salt").Comment("加密盐"),
 field.Bool("userRole").Comment("用户角色"),
 field.String("token").Comment("token"),
 field.Time("loginTime").Comment("登录时间"),
 field.Time("deadLineTime").Comment("有效期"),
 field.Time("updateTime").Comment("更新时间"),
 field.Time("createTime").Comment("创建时间")}
}
func (TPlatformAccount) Edges() []ent.Edge {
	return nil
}
func (TPlatformAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "t_platform_account"}}
}
