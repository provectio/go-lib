package slice

import "sort"

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}

func Dedup(slice []string) (result []string) {

	sort.Strings(slice)

	j := 0

	for i := 1; i < len(slice); i++ {
		if slice[j] == slice[i] {
			continue
		}

		j++

		slice[j] = slice[i]
	}

	result = slice[:j+1]

	return
}

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ReduceString(str string, num int) string {
	res := str

	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		res = str[0:num]
	}

	return res
}
