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
		MyBool, err := strconv.ParseBool(MyInterface.(string))
		if err != nil {
			return false
		}
		return MyBool
	default:
		return false
	}
}
