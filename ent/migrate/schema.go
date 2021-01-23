// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// ShortUrlsColumns holds the columns for the "short_urls" table.
	ShortUrlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Default: ""},
		{Name: "short_url", Type: field.TypeString, Default: ""},
		{Name: "long_url", Type: field.TypeString, Default: ""},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// ShortUrlsTable holds the schema information for the "short_urls" table.
	ShortUrlsTable = &schema.Table{
		Name:        "short_urls",
		Columns:     ShortUrlsColumns,
		PrimaryKey:  []*schema.Column{ShortUrlsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ShortUrlsTable,
	}
)

func init() {
}