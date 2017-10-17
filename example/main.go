package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/yinqiwen/fdns"
	"github.com/yinqiwen/gotoolkit/cip"
)

func main() {
	ccip, _ := cip.LoadIPSet("../../gotoolkit/cip/cnipset.txt", "CN")
	config := &fdns.Config{}
	config.IsCNIP = func(ip net.IP) bool {
		return ccip.IsInCountry(ip, "CN")
	}

	dns, _ := fdns.NewTrustedDNS(config)
	start := time.Now()
	rec, err := dns.LookupA(os.Args[1])
	log.Printf("Cost %v to get result:%v %v", time.Now().Sub(start), rec, err)
	// err = dns.Start()
	// if nil != err {
	// 	log.Fatal(err)
	// }
	// time.Sleep(1000000 * time.Second)
	r, _ := net.LookupHost("HUAWEI_Mate_9_Pro")
	log.Printf("###%v", r)

}
