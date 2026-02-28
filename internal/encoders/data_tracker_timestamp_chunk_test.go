package encoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestDataTrackerTimestampChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.DataTrackerTimestampChunkData
	}{
		{
			description: "basic timestamp",
			expected: []byte{
				6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
			},
			data: chunks.DataTrackerTimestampChunkData{
				Timestamp: 1234567890,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeDataTrackerTimestampChunk(testCase.data.Timestamp)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to encode chunk properly, expected: %v, actual: %v\n", testCase.expected, actual)
			}
		})
	}
}
