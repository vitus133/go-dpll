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

// getPinIDCmd represents the getPinID command
var getPinIDCmd = &cobra.Command{
	Use:   "getPinID",
	Short: "Get pin ID by clock ID and board label",
	Long: `This command will get pin ID by clock ID and board label:
	
	dpll-cli getPinID --clockID=0x507c6fffff1fb484 --boardLabel=GNSS-1PPS`,
	Run: func(cmd *cobra.Command, args []string) {
		clockID, _ := cmd.Flags().GetUint64("clockID")
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
			if pin.BoardLabel == boardLabel && pin.ClockID == clockID {
				fmt.Println(pin.ID)
				return
			}
		}
		fmt.Println("pin with clockID ", clockID, " and board label ", boardLabel, "not found")
	},
}

func init() {
	rootCmd.AddCommand(getPinIDCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPinIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getPinIDCmd.Flags().StringP("boardLabel", "l", "", "BoardLabel of the pin")
	getPinIDCmd.MarkFlagRequired("boardLabel")
	getPinIDCmd.Flags().Uint64P("clockID", "i", 0, "Clock ID")
	getPinIDCmd.MarkFlagRequired("clockID")
}
