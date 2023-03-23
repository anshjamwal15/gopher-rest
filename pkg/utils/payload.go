package utils

func Message(s bool, m string) map[string]interface{} {
	return map[string]interface{}{"status": s, "message": m}
}
