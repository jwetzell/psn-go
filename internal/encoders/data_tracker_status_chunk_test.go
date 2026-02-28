package encoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestDataTrackerStatusChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.DataTrackerStatusChunkData
	}{
		{
			description: "basic status",
			expected: []byte{
				3, 0, 4, 0, 0, 0, 128, 63,
			},
			data: chunks.DataTrackerStatusChunkData{
				Validity: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeDataTrackerStatusChunk(testCase.data.Validity)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to encode chunk properly, expected: %v, actual: %v\n", testCase.expected, actual)
			}
		})
	}
}
