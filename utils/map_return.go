package utils

func MapJson(code int64, msg string, data any) map[string]any {
	return map[string]any{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
