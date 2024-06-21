// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// Example1sColumns holds the columns for the "example1s" table.
	Example1sColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "example_value_1", Type: field.TypeString},
	}
	// Example1sTable holds the schema information for the "example1s" table.
	Example1sTable = &schema.Table{
		Name:       "example1s",
		Columns:    Example1sColumns,
		PrimaryKey: []*schema.Column{Example1sColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		Example1sTable,
	}
)

func init() {
}
