package commons

var dict = make(map[string]interface{})

var profiles = map[string]uint16{
	"OWNERS":    1,
	"CUSTOMERS": 2,
	"ADMINS":    3,
}
var assetsType = map[string]uint16{
	"frontalPictureProfile": 1,
	"frontalPictureId":      2,
	"posteriorPictureId":    3,
}

var accountStatus = map[string]uint8{
	"unverifiableIdentity": 1,
	"awaitingReview":       2,
	"approved":             3,
	"rejected":             4,
	"locked":               5,
}

func GetProfileDataDictionary() interface{} {
	dict["profiles"] = profiles
	return dict["profiles"]
}

func GetAssetDataDictionary() interface{} {
	dict["assets"] = assetsType
	return dict["assets"]
}

func GetAccStatusDictionary() interface{} {
	dict["status"] = accountStatus
	return dict["status"]
}

func GetProfileId(Key string, dictionary interface{}) uint16 {
	i := dictionary.(map[string]uint16)
	return i[Key]
}

func GetAssetTypeId(Key string, dictionary interface{}) uint16 {
	i := dictionary.(map[string]uint16)
	return i[Key]
}

func GetAccStatusId(Key string, dictionary interface{}) uint8 {
	i := dictionary.(map[string]uint8)
	return i[Key]
}
