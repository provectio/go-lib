package interfaces

import "strconv"

func toBool(MyInterface interface{}) bool {
	if MyInterface == nil {
		return false
	}
	switch i := MyInterface.(type) {
	case bool:
		return i
	case string:
		MyBool, err := strconv.ParseBool(i)
		if err != nil {
			return false
		}
		return MyBool
	case int:
		if i == 1 {
			return true
		}
		return false
	case float64:
		if i == 1 {
			return true
		}
		return false
	default:
		return false
	}
}
