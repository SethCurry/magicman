// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SethCurry/magicman/pkg/ent/predicate"
	"github.com/SethCurry/magicman/pkg/ent/ruling"
)

// RulingDelete is the builder for deleting a Ruling entity.
type RulingDelete struct {
	config
	hooks    []Hook
	mutation *RulingMutation
}

// Where appends a list predicates to the RulingDelete builder.
func (rd *RulingDelete) Where(ps ...predicate.Ruling) *RulingDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RulingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rd.hooks) == 0 {
		affected, err = rd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RulingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rd.mutation = mutation
			affected, err = rd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rd.hooks) - 1; i >= 0; i-- {
			if rd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RulingDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RulingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: ruling.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ruling.FieldID,
			},
		},
	}
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// RulingDeleteOne is the builder for deleting a single Ruling entity.
type RulingDeleteOne struct {
	rd *RulingDelete
}

// Exec executes the deletion query.
func (rdo *RulingDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ruling.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RulingDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}
