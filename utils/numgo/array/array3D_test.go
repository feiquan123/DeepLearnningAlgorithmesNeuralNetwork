package array

import (
	"localhost.com/deep-learnning-algorithmes/utils/assert"
	"testing"
)

func TestDefault3D(t *testing.T) {
	arr := Default3D(10, 10, 10, 100)
	if len(arr) != 10 {
		t.Fatal()
	}
	for _, v := range arr {
		for _, m := range v {
			for _, n := range m {
				if n != 100 {
					t.Fatal()
				}
			}
		}
	}

	//t.Log("TestDefault3D", arr)
}

func TestRandRange3D(t *testing.T) {
	arr := RandRange3D(10, 10, 10, -100, 0)
	if len(arr) != 10 {
		t.Fatal()
	}
	for _, v := range arr {
		if len(v) != 10 {
			t.Fatal()
		}
	}
	for _, v := range arr {
		for _, m := range v {
			for _, n := range m {
				if n < -100 || n > 0 {
					t.Errorf("the value %v should in [%v,%v]", m, -100, 0)
				}
			}
		}
	}

	//t.Log("TestRandRange3D:", arr)
}

func TestNOP3D(t *testing.T) {
	arr1 := [][][]float64{
		[][]float64{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 4}},
		[][]float64{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 4}},
	}
	arr2 := 10.0
	var exp, r [][][]float64

	// add
	exp = [][][]float64{
		[][]float64{[]float64{11, 12, 13, 14}, []float64{11, 12, 13, 14}},
		[][]float64{[]float64{11, 12, 13, 14}, []float64{11, 12, 13, 14}},
	}
	r = NOP3D(arr1, arr2, "+")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// dsc
	exp = [][][]float64{
		[][]float64{[]float64{-9, -8, -7, -6}, []float64{-9, -8, -7, -6}},
		[][]float64{[]float64{-9, -8, -7, -6}, []float64{-9, -8, -7, -6}},
	}
	r = NOP3D(arr1, arr2, "-")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// multi
	exp = [][][]float64{
		[][]float64{[]float64{10, 20, 30, 40}, []float64{10, 20, 30, 40}},
		[][]float64{[]float64{10, 20, 30, 40}, []float64{10, 20, 30, 40}},
	}
	r = NOP3D(arr1, arr2, "*")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// div
	exp = [][][]float64{
		[][]float64{[]float64{0.1, 0.2, 0.3, 0.4}, []float64{0.1, 0.2, 0.3, 0.4}},
		[][]float64{[]float64{0.1, 0.2, 0.3, 0.4}, []float64{0.1, 0.2, 0.3, 0.4}},
	}
	r = NOP3D(arr1, arr2, "/")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}
}

func TestOP3D(t *testing.T) {
	arr1 := [][][]float64{
		[][]float64{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 4}},
		[][]float64{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 4}},
	}
	arr2 := [][][]float64{
		[][]float64{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 4}},
		[][]float64{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 4}},
	}
	var exp, r [][][]float64

	// add
	exp = [][][]float64{
		[][]float64{[]float64{2, 4, 6, 8}, []float64{2, 4, 6, 8}},
		[][]float64{[]float64{2, 4, 6, 8}, []float64{2, 4, 6, 8}},
	}
	r = OP3D(arr1, arr2, "+")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// dsc
	exp = [][][]float64{
		[][]float64{[]float64{0, 0, 0, 0}, []float64{0, 0, 0, 0}},
		[][]float64{[]float64{0, 0, 0, 0}, []float64{0, 0, 0, 0}},
	}
	r = OP3D(arr1, arr2, "-")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// multi
	exp = [][][]float64{
		[][]float64{[]float64{1, 4, 9, 16}, []float64{1, 4, 9, 16}},
		[][]float64{[]float64{1, 4, 9, 16}, []float64{1, 4, 9, 16}},
	}

	r = OP3D(arr1, arr2, "*")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}

	// div
	exp = [][][]float64{
		[][]float64{[]float64{1, 1, 1, 1}, []float64{1, 1, 1, 1}},
		[][]float64{[]float64{1, 1, 1, 1}, []float64{1, 1, 1, 1}},
	}
	r = OP3D(arr1, arr2, "/")
	if !assert.IsEqual3DFloat64(exp, r) {
		t.Errorf("Except %v , get %v", exp, r)
	}
}

func TestSharp3D(t *testing.T) {
	arr1 := Default3D(10, 5, 9, 100)

	exp := [] int{10, 5, 9}
	if !assert.IsEqual1DInt(Sharp3D(arr1), exp) {
		t.Errorf("Except %v , get %v", 10, Sharp3D(arr1)[0])
	}
}
