package array

// get default 3D array[l1,l2,l3]
func Default3D(l1, l2, l3 int, v float64) (arr [][][]float64) {
	for i := 0; i < l1; i++ {
		arr = append(arr, Default2D(l2, l3, v))
	}
	return arr
}

// generate n value in [vmin, vmax),return 2D array[i,j,k]
func RandRange3D(l1, l2, l3 int, vmin, vmax float64) (a [][][]float64) {
	for i := 0; i < l1; i++ {
		a = append(a, RandRange2D(l2, l3, vmin, vmax))
	}
	return a
}

// array normal op, one 3D arr + n
func NOP3D(arr1 [][][]float64, n float64, op string) (a [][][]float64) {
	switch op {
	case "+":
		for _, v := range arr1 {
			a = append(a, NOP2D(v, n, op))
		}
	case "-":
		for _, v := range arr1 {
			a = append(a, NOP2D(v, n, op))
		}
	case "*":
		for _, v := range arr1 {
			a = append(a, NOP2D(v, n, op))
		}
	case "/":
		for _, v := range arr1 {
			a = append(a, NOP2D(v, n, op))
		}
	default:
		panic("unknown array opations.")
	}
	return a
}

// two 3D array op
func OP3D(arr1 [][][]float64, arr2 [][][]float64, op string) (a [][][]float64) {
	if len(arr1) < len(arr2) {
		arr1, arr2 = arr2, arr1
	}

	switch op {
	case "+":
		var (
			i int
			v [][]float64
		)
		for i, v = range arr2 {
			a = append(a, OP2D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	case "-":
		var (
			i int
			v [][]float64
		)
		for i, v = range arr2 {
			a = append(a, OP2D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	case "*":
		var (
			i int
			v [][]float64
		)
		for i, v = range arr2 {
			a = append(a, OP2D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	case "/":
		var (
			i int
			v [][]float64
		)
		for i, v = range arr2 {
			a = append(a, OP2D(v, arr1[i], op))
		}
		i++
		a = append(a, arr1[i:]...)
	default:
		panic("unknown array opations.")
	}
	return a
}

// get sharp
func Sharp3D(arr1 [][][]float64) ([]int) {
	if len(arr1) == 0 {
		return []int{0, 0}
	}

	return []int{len(arr1), len(arr1[0]), len(arr1[0][0])}
}
