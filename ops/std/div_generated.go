package stdops

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

import (
	"context"
	"runtime/trace"

	gctx "gorgonia.org/gorgonia/internal/context"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// divOp is the base op for elementwise division.
type divOp struct{ binop }

// String implements fmt.Stringer.
func (op divOp) String() string { return "÷" }

// Do performs elementwise division.
func (op divOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Div(a, b, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise division but with a preallocated return value.
// PreallocDo allows div to implement ops.PreallocOp.
func (op divOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Div(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// divVV is a tensor-tensor elementwise division.
type divVV struct {
	divOp
	binopVV
}

// divVS is a tensor-scalar elementwise division.
type divVS struct {
	divOp
	binopVS
}

// String implements fmt.Stringer.
func (op divVS) String() string { return "÷·" }

// divSV is a scalar-tensor elementwise division.
type divSV struct {
	divOp
	binopSV
}

// String implements fmt.Stringer.
func (op divSV) String() string { return "·÷" }
