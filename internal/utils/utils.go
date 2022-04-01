package utils

import (
	"componentmod/internal/utils/log"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
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

// []string to []interfacer{}
func ChangeStringToInterfaceArr(old []string) []interface{} {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}
	return new
}

// []int to []string
func ChangeIntToStringArr(old []int) []string {
	var new []string
	for _, i := range old {
		new = append(new, strconv.Itoa(i))
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

//file name (uuid + timestamp)
func GetUuidAndTimestamp() string {

	uuid := uuid.New()
	key := uuid.String()

	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)

	return key + timestamp
}

// save picture
func SavePicture(fileName, buffer string) error {
	picData, _ := base64.StdEncoding.DecodeString(buffer) //成图片文件并把文件写入到buffer
	return ioutil.WriteFile(fileName, picData, 0666)      //buffer输出到jpg文件中（不做处理，直接写到文件）
}
