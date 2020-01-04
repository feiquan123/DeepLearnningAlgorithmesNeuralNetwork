package chapter1_perceptron_model

import (
	"localhost.com/deep-learnning-algorithmes/utils/numgo/array"
)

// Perceptron model
func Perceptron() {
	w := []float64{2, -0.9, 1}
	b := -1
	h := NewHyperPlanel(w, float64(b))

	points, labels := h.GenerateRandomTestPoints(100)

	train(points, labels, 0.1)
}

// points (x1[],x2[],x3[])  ys [+1 || -1]
func train(points [][]float64, labels []float64, lrt float64) (w []float64, b float64) {
	w = []float64{0, 0, 0}
	b = 0
	iter := 0

	errExist := true
	for errExist {
		errExist = false
		for i := 0; i < array.Sharp2D(points)[0]; i++ {
			yt := array.Dot([][]float64{w}, array.Transposition2D([][]float64{points[i]}))
			y := yt[0][0] + b

			if labels[i]*y <= 0 { // err point
				drt_point := array.NOP1D(points[i], lrt*labels[i], "*")
				w = array.OP1D(w, drt_point, "+")
				b += lrt * labels[i]
				iter++
				logger.Infof("iter:%d		w:[%f %f %f]		b:%f", iter, w[0], w[1], w[2], b)
				errExist = true // err point exist , go on train
				// logger.Debug(i)
				break
			}
		}
	}
	return w, b
}
