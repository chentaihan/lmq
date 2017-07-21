package util

const (
	HTTP_SUCCESS = 0
	HTTP_PARAM_ERROR = 1
	HTTP_FAILED = 2
)

var codeString = map[int]string{
	HTTP_SUCCESS 		: "success",
	HTTP_PARAM_ERROR 	: "param error",
}

func GetCodeString(code int) string{
	if str,ok := codeString[code]; ok{
		return str
	}
	return ""
}