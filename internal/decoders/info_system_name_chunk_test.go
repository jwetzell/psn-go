package decoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoodInfoSystemNameChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    InfoSystemNameChunk
	}{
		{
			description: "InfoSystemNameChunk",
			bytes: []byte{
				1, 0, 10, 0, 80, 83, 78, 32, 83, 101, 114, 118, 101, 114,
			},
			expected: InfoSystemNameChunk{
				Chunk: Chunk{
					ChunkData: []byte{80, 83, 78, 32, 83, 101, 114, 118, 101, 114},
					Header:    ChunkHeader{DataLen: 10, Id: 1, HasSubchunks: false},
				},
				Data: InfoSystemNameChunkData{
					SystemName: "PSN Server",
				},
			},
		},
	}

	for _, testCase := range testCases {

		actual, err := DecodeInfoSystemNameChunk(testCase.bytes)

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
