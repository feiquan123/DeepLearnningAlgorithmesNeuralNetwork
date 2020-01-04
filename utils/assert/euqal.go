package assert


///////// interface ///////////////
func IsEqual1DInter(v1 []interface{}, v2 []interface{}) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i, v := range v1 {
		if v != v2[i] {
			return false
		}
	}
	return true
}

func IsEqual2DInter(v1 [][]interface{}, v2 [][]interface{}) bool {
	if len(v1) != len(v2) {
		return false
	}
	if len(v1[0]) != len(v2[0]) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		for j := 0; j < len(v1[0]); j++ {
			if v1[i][j] != v2[i][j] {
				return false
			}
		}
	}
	return true
}

func IsEqual3DInter(v1 [][][]interface{}, v2 [][][]interface{}) bool {
	if len(v1) != len(v2) {
		return false
	}
	if len(v1[0]) != len(v2[0]) {
		return false
	}
	if len(v1[0][0]) != len(v2[0][0]) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		for j := 0; j < len(v1[0]); j++ {
			for k := 0; k < len(v1[0][0]); k++ {
				if v1[i][j][k] != v2[i][j][k] {
					return false
				}
			}
		}
	}
	return true
}


///////// Float64 ///////////////
func IsEqual1DFloat64(v1 []float64, v2 []float64) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i, v := range v1 {
		if v != v2[i] {
			return false
		}
	}
	return true
}

func IsEqual2DFloat64(v1 [][]float64, v2 [][]float64) bool {
	if len(v1) != len(v2) {
		return false
	}
	if len(v1[0]) != len(v2[0]) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		for j := 0; j < len(v1[0]); j++ {
			if v1[i][j] != v2[i][j] {
				return false
			}
		}
	}
	return true
}

func IsEqual3DFloat64(v1 [][][]float64, v2 [][][]float64) bool {
	if len(v1) != len(v2) {
		return false
	}
	if len(v1[0]) != len(v2[0]) {
		return false
	}
	if len(v1[0][0]) != len(v2[0][0]) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		for j := 0; j < len(v1[0]); j++ {
			for k := 0; k < len(v1[0][0]); k++ {
				if v1[i][j][k] != v2[i][j][k] {
					return false
				}
			}
		}
	}
	return true
}

///////// Int ///////////////
func IsEqual1DInt(v1 []int, v2 []int) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i, v := range v1 {
		if v != v2[i] {
			return false
		}
	}
	return true
}

func IsEqual2DInt(v1 [][]int, v2 [][]int) bool {
	if len(v1) != len(v2) {
		return false
	}
	if len(v1[0]) != len(v2[0]) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		for j := 0; j < len(v1[0]); j++ {
			if v1[i][j] != v2[i][j] {
				return false
			}
		}
	}
	return true
}

func IsEqual3DInt(v1 [][][]int, v2 [][][]int) bool {
	if len(v1) != len(v2) {
		return false
	}
	if len(v1[0]) != len(v2[0]) {
		return false
	}
	if len(v1[0][0]) != len(v2[0][0]) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		for j := 0; j < len(v1[0]); j++ {
			for k := 0; k < len(v1[0][0]); k++ {
				if v1[i][j][k] != v2[i][j][k] {
					return false
				}
			}
		}
	}
	return true
}