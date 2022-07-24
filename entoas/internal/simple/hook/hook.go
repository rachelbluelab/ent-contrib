// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"entgo.io/contrib/entoas/internal/simple"
)

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *simple.CategoryMutation) (simple.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m simple.Mutation) (simple.Value, error) {
	mv, ok := m.(*simple.CategoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *simple.CategoryMutation", m)
	}
	return f(ctx, mv)
}

// The PetFunc type is an adapter to allow the use of ordinary
// function as Pet mutator.
type PetFunc func(context.Context, *simple.PetMutation) (simple.Value, error)

// Mutate calls f(ctx, m).
func (f PetFunc) Mutate(ctx context.Context, m simple.Mutation) (simple.Value, error) {
	mv, ok := m.(*simple.PetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *simple.PetMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *simple.UserMutation) (simple.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m simple.Mutation) (simple.Value, error) {
	mv, ok := m.(*simple.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *simple.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, simple.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m simple.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m simple.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m simple.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op simple.Op) Condition {
	return func(_ context.Context, m simple.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m simple.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m simple.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m simple.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk simple.Hook, cond Condition) simple.Hook {
	return func(next simple.Mutator) simple.Mutator {
		return simple.MutateFunc(func(ctx context.Context, m simple.Mutation) (simple.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, simple.Delete|simple.Create)
//
func On(hk simple.Hook, op simple.Op) simple.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, simple.Update|simple.UpdateOne)
//
func Unless(hk simple.Hook, op simple.Op) simple.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) simple.Hook {
	return func(simple.Mutator) simple.Mutator {
		return simple.MutateFunc(func(context.Context, simple.Mutation) (simple.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []simple.Hook {
//		return []simple.Hook{
//			Reject(simple.Delete|simple.Update),
//		}
//	}
//
func Reject(op simple.Op) simple.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []simple.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...simple.Hook) Chain {
	return Chain{append([]simple.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() simple.Hook {
	return func(mutator simple.Mutator) simple.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...simple.Hook) Chain {
	newHooks := make([]simple.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
