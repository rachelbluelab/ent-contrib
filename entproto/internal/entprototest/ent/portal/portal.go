// Code generated by ent, DO NOT EDIT.

package portal

const (
	// Label holds the string label denoting the portal type in the database.
	Label = "portal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// Table holds the table name of the portal in the database.
	Table = "portals"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "portals"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "portal_category"
)

// Columns holds all SQL columns for portal fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "portals"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"portal_category",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
