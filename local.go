package getmyip

import "net"

func GetLocalIP() []string {
	listInt, err := net.Interfaces()
	if err != nil {
		return make([]string, 0)
	}
	var listIP = make([]string, 0)
	for _, i2 := range listInt {
		ip, err := getIPByInterface(i2.Name)
		if err != nil {
			continue
		}
		listIP = append(listIP, ip...)
	}
	return listIP
}

func getIPByInterface(interfaceName string) ([]string, error) {
	itf, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return make([]string, 0), err
	}
	item, err := itf.Addrs()
	if err != nil {
		return make([]string, 0), err
	}
	var ips []net.IP
	for _, addr := range item {
		switch v := addr.(type) {
		case *net.IPNet:
			if !v.IP.IsLoopback() {
				if v.IP.To4() != nil { //Verify if IP is IPV4
					ips = append(ips, v.IP)
				}
			}
		}
	}

	var result = make([]string, len(ips))

	for i := range ips {
		result[i] = ips[i].String()
	}
	return result, nil
}
