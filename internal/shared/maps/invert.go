package maps

func Invert[K comparable, V comparable](mp map[K]V) map[V]K {
	if mp == nil {
		return nil
	}

	nMap := make(map[V]K, len(mp))

	for k, v := range mp {
		nMap[v] = k
	}

	return nMap
}
