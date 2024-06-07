package defaultvalue

func For[T comparable](value T, replacement T) T {
	var zero T
	if value == zero {
		return replacement
	}
	return value
}
