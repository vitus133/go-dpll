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

// setInputPinPrioCmd represents the setInputPinPrio command
var setInputPinPrioCmd = &cobra.Command{
	Use:   "setInputPinPrio",
	Short: "Sets input pin priority",
	Long: `Sets input pin priority:

User must provide pin-parent-device parent-id and priority. Can be specified multiple times.
For example:
dpll-cli setInputPinPrio -i 0 -p 1 -i 1 -p 255`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(ParentId) != len(Prio) {
			log.Fatalf("number of parent ids and priorities mismatch")
		}
		pp := dpll.PinPriorities{
			Id:         PinId,
			Priorities: make([]dpll.PinPriority, len(ParentId)),
		}
		for i, prio := range Prio {
			iPrio, err := strconv.ParseUint(prio, 10, 32)
			if err != nil {
				log.Fatalf("failed to parse priority")
			}
			pp.Priorities[i].Prio = uint32(iPrio)
			iId, err := strconv.ParseUint(ParentId[i], 10, 32)
			if err != nil {
				log.Fatalf("failed to parse parent ID")
			}
			pp.Priorities[i].PinParentId = uint32(iId)

		}
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		err = conn.PinInputPrioRequest(pp)
		if err != nil {
			log.Fatal(err)
		}
	},
}
var PinId uint32
var ParentId []string
var Prio []string

func init() {
	rootCmd.AddCommand(setInputPinPrioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setInputPinPrioCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	setInputPinPrioCmd.Flags().Uint32VarP(&PinId, "pinId", "d", 0, "Pin ID")
	setInputPinPrioCmd.MarkFlagRequired("pinId")
	setInputPinPrioCmd.Flags().StringSliceVarP(&ParentId, "parentId", "i", []string{}, "Pin parent ID")
	setInputPinPrioCmd.MarkFlagRequired("parentId")
	setInputPinPrioCmd.Flags().StringSliceVarP(&Prio, "prio", "p", []string{}, "Pin priority")
	setInputPinPrioCmd.MarkFlagRequired("prio")
}
