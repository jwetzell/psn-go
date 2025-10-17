package encoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/decoders"
)

func TestInfoSystemNameChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        decoders.InfoSystemNameChunkData
	}{
		{
			description: "InfoSystemNameChunk",
			expected: []byte{
				1, 0, 10, 0, 80, 83, 78, 32, 83, 101, 114, 118, 101, 114,
			},
			data: decoders.InfoSystemNameChunkData{
				SystemName: "PSN Server",
			},
		},
	}

	for _, testCase := range testCases {

		actual := EncodeInfoSystemNameChunk(testCase.data.SystemName)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to encode chunk properly", testCase.description)
			fmt.Printf("expected: %v\n", testCase.expected)
			fmt.Printf("actual: %v\n", actual)
		}
	}
}
