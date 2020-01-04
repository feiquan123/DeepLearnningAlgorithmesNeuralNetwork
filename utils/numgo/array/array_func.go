package array

/*
by hyper plane  func and x1[],x2[] generate x3[]
eg :
	w := []float64{1, 1, 1}
	b := 2.0
	col1 := []float64{1, 2, 3,4,5}
	col2:= []float64{1, 2, 3,6,7}
re:
	[]float64{-4,-6,-8,-12,-14}
*/
func GenerationLastDimeDatas(w []float64, b float64, col1, col2 []float64) (col3 []float64) {
	if len(col1) != len(col2) {
		panic("Dime x1 length unequal Dime x2")
	}
	for i := 0; i < len(col2); i++ {
		col3 = append(col3, - (w[0]*col1[i]+w[1]*col2[i]+b)/w[2])
	}
	return col3
}
