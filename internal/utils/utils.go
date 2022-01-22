package utils

import (
	"componentmod/internal/utils/log"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
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

func ValueIsInIntArray(a []int, b int) bool {
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

// string[] to []interfacer{}
func ChangeStringToInterfaceArr(old []string) []interface{} {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}
	return new
}

// 拿取.env 參數
func GetEnvParameterByName(pName string) string {
	err := godotenv.Load()
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), "Error loading .env file")
		log.Error(fmt.Sprintf("%+v", errData))
	}

	return os.Getenv(pName)
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
