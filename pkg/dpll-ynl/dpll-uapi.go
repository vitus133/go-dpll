// Definitions from <kernel-root>/include/uapi/linux/dpll.h and
// tools/net/ynl/generated/dpll-user.h

package dpll

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const DPLL_MCGRP_MONITOR = "monitor"
const DPLL_PHASE_OFFSET_DIVIDER = 1000
const DPLL_TEMP_DIVIDER = 1000
const (
	DPLL_A_TYPES = iota
	DPLL_A_ID
	DPLL_A_MODULE_NAME
	DPLL_A_PAD
	DPLL_A_CLOCK_ID
	DPLL_A_MODE
	DPLL_A_MODE_SUPPORTED
	DPLL_A_LOCK_STATUS
	DPLL_A_TEMP
	DPLL_A_TYPE

	__DPLL_A_MAX
	DPLL_A_MAX = __DPLL_A_MAX - 1
)

const (
	DPLL_A_PIN_TYPES = iota

	DPLL_A_PIN_ID
	DPLL_A_PIN_PARENT_ID
	DPLL_A_PIN_MODULE_NAME
	DPLL_A_PIN_PAD
	DPLL_A_PIN_CLOCK_ID
	DPLL_A_PIN_BOARD_LABEL
	DPLL_A_PIN_PANEL_LABEL
	DPLL_A_PIN_PACKAGE_LABEL
	DPLL_A_PIN_TYPE
	DPLL_A_PIN_DIRECTION
	DPLL_A_PIN_FREQUENCY
	DPLL_A_PIN_FREQUENCY_SUPPORTED
	DPLL_A_PIN_FREQUENCY_MIN
	DPLL_A_PIN_FREQUENCY_MAX
	DPLL_A_PIN_PRIO
	DPLL_A_PIN_STATE
	DPLL_A_PIN_CAPABILITIES
	DPLL_A_PIN_PARENT_DEVICE
	DPLL_A_PIN_PARENT_PIN
	DPLL_A_PIN_PHASE_ADJUST_MIN
	DPLL_A_PIN_PHASE_ADJUST_MAX
	DPLL_A_PIN_PHASE_ADJUST
	DPLL_A_PIN_PHASE_OFFSET
	DPLL_A_PIN_FRACTIONAL_FREQUENCY_OFFSET

	__DPLL_A_PIN_MAX
	DPLL_A_PIN_MAX = __DPLL_A_PIN_MAX - 1
)
const (
	DPLL_CMDS = iota
	DPLL_CMD_DEVICE_ID_GET
	DPLL_CMD_DEVICE_GET
	DPLL_CMD_DEVICE_SET
	DPLL_CMD_DEVICE_CREATE_NTF
	DPLL_CMD_DEVICE_DELETE_NTF
	DPLL_CMD_DEVICE_CHANGE_NTF
	DPLL_CMD_PIN_ID_GET
	DPLL_CMD_PIN_GET
	DPLL_CMD_PIN_SET
	DPLL_CMD_PIN_CREATE_NTF
	DPLL_CMD_PIN_DELETE_NTF
	DPLL_CMD_PIN_CHANGE_NTF

	__DPLL_CMD_MAX
	DPLL_CMD_MAX = (__DPLL_CMD_MAX - 1)
)

// GetLockStatus returns DPLL lock status as a string
func GetLockStatus(ls uint32) string {
	lockStatusMap := map[uint32]string{
		1: "unlocked",
		2: "locked",
		3: "locked-ho-acquired",
		4: "holdover",
	}
	status, found := lockStatusMap[ls]
	if found {
		return status
	}
	return ""
}

// GetDpllType returns DPLL type as a string
func GetDpllType(tp uint32) string {
	typeMap := map[int]string{
		1: "pps",
		2: "eec",
	}
	typ, found := typeMap[int(tp)]
	if found {
		return typ
	}
	return ""
}

// GetMode returns DPLL mode as a string
func GetMode(md uint32) string {
	modeMap := map[int]string{
		1: "manual",
		2: "automatic",
	}
	mode, found := modeMap[int(md)]
	if found {
		return mode
	}
	return ""
}

// DpllStatusHR represents human-readable DPLL status
type DpllStatusHR struct {
	Timestamp     time.Time `json:"timestamp"`
	Id            uint32    `json:"id"`
	ModuleName    string    `json:"moduleName"`
	Mode          string    `json:"mode"`
	ModeSupported string    `json:"modeSupported"`
	LockStatus    string    `json:"lockStatus"`
	ClockId       string    `json:"clockId"`
	Type          string    `json:"type"`
	Temp          float64   `json:"temp"`
}

// GetDpllStatusHR returns human-readable DPLL status
func GetDpllStatusHR(reply *DoDeviceGetReply, timestamp time.Time) ([]byte, error) {
	var modes []string
	for _, md := range reply.ModeSupported {
		modes = append(modes, GetMode(md))
	}
	hr := DpllStatusHR{
		Timestamp:     timestamp,
		Id:            reply.Id,
		ModuleName:    reply.ModuleName,
		Mode:          GetMode(reply.Mode),
		ModeSupported: fmt.Sprint(strings.Join(modes[:], ",")),
		LockStatus:    GetLockStatus(reply.LockStatus),
		ClockId:       fmt.Sprintf("0x%x", reply.ClockId),
		Type:          GetDpllType(reply.Type),
		Temp:          float64(reply.Temp) / DPLL_TEMP_DIVIDER,
	}
	return json.Marshal(hr)
}

