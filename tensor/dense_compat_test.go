package tensor

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/stretchr/testify/assert"
)

/*
GENERATED FILE. DO NOT EDIT
*/

var toMat64Tests = []struct {
	data   interface{}
	sliced interface{}
	shape  Shape
	dt     Dtype
}{
	{Range(Int, 0, 6), []int{0, 1, 3, 4}, Shape{2, 3}, Int},
	{Range(Int8, 0, 6), []int8{0, 1, 3, 4}, Shape{2, 3}, Int8},
	{Range(Int16, 0, 6), []int16{0, 1, 3, 4}, Shape{2, 3}, Int16},
	{Range(Int32, 0, 6), []int32{0, 1, 3, 4}, Shape{2, 3}, Int32},
	{Range(Int64, 0, 6), []int64{0, 1, 3, 4}, Shape{2, 3}, Int64},
	{Range(Uint, 0, 6), []uint{0, 1, 3, 4}, Shape{2, 3}, Uint},
	{Range(Uint8, 0, 6), []uint8{0, 1, 3, 4}, Shape{2, 3}, Uint8},
	{Range(Uint16, 0, 6), []uint16{0, 1, 3, 4}, Shape{2, 3}, Uint16},
	{Range(Uint32, 0, 6), []uint32{0, 1, 3, 4}, Shape{2, 3}, Uint32},
	{Range(Uint64, 0, 6), []uint64{0, 1, 3, 4}, Shape{2, 3}, Uint64},
	{Range(Float32, 0, 6), []float32{0, 1, 3, 4}, Shape{2, 3}, Float32},
	{Range(Float64, 0, 6), []float64{0, 1, 3, 4}, Shape{2, 3}, Float64},
	{Range(Complex64, 0, 6), []complex64{0, 1, 3, 4}, Shape{2, 3}, Complex64},
	{Range(Complex128, 0, 6), []complex128{0, 1, 3, 4}, Shape{2, 3}, Complex128},
}

func TestToMat64(t *testing.T) {
	assert := assert.New(t)
	for i, tmt := range toMat64Tests {
		T := New(WithBacking(tmt.data), WithShape(tmt.shape...))
		var m *mat64.Dense
		var err error
		if m, err = ToMat64(T); err != nil {
			t.Errorf("ToMat basic test %d failed : %v", i, err)
			continue
		}
		conv := anyToFloat64s(tmt.data)
		assert.Equal(conv, m.RawMatrix().Data, "i %d from %v", i, tmt.dt)

		if T, err = sliceDense(T, nil, makeRS(0, 2)); err != nil {
			t.Errorf("Slice failed %v", err)
			continue
		}
		if m, err = ToMat64(T); err != nil {
			t.Errorf("ToMat of slice test %d failed : %v", i, err)
			continue
		}
		conv = anyToFloat64s(tmt.sliced)
		assert.Equal(conv, m.RawMatrix().Data, "sliced test %d from %v", i, tmt.dt)
		t.Logf("Done")

		if tmt.dt == Float64 {
			T = New(WithBacking(tmt.data), WithShape(tmt.shape...))
			if m, err = ToMat64(T, UseUnsafe()); err != nil {
				t.Errorf("ToMat64 unsafe test %d failed: %v", i, err)
			}
			conv = anyToFloat64s(tmt.data)
			assert.Equal(conv, m.RawMatrix().Data, "float64 unsafe i %d from %v", i, tmt.dt)
			conv[0] = 1000
			assert.Equal(conv, m.RawMatrix().Data, "float64 unsafe i %d from %v", i, tmt.dt)
			conv[0] = 0 // reset for future tests that use the same backing
		}
	}
	// idiocy test
	T := New(Of(Float64), WithShape(2, 3, 4))
	_, err := ToMat64(T)
	if err == nil {
		t.Error("Expected an error when trying to convert a 3-T to *mat.Dense")
	}
}

func TestFromMat64(t *testing.T) {
	assert := assert.New(t)
	var m *mat64.Dense
	var T *Dense
	var backing []float64

	for i, tmt := range toMat64Tests {
		backing = Range(Float64, 0, 6).([]float64)
		m = mat64.NewDense(2, 3, backing)
		T = FromMat64(m)
		conv := anyToFloat64s(tmt.data)
		assert.Equal(conv, T.Data(), "test %d: []float64 from %v", i, tmt.dt)
		assert.True(T.Shape().Eq(tmt.shape))

		T = FromMat64(m, As(tmt.dt))
		assert.Equal(tmt.data, T.Data())
		assert.True(T.Shape().Eq(tmt.shape))

		if tmt.dt == Float64 {
			backing = Range(Float64, 0, 6).([]float64)
			m = mat64.NewDense(2, 3, backing)
			T = FromMat64(m, UseUnsafe())
			assert.Equal(backing, T.Float64s())
			assert.True(T.Shape().Eq(tmt.shape))
			backing[0] = 1000
			assert.Equal(backing, T.Float64s(), "test %d - unsafe float64", i)
		}
	}
}
