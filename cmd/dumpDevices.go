/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

// dumpDevicesCmd represents the dumpDevices command
var dumpDevicesCmd = &cobra.Command{
	Use:   "dumpDevices",
	Short: "Dump DPLL devices",
	Long:  `Dumps all DPLL devices in JSON format`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		replies, err := conn.DumpDeviceGet()
		if err != nil {
			log.Panic(err)
		}
		ts := time.Now().UTC()
		for _, reply := range replies {
			dev, err := dpll.GetDpllStatusHR(reply, ts)
			if err != nil {
				log.Panic(err)
			}
			fmt.Printf("%s\n", string(dev))
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpDevicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpDevicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpDevicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
