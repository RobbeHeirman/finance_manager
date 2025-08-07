package data_structures

func GetMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
