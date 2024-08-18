/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

// setPhaseAdjustCmd represents the setPhaseAdjust command
var setPhaseAdjustCmd = &cobra.Command{
	Use:   "setPhaseAdjust",
	Short: "set phase adjustment of a pin",
	Long:  `This will set phase adjustment of a pin specified by id`,
	Run: func(cmd *cobra.Command, args []string) {
		pinId, _ := cmd.Flags().GetUint32("pinId")
		phaseAdjust, _ := cmd.Flags().GetInt32("phaseAdjust")
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		err = conn.PinPhaseAdjust(dpll.PinPhaseAdjustRequest{Id: pinId, PhaseAdjust: phaseAdjust})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setPhaseAdjustCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setPhaseAdjustCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setPhaseAdjustCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setPhaseAdjustCmd.Flags().Uint32P("pinId", "i", 0, "Pin ID")
	setPhaseAdjustCmd.MarkFlagRequired("pinId")
	setPhaseAdjustCmd.Flags().Int32P("phaseAdjust", "p", 0, "Phase adjustment in ps")
	setPhaseAdjustCmd.MarkFlagRequired("phaseAdjust")
}
