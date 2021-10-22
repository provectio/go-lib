package interfaces

import "strconv"

func InterfaceToBool(MyInterface interface{}) bool {
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

func InterfaceToFloat64(MyInterface interface{}) float64 {
	if MyInterface == nil {
		return 0
	}
	switch i := MyInterface.(type) {
	case float64:
		return i
	case int:
		return float64(i)
	case string:
		MyFloat64, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0
		}
		return MyFloat64
	case bool:
		if i {
			return 1
		}
		return 0
	default:
		return 0
	}
}
