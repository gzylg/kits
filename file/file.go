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
	p, _ := filepath.Abs(os.Args[0])
	return p
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
	for _, p := range paths {
		if fullpath = filepath.Join(p, filename); Exists(fullpath) {
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

// WriteStr 写出string数据到文件。 若路径不存在，则自动创建。
func WriteStr(FilePath string, s string) (err error) {
	return WriteByte(FilePath, []byte(s))
}

func Base64Type(t string) string {
	switch t {
	case "txt":
		return "data:text/plain;base64,"
	case "doc":
		return "data:application/msword;base64,"
	case "docx":
		return "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,"
	case "xls":
		return "data:application/vnd.ms-excel;base64,"
	case "xlsx":
		return "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,"
	case "pdf":
		return "data:application/pdf;base64,"
	case "pptx":
		return "data:application/vnd.openxmlformats-officedocument.presentationml.presentation;base64,"
	case "ppt":
		return "data:application/vnd.ms-powerpoint;base64,"
	case "png":
		return "data:image/png;base64,"
	case "jpg":
		return "data:image/jpeg;base64,"
	case "gif":
		return "data:image/gif;base64,"
	case "svg":
		return "data:image/svg+xml;base64,"
	case "ico":
		return "data:image/x-icon;base64,"
	case "bmp":
		return "data:image/bmp;base64,"
	}

	return ""
}

// GetFilesFormPath 枚举指定目录下的文件，返回文件列表
func GetFilesFormPath(p string) ([]string, error) {
	// 读取目录
	files, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}

	// 枚举文件
	var list []string
	for _, file := range files {
		list = append(list, file.Name())
	}

	return list, nil
}
