// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/drrrMikado/shorten/ent/shorturl"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ShortUrlCreate is the builder for creating a ShortUrl entity.
type ShortUrlCreate struct {
	config
	mutation *ShortUrlMutation
	hooks    []Hook
}

// SetKey sets the "key" field.
func (suc *ShortUrlCreate) SetKey(s string) *ShortUrlCreate {
	suc.mutation.SetKey(s)
	return suc
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (suc *ShortUrlCreate) SetNillableKey(s *string) *ShortUrlCreate {
	if s != nil {
		suc.SetKey(*s)
	}
	return suc
}

// SetShortURL sets the "short_url" field.
func (suc *ShortUrlCreate) SetShortURL(s string) *ShortUrlCreate {
	suc.mutation.SetShortURL(s)
	return suc
}

// SetNillableShortURL sets the "short_url" field if the given value is not nil.
func (suc *ShortUrlCreate) SetNillableShortURL(s *string) *ShortUrlCreate {
	if s != nil {
		suc.SetShortURL(*s)
	}
	return suc
}

// SetLongURL sets the "long_url" field.
func (suc *ShortUrlCreate) SetLongURL(s string) *ShortUrlCreate {
	suc.mutation.SetLongURL(s)
	return suc
}

// SetNillableLongURL sets the "long_url" field if the given value is not nil.
func (suc *ShortUrlCreate) SetNillableLongURL(s *string) *ShortUrlCreate {
	if s != nil {
		suc.SetLongURL(*s)
	}
	return suc
}

// SetCreateAt sets the "create_at" field.
func (suc *ShortUrlCreate) SetCreateAt(t time.Time) *ShortUrlCreate {
	suc.mutation.SetCreateAt(t)
	return suc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (suc *ShortUrlCreate) SetNillableCreateAt(t *time.Time) *ShortUrlCreate {
	if t != nil {
		suc.SetCreateAt(*t)
	}
	return suc
}

// SetUpdateAt sets the "update_at" field.
func (suc *ShortUrlCreate) SetUpdateAt(t time.Time) *ShortUrlCreate {
	suc.mutation.SetUpdateAt(t)
	return suc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (suc *ShortUrlCreate) SetNillableUpdateAt(t *time.Time) *ShortUrlCreate {
	if t != nil {
		suc.SetUpdateAt(*t)
	}
	return suc
}

// Mutation returns the ShortUrlMutation object of the builder.
func (suc *ShortUrlCreate) Mutation() *ShortUrlMutation {
	return suc.mutation
}

// Save creates the ShortUrl in the database.
func (suc *ShortUrlCreate) Save(ctx context.Context) (*ShortUrl, error) {
	var (
		err  error
		node *ShortUrl
	)
	suc.defaults()
	if len(suc.hooks) == 0 {
		if err = suc.check(); err != nil {
			return nil, err
		}
		node, err = suc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShortUrlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suc.check(); err != nil {
				return nil, err
			}
			suc.mutation = mutation
			node, err = suc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suc.hooks) - 1; i >= 0; i-- {
			mut = suc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (suc *ShortUrlCreate) SaveX(ctx context.Context) *ShortUrl {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (suc *ShortUrlCreate) defaults() {
	if _, ok := suc.mutation.Key(); !ok {
		v := shorturl.DefaultKey
		suc.mutation.SetKey(v)
	}
	if _, ok := suc.mutation.ShortURL(); !ok {
		v := shorturl.DefaultShortURL
		suc.mutation.SetShortURL(v)
	}
	if _, ok := suc.mutation.LongURL(); !ok {
		v := shorturl.DefaultLongURL
		suc.mutation.SetLongURL(v)
	}
	if _, ok := suc.mutation.CreateAt(); !ok {
		v := shorturl.DefaultCreateAt()
		suc.mutation.SetCreateAt(v)
	}
	if _, ok := suc.mutation.UpdateAt(); !ok {
		v := shorturl.DefaultUpdateAt()
		suc.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *ShortUrlCreate) check() error {
	if _, ok := suc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New("ent: missing required field \"key\"")}
	}
	if v, ok := suc.mutation.Key(); ok {
		if err := shorturl.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf("ent: validator failed for field \"key\": %w", err)}
		}
	}
	if _, ok := suc.mutation.ShortURL(); !ok {
		return &ValidationError{Name: "short_url", err: errors.New("ent: missing required field \"short_url\"")}
	}
	if v, ok := suc.mutation.ShortURL(); ok {
		if err := shorturl.ShortURLValidator(v); err != nil {
			return &ValidationError{Name: "short_url", err: fmt.Errorf("ent: validator failed for field \"short_url\": %w", err)}
		}
	}
	if _, ok := suc.mutation.LongURL(); !ok {
		return &ValidationError{Name: "long_url", err: errors.New("ent: missing required field \"long_url\"")}
	}
	if v, ok := suc.mutation.LongURL(); ok {
		if err := shorturl.LongURLValidator(v); err != nil {
			return &ValidationError{Name: "long_url", err: fmt.Errorf("ent: validator failed for field \"long_url\": %w", err)}
		}
	}
	if _, ok := suc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New("ent: missing required field \"create_at\"")}
	}
	if _, ok := suc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New("ent: missing required field \"update_at\"")}
	}
	return nil
}

func (suc *ShortUrlCreate) sqlSave(ctx context.Context) (*ShortUrl, error) {
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (suc *ShortUrlCreate) createSpec() (*ShortUrl, *sqlgraph.CreateSpec) {
	var (
		_node = &ShortUrl{config: suc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shorturl.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shorturl.FieldID,
			},
		}
	)
	if value, ok := suc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := suc.mutation.ShortURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldShortURL,
		})
		_node.ShortURL = value
	}
	if value, ok := suc.mutation.LongURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldLongURL,
		})
		_node.LongURL = value
	}
	if value, ok := suc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := suc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	return _node, _spec
}

// ShortUrlCreateBulk is the builder for creating many ShortUrl entities in bulk.
type ShortUrlCreateBulk struct {
	config
	builders []*ShortUrlCreate
}

// Save creates the ShortUrl entities in the database.
func (sucb *ShortUrlCreateBulk) Save(ctx context.Context) ([]*ShortUrl, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*ShortUrl, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShortUrlMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *ShortUrlCreateBulk) SaveX(ctx context.Context) []*ShortUrl {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
