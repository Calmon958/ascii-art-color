package ascii

func IndexOf(str, substr string) int {
	n := len(str)
	m := len(substr)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && str[i+j] == substr[j] {
			j++
		}
		if j == m {
			return i
		}
	}

	return -1
}
