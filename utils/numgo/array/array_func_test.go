package array

import (
	"localhost.com/deep-learnning-algorithmes/utils/assert"
	"testing"
)

func TestGenerationLastDimeDatas(t *testing.T) {
	w := []float64{1, 1, 1}
	b := 2.0
	col1 := []float64{1, 2, 3, 4, 5}
	col2 := []float64{1, 2, 3, 6, 7}

	exp := []float64{-4, -6, -8, -12, -14}

	if ! assert.IsEqual1DFloat64(exp, GenerationLastDimeDatas(w, b, col1, col2)) {
		t.Fatal()
	}

}
