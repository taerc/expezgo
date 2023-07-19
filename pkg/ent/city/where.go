// Code generated by ent, DO NOT EDIT.

package city

import (
	"expezgo/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.City {
	return predicate.City(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.City {
	return predicate.City(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.City {
	return predicate.City(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.City {
	return predicate.City(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.City {
	return predicate.City(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.City {
	return predicate.City(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.City {
	return predicate.City(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.City {
	return predicate.City(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.City {
	return predicate.City(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.City {
	return predicate.City(sql.FieldEQ(FieldName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v uint32) predicate.City {
	return predicate.City(sql.FieldEQ(FieldType, v))
}

// Pid applies equality check predicate on the "pid" field. It's identical to PidEQ.
func Pid(v uint32) predicate.City {
	return predicate.City(sql.FieldEQ(FieldPid, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.City {
	return predicate.City(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.City {
	return predicate.City(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.City {
	return predicate.City(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.City {
	return predicate.City(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.City {
	return predicate.City(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.City {
	return predicate.City(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.City {
	return predicate.City(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.City {
	return predicate.City(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.City {
	return predicate.City(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.City {
	return predicate.City(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.City {
	return predicate.City(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.City {
	return predicate.City(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.City {
	return predicate.City(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v uint32) predicate.City {
	return predicate.City(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v uint32) predicate.City {
	return predicate.City(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...uint32) predicate.City {
	return predicate.City(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...uint32) predicate.City {
	return predicate.City(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v uint32) predicate.City {
	return predicate.City(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v uint32) predicate.City {
	return predicate.City(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v uint32) predicate.City {
	return predicate.City(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v uint32) predicate.City {
	return predicate.City(sql.FieldLTE(FieldType, v))
}

// PidEQ applies the EQ predicate on the "pid" field.
func PidEQ(v uint32) predicate.City {
	return predicate.City(sql.FieldEQ(FieldPid, v))
}

// PidNEQ applies the NEQ predicate on the "pid" field.
func PidNEQ(v uint32) predicate.City {
	return predicate.City(sql.FieldNEQ(FieldPid, v))
}

// PidIn applies the In predicate on the "pid" field.
func PidIn(vs ...uint32) predicate.City {
	return predicate.City(sql.FieldIn(FieldPid, vs...))
}

// PidNotIn applies the NotIn predicate on the "pid" field.
func PidNotIn(vs ...uint32) predicate.City {
	return predicate.City(sql.FieldNotIn(FieldPid, vs...))
}

// PidIsNil applies the IsNil predicate on the "pid" field.
func PidIsNil() predicate.City {
	return predicate.City(sql.FieldIsNull(FieldPid))
}

// PidNotNil applies the NotNil predicate on the "pid" field.
func PidNotNil() predicate.City {
	return predicate.City(sql.FieldNotNull(FieldPid))
}

// HasProvinces applies the HasEdge predicate on the "provinces" edge.
func HasProvinces() predicate.City {
	return predicate.City(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProvincesTable, ProvincesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProvincesWith applies the HasEdge predicate on the "provinces" edge with a given conditions (other predicates).
func HasProvincesWith(preds ...predicate.Province) predicate.City {
	return predicate.City(func(s *sql.Selector) {
		step := newProvincesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCounties applies the HasEdge predicate on the "counties" edge.
func HasCounties() predicate.City {
	return predicate.City(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CountiesTable, CountiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCountiesWith applies the HasEdge predicate on the "counties" edge with a given conditions (other predicates).
func HasCountiesWith(preds ...predicate.County) predicate.City {
	return predicate.City(func(s *sql.Selector) {
		step := newCountiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.City) predicate.City {
	return predicate.City(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.City) predicate.City {
	return predicate.City(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.City) predicate.City {
	return predicate.City(func(s *sql.Selector) {
		p(s.Not())
	})
}
