package chapter1_perceptron_model

import (
	"localhost.com/deep-learnning-algorithmes/utils/numgo/array"
)

// HyperPlanel hyper planel struct
type HyperPlanel struct {
	W []float64 // 法线
	B float64   // 偏置值
}

// NewHyperPlanel create a Hyplanel
func NewHyperPlanel(w []float64, b float64) *HyperPlanel {
	return &HyperPlanel{
		W: w,
		B: b,
	}
}

// GenerateRandomTestPoints xs (x1[],x2[],x3[])  ys [+1 || -1]
func (h *HyperPlanel) GenerateRandomTestPoints(n int) (points [][]float64, labels []float64) {
	// positivepointss
	t := array.RandRange1D(n, 0, 100)
	xs1 := array.RandRange1D(n, 0, 100)
	ys1 := array.RandRange1D(n, 0, 100)
	zs1 := array.GenerationLastDimeDatas(h.W, h.B, xs1, ys1)
	zs1 = array.OP1D(zs1, t, "+")

	// negativepointss
	t = array.RandRange1D(n, 0, 100)
	xs2 := array.RandRange1D(n, 0, 100)
	ys2 := array.RandRange1D(n, 0, 100)
	zs2 := array.GenerationLastDimeDatas(h.W, h.B, xs2, ys2)
	zs2 = array.OP1D(zs2, t, "-")

	for i := 0; i < n; i++ {
		points = append(points, []float64{xs1[i], ys1[i], zs1[i]})
	}

	for i := 0; i < n; i++ {
		points = append(points, []float64{xs2[i], ys2[i], zs2[i]})
	}

	labels = append(labels, append(array.Default1D(n, 1), array.Default1D(n, -1)...)...)

	return points, labels
}
