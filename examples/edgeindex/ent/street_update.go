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
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/predicate"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"
	"github.com/facebookincubator/ent/schema/field"
)

// StreetUpdate is the builder for updating Street entities.
type StreetUpdate struct {
	config
	hooks      []Hook
	mutation   *StreetMutation
	predicates []predicate.Street
}

// Where adds a new predicate for the builder.
func (su *StreetUpdate) Where(ps ...predicate.Street) *StreetUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetName sets the name field.
func (su *StreetUpdate) SetName(s string) *StreetUpdate {
	su.mutation.SetName(s)
	return su
}

// SetCityID sets the city edge to City by id.
func (su *StreetUpdate) SetCityID(id int) *StreetUpdate {
	su.mutation.SetCityID(id)
	return su
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (su *StreetUpdate) SetNillableCityID(id *int) *StreetUpdate {
	if id != nil {
		su = su.SetCityID(*id)
	}
	return su
}

// SetCity sets the city edge to City.
func (su *StreetUpdate) SetCity(c *City) *StreetUpdate {
	return su.SetCityID(c.ID)
}

// Mutation returns the StreetMutation object of the builder.
func (su *StreetUpdate) Mutation() *StreetMutation {
	return su.mutation
}

// ClearCity clears the city edge to City.
func (su *StreetUpdate) ClearCity() *StreetUpdate {
	su.mutation.ClearCity()
	return su
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *StreetUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StreetMutation)
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
func (su *StreetUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StreetUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StreetUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StreetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   street.Table,
			Columns: street.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: street.FieldID,
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
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: street.FieldName,
		})
	}
	if su.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
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
			err = &NotFoundError{street.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StreetUpdateOne is the builder for updating a single Street entity.
type StreetUpdateOne struct {
	config
	hooks    []Hook
	mutation *StreetMutation
}

// SetName sets the name field.
func (suo *StreetUpdateOne) SetName(s string) *StreetUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetCityID sets the city edge to City by id.
func (suo *StreetUpdateOne) SetCityID(id int) *StreetUpdateOne {
	suo.mutation.SetCityID(id)
	return suo
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (suo *StreetUpdateOne) SetNillableCityID(id *int) *StreetUpdateOne {
	if id != nil {
		suo = suo.SetCityID(*id)
	}
	return suo
}

// SetCity sets the city edge to City.
func (suo *StreetUpdateOne) SetCity(c *City) *StreetUpdateOne {
	return suo.SetCityID(c.ID)
}

// Mutation returns the StreetMutation object of the builder.
func (suo *StreetUpdateOne) Mutation() *StreetMutation {
	return suo.mutation
}

// ClearCity clears the city edge to City.
func (suo *StreetUpdateOne) ClearCity() *StreetUpdateOne {
	suo.mutation.ClearCity()
	return suo
}

// Save executes the query and returns the updated entity.
func (suo *StreetUpdateOne) Save(ctx context.Context) (*Street, error) {

	var (
		err  error
		node *Street
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StreetMutation)
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
func (suo *StreetUpdateOne) SaveX(ctx context.Context) *Street {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *StreetUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StreetUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StreetUpdateOne) sqlSave(ctx context.Context) (s *Street, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   street.Table,
			Columns: street.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: street.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Street.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: street.FieldName,
		})
	}
	if suo.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Street{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{street.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
