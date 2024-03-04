/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mdlayher/genetlink"
	"github.com/spf13/cobra"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitors DPLL events",
	Long: `This commands blocks until cancelled with ^C. 
	
	Dumps all DPLL notifications in JSON format`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

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
			timestamp := time.Now().UTC()
			for _, msg := range msgs {
				switch msg.Header.Command {
				case dpll.DPLL_CMD_DEVICE_CHANGE_NTF:
					ntfs, err := dpll.ParseDeviceReplies([]genetlink.Message{msg})
					if err != nil {
						log.Panic(err)
					}
					for _, ntf := range ntfs {
						dev, err := dpll.GetDpllStatusHR(ntf, timestamp)
						if err != nil {
							log.Panic(err)
						}
						fmt.Println(string(dev))
					}
				case dpll.DPLL_CMD_PIN_CHANGE_NTF:
					ntfs, err := dpll.ParsePinReplies([]genetlink.Message{msg})
					if err != nil {
						log.Panic(err)
					}
					for _, ntf := range ntfs {
						hr, err := dpll.GetPinInfoHR(ntf, timestamp)
						if err != nil {
							log.Panic(err)
						}
						fmt.Println(string(hr))
					}
				default:
					log.Println("unsupported dpll message")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
