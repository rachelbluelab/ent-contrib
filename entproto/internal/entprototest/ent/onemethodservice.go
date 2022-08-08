// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/onemethodservice"
	"entgo.io/ent/dialect/sql"
)

// OneMethodService is the model entity for the OneMethodService schema.
type OneMethodService struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OneMethodService) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case onemethodservice.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OneMethodService", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OneMethodService fields.
func (oms *OneMethodService) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case onemethodservice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oms.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this OneMethodService.
// Note that you need to call OneMethodService.Unwrap() before calling this method if this OneMethodService
// was returned from a transaction, and the transaction was committed or rolled back.
func (oms *OneMethodService) Update() *OneMethodServiceUpdateOne {
	return (&OneMethodServiceClient{config: oms.config}).UpdateOne(oms)
}

// Unwrap unwraps the OneMethodService entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oms *OneMethodService) Unwrap() *OneMethodService {
	_tx, ok := oms.config.driver.(*txDriver)
	if !ok {
		panic("ent: OneMethodService is not a transactional entity")
	}
	oms.config.driver = _tx.drv
	return oms
}

// String implements the fmt.Stringer.
func (oms *OneMethodService) String() string {
	var builder strings.Builder
	builder.WriteString("OneMethodService(")
	builder.WriteString(fmt.Sprintf("id=%v", oms.ID))
	builder.WriteByte(')')
	return builder.String()
}

// OneMethodServices is a parsable slice of OneMethodService.
type OneMethodServices []*OneMethodService

func (oms OneMethodServices) config(cfg config) {
	for _i := range oms {
		oms[_i].config = cfg
	}
}
