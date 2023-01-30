package commons

type profiles struct {
	profile uint
}

func GetDataDictionary(path string) interface{} {
	dict := make(map[string]interface{})

	profiles := map[string]int{"ADMIN": 1, "CUSTOMER": 2, "OWNERS": 3}

	dict["profiles"] = profiles

	return dict[path]
}
