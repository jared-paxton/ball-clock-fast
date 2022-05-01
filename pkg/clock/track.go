package clock

import (
	"fmt"
)

func (track *ballTrack) addBall(b int) []int {
	if len(track.balls) == track.max {
		returnedBalls := reverseBallOrder(track.balls)
		track.balls = track.balls[:0]
		return returnedBalls
	}

	track.balls = append(track.balls, b)
	return nil
}

func reverseBallOrder(orig []int) []int {
	reversed := make([]int, len(orig))
	copy(reversed, orig)

	half := len(orig) / 2
	for i := 0; i < half; i++ {
		j := len(orig) - i - 1
		reversed[i], reversed[j] = orig[j], orig[i]
	}

	return reversed
}

func newTrack(name string, max int) *ballTrack {
	balls := make([]int, 0, max)
	return &ballTrack{
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