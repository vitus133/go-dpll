//go:build linux

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"

	"github.com/mdlayher/netlink"
	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
)

const (
	IFLA_DPLL_PIN = 65
)

// getNdPinCmd represents the getNdPin command
var getNdPinCmd = &cobra.Command{
	Use:   "getNdPin",
	Short: "Get DPLL pin information for a network device",
	Long: `Retrieve DPLL pin information associated with a given network interface.

Example:
	dpll-cli getNdPin --ifname=eth0
`,
	Run: func(cmd *cobra.Command, args []string) {
		ifname, _ := cmd.Flags().GetString("ifname")
		// Placeholder implementation; wiring to dpll queries will follow.
		fmt.Println("getNdPin called for interface:", ifname)
		dpllPinID, found, err := GetNetdevDpllPin(ifname)
		if err != nil {
			fmt.Printf("Error getting DPLL pin information: %v\n", err)
			return
		}
		if !found {
			fmt.Printf("No DPLL pin found for interface: %s\n", ifname)
			return
		}
		fmt.Printf("DPLL pin ID: %d\n", dpllPinID)
	},
}

func GetNetdevDpllPin(ifname string) (uint32, bool, error) {
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		return 0, false, fmt.Errorf("could not find interface %s: %w", ifname, err)
	}

	conn, err := netlink.Dial(unix.NETLINK_ROUTE, nil)
	if err != nil {
		return 0, false, fmt.Errorf("failed to dial netlink: %w", err)
	}
	defer conn.Close()

	b := make([]byte, unix.SizeofIfInfomsg)
	binary.NativeEndian.PutUint32(b[4:8], uint32(iface.Index))

	msg := netlink.Message{
		Header: netlink.Header{
			Type:  unix.RTM_GETLINK,
			Flags: netlink.Request,
		},
		Data: b,
	}

	replyMsgs, err := conn.Execute(msg)
	if err != nil {
		return 0, false, fmt.Errorf("failed to execute netlink request: %w", err)
	}

	if len(replyMsgs) != 1 {
		return 0, false, fmt.Errorf("expected 1 reply message, got %d", len(replyMsgs))
	}
	reply := replyMsgs[0]

	if reply.Header.Type != unix.RTM_NEWLINK {
		return 0, false, fmt.Errorf("expected RTM_NEWLINK, got %d", reply.Header.Type)
	}
	if len(reply.Data) < unix.SizeofIfInfomsg {
		return 0, false, errors.New("reply message is too short")
	}

	attrData := reply.Data[unix.SizeofIfInfomsg:]

	ad, err := netlink.NewAttributeDecoder(attrData)
	if err != nil {
		return 0, false, fmt.Errorf("failed to create attribute decoder: %w", err)
	}

	var dpllPinID uint32
	var found bool

	for ad.Next() {
		if ad.Type() == IFLA_DPLL_PIN {
			ad.Nested(func(nad *netlink.AttributeDecoder) error {
				for nad.Next() {
					if nad.Type() == dpll.DpllPinID {
						dpllPinID = nad.Uint32()
						found = true
						return nil
					}
				}
				return nad.Err()
			})

			if found {
				break
			}
		}
	}

	if err := ad.Err(); err != nil {
		return 0, false, fmt.Errorf("attribute decoding failed: %w", err)
	}

	return dpllPinID, found, nil
}
func init() {
	rootCmd.AddCommand(getNdPinCmd)

	getNdPinCmd.Flags().String("ifname", "", "Network interface name (e.g., eth0)")
	getNdPinCmd.MarkFlagRequired("ifname")
}
