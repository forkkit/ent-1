// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv1"
)

// The CarFunc type is an adapter to allow the use of ordinary
// function as Car mutator.
type CarFunc func(context.Context, *entv1.CarMutation) (entv1.Value, error)

// Mutate calls f(ctx, m).
func (f CarFunc) Mutate(ctx context.Context, m entv1.Mutation) (entv1.Value, error) {
	mv, ok := m.(*entv1.CarMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entv1.CarMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *entv1.UserMutation) (entv1.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m entv1.Mutation) (entv1.Value, error) {
	mv, ok := m.(*entv1.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entv1.UserMutation", m)
	}
	return f(ctx, mv)
}

// On executes the given hook only of the given operation.
//
//	hook.On(Log, entv1.Delete|entv1.Create)
//
func On(hk entv1.Hook, op entv1.Op) entv1.Hook {
	return func(next entv1.Mutator) entv1.Mutator {
		return entv1.MutateFunc(func(ctx context.Context, m entv1.Mutation) (entv1.Value, error) {
			if m.Op().Is(op) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []entv1.Hook {
//		return []entv1.Hook{
//			Reject(entv1.Delete|entv1.Update),
//		}
//	}
//
func Reject(op entv1.Op) entv1.Hook {
	return func(next entv1.Mutator) entv1.Mutator {
		return entv1.MutateFunc(func(ctx context.Context, m entv1.Mutation) (entv1.Value, error) {
			if m.Op().Is(op) {
				return nil, fmt.Errorf("%s operation is not allowed", m.Op())
			}
			return next.Mutate(ctx, m)
		})
	}
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []entv1.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...entv1.Hook) Chain {
	return Chain{append([]entv1.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() entv1.Hook {
	return func(mutator entv1.Mutator) entv1.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...entv1.Hook) Chain {
	newHooks := make([]entv1.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
