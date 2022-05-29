package series

func All(n int, s string) []string {
	if n > len(s) {
		return []string{}
	}
	retArr := make([]string, len(s)-n+1)

	for i := 0; i+n <= len(s); i++ {
		retArr[i] = s[i : i+n]
	}

	return retArr
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
