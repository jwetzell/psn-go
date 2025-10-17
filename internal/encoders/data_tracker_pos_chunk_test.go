package encoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/decoders"
)

func TestDataTrackerPosChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        decoders.DataTrackerXYZChunkData
	}{
		{
			description: "basic position",
			expected: []byte{
				0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64,
			},
			data: decoders.DataTrackerXYZChunkData{
				X: 1,
				Y: 2,
				Z: 3,
			},
		},
	}

	for _, testCase := range testCases {

		actual := EncodeDataTrackerPosChunk(testCase.data.X, testCase.data.Y, testCase.data.Z)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to encode chunk properly", testCase.description)
			fmt.Printf("expected: %v\n", testCase.expected)
			fmt.Printf("actual: %v\n", actual)
		}
	}
}
