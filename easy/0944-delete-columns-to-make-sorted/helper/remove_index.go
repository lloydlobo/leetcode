package helper

func RemoveIndex(strs []string, i int) []string {
	n := len(strs) - 1
	var s = make([]string, n)

	key := 0
	for k := 0; k < n; k++ {
		if k != i {
			s[k] = strs[key]
			key++
		} else {
			key++
		}
	}

	return s
}

// https://stackoverflow.com/a/57213476
func RemoveIndexConcise(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
