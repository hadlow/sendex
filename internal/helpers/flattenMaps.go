package helpers

func FlattenMaps(maps []map[string]string) (result map[string]string) {
	result = make(map[string]string)

	for _, m := range maps {
		for key, value := range m {
			result[key] = value
		}
	}

	return
}
