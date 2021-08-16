package iteration

func Repeat(a string) string {
	var world string
	for i := 0; i < 999999; i++ {
		world += a
	}
	return world
}
