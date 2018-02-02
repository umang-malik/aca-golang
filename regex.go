package regex

import (
	"strconv"
	"strings"
	"unicode"
)

func filter(s string) string {
	var str string
	for _, r := range s {
		if unicode.IsLetter(r) || r == '_' || unicode.IsNumber(r) {
			str += string(r)
		}
	}
	return str
}

func isIp(s string) bool {
	values := strings.Split(s, ".")
	if len(values) != 4 {
		return false
	} else {
		for _, r := range values {
			if temp, err := strconv.Atoi(r); err == nil {
				if temp > 255 || temp < 0 {
					return false
				}
			}
		}
		return true
	}
}

//func main() {
//s := "h1233ABCDabcd___--#$#@"
//fmt.Println(filter(s))
//fmt.Println(isIp("123.22.43.00"))
//fmt.Println(isIp("4314.23.123.3333"))
//}
