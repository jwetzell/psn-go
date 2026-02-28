package decoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestGoodDataTrackerTimestampChunk(t *testing.T) {

	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.DataTrackerTimestampChunk
	}{
		{
			description: "DataTrackerTimestampChunk",
			bytes:       []byte{6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0},
			expected: chunks.DataTrackerTimestampChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{210, 2, 150, 73, 0, 0, 0, 0},
					Header:    chunks.ChunkHeader{DataLen: 8, Id: 6, HasSubchunks: false},
				},
				Data: chunks.DataTrackerTimestampChunkData{
					Timestamp: 1234567890,
				},
			},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeDataTrackerTimestampChunk(testCase.bytes)

			if err != nil {
				t.Errorf("failed to decode chunk properly, error: %v", err)
			}

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to decode chunk properly, expected: %+v, actual: %+v", testCase.expected, actual)
			}
		})
	}

}
