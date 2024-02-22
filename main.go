package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mdlayher/genetlink"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

func main() {
	conn, err := dpll.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Dumping devices")
	replies, err := conn.DumpDeviceGet()
	if err != nil {
		log.Panic(err)
	}
	for _, reply := range replies {
		log.Println(dpll.GetDpllStatusHR(reply))
	}
	log.Println("dumping pins")
	pinReplies, err := conn.DumpPinGet()
	if err != nil {
		log.Panic(err)
	}
	for _, pinInfo := range pinReplies {
		pinInfo, err := dpll.GetPinInfoHR(pinInfo)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(pinInfo))
	}
	log.Println("monitoring notifications")
	mcastId, found := conn.GetMcastGroupId(dpll.DPLL_MCGRP_MONITOR)
	if !found {
		log.Panic("multicast ID ", dpll.DPLL_MCGRP_MONITOR, " not found")
	}

	c := conn.GetGenetlinkConn()
	err = c.JoinGroup(mcastId)
	if err != nil {
		log.Panic(err)
	}

	for {
		c.SetReadDeadline(time.Now().Add(3 * time.Second))
		msgs, _, err := c.Receive()
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				continue
			}
			log.Println(err)
			continue
		}
		for _, msg := range msgs {
			switch msg.Header.Command {
			case dpll.DPLL_CMD_DEVICE_CHANGE_NTF:
				ntfs, err := dpll.ParseDeviceReplies([]genetlink.Message{msg})
				if err != nil {
					log.Panic(err)
				}
				for _, ntf := range ntfs {
					log.Println("")
					fmt.Println(dpll.GetDpllStatusHR(ntf))
				}
			case dpll.DPLL_CMD_PIN_CHANGE_NTF:
				ntfs, err := dpll.ParsePinReplies([]genetlink.Message{msg})
				if err != nil {
					log.Panic(err)
				}
				for _, ntf := range ntfs {
					hr, err := dpll.GetPinInfoHR(ntf)
					if err != nil {
						log.Panic(err)
					}
					log.Println("")
					fmt.Println(string(hr))
				}
			default:
				log.Println("unsupported dpll message")

			}

		}
	}
}
