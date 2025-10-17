package decoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoodInfoTrackerChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    InfoTrackerChunk
	}{
		{
			description: "InfoTrackerChunk",
			bytes: []byte{
				1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			expected: InfoTrackerChunk{
				Chunk: Chunk{
					ChunkData: []byte{0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49},
					Header:    ChunkHeader{DataLen: 13, Id: 1, HasSubchunks: true},
				},
				Data: InfoTrackerChunkData{
					TrackerName: &InfoTrackerNameChunk{
						Chunk: Chunk{
							ChunkData: []byte{84, 114, 97, 99, 107, 101, 114, 32, 49},
							Header:    ChunkHeader{DataLen: 9, Id: 0, HasSubchunks: false},
						},
						Data: InfoTrackerNameChunkData{
							TrackerName: "Tracker 1",
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {

		actual, err := DecodeInfoTrackerChunk(testCase.bytes)

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
