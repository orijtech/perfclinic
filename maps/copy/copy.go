package maps

func copyIt(m map[int]int) map[int]int {
	cp := make(map[int]int)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}
