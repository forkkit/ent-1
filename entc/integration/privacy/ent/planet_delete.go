// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// PlanetDelete is the builder for deleting a Planet entity.
type PlanetDelete struct {
	config
	hooks      []Hook
	mutation   *PlanetMutation
	predicates []predicate.Planet
}

// Where adds a new predicate to the delete builder.
func (pd *PlanetDelete) Where(ps ...predicate.Planet) *PlanetDelete {
	pd.predicates = append(pd.predicates, ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PlanetDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PlanetDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PlanetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: planet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: planet.FieldID,
			},
		},
	}
	if ps := pd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// PlanetDeleteOne is the builder for deleting a single Planet entity.
type PlanetDeleteOne struct {
	pd *PlanetDelete
}

// Exec executes the deletion query.
func (pdo *PlanetDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{planet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PlanetDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
