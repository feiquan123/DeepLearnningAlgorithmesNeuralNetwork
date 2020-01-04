package array

import (
	"testing"

	"localhost.com/deep-learnning-algorithmes/utils/assert"
)

func TestDefault1D(t *testing.T) {
	arr := Default1D(10, 100)
	if len(arr) != 10 {
		t.Fatal()
	}
	for _, v := range arr {
		if v != 100 {
			t.Fatal()
		}
	}
	t.Log("TestDefault1D:", arr)
}

func TestRandRange1D(t *testing.T) {
	arr := RandRange1D(10, -100, 0)
	if len(arr) != 10 {
		t.Fatal()
	}
	for _, v := range arr {
		if v < -100 || v > 0 {
			t.Errorf("the value %v should in [%v,%v]", v, -100, 0)
		}
	}
	t.Log("TestRandRange1D:", arr)
}

func TestNOP1D(t *testing.T) {
	arr1 := []float64{1, 2, 3, 4}
	arr2 := 10.0
	var exp, r []float64

	// add
	exp = []float64{11, 12, 13, 14}
	r = NOP1D(arr1, arr2, "+")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// dsc
	exp = []float64{-9, -8, -7, -6}
	r = NOP1D(arr1, arr2, "-")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// multi
	exp = []float64{10, 20, 30, 40}
	r = NOP1D(arr1, arr2, "*")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// div
	exp = []float64{0.1, 0.2, 0.3, 0.4}
	r = NOP1D(arr1, arr2, "/")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}
}

func TestOP1D(t *testing.T) {
	arr1 := []float64{1, 2, 3, 4}
	arr2 := []float64{1, 2, 3, 4}
	var exp, r []float64

	// add
	exp = []float64{2, 4, 6, 8}
	r = OP1D(arr1, arr2, "+")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// dsc
	exp = []float64{0, 0, 0, 0}
	r = OP1D(arr1, arr2, "-")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// multi
	exp = []float64{1, 4, 9, 16}
	r = OP1D(arr1, arr2, "*")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// div
	exp = []float64{1, 1, 1, 1}
	r = OP1D(arr1, arr2, "/")
	if !assert.IsEqual1DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}
}

func TestSharp1D(t *testing.T) {
	arr1 := Default1D(10, 100)

	exp := []int{10}
	if !assert.IsEqual1DInt(Sharp1D(arr1), exp) {
		t.Errorf("Except %v , get %v", 10, Sharp1D(arr1)[0])
	}
}
