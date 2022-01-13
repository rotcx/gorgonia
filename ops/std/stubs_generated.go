package stdops

import (
	"context"

	"gorgonia.org/gorgonia/exprgraph"
	"gorgonia.org/gorgonia/internal/datatypes"
	"gorgonia.org/gorgonia/values/dual"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// DoDiff is the method that allows automatic differentiation of `sub` .
func (op subOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// DoDiff is the method that allows automatic differentiation of `mul` .
func (op mulOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// DoDiff is the method that allows automatic differentiation of `div` .
func (op divOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of pow.
func (op powOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `pow` .
func (op powOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of exp.
func (op expOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `exp` .
func (op expOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of ln.
func (op lnOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `ln` .
func (op lnOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of log2.
func (op log2Op) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `log2` .
func (op log2Op) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of neg.
func (op negOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `neg` .
func (op negOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of square.
func (op squareOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `square` .
func (op squareOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of sqrt.
func (op sqrtOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `sqrt` .
func (op sqrtOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of inv.
func (op invOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `inv` .
func (op invOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of invSqrt.
func (op invSqrtOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `invSqrt` .
func (op invSqrtOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of cube.
func (op cubeOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `cube` .
func (op cubeOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}

// SymDiff performs the symbolic differentiation of tanh.
func (op tanhOp) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented")
}

// DoDiff is the method that allows automatic differentiation of `tanh` .
func (op tanhOp) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) error {
	adv := exprgraph.T2T(inputs[0]).(*dual.Dual)
	bdv := exprgraph.T2T(inputs[1]).(*dual.Dual)
	cdv := exprgraph.T2T(output).(*dual.Dual)

	advd := adv.Deriv()
	bdvd := bdv.Deriv()

	_, _, _ = cdv, advd, bdvd
	panic("Not implemented")
}
