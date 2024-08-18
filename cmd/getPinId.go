/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

// getPinIdCmd represents the getPinId command
var getPinIdCmd = &cobra.Command{
	Use:   "getPinId",
	Short: "Get pin ID by clock ID and board label",
	Long: `This command will get pin ID by clock ID and board label:
	
	dpll-cli getPinId --clockId=0x507c6fffff1fb484 --boardLabel=GNSS-1PPS`,
	Run: func(cmd *cobra.Command, args []string) {
		clockId, _ := cmd.Flags().GetUint64("clockId")
		boardLabel, _ := cmd.Flags().GetString("boardLabel")

		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		pinReplies, err := conn.DumpPinGet()
		if err != nil {
			log.Panic(err)
		}
		for _, pin := range pinReplies {
			if pin.BoardLabel == boardLabel && pin.ClockId == clockId {
				fmt.Println(pin.Id)
				return
			}
		}
		fmt.Println("pin with clockId ", clockId, " and board label ", boardLabel, "not found")
	},
}

func init() {
	rootCmd.AddCommand(getPinIdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPinIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getPinIdCmd.Flags().StringP("boardLabel", "l", "", "BoardLabel of the pin")
	getPinIdCmd.MarkFlagRequired("boardLabel")
	getPinIdCmd.Flags().Uint64P("clockId", "i", 0, "Clock ID")
	getPinIdCmd.MarkFlagRequired("clockId")
}
