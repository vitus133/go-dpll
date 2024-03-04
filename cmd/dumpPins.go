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
var dumpPinsCmd = &cobra.Command{
	Use:   "dumpPins",
	Short: "Dump DPLL pins",
	Long:  `Dumps all DPLL pins in JSON format`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		pinReplies, err := conn.DumpPinGet()
		if err != nil {
			log.Panic(err)
		}
		timestamp := time.Now().UTC()
		for _, pinInfo := range pinReplies {
			pinInfo, err := dpll.GetPinInfoHR(pinInfo, timestamp)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println(string(pinInfo))
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpPinsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpPinsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpPinsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
