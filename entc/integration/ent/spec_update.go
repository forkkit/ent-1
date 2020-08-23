// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/card"
	"github.com/facebook/ent/entc/integration/ent/predicate"
	"github.com/facebook/ent/entc/integration/ent/spec"
	"github.com/facebook/ent/schema/field"
)

// SpecUpdate is the builder for updating Spec entities.
type SpecUpdate struct {
	config
	hooks      []Hook
	mutation   *SpecMutation
	predicates []predicate.Spec
}

// Where adds a new predicate for the builder.
func (su *SpecUpdate) Where(ps ...predicate.Spec) *SpecUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// AddCardIDs adds the card edge to Card by ids.
func (su *SpecUpdate) AddCardIDs(ids ...int) *SpecUpdate {
	su.mutation.AddCardIDs(ids...)
	return su
}

// AddCard adds the card edges to Card.
func (su *SpecUpdate) AddCard(c ...*Card) *SpecUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (su *SpecUpdate) Mutation() *SpecMutation {
	return su.mutation
}

// RemoveCardIDs removes the card edge to Card by ids.
func (su *SpecUpdate) RemoveCardIDs(ids ...int) *SpecUpdate {
	su.mutation.RemoveCardIDs(ids...)
	return su
}

// RemoveCard removes card edges to Card.
func (su *SpecUpdate) RemoveCard(c ...*Card) *SpecUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCardIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SpecUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpecUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpecUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpecUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SpecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   spec.Table,
			Columns: spec.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spec.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nodes := su.mutation.RemovedCardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
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
	if nodes := su.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spec.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SpecUpdateOne is the builder for updating a single Spec entity.
type SpecUpdateOne struct {
	config
	hooks    []Hook
	mutation *SpecMutation
}

// AddCardIDs adds the card edge to Card by ids.
func (suo *SpecUpdateOne) AddCardIDs(ids ...int) *SpecUpdateOne {
	suo.mutation.AddCardIDs(ids...)
	return suo
}

// AddCard adds the card edges to Card.
func (suo *SpecUpdateOne) AddCard(c ...*Card) *SpecUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (suo *SpecUpdateOne) Mutation() *SpecMutation {
	return suo.mutation
}

// RemoveCardIDs removes the card edge to Card by ids.
func (suo *SpecUpdateOne) RemoveCardIDs(ids ...int) *SpecUpdateOne {
	suo.mutation.RemoveCardIDs(ids...)
	return suo
}

// RemoveCard removes card edges to Card.
func (suo *SpecUpdateOne) RemoveCard(c ...*Card) *SpecUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCardIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SpecUpdateOne) Save(ctx context.Context) (*Spec, error) {

	var (
		err  error
		node *Spec
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpecUpdateOne) SaveX(ctx context.Context) *Spec {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SpecUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpecUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SpecUpdateOne) sqlSave(ctx context.Context) (s *Spec, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   spec.Table,
			Columns: spec.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spec.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Spec.ID for update")}
	}
	_spec.Node.ID.Value = id
	if nodes := suo.mutation.RemovedCardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
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
	if nodes := suo.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
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
	s = &Spec{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spec.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
