package ip

import (
	"fmt"
	"testing"
)

func Test_ip(t *testing.T) {
	ip1 := "192.168.1.1"
	int_ip1, _ := ConvertToIntIP(ip1)
	ip2 := "192.168.1.255"
	int_ip2, _ := ConvertToIntIP(ip2)
	ip3 := "192.168.2.0"
	int_ip3, _ := ConvertToIntIP(ip3)
	fmt.Println(int_ip1)
	fmt.Println(ipAddress(ip1))
	fmt.Println(int_ip2)
	fmt.Println(ipAddress(ip2))
	fmt.Println(int_ip3)
	fmt.Println(ipAddress(ip3))

	fmt.Println(sameIpAddress(ip1, ip2))
	fmt.Println(sameIpAddress(ip1, ip3))
	fmt.Println(sameIpAddress(ip2, ip3))
}
