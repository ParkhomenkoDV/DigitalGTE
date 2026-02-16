package utils

// All проверяет, все ли значения true
func All(sl ...bool) bool {
	for _, b := range sl {
		if !b {
			return false
		}
	}
	return true
}

// Any проверяет, есть ли хотя бы одно true
func Any(sl ...bool) bool {
	for _, b := range sl {
		if b {
			return true
		}
	}
	return false
}
