package decoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoodDataTrackerStatusChunk(t *testing.T) {

	testCases := []struct {
		description string
		bytes       []byte
		expected    DataTrackerStatusChunk
	}{
		{
			description: "DataTrackerStatusChunk",
			bytes:       []byte{3, 0, 4, 0, 0, 0, 128, 63},
			expected: DataTrackerStatusChunk{
				Chunk: Chunk{
					ChunkData: []byte{0, 0, 128, 63},
					Header:    ChunkHeader{DataLen: 4, Id: 3, HasSubchunks: false},
				},
				Data: DataTrackerStatusChunkData{
					Validity: 1.0,
				},
			},
		},
	}

	for _, testCase := range testCases {

		actual, err := DecodeDataTrackerStatusChunk(testCase.bytes)

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
