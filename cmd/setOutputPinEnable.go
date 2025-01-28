/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

// setOutputPinEnableCmd represents the setOutputPinEnable command
var setOutputPinEnableCmd = &cobra.Command{
	Use:   "setOutputPinEnable",
	Short: "Enable or disable an output pin",
	Long: `TEnables or disables the specified output pin:

User must provide pin-parent-device parent-id and enable flag. Can be specified multiple times.
For example:
dpll-cli setOutputPinEnable -d 50 -i 6 -i 7 -e=false
parentId (-e) can be set multiple times - the same operation will be applied to all.`,
	Run: func(cmd *cobra.Command, args []string) {
		en := dpll.PinOutputs{
			Id:      PinId,
			Outputs: make([]dpll.PinOutput, len(ParentId)),
		}
		for i, p := range ParentId {
			par, err := strconv.ParseUint(p, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			en.Outputs[i].PinParentId = uint32(par)
			en.Outputs[i].Enable = *En
		}
		log.Println(en)
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		err = conn.PinOutputControlRequest(en)
		if err != nil {
			log.Fatal(err)
		}
	},
}
var En *bool

func init() {
	rootCmd.AddCommand(setOutputPinEnableCmd)

	setOutputPinEnableCmd.Flags().Uint32VarP(&PinId, "pinId", "d", 0, "Pin ID")
	setInputPinPrioCmd.MarkFlagRequired("pinId")
	setOutputPinEnableCmd.Flags().StringSliceVarP(&ParentId, "parentId", "i", []string{}, "Pin parent ID")
	setInputPinPrioCmd.MarkFlagRequired("parentId")
	En = setOutputPinEnableCmd.Flags().BoolP("enable", "e", true, "enable or disable")
}
