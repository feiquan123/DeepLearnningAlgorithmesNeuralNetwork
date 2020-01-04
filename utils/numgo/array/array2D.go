package array

// get default 2D array[l1,l2]
func Default2D(l1, l2 int, v float64) (arr [][]float64) {
	for i := 0; i < l1; i++ {
		arr = append(arr, Default1D(l2, v))
	}
	return arr
}

// generate n value in [vmin, vmax),return 2D array[i,j]
func RandRange2D(l1, l2 int, vmin, vmax float64) (a [][]float64) {
	for i := 0; i < l1; i++ {
		a = append(a, RandRange1D(l2, vmin, vmax))
	}
	return a
}

// array normal op, one 2D arr + n
func NOP2D(arr1 [][]float64, n float64, op string) (a [][]float64) {
	switch op {
	case "+":
		for _, v := range arr1 {
			a = append(a, NOP1D(v, n, op))
		}
	case "-":
		for _, v := range arr1 {
			a = append(a, NOP1D(v, n, op))
		}
	case "*":
		for _, v := range arr1 {
			a = append(a, NOP1D(v, n, op))
		}
	case "/":
		for _, v := range arr1 {
			a = append(a, NOP1D(v, n, op))
		}
	default:
		panic("unknown array opations.")
	}
	return a
}

// two 2D array op
func OP2D(arr1 [][]float64, arr2 [][]float64, op string) (a [][]float64) {
	if len(arr1) < len(arr2) {
		arr1, arr2 = arr2, arr1
	}

	switch op {
	case "+":
		var (
			i int
			v []float64
		)
		for i, v = range arr2 {
			a = append(a, OP1D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	case "-":
		var (
			i int
			v []float64
		)
		for i, v = range arr2 {
			a = append(a, OP1D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	case "*":
		var (
			i int
			v []float64
		)
		for i, v = range arr2 {
			a = append(a, OP1D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	case "/":
		var (
			i int
			v []float64
		)
		for i, v = range arr2 {
			a = append(a, OP1D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	default:
		panic("unknown array opations.")
	}
	return a
}

// 2D array transposition
func Transposition2D(arr1 [][]float64) (a [][]float64) {
	if len(arr1) == 0 {
		return a
	}
	if len(arr1[0]) == 0 {
		return a
	}

	a = Default2D(len(arr1[0]), len(arr1), 0)

	for m := 0; m < len(arr1); m++ {
		for n := 0; n < len(arr1[0]); n++ {
			a[n][m] = arr1[m][n]
		}
	}

	return a
}

// two 2D array dot, arr1 . arr2
func Dot(arr1 [][]float64, arr2 [][]float64) (a [][]float64) {
	if len(arr1)+len(arr2) == 0 {
		return a
	}
	if len(arr1)+len(arr2) == 1 {
		panic("dot can not exist array length 0.")
	}
	if len(arr1) != len(arr2[0]) {
		panic("two array length unequal,can not dot.")
	}
	arr2 = Transposition2D(arr2)

	for i := 0; i < len(arr1); i++ {
		t_row := []float64{}
		for _, col := range arr2 {
			sum := 0.0
			for j := 0; j < len(arr1[i]); j++ {
				sum += arr1[i][j] * col[j]
			}
			t_row = append(t_row, sum)
		}
		a = append(a, t_row)
	}
	return a
}

// get sharp
func Sharp2D(arr1 [][]float64) ([]int) {
	if len(arr1) == 0 {
		return []int{0, 0}
	}

	return []int{len(arr1), len(arr1[0])}
}
