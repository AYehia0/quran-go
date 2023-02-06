package utils

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func StringInArray(s string, strings []string) bool {
	for _, str := range strings {
		if str == s {
			return true
		}
	}
	return false
}
