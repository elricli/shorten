// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/drrrMikado/shorten/ent/schema"
	"github.com/drrrMikado/shorten/ent/shorturl"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	shorturlFields := schema.ShortUrl{}.Fields()
	_ = shorturlFields
	// shorturlDescKey is the schema descriptor for key field.
	shorturlDescKey := shorturlFields[0].Descriptor()
	// shorturl.DefaultKey holds the default value on creation for the key field.
	shorturl.DefaultKey = shorturlDescKey.Default.(string)
	// shorturl.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	shorturl.KeyValidator = shorturlDescKey.Validators[0].(func(string) error)
	// shorturlDescShortURL is the schema descriptor for short_url field.
	shorturlDescShortURL := shorturlFields[1].Descriptor()
	// shorturl.DefaultShortURL holds the default value on creation for the short_url field.
	shorturl.DefaultShortURL = shorturlDescShortURL.Default.(string)
	// shorturl.ShortURLValidator is a validator for the "short_url" field. It is called by the builders before save.
	shorturl.ShortURLValidator = shorturlDescShortURL.Validators[0].(func(string) error)
	// shorturlDescLongURL is the schema descriptor for long_url field.
	shorturlDescLongURL := shorturlFields[2].Descriptor()
	// shorturl.DefaultLongURL holds the default value on creation for the long_url field.
	shorturl.DefaultLongURL = shorturlDescLongURL.Default.(string)
	// shorturl.LongURLValidator is a validator for the "long_url" field. It is called by the builders before save.
	shorturl.LongURLValidator = shorturlDescLongURL.Validators[0].(func(string) error)
	// shorturlDescCreateAt is the schema descriptor for create_at field.
	shorturlDescCreateAt := shorturlFields[3].Descriptor()
	// shorturl.DefaultCreateAt holds the default value on creation for the create_at field.
	shorturl.DefaultCreateAt = shorturlDescCreateAt.Default.(func() time.Time)
	// shorturlDescUpdateAt is the schema descriptor for update_at field.
	shorturlDescUpdateAt := shorturlFields[4].Descriptor()
	// shorturl.DefaultUpdateAt holds the default value on creation for the update_at field.
	shorturl.DefaultUpdateAt = shorturlDescUpdateAt.Default.(func() time.Time)
}
