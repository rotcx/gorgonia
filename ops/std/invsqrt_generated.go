package stdops

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

import (
	"context"
	"runtime/trace"

	gctx "gorgonia.org/gorgonia/internal/context"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// invSqrt is a elementwise 1/√x.
type invSqrtOp struct{ unop }

// String implements fmt.Stringer.
func (op invSqrtOp) String() string { return "1/√·" }

// Do performs elementwise 1/√x.
func (op invSqrtOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.InvSqrt(a, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise 1/√x but with a preallocated return value.
// PreallocDo allows add to implement ops.PreallocOp.
func (op invSqrtOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.InvSqrt(a, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}
