package file

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	rwLock    sync.RWMutex
	sizeBytes = 1024 * 1024 * 500 //500 M
)

func CreateFile(folderPath, fileName string, count int) *os.File {
	folderPath = strings.TrimRight(folderPath, "/")
	fileName = strings.TrimLeft(fileName, "/")
	if FileIsExist(folderPath) == false {
		errDir := os.MkdirAll(folderPath, 0755)
		if errDir != nil {
			log.Panic(errDir)
		}
	}

	filePath := fmt.Sprintf("%v/%v", folderPath, fileName)

	if count != 0 {
		strCount := strconv.Itoa(count)
		strLen := strings.LastIndex(fileName, ".")
		newFileName := ""
		if strLen > 0 {
			newFileName = fileName[:strLen] + strCount + fileName[strLen:]
		} else {
			newFileName = fileName + strCount
		}
		fmt.Println(newFileName)
		filePath = fmt.Sprintf("%v/%v", folderPath, newFileName)
	}

	// fmt.Println(filePath)

	rwLock.Lock() //創建file or folder 前先鎖
	defer rwLock.Unlock()
	if FileIsExist(filePath) == false {
		File, err := os.Create(filePath)
		if err != nil {
			log.Panic(err)
		}
		return File

	}

	if FileSizeOver(sizeBytes, filePath) {
		count++
		CreateFile(folderPath, fileName, count)
	} else {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Panic(err)
		}
		return f

	}
	return nil
}

func WriteFile(content, folderPath, fileName string, count int) {
	rwLock.Lock() //寫入 file 前先鎖
	defer rwLock.Unlock()
	f := CreateFile(folderPath, fileName, count)
	defer f.Close()
	_, err := f.WriteString(content + "\n")
	if err != nil {
		log.Panic(err)
	}
}

func FileIsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func FileSizeOver(sizeBytes int, filePath string) bool {
	file, _ := os.Stat(filePath)
	// if os.IsNotExist(err) {
	// 	return false
	// }

	if fileSize := file.Size(); (fileSize) > int64(sizeBytes) {
		return true
	}
	return false
}
