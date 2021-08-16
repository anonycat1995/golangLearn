package listSlice

func Sum(l [5]int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return sum
}
