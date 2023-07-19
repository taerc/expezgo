// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"expezgo/pkg/ent/city"
	"expezgo/pkg/ent/county"
	"expezgo/pkg/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CountyUpdate is the builder for updating County entities.
type CountyUpdate struct {
	config
	hooks     []Hook
	mutation  *CountyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CountyUpdate builder.
func (cu *CountyUpdate) Where(ps ...predicate.County) *CountyUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CountyUpdate) SetName(s string) *CountyUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetType sets the "type" field.
func (cu *CountyUpdate) SetType(u uint32) *CountyUpdate {
	cu.mutation.ResetType()
	cu.mutation.SetType(u)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *CountyUpdate) SetNillableType(u *uint32) *CountyUpdate {
	if u != nil {
		cu.SetType(*u)
	}
	return cu
}

// AddType adds u to the "type" field.
func (cu *CountyUpdate) AddType(u int32) *CountyUpdate {
	cu.mutation.AddType(u)
	return cu
}

// SetPid sets the "pid" field.
func (cu *CountyUpdate) SetPid(u uint32) *CountyUpdate {
	cu.mutation.SetPid(u)
	return cu
}

// SetNillablePid sets the "pid" field if the given value is not nil.
func (cu *CountyUpdate) SetNillablePid(u *uint32) *CountyUpdate {
	if u != nil {
		cu.SetPid(*u)
	}
	return cu
}

// ClearPid clears the value of the "pid" field.
func (cu *CountyUpdate) ClearPid() *CountyUpdate {
	cu.mutation.ClearPid()
	return cu
}

// SetCityID sets the "city" edge to the City entity by ID.
func (cu *CountyUpdate) SetCityID(id uint32) *CountyUpdate {
	cu.mutation.SetCityID(id)
	return cu
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (cu *CountyUpdate) SetNillableCityID(id *uint32) *CountyUpdate {
	if id != nil {
		cu = cu.SetCityID(*id)
	}
	return cu
}

// SetCity sets the "city" edge to the City entity.
func (cu *CountyUpdate) SetCity(c *City) *CountyUpdate {
	return cu.SetCityID(c.ID)
}

// Mutation returns the CountyMutation object of the builder.
func (cu *CountyUpdate) Mutation() *CountyMutation {
	return cu.mutation
}

// ClearCity clears the "city" edge to the City entity.
func (cu *CountyUpdate) ClearCity() *CountyUpdate {
	cu.mutation.ClearCity()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CountyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CountyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CountyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CountyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CountyUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := county.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "County.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CountyUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CountyUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CountyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(county.Table, county.Columns, sqlgraph.NewFieldSpec(county.FieldID, field.TypeUint32))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(county.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(county.FieldType, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedType(); ok {
		_spec.AddField(county.FieldType, field.TypeUint32, value)
	}
	if cu.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   county.CityTable,
			Columns: []string{county.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   county.CityTable,
			Columns: []string{county.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{county.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CountyUpdateOne is the builder for updating a single County entity.
type CountyUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CountyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (cuo *CountyUpdateOne) SetName(s string) *CountyUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetType sets the "type" field.
func (cuo *CountyUpdateOne) SetType(u uint32) *CountyUpdateOne {
	cuo.mutation.ResetType()
	cuo.mutation.SetType(u)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *CountyUpdateOne) SetNillableType(u *uint32) *CountyUpdateOne {
	if u != nil {
		cuo.SetType(*u)
	}
	return cuo
}

// AddType adds u to the "type" field.
func (cuo *CountyUpdateOne) AddType(u int32) *CountyUpdateOne {
	cuo.mutation.AddType(u)
	return cuo
}

// SetPid sets the "pid" field.
func (cuo *CountyUpdateOne) SetPid(u uint32) *CountyUpdateOne {
	cuo.mutation.SetPid(u)
	return cuo
}

// SetNillablePid sets the "pid" field if the given value is not nil.
func (cuo *CountyUpdateOne) SetNillablePid(u *uint32) *CountyUpdateOne {
	if u != nil {
		cuo.SetPid(*u)
	}
	return cuo
}

// ClearPid clears the value of the "pid" field.
func (cuo *CountyUpdateOne) ClearPid() *CountyUpdateOne {
	cuo.mutation.ClearPid()
	return cuo
}

// SetCityID sets the "city" edge to the City entity by ID.
func (cuo *CountyUpdateOne) SetCityID(id uint32) *CountyUpdateOne {
	cuo.mutation.SetCityID(id)
	return cuo
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (cuo *CountyUpdateOne) SetNillableCityID(id *uint32) *CountyUpdateOne {
	if id != nil {
		cuo = cuo.SetCityID(*id)
	}
	return cuo
}

// SetCity sets the "city" edge to the City entity.
func (cuo *CountyUpdateOne) SetCity(c *City) *CountyUpdateOne {
	return cuo.SetCityID(c.ID)
}

// Mutation returns the CountyMutation object of the builder.
func (cuo *CountyUpdateOne) Mutation() *CountyMutation {
	return cuo.mutation
}

// ClearCity clears the "city" edge to the City entity.
func (cuo *CountyUpdateOne) ClearCity() *CountyUpdateOne {
	cuo.mutation.ClearCity()
	return cuo
}

// Where appends a list predicates to the CountyUpdate builder.
func (cuo *CountyUpdateOne) Where(ps ...predicate.County) *CountyUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CountyUpdateOne) Select(field string, fields ...string) *CountyUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated County entity.
func (cuo *CountyUpdateOne) Save(ctx context.Context) (*County, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CountyUpdateOne) SaveX(ctx context.Context) *County {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CountyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CountyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CountyUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := county.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "County.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CountyUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CountyUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CountyUpdateOne) sqlSave(ctx context.Context) (_node *County, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(county.Table, county.Columns, sqlgraph.NewFieldSpec(county.FieldID, field.TypeUint32))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "County.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, county.FieldID)
		for _, f := range fields {
			if !county.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != county.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(county.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(county.FieldType, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedType(); ok {
		_spec.AddField(county.FieldType, field.TypeUint32, value)
	}
	if cuo.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   county.CityTable,
			Columns: []string{county.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   county.CityTable,
			Columns: []string{county.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &County{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{county.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
