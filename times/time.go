package times

import (
	"log"
	"time"
)

// ParseRFC3339TimeStrToTimestamp 将ISO8601时间转为当地时间戳
func ParseRFC3339TimeStrToTimestamp(layout, ios8601 string) int64 {
	result, err := time.ParseInLocation(layout, ios8601, time.Local)
	//如果错误则退出
	if err != nil {
		log.Println("ParseRFC3339TimeStrToTimestamp Err:", err)
		return -1
	}

	return result.Unix()
}
