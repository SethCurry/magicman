// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SethCurry/magicman/pkg/ent/card"
	"github.com/SethCurry/magicman/pkg/ent/subtype"
)

// SubTypeCreate is the builder for creating a SubType entity.
type SubTypeCreate struct {
	config
	mutation *SubTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (stc *SubTypeCreate) SetName(s string) *SubTypeCreate {
	stc.mutation.SetName(s)
	return stc
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (stc *SubTypeCreate) AddCardIDs(ids ...int) *SubTypeCreate {
	stc.mutation.AddCardIDs(ids...)
	return stc
}

// AddCards adds the "cards" edges to the Card entity.
func (stc *SubTypeCreate) AddCards(c ...*Card) *SubTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return stc.AddCardIDs(ids...)
}

// Mutation returns the SubTypeMutation object of the builder.
func (stc *SubTypeCreate) Mutation() *SubTypeMutation {
	return stc.mutation
}

// Save creates the SubType in the database.
func (stc *SubTypeCreate) Save(ctx context.Context) (*SubType, error) {
	var (
		err  error
		node *SubType
	)
	if len(stc.hooks) == 0 {
		if err = stc.check(); err != nil {
			return nil, err
		}
		node, err = stc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stc.check(); err != nil {
				return nil, err
			}
			stc.mutation = mutation
			if node, err = stc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(stc.hooks) - 1; i >= 0; i-- {
			if stc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (stc *SubTypeCreate) SaveX(ctx context.Context) *SubType {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stc *SubTypeCreate) Exec(ctx context.Context) error {
	_, err := stc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stc *SubTypeCreate) ExecX(ctx context.Context) {
	if err := stc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *SubTypeCreate) check() error {
	if _, ok := stc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (stc *SubTypeCreate) sqlSave(ctx context.Context) (*SubType, error) {
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (stc *SubTypeCreate) createSpec() (*SubType, *sqlgraph.CreateSpec) {
	var (
		_node = &SubType{config: stc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subtype.FieldID,
			},
		}
	)
	if value, ok := stc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subtype.FieldName,
		})
		_node.Name = value
	}
	if nodes := stc.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subtype.CardsTable,
			Columns: subtype.CardsPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubTypeCreateBulk is the builder for creating many SubType entities in bulk.
type SubTypeCreateBulk struct {
	config
	builders []*SubTypeCreate
}

// Save creates the SubType entities in the database.
func (stcb *SubTypeCreateBulk) Save(ctx context.Context) ([]*SubType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*SubType, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *SubTypeCreateBulk) SaveX(ctx context.Context) []*SubType {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stcb *SubTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := stcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stcb *SubTypeCreateBulk) ExecX(ctx context.Context) {
	if err := stcb.Exec(ctx); err != nil {
		panic(err)
	}
}
