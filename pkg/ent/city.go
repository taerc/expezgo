// Code generated by ent, DO NOT EDIT.

package ent

import (
	"expezgo/pkg/ent/city"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// City is the model entity for the City schema.
type City struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type uint32 `json:"type,omitempty"`
	// Pid holds the value of the "pid" field.
	Pid          uint32 `json:"pid,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*City) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case city.FieldID, city.FieldType, city.FieldPid:
			values[i] = new(sql.NullInt64)
		case city.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the City fields.
func (c *City) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case city.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint32(value.Int64)
		case city.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case city.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = uint32(value.Int64)
			}
		case city.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				c.Pid = uint32(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the City.
// This includes values selected through modifiers, order, etc.
func (c *City) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this City.
// Note that you need to call City.Unwrap() before calling this method if this City
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *City) Update() *CityUpdateOne {
	return NewCityClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the City entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *City) Unwrap() *City {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: City is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *City) String() string {
	var builder strings.Builder
	builder.WriteString("City(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", ")
	builder.WriteString("pid=")
	builder.WriteString(fmt.Sprintf("%v", c.Pid))
	builder.WriteByte(')')
	return builder.String()
}

// Cities is a parsable slice of City.
type Cities []*City
