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

// getPinCmd represents the getPin command
var getPinCmd = &cobra.Command{
	Use:   "getPin",
	Short: "Get pin properties by ID",
	Long: `Returns pin properties by pin ID
	
	dpll-cli getPin --pinId=23`,
	Run: func(cmd *cobra.Command, args []string) {
		pinId, err := cmd.Flags().GetUint32("pinId")
		if err != nil {
			log.Fatal(err)
		}
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		pinReply, err := conn.DoPinGet(dpll.DoPinGetRequest{Id: pinId})
		if err != nil {
			log.Panic(err)
		}
		timestamp := time.Now().UTC()

		pinInfo, err := dpll.GetPinInfoHR(pinReply, timestamp)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(pinInfo))

	},
}

func init() {
	rootCmd.AddCommand(getPinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getPinCmd.Flags().Uint32P("pinId", "i", 0, "Pin ID")
	getPinCmd.MarkFlagRequired("pinId")
}
