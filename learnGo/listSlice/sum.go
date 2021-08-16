package listSlice

func Sum(l [5]int) int {
	var sum int
	for _, v := range l {
		sum += v
	}
	return sum
}
