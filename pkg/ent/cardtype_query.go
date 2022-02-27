// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SethCurry/magicman/pkg/ent/card"
	"github.com/SethCurry/magicman/pkg/ent/cardtype"
	"github.com/SethCurry/magicman/pkg/ent/predicate"
)

// CardTypeQuery is the builder for querying CardType entities.
type CardTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CardType
	// eager-loading edges.
	withCards *CardQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CardTypeQuery builder.
func (ctq *CardTypeQuery) Where(ps ...predicate.CardType) *CardTypeQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit adds a limit step to the query.
func (ctq *CardTypeQuery) Limit(limit int) *CardTypeQuery {
	ctq.limit = &limit
	return ctq
}

// Offset adds an offset step to the query.
func (ctq *CardTypeQuery) Offset(offset int) *CardTypeQuery {
	ctq.offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *CardTypeQuery) Unique(unique bool) *CardTypeQuery {
	ctq.unique = &unique
	return ctq
}

// Order adds an order step to the query.
func (ctq *CardTypeQuery) Order(o ...OrderFunc) *CardTypeQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryCards chains the current query on the "cards" edge.
func (ctq *CardTypeQuery) QueryCards() *CardQuery {
	query := &CardQuery{config: ctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cardtype.Table, cardtype.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, cardtype.CardsTable, cardtype.CardsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CardType entity from the query.
// Returns a *NotFoundError when no CardType was found.
func (ctq *CardTypeQuery) First(ctx context.Context) (*CardType, error) {
	nodes, err := ctq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cardtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *CardTypeQuery) FirstX(ctx context.Context) *CardType {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CardType ID from the query.
// Returns a *NotFoundError when no CardType ID was found.
func (ctq *CardTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cardtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *CardTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CardType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CardType entity is not found.
// Returns a *NotFoundError when no CardType entities are found.
func (ctq *CardTypeQuery) Only(ctx context.Context) (*CardType, error) {
	nodes, err := ctq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cardtype.Label}
	default:
		return nil, &NotSingularError{cardtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *CardTypeQuery) OnlyX(ctx context.Context) *CardType {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CardType ID in the query.
// Returns a *NotSingularError when exactly one CardType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ctq *CardTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = &NotSingularError{cardtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *CardTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CardTypes.
func (ctq *CardTypeQuery) All(ctx context.Context) ([]*CardType, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ctq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ctq *CardTypeQuery) AllX(ctx context.Context) []*CardType {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CardType IDs.
func (ctq *CardTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ctq.Select(cardtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *CardTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *CardTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ctq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *CardTypeQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *CardTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ctq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *CardTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CardTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *CardTypeQuery) Clone() *CardTypeQuery {
	if ctq == nil {
		return nil
	}
	return &CardTypeQuery{
		config:     ctq.config,
		limit:      ctq.limit,
		offset:     ctq.offset,
		order:      append([]OrderFunc{}, ctq.order...),
		predicates: append([]predicate.CardType{}, ctq.predicates...),
		withCards:  ctq.withCards.Clone(),
		// clone intermediate query.
		sql:  ctq.sql.Clone(),
		path: ctq.path,
	}
}

// WithCards tells the query-builder to eager-load the nodes that are connected to
// the "cards" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *CardTypeQuery) WithCards(opts ...func(*CardQuery)) *CardTypeQuery {
	query := &CardQuery{config: ctq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctq.withCards = query
	return ctq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CardType.Query().
//		GroupBy(cardtype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ctq *CardTypeQuery) GroupBy(field string, fields ...string) *CardTypeGroupBy {
	group := &CardTypeGroupBy{config: ctq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.CardType.Query().
//		Select(cardtype.FieldName).
//		Scan(ctx, &v)
//
func (ctq *CardTypeQuery) Select(fields ...string) *CardTypeSelect {
	ctq.fields = append(ctq.fields, fields...)
	return &CardTypeSelect{CardTypeQuery: ctq}
}

func (ctq *CardTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ctq.fields {
		if !cardtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *CardTypeQuery) sqlAll(ctx context.Context) ([]*CardType, error) {
	var (
		nodes       = []*CardType{}
		_spec       = ctq.querySpec()
		loadedTypes = [1]bool{
			ctq.withCards != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CardType{config: ctq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ctq.withCards; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*CardType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Cards = []*Card{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*CardType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   cardtype.CardsTable,
				Columns: cardtype.CardsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(cardtype.CardsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, ctq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "cards": %w`, err)
		}
		query.Where(card.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "cards" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Cards = append(nodes[i].Edges.Cards, n)
			}
		}
	}

	return nodes, nil
}

func (ctq *CardTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *CardTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ctq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ctq *CardTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardtype.Table,
			Columns: cardtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardtype.FieldID,
			},
		},
		From:   ctq.sql,
		Unique: true,
	}
	if unique := ctq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ctq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cardtype.FieldID)
		for i := range fields {
			if fields[i] != cardtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *CardTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(cardtype.Table)
	columns := ctq.fields
	if len(columns) == 0 {
		columns = cardtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CardTypeGroupBy is the group-by builder for CardType entities.
type CardTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *CardTypeGroupBy) Aggregate(fns ...AggregateFunc) *CardTypeGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ctgb *CardTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ctgb.path(ctx)
	if err != nil {
		return err
	}
	ctgb.sql = query
	return ctgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ctgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("ent: CardTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ctgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ctgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) StringX(ctx context.Context) string {
	v, err := ctgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("ent: CardTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ctgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ctgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) IntX(ctx context.Context) int {
	v, err := ctgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("ent: CardTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ctgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ctgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ctgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("ent: CardTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ctgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ctgb *CardTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ctgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ctgb *CardTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := ctgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ctgb *CardTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ctgb.fields {
		if !cardtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ctgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ctgb *CardTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ctgb.sql.Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ctgb.fields)+len(ctgb.fns))
		for _, f := range ctgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ctgb.fields...)...)
}

// CardTypeSelect is the builder for selecting fields of CardType entities.
type CardTypeSelect struct {
	*CardTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cts *CardTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	cts.sql = cts.CardTypeQuery.sqlQuery(ctx)
	return cts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cts *CardTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("ent: CardTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cts *CardTypeSelect) StringsX(ctx context.Context) []string {
	v, err := cts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cts *CardTypeSelect) StringX(ctx context.Context) string {
	v, err := cts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("ent: CardTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cts *CardTypeSelect) IntsX(ctx context.Context) []int {
	v, err := cts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cts *CardTypeSelect) IntX(ctx context.Context) int {
	v, err := cts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("ent: CardTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cts *CardTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cts *CardTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := cts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("ent: CardTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cts *CardTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := cts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cts *CardTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = fmt.Errorf("ent: CardTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cts *CardTypeSelect) BoolX(ctx context.Context) bool {
	v, err := cts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cts *CardTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cts.sql.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}