package util

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

func FileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// FilePutContents file_put_contents
func FilePutContents(filename string, data []byte) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}


func Mkdir(name string, mode os.FileMode) bool {
	err := os.Mkdir(name, mode)
	return err == nil
}

func MkdirAll(filePath string,mode os.FileMode)  error  {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath,mode)
		return err
	}
	return nil
}

func IsExist(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func UrlDecode(s string) string {
	r, err := url.QueryUnescape(s)
	if err != nil {
		return ""
	}
	return r
}

func UrlEncode(s string) string {
	return url.QueryEscape(s)
}
