package encoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestDataTrackerPosChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.DataTrackerXYZChunkData
	}{
		{
			description: "basic position",
			expected: []byte{
				0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64,
			},
			data: chunks.DataTrackerXYZChunkData{
				X: 1,
				Y: 2,
				Z: 3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeDataTrackerPosChunk(testCase.data.X, testCase.data.Y, testCase.data.Z)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to encode chunk properly, expected: %v, actual: %v\n", testCase.expected, actual)
			}
		})
	}
}
