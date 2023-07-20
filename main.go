package main

import (
	"fmt"
	"log"
	"net"

	pci "github.com/u-root/u-root/pkg/pci"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

func main() {
	pcii, err := pci.OnePCI("/sys/class/net/ens2f0/device/")
	if err != nil {
		log.Println(err)
	}

	const offset int64 = 340
	data, err := pcii.ReadConfigRegister(offset, 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("ens2f0 clock ID: %d\n", data)

	pcii, err = pci.OnePCI("/sys/class/net/ens2f1/device/")
	if err != nil {
		log.Println(err)
	}

	data, err = pcii.ReadConfigRegister(offset, 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("ens2f1 clock ID: %d\n", data)

	var iface = "ens2f0"
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, ifc := range interfaces {
		if ifc.Name == iface {
			iface = ifc.HardwareAddr.String()
			log.Println(len(ifc.HardwareAddr))
			log.Println(ifc.HardwareAddr)
		}
	}

	mac, err := net.ParseMAC(iface)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mac)
	conn, err := dpll.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mcastId, found := conn.GetMcastGroupId(dpll.DPLL_MCGRP_MONITOR)
	if !found {
		log.Panic("multicast ID ", dpll.DPLL_MCGRP_MONITOR, " not found")
	}

	replies, err := conn.DumpDeviceGet()
	if err != nil {
		log.Panic(err)
	}
	for _, reply := range replies {
		log.Println(dpll.GetDpllStatusHR(reply))
	}

	c := conn.GetGenetlinkConn()
	err = c.JoinGroup(mcastId)
	if err != nil {
		log.Panic(err)
	}
	for {
		msgs, _, err := c.Receive()
		if err != nil {
			log.Panic(err)
		}
		ntfs, err := dpll.ParseDeviceReplies(msgs)
		if err != nil {
			log.Panic(err)
		}
		for _, ntf := range ntfs {
			log.Println(dpll.GetDpllStatusHR(ntf))
		}
	}
}
