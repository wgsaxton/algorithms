package module01

func NumInList(list []int, num int) bool {
	for i := range list {
		if i == num {
			return true
		}
	}

	return false
}