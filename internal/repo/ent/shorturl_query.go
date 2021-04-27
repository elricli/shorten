// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/drrrMikado/shorten/internal/repo/ent/predicate"
	"github.com/drrrMikado/shorten/internal/repo/ent/shorturl"
)

// ShortUrlQuery is the builder for querying ShortUrl entities.
type ShortUrlQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShortUrl
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShortUrlQuery builder.
func (suq *ShortUrlQuery) Where(ps ...predicate.ShortUrl) *ShortUrlQuery {
	suq.predicates = append(suq.predicates, ps...)
	return suq
}

// Limit adds a limit step to the query.
func (suq *ShortUrlQuery) Limit(limit int) *ShortUrlQuery {
	suq.limit = &limit
	return suq
}

// Offset adds an offset step to the query.
func (suq *ShortUrlQuery) Offset(offset int) *ShortUrlQuery {
	suq.offset = &offset
	return suq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (suq *ShortUrlQuery) Unique(unique bool) *ShortUrlQuery {
	suq.unique = &unique
	return suq
}

// Order adds an order step to the query.
func (suq *ShortUrlQuery) Order(o ...OrderFunc) *ShortUrlQuery {
	suq.order = append(suq.order, o...)
	return suq
}

// First returns the first ShortUrl entity from the query.
// Returns a *NotFoundError when no ShortUrl was found.
func (suq *ShortUrlQuery) First(ctx context.Context) (*ShortUrl, error) {
	nodes, err := suq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shorturl.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (suq *ShortUrlQuery) FirstX(ctx context.Context) *ShortUrl {
	node, err := suq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShortUrl ID from the query.
// Returns a *NotFoundError when no ShortUrl ID was found.
func (suq *ShortUrlQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = suq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shorturl.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (suq *ShortUrlQuery) FirstIDX(ctx context.Context) int {
	id, err := suq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShortUrl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ShortUrl entity is not found.
// Returns a *NotFoundError when no ShortUrl entities are found.
func (suq *ShortUrlQuery) Only(ctx context.Context) (*ShortUrl, error) {
	nodes, err := suq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shorturl.Label}
	default:
		return nil, &NotSingularError{shorturl.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (suq *ShortUrlQuery) OnlyX(ctx context.Context) *ShortUrl {
	node, err := suq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShortUrl ID in the query.
// Returns a *NotSingularError when exactly one ShortUrl ID is not found.
// Returns a *NotFoundError when no entities are found.
func (suq *ShortUrlQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = suq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = &NotSingularError{shorturl.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (suq *ShortUrlQuery) OnlyIDX(ctx context.Context) int {
	id, err := suq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShortUrls.
func (suq *ShortUrlQuery) All(ctx context.Context) ([]*ShortUrl, error) {
	if err := suq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return suq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (suq *ShortUrlQuery) AllX(ctx context.Context) []*ShortUrl {
	nodes, err := suq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShortUrl IDs.
func (suq *ShortUrlQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := suq.Select(shorturl.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (suq *ShortUrlQuery) IDsX(ctx context.Context) []int {
	ids, err := suq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (suq *ShortUrlQuery) Count(ctx context.Context) (int, error) {
	if err := suq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return suq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (suq *ShortUrlQuery) CountX(ctx context.Context) int {
	count, err := suq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (suq *ShortUrlQuery) Exist(ctx context.Context) (bool, error) {
	if err := suq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return suq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (suq *ShortUrlQuery) ExistX(ctx context.Context) bool {
	exist, err := suq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShortUrlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (suq *ShortUrlQuery) Clone() *ShortUrlQuery {
	if suq == nil {
		return nil
	}
	return &ShortUrlQuery{
		config:     suq.config,
		limit:      suq.limit,
		offset:     suq.offset,
		order:      append([]OrderFunc{}, suq.order...),
		predicates: append([]predicate.ShortUrl{}, suq.predicates...),
		// clone intermediate query.
		sql:  suq.sql.Clone(),
		path: suq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ShortUrl.Query().
//		GroupBy(shorturl.FieldKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (suq *ShortUrlQuery) GroupBy(field string, fields ...string) *ShortUrlGroupBy {
	group := &ShortUrlGroupBy{config: suq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := suq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return suq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//	}
//
//	client.ShortUrl.Query().
//		Select(shorturl.FieldKey).
//		Scan(ctx, &v)
//
func (suq *ShortUrlQuery) Select(field string, fields ...string) *ShortUrlSelect {
	suq.fields = append([]string{field}, fields...)
	return &ShortUrlSelect{ShortUrlQuery: suq}
}

func (suq *ShortUrlQuery) prepareQuery(ctx context.Context) error {
	for _, f := range suq.fields {
		if !shorturl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if suq.path != nil {
		prev, err := suq.path(ctx)
		if err != nil {
			return err
		}
		suq.sql = prev
	}
	return nil
}

func (suq *ShortUrlQuery) sqlAll(ctx context.Context) ([]*ShortUrl, error) {
	var (
		nodes = []*ShortUrl{}
		_spec = suq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ShortUrl{config: suq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, suq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (suq *ShortUrlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := suq.querySpec()
	return sqlgraph.CountNodes(ctx, suq.driver, _spec)
}

func (suq *ShortUrlQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := suq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (suq *ShortUrlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shorturl.Table,
			Columns: shorturl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shorturl.FieldID,
			},
		},
		From:   suq.sql,
		Unique: true,
	}
	if unique := suq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := suq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shorturl.FieldID)
		for i := range fields {
			if fields[i] != shorturl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := suq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := suq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := suq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := suq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (suq *ShortUrlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(suq.driver.Dialect())
	t1 := builder.Table(shorturl.Table)
	selector := builder.Select(t1.Columns(shorturl.Columns...)...).From(t1)
	if suq.sql != nil {
		selector = suq.sql
		selector.Select(selector.Columns(shorturl.Columns...)...)
	}
	for _, p := range suq.predicates {
		p(selector)
	}
	for _, p := range suq.order {
		p(selector)
	}
	if offset := suq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := suq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShortUrlGroupBy is the group-by builder for ShortUrl entities.
type ShortUrlGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sugb *ShortUrlGroupBy) Aggregate(fns ...AggregateFunc) *ShortUrlGroupBy {
	sugb.fns = append(sugb.fns, fns...)
	return sugb
}

// Scan applies the group-by query and scans the result into the given value.
func (sugb *ShortUrlGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sugb.path(ctx)
	if err != nil {
		return err
	}
	sugb.sql = query
	return sugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sugb.fields) > 1 {
		return nil, errors.New("ent: ShortUrlGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) StringsX(ctx context.Context) []string {
	v, err := sugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) StringX(ctx context.Context) string {
	v, err := sugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sugb.fields) > 1 {
		return nil, errors.New("ent: ShortUrlGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) IntsX(ctx context.Context) []int {
	v, err := sugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) IntX(ctx context.Context) int {
	v, err := sugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sugb.fields) > 1 {
		return nil, errors.New("ent: ShortUrlGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sugb.fields) > 1 {
		return nil, errors.New("ent: ShortUrlGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sugb *ShortUrlGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sugb *ShortUrlGroupBy) BoolX(ctx context.Context) bool {
	v, err := sugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sugb *ShortUrlGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sugb.fields {
		if !shorturl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sugb *ShortUrlGroupBy) sqlQuery() *sql.Selector {
	selector := sugb.sql
	columns := make([]string, 0, len(sugb.fields)+len(sugb.fns))
	columns = append(columns, sugb.fields...)
	for _, fn := range sugb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sugb.fields...)
}

// ShortUrlSelect is the builder for selecting fields of ShortUrl entities.
type ShortUrlSelect struct {
	*ShortUrlQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sus *ShortUrlSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sus.prepareQuery(ctx); err != nil {
		return err
	}
	sus.sql = sus.ShortUrlQuery.sqlQuery(ctx)
	return sus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sus *ShortUrlSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sus.fields) > 1 {
		return nil, errors.New("ent: ShortUrlSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sus *ShortUrlSelect) StringsX(ctx context.Context) []string {
	v, err := sus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sus *ShortUrlSelect) StringX(ctx context.Context) string {
	v, err := sus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sus.fields) > 1 {
		return nil, errors.New("ent: ShortUrlSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sus *ShortUrlSelect) IntsX(ctx context.Context) []int {
	v, err := sus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sus *ShortUrlSelect) IntX(ctx context.Context) int {
	v, err := sus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sus.fields) > 1 {
		return nil, errors.New("ent: ShortUrlSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sus *ShortUrlSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sus *ShortUrlSelect) Float64X(ctx context.Context) float64 {
	v, err := sus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sus.fields) > 1 {
		return nil, errors.New("ent: ShortUrlSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sus *ShortUrlSelect) BoolsX(ctx context.Context) []bool {
	v, err := sus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sus *ShortUrlSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shorturl.Label}
	default:
		err = fmt.Errorf("ent: ShortUrlSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sus *ShortUrlSelect) BoolX(ctx context.Context) bool {
	v, err := sus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sus *ShortUrlSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sus.sqlQuery().Query()
	if err := sus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sus *ShortUrlSelect) sqlQuery() sql.Querier {
	selector := sus.sql
	selector.Select(selector.Columns(sus.fields...)...)
	return selector
}
