package clock

import (
	"fmt"
)

func newTrack(name string, max int) ballTrack {
	balls := make([]uint8, 0, max)
	return ballTrack{
		name:  name,
		balls: balls,
		max:   max,
	}
}

func (track *ballTrack) equals(otherTrack *ballTrack) bool {
	if len(track.balls) != len(otherTrack.balls) {
		return false
	}

	for i := 0; i < len(track.balls); i++ {
		if (track.balls)[i] != (otherTrack.balls)[i] {
			return false
		}
	}

	return true
}

// print is a helper function for visualizing the track (debugging)
func (track *ballTrack) print(numToAdd int) {
	fmt.Printf("%s:\n", track.name)
	for i := 0; i < len(track.balls)+numToAdd; i++ {
		fmt.Printf("* ")
	}
	fmt.Println()
	for i := 1; i <= track.max+numToAdd; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for _, ball := range track.balls {
		fmt.Printf("[%d] ", ball)
	}
	fmt.Println()
}
