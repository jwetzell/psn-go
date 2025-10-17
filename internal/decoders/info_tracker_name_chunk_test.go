package decoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestGoodInfoTrackerNameChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.InfoTrackerNameChunk
	}{
		{
			description: "InfoTrackerNameChunk",
			bytes: []byte{
				0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			expected: chunks.InfoTrackerNameChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{84, 114, 97, 99, 107, 101, 114, 32, 49},
					Header:    chunks.ChunkHeader{DataLen: 9, Id: 0, HasSubchunks: false},
				},
				Data: chunks.InfoTrackerNameChunkData{
					TrackerName: "Tracker 1",
				},
			},
		},
	}

	for _, testCase := range testCases {

		actual, err := DecodeInfoTrackerNameChunk(testCase.bytes)

		if err != nil {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Println(err)
		}

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Printf("expected: %+v\n", testCase.expected)
			fmt.Printf("actual: %+v\n", actual)
		}
	}
}
