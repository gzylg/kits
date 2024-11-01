package ip

import (
	"errors"
	"fmt"
	"github.com/gzylg/kits/random"
	"github.com/imroc/req/v3"
	"net"
	"strconv"
	"strings"
)

// GetExternalIP 获取公网ip，使用：https://ip.shuzilm.cn/ip?pkg=com.xxxxx.xxxxx
func GetExternalIP() (string, error) {
	return getExternalIP(fmt.Sprintf("https://ip.shuzilm.cn/ip?pkg=com.%s.%s", random.Str(5), random.Str(5)))
}

// GetExternalIP1 获取公网ip，使用：http://v4.ip.zxinc.org/getip
func GetExternalIP1() (string, error) {
	return getExternalIP("http://v4.ip.zxinc.org/getip")
}

// GetExternalIP2 获取公网ip，使用：https://ip.3322.net
func GetExternalIP2() (string, error) {
	return getExternalIP("https://ip.3322.net")
}

// GetExternalIP3 获取公网ip，使用：https://4.ipw.cn
func GetExternalIP3() (string, error) {
	return getExternalIP("https://4.ipw.cn")
}

// GetExternalIP4 获取公网ip，使用：https://v4.myip.la
// Deprecated: 失效
func GetExternalIP4() (string, error) {
	return getExternalIP("https://v4.myip.la")
}

/*
	TODO:
	  V4
		https://ddns.oray.com/checkip
		https://myip4.ipip.net
		https://www.taobao.com/help/getip.php
		http://txt.go.sohu.com/ip/soip
	  V6
		https://ipv6.ddnspod.com
		https://6.ipw.cn
		http://v6.ip.zxinc.org/getip
		https://speed.neu6.edu.cn/getIP.php
		https://v6.ident.me
		https://v6.myip.la
*/

func getExternalIP(uri string) (string, error) {
	var (
		client = req.C().R()
	)
	resp, err := client.Get(uri)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

//// GetExternalIP3 获取本机外网IP 第 1 快
//func GetExternalIP3() (ip string, err error) {
//	var (
//		// NewWithClient：每次重定向都将进入 redirectPostOn302 函数
//		//req  = resty.NewWithClient(&http.Client{CheckRedirect: redirectPostOn302}).SetTimeout(time.Second * 5).R()
//		req  = resty.New().SetTimeout(time.Second * 2).R()
//		resp *resty.Response
//		uri  = "http://txt.go.sohu.com/ip/soip"
//	)
//	if resp, err = req.Get(uri); err != nil {
//		return "", err
//	}
//	if resp.StatusCode() != 200 {
//		return "", errors.New(fmt.Sprintf("get ip failed: response status code is %d", resp.StatusCode()))
//	}
//
//	ip, ok := getAndCheckIP(string(resp.Body()), `(\d+.\d+.\d+.\d+)`)
//	if !ok {
//		return "", errors.New("get ip failed: can not find ip")
//	}
//	return ip, nil
//}
//
//func getAndCheckIP(str, reStr string) (ip string, ok bool) {
//	re := regexp.MustCompile(reStr)
//	matched := re.FindAllStringSubmatch(str, -1)
//
//	if len(matched) != 1 {
//		return "", false
//	}
//
//	for _, match := range matched {
//		if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", match[1]); m {
//			ip = match[1]
//			return ip, true
//		}
//	}
//	return "", false
//}

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

// Ip2Uint32 net.IP到整型
func Ip2Uint32(ipnr net.IP) uint32 {
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

// IpStr2Uint32 string Ip到整型
func IpStr2Uint32(ip string) uint32 {
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
