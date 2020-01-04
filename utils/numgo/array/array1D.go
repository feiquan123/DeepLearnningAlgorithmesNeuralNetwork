package array

import (
	"math/rand"
	"time"
)

// get default 1D array[l1]
func Default1D(l int, v float64) (arr []float64) {
	for i := 0; i < l; i++ {
		arr = append(arr, v)
	}
	return arr
}

// generate n value in [vmin, vmax),return 1D array[l1]
func RandRange1D(l1 int, vmin, vmax float64) (a []float64) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l1; i++ {
		a = append(a, rand.Float64()*(vmax-vmin)+vmin)
	}
	return a
}

// array normal op, one 1D arr + n
func NOP1D(arr1 []float64, n float64, op string) (a []float64) {
	switch op {
	case "+":
		for _, v := range arr1 {
			a = append(a, v+n)
		}
	case "-":
		for _, v := range arr1 {
			a = append(a, v-n)
		}
	case "*":
		for _, v := range arr1 {
			a = append(a, v*n)
		}
	case "/":
		for _, v := range arr1 {
			a = append(a, v/n)
		}
	default:
		panic("unknown array opations.")
	}
	return a
}

// two 1D array op
func OP1D(arr1 []float64, arr2 []float64, op string) (a []float64) {
	if len(arr1) < len(arr2) {
		arr1, arr2 = arr2, arr1
	}

	switch op {
	case "+":
		var (
			i int
			v float64
		)
		for i, v = range arr2 {
			a = append(a, v+arr1[i])
		}
		i++
		a = append(a, arr1[i:]...)
	case "-":
		var (
			i int
			v float64
		)
		for i, v = range arr2 {
			a = append(a, v-arr1[i])
		}
		i++
		a = append(a, arr1[i:]...)
	case "*":
		var (
			i int
			v float64
		)
		for i, v = range arr2 {
			a = append(a, v*arr1[i])
		}
		i++
		a = append(a, arr1[i:]...)
	case "/":
		var (
			i int
			v float64
		)
		for i, v = range arr2 {
			a = append(a, v/arr1[i])
		}
		i++
		a = append(a, arr1[i:]...)
	default:
		panic("unknown array opations.")
	}
	return a
}

// get sharp
func Sharp1D(arr1 []float64) (a []int) {
	return []int{len(arr1)}
}
