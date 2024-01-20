package helpers

func FormatDataResponse(message string, data map[string]interface{}) (map[string]interface{}) {
	return map[string]interface{}{"message": message, "data": data}
}