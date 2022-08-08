// Code generated by ent, DO NOT EDIT.

package pets

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/contrib/entoas/internal/pets/pet"
	"entgo.io/contrib/entoas/internal/pets/user"
	"entgo.io/ent/dialect/sql"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Nicknames holds the value of the "nicknames" field.
	Nicknames []string `json:"nicknames,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges     PetEdges `json:"edges"`
	user_pets *int
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*Pet `json:"friends,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e PetEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[0] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e PetEdges) FriendsOrErr() ([]*Pet, error) {
	if e.loadedTypes[2] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pet.FieldNicknames:
			values[i] = new([]byte)
		case pet.FieldID, pet.FieldAge:
			values[i] = new(sql.NullInt64)
		case pet.FieldName:
			values[i] = new(sql.NullString)
		case pet.ForeignKeys[0]: // user_pets
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case pet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case pet.FieldNicknames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field nicknames", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Nicknames); err != nil {
					return fmt.Errorf("unmarshal field nicknames: %w", err)
				}
			}
		case pet.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pe.Age = int(value.Int64)
			}
		case pet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_pets", value)
			} else if value.Valid {
				pe.user_pets = new(int)
				*pe.user_pets = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCategories queries the "categories" edge of the Pet entity.
func (pe *Pet) QueryCategories() *CategoryQuery {
	return (&PetClient{config: pe.config}).QueryCategories(pe)
}

// QueryOwner queries the "owner" edge of the Pet entity.
func (pe *Pet) QueryOwner() *UserQuery {
	return (&PetClient{config: pe.config}).QueryOwner(pe)
}

// QueryFriends queries the "friends" edge of the Pet entity.
func (pe *Pet) QueryFriends() *PetQuery {
	return (&PetClient{config: pe.config}).QueryFriends(pe)
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("pets: Pet is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("nicknames=")
	builder.WriteString(fmt.Sprintf("%v", pe.Nicknames))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", pe.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
