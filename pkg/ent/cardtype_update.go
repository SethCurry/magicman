// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SethCurry/magicman/pkg/ent/card"
	"github.com/SethCurry/magicman/pkg/ent/cardtype"
	"github.com/SethCurry/magicman/pkg/ent/predicate"
)

// CardTypeUpdate is the builder for updating CardType entities.
type CardTypeUpdate struct {
	config
	hooks    []Hook
	mutation *CardTypeMutation
}

// Where appends a list predicates to the CardTypeUpdate builder.
func (ctu *CardTypeUpdate) Where(ps ...predicate.CardType) *CardTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetName sets the "name" field.
func (ctu *CardTypeUpdate) SetName(s string) *CardTypeUpdate {
	ctu.mutation.SetName(s)
	return ctu
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (ctu *CardTypeUpdate) AddCardIDs(ids ...int) *CardTypeUpdate {
	ctu.mutation.AddCardIDs(ids...)
	return ctu
}

// AddCards adds the "cards" edges to the Card entity.
func (ctu *CardTypeUpdate) AddCards(c ...*Card) *CardTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.AddCardIDs(ids...)
}

// Mutation returns the CardTypeMutation object of the builder.
func (ctu *CardTypeUpdate) Mutation() *CardTypeMutation {
	return ctu.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (ctu *CardTypeUpdate) ClearCards() *CardTypeUpdate {
	ctu.mutation.ClearCards()
	return ctu
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (ctu *CardTypeUpdate) RemoveCardIDs(ids ...int) *CardTypeUpdate {
	ctu.mutation.RemoveCardIDs(ids...)
	return ctu
}

// RemoveCards removes "cards" edges to Card entities.
func (ctu *CardTypeUpdate) RemoveCards(c ...*Card) *CardTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.RemoveCardIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CardTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctu.hooks) == 0 {
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CardTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CardTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CardTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctu *CardTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardtype.Table,
			Columns: cardtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardtype.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cardtype.FieldName,
		})
	}
	if ctu.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: cardtype.CardsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedCardsIDs(); len(nodes) > 0 && !ctu.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: cardtype.CardsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: cardtype.CardsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CardTypeUpdateOne is the builder for updating a single CardType entity.
type CardTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardTypeMutation
}

// SetName sets the "name" field.
func (ctuo *CardTypeUpdateOne) SetName(s string) *CardTypeUpdateOne {
	ctuo.mutation.SetName(s)
	return ctuo
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (ctuo *CardTypeUpdateOne) AddCardIDs(ids ...int) *CardTypeUpdateOne {
	ctuo.mutation.AddCardIDs(ids...)
	return ctuo
}

// AddCards adds the "cards" edges to the Card entity.
func (ctuo *CardTypeUpdateOne) AddCards(c ...*Card) *CardTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.AddCardIDs(ids...)
}

// Mutation returns the CardTypeMutation object of the builder.
func (ctuo *CardTypeUpdateOne) Mutation() *CardTypeMutation {
	return ctuo.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (ctuo *CardTypeUpdateOne) ClearCards() *CardTypeUpdateOne {
	ctuo.mutation.ClearCards()
	return ctuo
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (ctuo *CardTypeUpdateOne) RemoveCardIDs(ids ...int) *CardTypeUpdateOne {
	ctuo.mutation.RemoveCardIDs(ids...)
	return ctuo
}

// RemoveCards removes "cards" edges to Card entities.
func (ctuo *CardTypeUpdateOne) RemoveCards(c ...*Card) *CardTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.RemoveCardIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CardTypeUpdateOne) Select(field string, fields ...string) *CardTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CardType entity.
func (ctuo *CardTypeUpdateOne) Save(ctx context.Context) (*CardType, error) {
	var (
		err  error
		node *CardType
	)
	if len(ctuo.hooks) == 0 {
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CardTypeUpdateOne) SaveX(ctx context.Context) *CardType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CardTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CardTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctuo *CardTypeUpdateOne) sqlSave(ctx context.Context) (_node *CardType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardtype.Table,
			Columns: cardtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardtype.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CardType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cardtype.FieldID)
		for _, f := range fields {
			if !cardtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cardtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cardtype.FieldName,
		})
	}
	if ctuo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: cardtype.CardsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedCardsIDs(); len(nodes) > 0 && !ctuo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: cardtype.CardsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: cardtype.CardsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CardType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
