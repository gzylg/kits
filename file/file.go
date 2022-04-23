package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// SelfPath 获取运行文件的绝对路径
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// SelfDir 获取运行文件的目录
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// Exists 检查文件名或目录是否存在
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// SearchFile 在传入的目录列表中搜索文件
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		if fullpath = filepath.Join(path, filename); Exists(fullpath) {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}

// WriteByte 写出byte数据到文件。 若路径不存在，则自动创建。
func WriteByte(FilePath string, data []byte) (err error) {
	if err = os.MkdirAll(path.Dir(FilePath), os.ModePerm); err != nil {
		return
	}
	return ioutil.WriteFile(FilePath, data, 0644)
}

// WriteByte 写出string数据到文件。 若路径不存在，则自动创建。
func WriteStr(FilePath string, s string) (err error) {
	return WriteByte(FilePath, []byte(s))
}
