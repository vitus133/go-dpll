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

// dumpPinsCmd represents the dumpPins command
var dumpPinsPneByOneCmd = &cobra.Command{
	Use:   "dumpPinsOneByOne",
	Short: "Dump DPLL pins one by one",
	Long:  `Dumps all DPLL pins one by one in JSON format`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		for id := 0; ; id++ {
			pinReply, err := conn.DoPinGet(dpll.DoPinGetRequest{Id: uint32(id)})
			if err != nil {
				return
			}
			timestamp := time.Now().UTC()
			pinInfo, err := dpll.GetPinInfoHR(pinReply, timestamp)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println(string(pinInfo))
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpPinsPneByOneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpPinsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpPinsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
