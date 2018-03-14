package store

var accessCode string

func GetAccessCode() string {
	return accessCode
}

func SetAccessCode(code string) {
	accessCode = code
}
