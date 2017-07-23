package util

const (
	HTTP_SUCCESS = 0
	HTTP_PARAM_ERROR = 1
	HTTP_FAILED = 2
	HTTP_PARAM_MODULE_NOT_EXIST = 3
	HTTP_SAVEMESSAGE_FAILED = 4
)

var codeString = map[int]string{
	HTTP_SUCCESS 				: "success",
	HTTP_PARAM_ERROR 			: "param error",
	HTTP_FAILED					: "failed",
	HTTP_SAVEMESSAGE_FAILED		: "save message failed",
	HTTP_PARAM_MODULE_NOT_EXIST : "module not exist",
}

func GetCodeString(code int) string{
	if str,ok := codeString[code]; ok{
		return str
	}
	return ""
}