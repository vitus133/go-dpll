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
	"k8s.io/utils/cpuset"
)

// getPinCmd represents the getPin command
var getPinCmd = &cobra.Command{
	Use:   "getPin",
	Short: "Get pin properties by ID",
	Long: `Returns pin properties by pin ID
	
	dpll-cli getPin --pinId=23. This also accepts a number of IDs in a CpuSet format, e.g. 1,3-8,19`,
	Run: func(cmd *cobra.Command, args []string) {
		inp, err := cmd.Flags().GetString("pinId")
		if err != nil {
			log.Fatal(err)
		}
		ids, err := parsePinIds(inp)
		if err != nil {
			log.Fatal(err)
		}
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		for _, id := range ids {
			pinReply, err := conn.DoPinGet(dpll.DoPinGetRequest{Id: uint32(id)})
			if err != nil {
				log.Panic(err)
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
	rootCmd.AddCommand(getPinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getPinCmd.Flags().StringP("pinId", "p", "", "Pin ID(s)")
	getPinCmd.MarkFlagRequired("pinId")
}

func parsePinIds(input string) ([]int, error) {
	l, err := cpuset.Parse(input)
	if err != nil {
		return []int{}, err
	}
	return l.List(), nil
}
