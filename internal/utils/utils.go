package utils

//array indexof
func ValueIsInArray(a []string, b string) bool {
	for _, value := range a {
		if value == b {
			return true
		}
	}

	return false
}

func GetArrayIndexOf(a []string, b string) int {
	for index, value := range a {
		if value == b {
			return index
		}
	}

	return -1
}
