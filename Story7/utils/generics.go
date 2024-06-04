package utils

func All[T any](conditions []func(T) bool, value T) bool {

	len := len(conditions)
	for idx := 0; idx < len; idx++ {
		f := conditions[idx]
		if !f(value) {
			return false
		}
	}

	return true
}

func Any[T any](conditions []func(T) bool, value T) bool {

	len := len(conditions)
	for idx := 0; idx < len; idx++ {
		f := conditions[idx]
		if f(value) {
			return true
		}
	}

	return false
}
