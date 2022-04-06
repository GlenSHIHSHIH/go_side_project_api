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

func filterBase64(buffer string) string {
	filter := "base64,"
	return buffer[strings.Index(buffer, filter)+len(filter):]
}

// save picture
func SavePicture(fileName, buffer string) error {
	data := filterBase64(buffer)
	picData, _ := base64.StdEncoding.DecodeString(data) //成图片文件并把文件写入到buffer
	return ioutil.WriteFile(fileName, picData, 0666)    //buffer输出到jpg文件中（不做处理，直接写到文件）
}

//判断图片base64流字节大小

//如何使用java判断图片base64字节流的大小，以及计算后是多少KB。

/**
 *通过图片base64流判断图片等于多少字节
 *image 图片流
 */
func GetImageSize(image string) int {
	str := filterBase64(image)             // 1.需要计算文件流大小，首先把头部的data:image/png;base64,（注意有逗号）去掉。
	equalIndex := strings.Index(str, "==") //2.找到等号，把等号也去掉
	if equalIndex >= 0 {
		str = str[0:equalIndex]
	}
	strLength := len(str)     //3.原来的字符流大小，单位为字节
	size := strLength * 6 / 8 //4.计算后得到的文件流大小，单位为字节
	return size
}
