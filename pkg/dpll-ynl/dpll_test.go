package dpll

import (
	"encoding/json"
	"math"
	"os"
	"testing"
	"time"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
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

// TestParsePinReplies_DpllPinFractionalFrequencyOffsetPPT verifies that
// ParsePinReplies correctly decodes the top-level DpllPinFractionalFrequencyOffsetPPT
// attribute (FFO in parts per trillion) into PinInfo.FractionalFrequencyOffsetPPT.
// Kernel nla_put_sint may encode as 4 or 8 bytes; both must decode.
func TestParsePinReplies_DpllPinFractionalFrequencyOffsetPPT(t *testing.T) {
	tests := []struct {
		name    string
		encode  func(*netlink.AttributeEncoder)
		wantPPT int64
	}{
		{
			name: "int32-width positive",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int32(DpllPinFractionalFrequencyOffsetPPT, 12345)
			},
			wantPPT: 12345,
		},
		{
			name: "int32-width zero",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int32(DpllPinFractionalFrequencyOffsetPPT, 0)
			},
			wantPPT: 0,
		},
		{
			name: "int32-width negative",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int32(DpllPinFractionalFrequencyOffsetPPT, -999)
			},
			wantPPT: -999,
		},
		{
			name: "int64-width value beyond int32",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int64(DpllPinFractionalFrequencyOffsetPPT, 3_000_000_000)
			},
			wantPPT: 3_000_000_000,
		},
		{
			name: "int64-width negative",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int64(DpllPinFractionalFrequencyOffsetPPT, -3_000_000_000)
			},
			wantPPT: -3_000_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ae := netlink.NewAttributeEncoder()
			ae.Uint32(DpllPinID, 1)
			tt.encode(ae)
			payload, err := ae.Encode()
			assert.NoError(t, err, "encode attributes")

			msgs := []genetlink.Message{{Data: payload}}
			replies, err := ParsePinReplies(msgs)
			assert.NoError(t, err)
			assert.Len(t, replies, 1)
			assert.Equal(t, tt.wantPPT, replies[0].FractionalFrequencyOffsetPPT,
				"FractionalFrequencyOffsetPPT should match encoded value")
		})
	}
}

// TestParsePinReplies_DpllPinFractionalFrequencyOffset verifies that the
// top-level DpllPinFractionalFrequencyOffset attribute decodes via decodeSint,
// handling both 4- and 8-byte SINT payloads.
func TestParsePinReplies_DpllPinFractionalFrequencyOffset(t *testing.T) {
	tests := []struct {
		name   string
		encode func(*netlink.AttributeEncoder)
		want   int
	}{
		{
			name: "int32-width positive",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int32(DpllPinFractionalFrequencyOffset, 5)
			},
			want: 5,
		},
		{
			name: "int32-width negative",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int32(DpllPinFractionalFrequencyOffset, -42)
			},
			want: -42,
		},
		{
			name: "int64-width value beyond int32",
			encode: func(ae *netlink.AttributeEncoder) {
				ae.Int64(DpllPinFractionalFrequencyOffset, 7_000_000_000)
			},
			want: 7_000_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ae := netlink.NewAttributeEncoder()
			ae.Uint32(DpllPinID, 1)
			tt.encode(ae)
			payload, err := ae.Encode()
			assert.NoError(t, err, "encode attributes")

			msgs := []genetlink.Message{{Data: payload}}
			replies, err := ParsePinReplies(msgs)
			assert.NoError(t, err)
			assert.Len(t, replies, 1)
			assert.Equal(t, tt.want, replies[0].FractionalFrequencyOffset,
				"FractionalFrequencyOffset should match encoded value")
		})
	}
}

// TestParsePinReplies_ParentDeviceNewFields verifies that ParsePinReplies
// decodes operstate, FFO, and FFO-PPT from the pin-parent-device nest via
// decodeSint (FFO-PPT encoded as an 8-byte SINT beyond int32 range).
func TestParsePinReplies_ParentDeviceNewFields(t *testing.T) {
	const parentID = uint32(7)
	const operstate = uint32(PinOperstateActive)
	const ffo = int32(-42)
	const ffoPPT = int64(3_000_000_000)

	ae := netlink.NewAttributeEncoder()
	ae.Uint32(DpllPinID, 1)
	ae.Nested(DpllPinParentDevice, func(nae *netlink.AttributeEncoder) error {
		nae.Uint32(DpllPinParentID, parentID)
		nae.Uint32(DpllPinOperstate, operstate)
		nae.Int32(DpllPinFractionalFrequencyOffset, ffo)
		nae.Int64(DpllPinFractionalFrequencyOffsetPPT, ffoPPT)
		return nil
	})
	payload, err := ae.Encode()
	assert.NoError(t, err)

	msgs := []genetlink.Message{{Data: payload}}
	replies, err := ParsePinReplies(msgs)
	assert.NoError(t, err)
	assert.Len(t, replies, 1)
	assert.Len(t, replies[0].ParentDevice, 1)
	pd := replies[0].ParentDevice[0]
	assert.Equal(t, parentID, pd.ParentID)
	assert.Equal(t, operstate, pd.Operstate)
	assert.Equal(t, int(ffo), pd.FractionalFrequencyOffset)
	assert.Equal(t, ffoPPT, pd.FractionalFrequencyOffsetPPT)
}

// TestDecodeSint_InvalidLength verifies that decodeSint rejects SINT payloads
// that are neither 4 nor 8 bytes wide.
func TestDecodeSint_InvalidLength(t *testing.T) {
	for _, size := range []int{1, 2, 3, 5, 6, 7} {
		ae := netlink.NewAttributeEncoder()
		ae.Bytes(DpllPinFractionalFrequencyOffsetPPT, make([]byte, size))
		payload, err := ae.Encode()
		assert.NoError(t, err)

		msgs := []genetlink.Message{{Data: payload}}
		replies, err := ParsePinReplies(msgs)
		assert.Error(t, err, "size %d: expected decodeSint to reject non-sint payload", size)
		assert.Nil(t, replies)
	}
}
