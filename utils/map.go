package utils

func PutIfAbsent[k comparable, v any](m map[k]v, key k, val v) bool {
	_, ok := m[key]
	if !ok {
		m[key] = val
	}

	return !ok
}
