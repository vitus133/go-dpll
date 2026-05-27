package dpll

import (
	"encoding/json"
	"math"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// --- US1: device frequency-monitor ---

func TestGetFrequencyMonitor(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input    uint32
		expected string
	}{
		{1, "enabled"},
		{2, "disabled"},
		{0, ""},
		{99, ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.expected, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, GetFrequencyMonitor(tc.input))
		})
	}
}

func TestGetDpllStatusHR_FrequencyMonitor(t *testing.T) {
	t.Parallel()
	reply := &DoDeviceGetReply{
		FrequencyMonitor: 1,
	}
	data, err := GetDpllStatusHR(reply, time.Time{})
	assert.NoError(t, err)
	var out map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &out))
	assert.Equal(t, "enabled", out["frequencyMonitor"])
}

func TestGetDpllStatusHR_FrequencyMonitor_Omitted(t *testing.T) {
	t.Parallel()
	reply := &DoDeviceGetReply{FrequencyMonitor: 0}
	data, err := GetDpllStatusHR(reply, time.Time{})
	assert.NoError(t, err)
	var out map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &out))
	_, present := out["frequencyMonitor"]
	assert.False(t, present, "frequencyMonitor should be omitted when zero")
}

// --- US2: pin measured-frequency + operstate ---

func TestGetPinOperstate(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input    uint32
		expected string
	}{
		{PinOperstateActive, "active"},
		{PinOperstateStandby, "standby"},
		{PinOperstateNoSignal, "no-signal"},
		{PinOperstateQualFailed, "qual-failed"},
		{0, ""},
		{99, ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.expected+"_"+string(rune('0'+tc.input)), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, GetPinOperstate(tc.input))
		})
	}
}

func TestGetPinInfoHR_NewFields(t *testing.T) {
	t.Parallel()
	pin := &PinInfo{
		MeasuredFrequency: 1000001,
		Operstate:         PinOperstateActive,
	}
	data, err := GetPinInfoHR(pin, time.Time{})
	assert.NoError(t, err)
	var out map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &out))
	assert.InDelta(t, 1000.001, out["measuredFrequencyHz"], 1e-9)
	assert.Equal(t, "active", out["operstate"])
}

func TestGetPinInfoHR_NewFields_Omitted(t *testing.T) {
	t.Parallel()
	pin := &PinInfo{MeasuredFrequency: 0, Operstate: 0}
	data, err := GetPinInfoHR(pin, time.Time{})
	assert.NoError(t, err)
	var out map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &out))
	_, hasMF := out["measuredFrequencyHz"]
	_, hasOS := out["operstate"]
	assert.False(t, hasMF, "measuredFrequencyHz should be omitted when zero")
	assert.False(t, hasOS, "operstate should be omitted when zero")
}

// --- US3: pin-parent-device nested attrs ---

func TestGetPinInfoHR_ParentDeviceNewFields(t *testing.T) {
	t.Parallel()
	pin := &PinInfo{
		ParentDevice: []PinParentDevice{
			{
				Operstate:                    PinOperstateStandby,
				FractionalFrequencyOffset:    -3,
				FractionalFrequencyOffsetPPT: -3000000,
			},
		},
	}
	data, err := GetPinInfoHR(pin, time.Time{})
	assert.NoError(t, err)
	var out struct {
		ParentDevice []struct {
			Operstate                    string  `json:"operstate"`
			FractionalFrequencyOffset    float64 `json:"fractionalFrequencyOffset"`
			FractionalFrequencyOffsetPPT float64 `json:"fractionalFrequencyOffsetPPT"`
		} `json:"pinParentDevice"`
	}
	assert.NoError(t, json.Unmarshal(data, &out))
	assert.Len(t, out.ParentDevice, 1)
	pd := out.ParentDevice[0]
	assert.Equal(t, "standby", pd.Operstate)
	assert.Equal(t, float64(-3), pd.FractionalFrequencyOffset)
	assert.Equal(t, float64(-3000000), pd.FractionalFrequencyOffsetPPT)
}

func TestGetPinInfoHR_ParentDeviceNewFields_Omitted(t *testing.T) {
	t.Parallel()
	pin := &PinInfo{
		ParentDevice: []PinParentDevice{
			{Operstate: 0, FractionalFrequencyOffset: 0, FractionalFrequencyOffsetPPT: 0},
		},
	}
	data, err := GetPinInfoHR(pin, time.Time{})
	assert.NoError(t, err)
	var raw map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &raw))
	devs := raw["pinParentDevice"].([]interface{})
	pd := devs[0].(map[string]interface{})
	_, hasOS := pd["operstate"]
	_, hasFFO := pd["fractionalFrequencyOffset"]
	_, hasFFOPPT := pd["fractionalFrequencyOffsetPPT"]
	assert.False(t, hasOS, "operstate should be omitted when zero")
	assert.False(t, hasFFO, "fractionalFrequencyOffset should be omitted when zero")
	assert.False(t, hasFFOPPT, "fractionalFrequencyOffsetPPT should be omitted when zero")
}

// --- existing test ---

func Test_EncodePinControl(t *testing.T) {
	assert.New(t)
	// test phase adjustment
	pc := PinParentDeviceCtl{
		ID: 88,
		PhaseAdjust: func() *int32 {
			t := int32(math.MinInt32)
			return &t
		}(),
	}
	b, err := EncodePinControl(pc)
	assert.NoError(t, err, "failed to encode phase adjustment")
	expected, err := os.ReadFile("testdata/encoded-phase-adjustment")
	assert.NoError(t, err, "failed to read testdata for phase adjustment")
	assert.Equal(t, expected, b, "encoded data is different from the desired phase adjustment data")

	// Test priority settings
	pc.PhaseAdjust = nil
	pc.PinParentCtl = []PinControl{
		{
			PinParentID: 8,
			Prio: func() *uint32 {
				t := uint32(math.MaxUint32)
				return &t
			}(),
		},
		{
			PinParentID: 8,
			Prio: func() *uint32 {
				t := uint32(math.MaxUint8)
				return &t
			}(),
		},
	}
	b, err = EncodePinControl(pc)
	assert.NoError(t, err, "failed to encode pin priority setting")
	expected, err = os.ReadFile("testdata/encoded-prio")
	assert.NoError(t, err, "failed to read testdata for priority setting")
	assert.Equal(t, expected, b, "encoded data is different from the desired pin priority data")

	// Test setting the connection state
	pc.PinParentCtl = []PinControl{
		{
			PinParentID: 0,
			Prio:        nil,
			Direction: func() *uint32 {
				t := uint32(2)
				return &t
			}(),
		},
		{
			PinParentID: 1,
			Prio:        nil,
			Direction: func() *uint32 {
				t := uint32(1)
				return &t
			}(),
		},
	}
	b, err = EncodePinControl(pc)
	assert.NoError(t, err, "failed to encode pin connection setting")
	expected, err = os.ReadFile("testdata/encoded-connection")
	assert.NoError(t, err, "failed to read testdata for connection setting")
	assert.Equal(t, expected, b, "encoded data is different from the desired pin connection setting data")

}
