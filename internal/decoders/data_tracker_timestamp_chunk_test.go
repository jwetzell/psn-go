package decoders

import (
	"fmt"
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

		actual, err := DecodeDataTrackerTimestampChunk(testCase.bytes)

		if err != nil {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Println(err)
		}

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Printf("expected: %v\n", testCase.expected)
			fmt.Printf("actual: %v\n", actual)
		}
	}

}
