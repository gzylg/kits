package ip

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// // 第 3 快
// func GetExternalIP1() (ip string, err error) {
// 	agent := yclient.New()
// 	_, ip, _ = agent.Get("https://myexternalip.com/raw").End()
// 	if len(agent.Errors) > 0 {
// 		return "", agent.Errors[0]
// 	}
// 	ip, ok := getAndCheckIP(ip, `(\d+.\d+.\d+.\d+)`)
// 	if !ok {
// 		return "", errors.New("can not get Ip.")
// 	}
// 	return ip, nil
// }

// // 第 2 快
// func GetExternalIP2() (ip string, err error) {
// 	agent := yclient.New()
// 	_, str, _ := agent.Get("http://www.net.cn/static/customercare/yourip.asp").End()
// 	if len(agent.Errors) > 0 {
// 		return "", agent.Errors[0]
// 	}
// 	ip, ok := getAndCheckIP(str, `>(\d+.\d+.\d+.\d+)<`)
// 	if !ok {
// 		return "", errors.New("can not get Ip.")
// 	}
// 	return ip, nil
// }

// GetExternalIP3 获取本机外网IP 第 1 快
func GetExternalIP3() (ip string, err error) {
	var (
		// NewWithClient：每次重定向都将进入 redirectPostOn302 函数
		//req  = resty.NewWithClient(&http.Client{CheckRedirect: redirectPostOn302}).SetTimeout(time.Second * 5).R()
		req  = resty.New().SetTimeout(time.Second * 2).R()
		resp *resty.Response
		uri  = "http://txt.go.sohu.com/ip/soip"
	)
	if resp, err = req.Get(uri); err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", errors.New(fmt.Sprintf("get ip failed: response status code is %d", resp.StatusCode()))
	}

	ip, ok := getAndCheckIP(string(resp.Body()), `(\d+.\d+.\d+.\d+)`)
	if !ok {
		return "", errors.New("get ip failed: can not find ip")
	}
	return ip, nil
}

func getAndCheckIP(str, reStr string) (ip string, ok bool) {
	re := regexp.MustCompile(reStr)
	matched := re.FindAllStringSubmatch(str, -1)

	if len(matched) != 1 {
		return "", false
	}

	for _, match := range matched {
		if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", match[1]); m {
			ip = match[1]
			return ip, true
		}
	}
	return "", false
}

// Uint32ToIP 整型到net.IP
func Uint32ToIP(intIP uint32) net.IP {
	var bytes [4]byte
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// Uint32ToIPStr 整型到string
func Uint32ToIPStr(intIP uint32) string {
	return Uint32ToIP(intIP).String()
}

// IPToUint32 net.IP到整型
func IPToUint32(ipnr net.IP) uint32 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}

// IPStrToUint32 string Ip到整型
func IPStrToUint32(ip string) uint32 {
	bits := strings.Split(ip, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}

// GetLocalIP 获取内网ip
func GetLocalIP() (string, error) {
	info, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range info {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", errors.New("valid local IP not found")
}
