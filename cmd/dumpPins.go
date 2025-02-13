/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

var rawOutput *bool

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
		if *rawOutput {
			raw, err := json.Marshal(pinReplies)
			if err != nil {
				log.Panic(err)
			}
			fmt.Printf("%s\n", string(raw))
		} else {

			timestamp := time.Now().UTC()
			for _, pinInfo := range pinReplies {
				pinInfo, err := dpll.GetPinInfoHR(pinInfo, timestamp)
				if err != nil {
					log.Panic(err)
				}
				fmt.Printf("%s\n", string(pinInfo))
			}
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
	rawOutput = dumpPinsCmd.Flags().BoolP("raw", "r", false, "Print json in the raw format")

}