// PinInfo is used with the DoPinGet method.
type PinInfoHR struct {
	Timestamp                 time.Time           `json:"timestamp"`
	Id                        uint32              `json:"id"`
	ClockId                   string              `json:"clockId"`
	BoardLabel                string              `json:"boardLabel"`
	PanelLabel                string              `json:"panelLabel"`
	PackageLabel              string              `json:"packageLabel"`
	Type                      string              `json:"type"`
	Frequency                 uint64              `json:"frequency"`
	FrequencySupported        []FrequencyRange    `json:"frequencySupported"`
	Capabilities              string              `json:"capabilities"`
	ParentDevice              []PinParentDeviceHR `json:"pinParentDevice"`
	ParentPin                 []PinParentPinHR    `json:"pinParentPin"`
	PhaseAdjustMin            int32               `json:"phaseAdjustMin"`
	PhaseAdjustMax            int32               `json:"phaseAdjustMax"`
	PhaseAdjust               int32               `json:"phaseAdjust"`
	FractionalFrequencyOffset int                 `json:"fractionalFrequencyOffset"`
	ModuleName                string              `json:"moduleName"`
}

// PinParentDeviceHR contains nested netlink attributes.
type PinParentDeviceHR struct {
	ParentId      uint32  `json:"parentId"`
	Direction     string  `json:"direction"`
	Prio          uint32  `json:"prio"`
	State         string  `json:"state"`
	PhaseOffsetPs float64 `json:"phaseOffsetPs"`
}

// PinParentPin contains nested netlink attributes.
type PinParentPinHR struct {
	ParentId uint32 `json:"parentId"`
	State    string `json:"parentState"`
}

const (
	DPLL_A_PIN_STATE_CONNECTED    = 1
	DPLL_A_PIN_STATE_DISCONNECTED = 2
	DPLL_A_PIN_STATE_SELECTABLE   = 3
)

// GetPinState returns DPLL pin state as a string
func GetPinState(s uint32) string {
	stateMap := map[int]string{
		DPLL_A_PIN_STATE_CONNECTED:    "connected",
		DPLL_A_PIN_STATE_DISCONNECTED: "disconnected",
		DPLL_A_PIN_STATE_SELECTABLE:   "selectable",
	}
	r, found := stateMap[int(s)]
	if found {
		return r
	}
	return ""
}

// GetPinType returns DPLL pin type as a string
func GetPinType(tp uint32) string {
	typeMap := map[int]string{
		1: "mux",
		2: "ext",
		3: "synce-eth-port",
		4: "int-oscillator",
		5: "gnss",
	}
	typ, found := typeMap[int(tp)]
	if found {
		return typ
	}
	return ""
}

// GetPinDirection returns DPLL pin direction as a string
func GetPinDirection(d uint32) string {
	directionMap := map[int]string{
		1: "input",
		2: "output",
	}
	dir, found := directionMap[int(d)]
	if found {
		return dir
	}
	return ""
}

// GetPinCapabilities returns DPLL pin capabilities as a csv
func GetPinCapabilities(c uint32) string {
	cMap := map[int]string{
		0: "",
		1: "direction-can-change",
		2: "priority-can-change",
		3: "direction-can-change,priority-can-change",
		4: "state-can-change",
		5: "state-can-change,direction-can-change",
		6: "state-can-change,priority-can-change",
		7: "state-can-change,direction-can-change,priority-can-change",
	}
	cap, found := cMap[int(c)]
	if found {
		return cap
	}
	return ""
}

// GetPinInfoHR returns human-readable pin status
func GetPinInfoHR(reply *PinInfo, timestamp time.Time) ([]byte, error) {
	hr := PinInfoHR{
		Timestamp:                 timestamp,
		Id:                        reply.Id,
		ClockId:                   fmt.Sprintf("0x%x", reply.ClockId),
		BoardLabel:                reply.BoardLabel,
		PanelLabel:                reply.PanelLabel,
		PackageLabel:              reply.PackageLabel,
		Type:                      GetPinType(reply.Type),
		Frequency:                 reply.Frequency,
		FrequencySupported:        make([]FrequencyRange, 0),
		PhaseAdjustMin:            reply.PhaseAdjustMin,
		PhaseAdjustMax:            reply.PhaseAdjustMax,
		PhaseAdjust:               reply.PhaseAdjust,
		FractionalFrequencyOffset: reply.FractionalFrequencyOffset,
		ModuleName:                reply.ModuleName,
		ParentDevice:              make([]PinParentDeviceHR, 0),
		ParentPin:                 make([]PinParentPinHR, 0),
		Capabilities:              GetPinCapabilities(reply.Capabilities),
	}
	for i := 0; i < len(reply.ParentDevice); i++ {
		hr.ParentDevice = append(hr.ParentDevice, PinParentDeviceHR{
			ParentId:      reply.ParentDevice[i].ParentId,
			Direction:     GetPinDirection(reply.ParentDevice[i].Direction),
			Prio:          reply.ParentDevice[i].Prio,
			State:         GetPinState(reply.ParentDevice[i].State),
			PhaseOffsetPs: float64(reply.ParentDevice[i].PhaseOffset) / DPLL_PHASE_OFFSET_DIVIDER,
		})

	}
	for i := 0; i < len(reply.ParentPin); i++ {
		hr.ParentPin = append(hr.ParentPin, PinParentPinHR{
			ParentId: reply.ParentPin[i].ParentId,
			State:    GetPinState(reply.ParentPin[i].State),
		})
	}
	for i := 0; i < len(reply.FrequencySupported); i++ {
		hr.FrequencySupported = append(hr.FrequencySupported, FrequencyRange{
			FrequencyMin: reply.FrequencySupported[i].FrequencyMin,
			FrequencyMax: reply.FrequencySupported[i].FrequencyMax,
		})
	}
	return json.Marshal(hr)
}
