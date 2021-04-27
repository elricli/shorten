// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/drrrMikado/shorten/internal/repo/ent/predicate"
	"github.com/drrrMikado/shorten/internal/repo/ent/shorturl"
)

// ShortUrlUpdate is the builder for updating ShortUrl entities.
type ShortUrlUpdate struct {
	config
	hooks    []Hook
	mutation *ShortUrlMutation
}

// Where adds a new predicate for the ShortUrlUpdate builder.
func (suu *ShortUrlUpdate) Where(ps ...predicate.ShortUrl) *ShortUrlUpdate {
	suu.mutation.predicates = append(suu.mutation.predicates, ps...)
	return suu
}

// SetKey sets the "key" field.
func (suu *ShortUrlUpdate) SetKey(s string) *ShortUrlUpdate {
	suu.mutation.SetKey(s)
	return suu
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (suu *ShortUrlUpdate) SetNillableKey(s *string) *ShortUrlUpdate {
	if s != nil {
		suu.SetKey(*s)
	}
	return suu
}

// SetURL sets the "url" field.
func (suu *ShortUrlUpdate) SetURL(s string) *ShortUrlUpdate {
	suu.mutation.SetURL(s)
	return suu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (suu *ShortUrlUpdate) SetNillableURL(s *string) *ShortUrlUpdate {
	if s != nil {
		suu.SetURL(*s)
	}
	return suu
}

// SetPv sets the "pv" field.
func (suu *ShortUrlUpdate) SetPv(u uint64) *ShortUrlUpdate {
	suu.mutation.ResetPv()
	suu.mutation.SetPv(u)
	return suu
}

// SetNillablePv sets the "pv" field if the given value is not nil.
func (suu *ShortUrlUpdate) SetNillablePv(u *uint64) *ShortUrlUpdate {
	if u != nil {
		suu.SetPv(*u)
	}
	return suu
}

// AddPv adds u to the "pv" field.
func (suu *ShortUrlUpdate) AddPv(u uint64) *ShortUrlUpdate {
	suu.mutation.AddPv(u)
	return suu
}

// ClearPv clears the value of the "pv" field.
func (suu *ShortUrlUpdate) ClearPv() *ShortUrlUpdate {
	suu.mutation.ClearPv()
	return suu
}

// SetExpire sets the "expire" field.
func (suu *ShortUrlUpdate) SetExpire(t time.Time) *ShortUrlUpdate {
	suu.mutation.SetExpire(t)
	return suu
}

// SetNillableExpire sets the "expire" field if the given value is not nil.
func (suu *ShortUrlUpdate) SetNillableExpire(t *time.Time) *ShortUrlUpdate {
	if t != nil {
		suu.SetExpire(*t)
	}
	return suu
}

// ClearExpire clears the value of the "expire" field.
func (suu *ShortUrlUpdate) ClearExpire() *ShortUrlUpdate {
	suu.mutation.ClearExpire()
	return suu
}

// SetCreateAt sets the "create_at" field.
func (suu *ShortUrlUpdate) SetCreateAt(t time.Time) *ShortUrlUpdate {
	suu.mutation.SetCreateAt(t)
	return suu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (suu *ShortUrlUpdate) SetNillableCreateAt(t *time.Time) *ShortUrlUpdate {
	if t != nil {
		suu.SetCreateAt(*t)
	}
	return suu
}

// SetUpdateAt sets the "update_at" field.
func (suu *ShortUrlUpdate) SetUpdateAt(t time.Time) *ShortUrlUpdate {
	suu.mutation.SetUpdateAt(t)
	return suu
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (suu *ShortUrlUpdate) SetNillableUpdateAt(t *time.Time) *ShortUrlUpdate {
	if t != nil {
		suu.SetUpdateAt(*t)
	}
	return suu
}

// Mutation returns the ShortUrlMutation object of the builder.
func (suu *ShortUrlUpdate) Mutation() *ShortUrlMutation {
	return suu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (suu *ShortUrlUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(suu.hooks) == 0 {
		if err = suu.check(); err != nil {
			return 0, err
		}
		affected, err = suu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShortUrlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suu.check(); err != nil {
				return 0, err
			}
			suu.mutation = mutation
			affected, err = suu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(suu.hooks) - 1; i >= 0; i-- {
			mut = suu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (suu *ShortUrlUpdate) SaveX(ctx context.Context) int {
	affected, err := suu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (suu *ShortUrlUpdate) Exec(ctx context.Context) error {
	_, err := suu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suu *ShortUrlUpdate) ExecX(ctx context.Context) {
	if err := suu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suu *ShortUrlUpdate) check() error {
	if v, ok := suu.mutation.Key(); ok {
		if err := shorturl.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf("ent: validator failed for field \"key\": %w", err)}
		}
	}
	if v, ok := suu.mutation.URL(); ok {
		if err := shorturl.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	return nil
}

func (suu *ShortUrlUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shorturl.Table,
			Columns: shorturl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shorturl.FieldID,
			},
		},
	}
	if ps := suu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldKey,
		})
	}
	if value, ok := suu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldURL,
		})
	}
	if value, ok := suu.mutation.Pv(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: shorturl.FieldPv,
		})
	}
	if value, ok := suu.mutation.AddedPv(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: shorturl.FieldPv,
		})
	}
	if suu.mutation.PvCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: shorturl.FieldPv,
		})
	}
	if value, ok := suu.mutation.Expire(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldExpire,
		})
	}
	if suu.mutation.ExpireCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: shorturl.FieldExpire,
		})
	}
	if value, ok := suu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldCreateAt,
		})
	}
	if value, ok := suu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldUpdateAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, suu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shorturl.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ShortUrlUpdateOne is the builder for updating a single ShortUrl entity.
type ShortUrlUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShortUrlMutation
}

// SetKey sets the "key" field.
func (suuo *ShortUrlUpdateOne) SetKey(s string) *ShortUrlUpdateOne {
	suuo.mutation.SetKey(s)
	return suuo
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (suuo *ShortUrlUpdateOne) SetNillableKey(s *string) *ShortUrlUpdateOne {
	if s != nil {
		suuo.SetKey(*s)
	}
	return suuo
}

// SetURL sets the "url" field.
func (suuo *ShortUrlUpdateOne) SetURL(s string) *ShortUrlUpdateOne {
	suuo.mutation.SetURL(s)
	return suuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (suuo *ShortUrlUpdateOne) SetNillableURL(s *string) *ShortUrlUpdateOne {
	if s != nil {
		suuo.SetURL(*s)
	}
	return suuo
}

// SetPv sets the "pv" field.
func (suuo *ShortUrlUpdateOne) SetPv(u uint64) *ShortUrlUpdateOne {
	suuo.mutation.ResetPv()
	suuo.mutation.SetPv(u)
	return suuo
}

// SetNillablePv sets the "pv" field if the given value is not nil.
func (suuo *ShortUrlUpdateOne) SetNillablePv(u *uint64) *ShortUrlUpdateOne {
	if u != nil {
		suuo.SetPv(*u)
	}
	return suuo
}

// AddPv adds u to the "pv" field.
func (suuo *ShortUrlUpdateOne) AddPv(u uint64) *ShortUrlUpdateOne {
	suuo.mutation.AddPv(u)
	return suuo
}

// ClearPv clears the value of the "pv" field.
func (suuo *ShortUrlUpdateOne) ClearPv() *ShortUrlUpdateOne {
	suuo.mutation.ClearPv()
	return suuo
}

// SetExpire sets the "expire" field.
func (suuo *ShortUrlUpdateOne) SetExpire(t time.Time) *ShortUrlUpdateOne {
	suuo.mutation.SetExpire(t)
	return suuo
}

// SetNillableExpire sets the "expire" field if the given value is not nil.
func (suuo *ShortUrlUpdateOne) SetNillableExpire(t *time.Time) *ShortUrlUpdateOne {
	if t != nil {
		suuo.SetExpire(*t)
	}
	return suuo
}

// ClearExpire clears the value of the "expire" field.
func (suuo *ShortUrlUpdateOne) ClearExpire() *ShortUrlUpdateOne {
	suuo.mutation.ClearExpire()
	return suuo
}

// SetCreateAt sets the "create_at" field.
func (suuo *ShortUrlUpdateOne) SetCreateAt(t time.Time) *ShortUrlUpdateOne {
	suuo.mutation.SetCreateAt(t)
	return suuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (suuo *ShortUrlUpdateOne) SetNillableCreateAt(t *time.Time) *ShortUrlUpdateOne {
	if t != nil {
		suuo.SetCreateAt(*t)
	}
	return suuo
}

// SetUpdateAt sets the "update_at" field.
func (suuo *ShortUrlUpdateOne) SetUpdateAt(t time.Time) *ShortUrlUpdateOne {
	suuo.mutation.SetUpdateAt(t)
	return suuo
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (suuo *ShortUrlUpdateOne) SetNillableUpdateAt(t *time.Time) *ShortUrlUpdateOne {
	if t != nil {
		suuo.SetUpdateAt(*t)
	}
	return suuo
}

// Mutation returns the ShortUrlMutation object of the builder.
func (suuo *ShortUrlUpdateOne) Mutation() *ShortUrlMutation {
	return suuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suuo *ShortUrlUpdateOne) Select(field string, fields ...string) *ShortUrlUpdateOne {
	suuo.fields = append([]string{field}, fields...)
	return suuo
}

// Save executes the query and returns the updated ShortUrl entity.
func (suuo *ShortUrlUpdateOne) Save(ctx context.Context) (*ShortUrl, error) {
	var (
		err  error
		node *ShortUrl
	)
	if len(suuo.hooks) == 0 {
		if err = suuo.check(); err != nil {
			return nil, err
		}
		node, err = suuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShortUrlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suuo.check(); err != nil {
				return nil, err
			}
			suuo.mutation = mutation
			node, err = suuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suuo.hooks) - 1; i >= 0; i-- {
			mut = suuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suuo *ShortUrlUpdateOne) SaveX(ctx context.Context) *ShortUrl {
	node, err := suuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suuo *ShortUrlUpdateOne) Exec(ctx context.Context) error {
	_, err := suuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suuo *ShortUrlUpdateOne) ExecX(ctx context.Context) {
	if err := suuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suuo *ShortUrlUpdateOne) check() error {
	if v, ok := suuo.mutation.Key(); ok {
		if err := shorturl.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf("ent: validator failed for field \"key\": %w", err)}
		}
	}
	if v, ok := suuo.mutation.URL(); ok {
		if err := shorturl.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	return nil
}

func (suuo *ShortUrlUpdateOne) sqlSave(ctx context.Context) (_node *ShortUrl, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shorturl.Table,
			Columns: shorturl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shorturl.FieldID,
			},
		},
	}
	id, ok := suuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ShortUrl.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shorturl.FieldID)
		for _, f := range fields {
			if !shorturl.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shorturl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldKey,
		})
	}
	if value, ok := suuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldURL,
		})
	}
	if value, ok := suuo.mutation.Pv(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: shorturl.FieldPv,
		})
	}
	if value, ok := suuo.mutation.AddedPv(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: shorturl.FieldPv,
		})
	}
	if suuo.mutation.PvCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: shorturl.FieldPv,
		})
	}
	if value, ok := suuo.mutation.Expire(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldExpire,
		})
	}
	if suuo.mutation.ExpireCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: shorturl.FieldExpire,
		})
	}
	if value, ok := suuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldCreateAt,
		})
	}
	if value, ok := suuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldUpdateAt,
		})
	}
	_node = &ShortUrl{config: suuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shorturl.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
