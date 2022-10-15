package tid

import (
	"github.com/gzylg/kits/now"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	DateLayout string   // 日期格式化模板，如：20060102
	TimeUnit   TimeUnit // 时间单位
	TidLength  int      // tid的总长度
}

func New(c Config) *timeId {
	// 默认日期格式化模板
	if c.DateLayout == "" {
		c.DateLayout = "060102"
	}
	// 默认事件单位
	if c.TimeUnit == 0 {
		c.TimeUnit = TimeUnitMilliSecond
	}

	// * 计算当前配置中 日期+时间 的字符串长度
	lenTid := len(c.DateLayout) + timeLength[c.TimeUnit]
	if c.TidLength-lenTid < 1 { // 如果传入的总长度-计算出的日期+时间的字符串长度<1
		c.TidLength = lenTid + 1 // 重新设置总长度
	}

	return &timeId{conf: &c, lock: new(sync.Mutex)}
}

type timeId struct {
	date       string
	passedTime string
	count      int

	timestamp int64 // 记录当前时间的时间戳，配置中单位为哪个就记录哪种类型的时间戳
	lock      *sync.Mutex

	conf *Config
}
type TimeUnit uint32

const (
	TimeUnitSecond      TimeUnit = 1 << iota // 1 秒为单位
	TimeUnitMilliSecond                      // 2 毫秒为单位
	TimeUnitMicroSecond                      // 4 微秒为单位
	TimeUnitNanoSecond                       // 8 纳秒为单位
)

// 每种时间单位转换成字符串后的最大长度
var timeLength = map[TimeUnit]int{
	TimeUnitSecond:      5,  // 每天最大 86399 秒，长度：5
	TimeUnitMilliSecond: 8,  // 每天最大 86399999 毫秒，长度：8
	TimeUnitMicroSecond: 11, // 每天最大 86399999999 微秒，长度：11
	TimeUnitNanoSecond:  14, // 每天最大 86399999999999 纳秒，长度：14
}

func (tid *timeId) General() string {
	// * 加锁
	tid.lock.Lock()
	defer tid.lock.Unlock()

	timeNow := time.Now()
	// * 生成日期
	tid.date = timeNow.Format(tid.conf.DateLayout)

	// * 生成时间
	beginningOfDay := now.BeginningOfDay()
	var passedTime, timestamp int64
	switch tid.conf.TimeUnit {
	case TimeUnitSecond:
		timestamp = timeNow.Unix()
		passedTime = timestamp - beginningOfDay.Unix()
	case TimeUnitMilliSecond:
		timestamp = timeNow.UnixMilli()
		passedTime = timeNow.UnixMilli() - beginningOfDay.UnixMilli()
	case TimeUnitMicroSecond:
		timestamp = timeNow.UnixMicro()
		passedTime = timeNow.UnixMicro() - beginningOfDay.UnixMicro()
	case TimeUnitNanoSecond:
		timestamp = timeNow.UnixNano()
		passedTime = timeNow.UnixNano() - beginningOfDay.UnixNano()
	}
	tid.passedTime = strconv.FormatInt(passedTime, 10)
	// 补全位数
	for {
		if len(tid.passedTime) == timeLength[tid.conf.TimeUnit] {
			break
		}
		tid.passedTime = "0" + tid.passedTime
	}

	// * 计算次数
	if timestamp != tid.timestamp { // 如果当前时间不等于上次记录的时候
		tid.count = 0 // 将计数归零
	} else { // 否则
		tid.count++ // 计数递增
	}

	tid.timestamp = timestamp // 完成后，将记录中的时间戳改为当前时间戳

	// * 计算次数长度
	lastLen := tid.conf.TidLength - len(tid.date) - len(tid.passedTime) // 总长度-日期长度-时间长度=计数长度
	countStr := strconv.Itoa(tid.count)

	if lastLen == len(countStr) {
		return tid.date + tid.passedTime + countStr
	}

	if lastLen > len(countStr) {
		for {
			countStr = "0" + countStr
			if len(countStr) == lastLen {
				return tid.date + tid.passedTime + countStr
			}
		}
	}

	c := len(countStr) - lastLen
	a := countStr[len(countStr)-c:]
	return tid.date + tid.passedTime + a
}
