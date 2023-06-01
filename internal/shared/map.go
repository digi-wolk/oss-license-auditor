package shared

func MapContainsKey(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}
