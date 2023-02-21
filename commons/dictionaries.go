package commons

var dict = make(map[string]interface{})
var profiles = map[string]uint16{"OWNERS": 1, "CUSTOMERS": 2, "ADMINS": 3}

func GetProfileDataDictionary() interface{} {
	dict["profiles"] = profiles
	return dict["profiles"]
}

func GetProfileId(Key string, dictionary interface{}) uint16 {
	i := dictionary.(map[string]uint16)
	return i[Key]
}
