// Definitions from <kernel-root>/include/uapi/linux/dpll.h and
// tools/net/ynl/generated/dpll-user.h

package dpll

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const DpllMCGRPMonitor = "monitor"
const DpllPhaseOffsetDivider = 1000
const DpllTemperatureDivider = 1000

const (
	// attribute-set dpll_a
	DpllAttributes = iota
	DpllID
	DpllModuleName
	DpllAttPadding
	DpllClockID
	DpllMode
	DpllModeSupported
	DpllLockStatus
	DpllTemp
	DpllType
	DpllLockStatusError
	DpllClockQualityLevel
)

const (
	// attribute-set dpll_a_pin
	DpllPinTypeS = iota

	DpllPinId
	DpllPinParentId
	DpllPinModuleName
	DpllPinPadding
	DpllPinClockId
	DpllPinBoardLabel
	DpllPinPanelLabel
	DpllPinPackageLabel
	DpllPinType
	DpllPinDirection
	DpllPinFrequency
	DpllPinFrequencySupported
	DpllPinFrequencyMin
	DpllPinFrequencyMax
	DpllPinPrio
	DPLL_A_PIN_STATE
	DpllPinCapabilities
	DpllPinParentDevice
	DpllPinParentPin
	DpllPinPhaseAdjustMin
	DpllPinPhaseAdjustMax
	DpllPinPhaseAdjust
	DpllPinPhaseOffset
	DpllPinFractionalFrequencyOffset
	DpllPinEsyncFrequency
	DpllPinEsyncFrequencySupported
	DpllPinEsyncPulse
)
const (
	DpllCmds = iota
	DpllCmdDeviceIdGet
	DpllCmdDeviceGet
	DpllCmdDeviceSet
	DpllCmdDeviceCreateNtf
	DpllCmdDeviceDeleteNtf
	DpllCmdDeviceChangeNtf
	DpllCmdPinIDGet
	DpllCmdPinGet
	DpllCmdPinSet
	DpllCmdPinCreateNtf
	DpllCmdPinDeleteNtf
	DpllCmdPinChangeNtf
)

const (
	DpllLockStatusAttribute = iota
	DpllLockStatusUnlocked
	DpllLockStatusLocked
	DpllLockStatusLockedHoldoverAcquired
	DpllLockStatusHoldover
)

const (
	LockStatusErrorTypes = iota
	LockStatusErrorNone
	LockStatusErrorUndefined
	// dpll device lock status was changed because of associated
	// media got down.
	// This may happen for example if dpll device was previously
	// locked on an input pin of type PIN_TYPE_SYNCE_ETH_PORT.
	LockStatusErrorMediaDown
	// the FFO (Fractional Frequency Offset) between the RX and TX
	// symbol rate on the media got too high.
	// This may happen for example if dpll device was previously
	// locked on an input pin of type PIN_TYPE_SYNCE_ETH_PORT.
	LockStatusReeoeFFOTooHigh
)

const (
	ClockQualityLevel = iota
	ClockQualityLevelITUOpt1PRC
	ClockQualityLevelITUOpt1SSU_A
	ClockQualityLevelITUOpt1SSU_B
	ClockQualityLevelITUOpt1EEC1
	ClockQualityLevelITUOpt1PRTC
	ClockQualityLevelITUOpt1EPRTC
	ClockQualityLevelITUOpt1EEEC
	ClockQualityLevelItuOpt1EPRC
)

const (
	DpllTypeAttribute = iota
	// dpll produces Pulse-Per-Second signal
	DpllTypePPS
	// dpll drives the Ethernet Equipment Clock
	DpllTypeEEC
)

// GetLockStatus returns DPLL lock status as a string
func GetLockStatus(ls uint32) string {
	lockStatusMap := map[uint32]string{
		DpllLockStatusUnlocked:               "unlocked",
		DpllLockStatusLocked:                 "locked",
		DpllLockStatusLockedHoldoverAcquired: "locked-ho-acquired",
		DpllLockStatusHoldover:               "holdover",
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
		DpllTypePPS: "pps",
		DpllTypeEEC: "eec",
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
		Temp:          float64(reply.Temp) / DpllTemperatureDivider,
	}
	return json.Marshal(hr)
}

// PinInfoHR is used with the DoPinGet method.
type PinInfoHR struct {
	Timestamp                 time.Time           `json:"timestamp"`
	Id                        uint32              `json:"id"`
	ModuleName                string              `json:"moduleName"`
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
	EsyncFrequency            int64               `json:"esyncFrequency"`
	EsyncFrequencySupported   []FrequencyRange    `json:"esyncFrequencySupported"`
	EsyncPulse                int64               `json:"esyncPulse"`
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
	PinStateConnected    = 1
	PinStateDisconnected = 2
	PinStateSelectable   = 3
)

// GetPinState returns DPLL pin state as a string
func GetPinState(s uint32) string {
	stateMap := map[int]string{
		PinStateConnected:    "connected",
		PinStateDisconnected: "disconnected",
		PinStateSelectable:   "selectable",
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
		EsyncFrequency:            reply.EsyncFrequency,
		EsyncFrequencySupported:   make([]FrequencyRange, 0),
		EsyncPulse:                int64(reply.EsyncPulse),
	}
	for i := 0; i < len(reply.ParentDevice); i++ {
		hr.ParentDevice = append(hr.ParentDevice, PinParentDeviceHR{
			ParentId:      reply.ParentDevice[i].ParentId,
			Direction:     GetPinDirection(reply.ParentDevice[i].Direction),
			Prio:          reply.ParentDevice[i].Prio,
			State:         GetPinState(reply.ParentDevice[i].State),
			PhaseOffsetPs: float64(reply.ParentDevice[i].PhaseOffset) / DpllPhaseOffsetDivider,
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
	for i := 0; i < len(reply.EsyncFrequencySupported); i++ {
		hr.EsyncFrequencySupported = append(hr.EsyncFrequencySupported, FrequencyRange{
			FrequencyMin: reply.EsyncFrequencySupported[i].FrequencyMin,
			FrequencyMax: reply.EsyncFrequencySupported[i].FrequencyMax,
		})
	}
	return json.Marshal(hr)
}
