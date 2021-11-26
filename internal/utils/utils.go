package utils

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

//array indexof
func ValueIsInArray(a []string, b string) bool {
	for _, value := range a {
		if value == b {
			return true
		}
	}

	return false
}

func GetArrayIndexOf(a []string, b string) int {
	for index, value := range a {
		if value == b {
			return index
		}
	}

	return -1
}

func ChangeGjsonArrayToString(Result []gjson.Result) string {
	var data string

	for _, val := range Result {
		data += "," + fmt.Sprintf("%v", val)
	}

	data = strings.TrimPrefix(data, ",")
	return data
}
