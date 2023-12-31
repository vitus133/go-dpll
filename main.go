package main

import (
	"log"
	"strings"
	"time"

	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

func main() {
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
		c.SetReadDeadline(time.Now().Add(3 * time.Second))
		msgs, _, err := c.Receive()
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				continue
			}
			log.Println(err)
			continue
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
