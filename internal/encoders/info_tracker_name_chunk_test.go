package encoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestInfoTrackerNameChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.InfoTrackerNameChunkData
	}{
		{
			description: "InfoTrackerNameChunk",
			expected: []byte{
				0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			data: chunks.InfoTrackerNameChunkData{
				TrackerName: "Tracker 1",
			},
		},
	}

	for _, testCase := range testCases {

		actual := EncodeInfoTrackerNameChunk(testCase.data.TrackerName)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to encode chunk properly", testCase.description)
			fmt.Printf("expected: %v\n", testCase.expected)
			fmt.Printf("actual: %v\n", actual)
		}
	}
}
