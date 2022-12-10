package gosupport

import "net"

// 获取本机ip,一般取第0个单元作为本机ip,示例：ip := gosupport.GetLocalIPV1()[0]
func GetLocalIPV1() []string {
	ret := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ret
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ret = append(ret, ipnet.IP.String())
			}
		}
	}
	return ret
}
