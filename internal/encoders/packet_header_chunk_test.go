package encoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestPacketHeaderChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.PacketHeaderChunkData
	}{
		{
			description: "PacketHeaderChunk",
			expected: []byte{
				0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123,
			},
			data: chunks.PacketHeaderChunkData{
				PacketTimestamp:  1234567890,
				VersionHigh:      2,
				VersionLow:       3,
				FrameId:          1,
				FramePacketCount: 123,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodePacketHeaderChunk(testCase.data.PacketTimestamp, testCase.data.VersionHigh, testCase.data.VersionLow, testCase.data.FrameId, testCase.data.FramePacketCount)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to encode chunk properly, expected: %v, actual: %v\n", testCase.expected, actual)
			}
		})
	}
}
