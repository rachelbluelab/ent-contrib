// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/validmessage"
	"entgo.io/ent/dialect/sql"
)

// ValidMessage is the model entity for the ValidMessage schema.
type ValidMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Ts holds the value of the "ts" field.
	Ts time.Time `json:"ts,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ValidMessage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case validmessage.FieldID:
			values[i] = &sql.NullInt64{}
		case validmessage.FieldName:
			values[i] = &sql.NullString{}
		case validmessage.FieldTs:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ValidMessage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ValidMessage fields.
func (vm *ValidMessage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case validmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vm.ID = int(value.Int64)
		case validmessage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				vm.Name = value.String
			}
		case validmessage.FieldTs:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ts", values[i])
			} else if value.Valid {
				vm.Ts = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ValidMessage.
// Note that you need to call ValidMessage.Unwrap() before calling this method if this ValidMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (vm *ValidMessage) Update() *ValidMessageUpdateOne {
	return (&ValidMessageClient{config: vm.config}).UpdateOne(vm)
}

// Unwrap unwraps the ValidMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vm *ValidMessage) Unwrap() *ValidMessage {
	tx, ok := vm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ValidMessage is not a transactional entity")
	}
	vm.config.driver = tx.drv
	return vm
}

// String implements the fmt.Stringer.
func (vm *ValidMessage) String() string {
	var builder strings.Builder
	builder.WriteString("ValidMessage(")
	builder.WriteString(fmt.Sprintf("id=%v", vm.ID))
	builder.WriteString(", name=")
	builder.WriteString(vm.Name)
	builder.WriteString(", ts=")
	builder.WriteString(vm.Ts.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ValidMessages is a parsable slice of ValidMessage.
type ValidMessages []*ValidMessage

func (vm ValidMessages) config(cfg config) {
	for _i := range vm {
		vm[_i].config = cfg
	}
}
