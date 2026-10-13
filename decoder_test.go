package psn_test

import (
	"fmt"
	"testing"

	"github.com/jwetzell/psn-go"
	"github.com/stretchr/testify/require"
)

func TestGoodDecoder(t *testing.T) {

	testCases := []struct {
		name               string
		bytes              [][]byte
		expectedSystemName string
		expectedTrackers   map[uint16]*psn.Tracker
	}{
		{
			bytes: [][]byte{
				[]byte{
					85, 103, 124, 128, 0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 1, 1, 0, 104, 128, 1, 0, 100, 128, 0, 0,
					12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 2, 0, 12,
					0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64,
					0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
				},
				[]byte{
					86, 103, 51, 128, 0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 1, 1, 0, 10, 0, 80, 83, 78, 32, 83, 101,
					114, 118, 101, 114, 2, 0, 17, 128, 1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
				},
			},
			expectedSystemName: "PSN Server",
			expectedTrackers: map[uint16]*psn.Tracker{
				1: &psn.Tracker{
					Pos: &psn.XYZData{
						X: float32(1),
						Y: float32(2),
						Z: float32(3),
					},
					Speed: &psn.XYZData{
						X: float32(1),
						Y: float32(2),
						Z: float32(3),
					},
					Ori: &psn.XYZData{
						X: float32(1),
						Y: float32(2),
						Z: float32(3),
					},
					Validity: new(float32(1)),
					Accel: &psn.XYZData{
						X: float32(1),
						Y: float32(2),
						Z: float32(3),
					},
					TrgtPos: &psn.XYZData{
						X: float32(1),
						Y: float32(2),
						Z: float32(3),
					},
					Timestamp: new(uint64(1234567890)),
					Name:      "Tracker 1",
					Id:        1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			decoder := psn.NewDecoder()
			for _, b := range tc.bytes {
				err := decoder.Decode(b)
				require.NoError(t, err)
			}
			fmt.Printf("Decoder: %+v\n", decoder)
			require.Equal(t, tc.expectedSystemName, decoder.SystemName)
			require.Equal(t, len(tc.expectedTrackers), len(decoder.Trackers))

			for id, expectedTracker := range tc.expectedTrackers {
				actualTracker, ok := decoder.Trackers[id]
				require.True(t, ok, "expected tracker with ID %d not found", id)
				require.Equal(t, expectedTracker, actualTracker)
			}
		})
	}
}
