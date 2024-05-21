package utils

func PutIfAbsent[k comparable, v any](m map[k]v, key k, val v) bool {
	_, ok := m[key]
	if !ok {
		m[key] = val
	}

	return !ok
}

func ContainsValue[Key comparable, Value comparable](m map[Key]Value, val Value) bool {
	for _, v := range m {
		if val == v {
			return true
		}
	}

	return false
}
