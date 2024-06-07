package extslices

func LastElem[S ~[]E, E any](s S) (value E, ok bool) {
	if len(s) == 0 {
		return value, false
	}
	lastIdx := len(s) - 1
	return s[lastIdx], true
}
