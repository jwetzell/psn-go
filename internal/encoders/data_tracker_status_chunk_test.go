package encoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/decoders"
)

func TestDataTrackerStatusChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        decoders.DataTrackerStatusChunkData
	}{
		{
			description: "basic status",
			expected: []byte{
				3, 0, 4, 0, 0, 0, 128, 63,
			},
			data: decoders.DataTrackerStatusChunkData{
				Validity: 1,
			},
		},
	}

	for _, testCase := range testCases {

		actual := EncodeDataTrackerStatusChunk(testCase.data.Validity)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to encode chunk properly", testCase.description)
			fmt.Printf("expected: %v\n", testCase.expected)
			fmt.Printf("actual: %v\n", actual)
		}
	}
}
