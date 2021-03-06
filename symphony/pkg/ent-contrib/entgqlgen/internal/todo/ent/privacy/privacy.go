// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/symphony/pkg/ent-contrib/entgqlgen/internal/todo/ent"
)

var (
	// Allow may be returned by read/write rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by read/write rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by read/write rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type (
	// ReadPolicy combines multiple read rules into a single policy.
	ReadPolicy []ReadRule

	// ReadRule defines the interface deciding whether a read is allowed.
	ReadRule interface {
		EvalRead(context.Context, ent.Value) error
	}
)

// EvalRead evaluates a load against a read policy.
func (policy ReadPolicy) EvalRead(ctx context.Context, v ent.Value) error {
	for _, rule := range policy {
		switch err := rule.EvalRead(ctx, v); {
		case err == nil || errors.Is(err, Skip):
		case errors.Is(err, Allow):
			return nil
		default:
			return err
		}
	}
	return nil
}

// ReadRuleFunc type is an adapter to allow the use of
// ordinary functions as read rules.
type ReadRuleFunc func(context.Context, ent.Value) error

// Eval calls f(ctx, v).
func (f ReadRuleFunc) EvalRead(ctx context.Context, v ent.Value) error {
	return f(ctx, v)
}

type (
	// WritePolicy combines multiple write rules into a single policy.
	WritePolicy []WriteRule

	// WriteRule defines the interface deciding whether a write is allowed.
	WriteRule interface {
		EvalWrite(context.Context, ent.Mutation) error
	}
)

// EvalWrite evaluates a mutation against a write policy.
func (policy WritePolicy) EvalWrite(ctx context.Context, m ent.Mutation) error {
	for _, rule := range policy {
		switch err := rule.EvalWrite(ctx, m); {
		case err == nil || errors.Is(err, Skip):
		case errors.Is(err, Allow):
			return nil
		default:
			return err
		}
	}
	return nil
}

// WriteRuleFunc type is an adapter to allow the use of
// ordinary functions as write rules.
type WriteRuleFunc func(context.Context, ent.Mutation) error

// Eval calls f(ctx, m).
func (f WriteRuleFunc) EvalWrite(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups read and write policies.
type Policy struct {
	Read  ReadPolicy
	Write WritePolicy
}

// EvalRead forwards evaluation to read policy.
func (policy Policy) EvalRead(ctx context.Context, v ent.Value) error {
	return policy.Read.EvalRead(ctx, v)
}

// EvalWrite forwards evaluation to write policy.
func (policy Policy) EvalWrite(ctx context.Context, m ent.Mutation) error {
	return policy.Write.EvalWrite(ctx, m)
}

// ReadWriteRule is the interface that groups read and write rules.
type ReadWriteRule interface {
	ReadRule
	WriteRule
}

// AlwaysAllowRule returns a read/write rule that returns an allow decision.
func AlwaysAllowRule() ReadWriteRule {
	return fixedDecisionRule{Allow}
}

// AlwaysDenyRule returns a read/write rule that returns a deny decision.
func AlwaysDenyRule() ReadWriteRule {
	return fixedDecisionRule{Deny}
}

type fixedDecisionRule struct{ err error }

func (f fixedDecisionRule) EvalRead(context.Context, ent.Value) error     { return f.err }
func (f fixedDecisionRule) EvalWrite(context.Context, ent.Mutation) error { return f.err }

// The TodoReadRuleFunc type is an adapter to allow the use of ordinary
// functions as a read rule.
type TodoReadRuleFunc func(context.Context, *ent.Todo) error

// EvalRead calls f(ctx, v).
func (f TodoReadRuleFunc) EvalRead(ctx context.Context, v ent.Value) error {
	if v, ok := v.(*ent.Todo); ok {
		return f(ctx, v)
	}
	return Denyf("ent/privacy: unexpected value type %T, expect *ent.Todo", v)
}

// The TodoWriteRuleFunc type is an adapter to allow the use of ordinary
// functions as a write rule.
type TodoWriteRuleFunc func(context.Context, *ent.TodoMutation) error

// EvalWrite calls f(ctx, m).
func (f TodoWriteRuleFunc) EvalWrite(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TodoMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TodoMutation", m)
}
