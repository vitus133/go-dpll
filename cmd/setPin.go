/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math"

	"github.com/spf13/cobra"
	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

// setPinCmd represents the setPin command
var setPinCmd = &cobra.Command{
	Use:   "setPin",
	Short: "Set pin attributes",
	Long: `Set pin attributes such as delay compensation or pin parent device 
attributes, such as priority, direction or output mode. For example:

To set pin ID 50, parent IDs 6 and 7 to disconnected:
	./dpll-cli setPin -i 50 -d 6 -d 7 -s 2,2
(Pin state: 
	PinStateConnected    = 1
	PinStateDisconnected = 2
	PinStateSelectable   = 3)
To set pin priority of pin ID 44 pin parent IDs 6 and 7 to 0:
	./dpll-cli setPin -i 44 -d 6,7 -p 0,0
To set phase adjust of pin ID 44 to 600ns:
	./dpll-cli setPin -i 44 -j 600
Flags "prio" and "pinState" are mutually exclusive.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setPin called")
		if (len(prio) > 0 && len(parentId) != len(prio)) ||
			(len(pinState) > 0 && len(parentId) != len(pinState)) {
			log.Fatalf("number of parent IDs doesn't match pin priorities or states!")
		}
		pc := dpll.PinParentDeviceCtl{
			Id:           *pinId,
			PinParentCtl: make([]dpll.PinControl, len(parentId)),
		}
		if *phaseAdjust != math.MaxInt32 {
			pc.PhaseAdjust = phaseAdjust
		}
		if *esyncFrequency != 0 {
			pc.EsyncFrequency = esyncFrequency
		}
		if *frequency != 0 {
			pc.Frequency = frequency
		}
		for i, pid := range parentId {
			pc.PinParentCtl[i].PinParentId = uint32(pid)
			if len(pinState) > 0 {
				pc.PinParentCtl[i].State = func(st uint) *uint32 {
					state := uint32(st)
					return &state
				}(pinState[i])
			} else {
				pc.PinParentCtl[i].Prio = func(p uint) *uint32 {
					prio := uint32(p)
					return &prio
				}(prio[i])
			}
		}
		conn, err := dpll.Dial(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		data, err := dpll.EncodePinControl(pc)
		if err != nil {
			log.Fatal(err)
		}
		err = conn.SendCommand(dpll.DpllCmdPinSet, data)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var pinId *uint32
var phaseAdjust *int32
var esyncFrequency *uint64
var frequency *uint64
var parentId []uint
var prio []uint
var pinState []uint

func init() {
	rootCmd.AddCommand(setPinCmd)
	pinId = setPinCmd.Flags().Uint32P("pinId", "i", 0, "Pin ID")
	setPinCmd.MarkFlagRequired("pinId")
	phaseAdjust = setPinCmd.Flags().Int32P("phaseAdjust", "j", math.MaxInt32, "Phase adjustment in ps")
	setPinCmd.Flags().UintSliceVarP(&parentId, "parentId", "d", []uint{}, "Pin parent ID(s)")
	setPinCmd.Flags().UintSliceVarP(&prio, "prio", "p", []uint{}, "Pin priorit(y/ies)")
	setPinCmd.Flags().UintSliceVarP(&pinState, "pinState", "s", []uint{}, "Pin State(s)")
	setPinCmd.MarkFlagsMutuallyExclusive("prio", "pinState")
	esyncFrequency = setPinCmd.Flags().Uint64P("esyncFrequency", "e", 0, "E-Sync frequency in Hz")
	frequency = setPinCmd.Flags().Uint64P("frequency", "f", 0, "Frequency in Hz")
}
