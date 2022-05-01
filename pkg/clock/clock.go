package clock

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// State prints the state of a ball clock with the given number of balls
// and ran for the specified minutes in a JSON format.
func State(numBalls, minToRun int) error {
	err := validateInput(numBalls)
	if err != nil {
		return err
	}

	clock := determineClockState(numBalls, minToRun)
	jsonOutput, err := clock.marshallJSON()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println(string(jsonOutput))
	fmt.Println()
	return nil
}

// CycleDays prints the number of days it takes for the ordering of the balls in the clock
// to return to the same order in its initial state, given the number of balls.
func CycleDays(numBalls int) error {
	err := validateInput(numBalls)
	if err != nil {
		return err
	}

	start := time.Now()
	days := determineCycleDays(numBalls)
	duration := time.Since(start)

	fmt.Printf("%d balls cycle after %d days\n", numBalls, days)
	fmt.Printf("Completed in %d milliseconds (%.3f seconds)\n", duration.Milliseconds(), duration.Seconds())
	return nil
}

func determineClockState(numBalls, minToRun int) *ballClock {
	clock := new(numBalls)
	clock.incrementMultipleMin(minToRun)
	return clock
}

func (c *ballClock) marshallJSON() ([]byte, error) {
	clock := clockJSON{
		OneMinTrack:  c.tracks[minutePos].balls,
		FiveMinTrack: c.tracks[fiveMinPos].balls,
		HourTrack:    c.tracks[hrPos].balls,
		Queue:        c.queue.balls,
	}

	js, err := json.Marshal(clock)
	if err != nil {
		return js, err
	}

	return js, nil
}

func determineCycleDays(numBalls int) int {
	clock := new(numBalls)
	initialClock := *clock

	// No need to check if the states are equal before the calculated minimum
	clock.incrementMultipleMin(minMinutesToRepeat - 1)

	min := minMinutesToRepeat
	for !clock.equals(&initialClock) {
		clock.incrementOneMin()
		min++
	}

	return minutesToDays(min)
}

func validateInput(numBalls int) error {
	if numBalls < minBalls || numBalls > maxBalls {
		return fmt.Errorf("number of balls must be between %d and %d", minBalls, maxBalls)
	}
	return nil
}

func new(numBalls int) *ballClock {
	t := make([]ballTrack, 0, 3)
	t = append(t, *newTrack(oneMinTrackName, oneMinTrackMax))
	t = append(t, *newTrack(fiveMinTrackName, fiveMinTrackMax))
	t = append(t, *newTrack(hourTrackName, hourTrackMax))

	return &ballClock{
		tracks: t,
		queue:  *newQueue(numBalls),
	}
}

func (c *ballClock) incrementMultipleMin(minutes int) {
	for i := 0; i < minutes; i++ {
		c.incrementOneMin()
	}
}

// incrementOneMin implements the core functionality of the ball clock
// simulation. It increments the clock by one minute, and modifies the
// state of the clock accordingly.
func (c *ballClock) incrementOneMin() {
	nextBall := c.queue.removeBall()

	for i := range c.tracks {
		returningBalls := c.tracks[i].addBall(nextBall)
		if returningBalls == nil {
			return
		}
		c.queue.addBalls(&returningBalls)
	}

	c.queue.addBall(nextBall)
}

func (c *ballClock) equals(otherClock *ballClock) bool {
	return c.queue.equals(&otherClock.queue)
}

// time is a helper function for debugging and testing
func (c *ballClock) time() string {
	hour := len(c.tracks[hrPos].balls) + 1
	fiveMin := len(c.tracks[fiveMinPos].balls) * 5
	minute := len(c.tracks[minutePos].balls) + fiveMin

	hourStr := strconv.Itoa(hour)
	if hour < 10 {
		hourStr = fmt.Sprintf("0%d", hour)
	}
	minuteStr := strconv.Itoa(minute)
	if minute < 10 {
		minuteStr = fmt.Sprintf("0%d", minute)
	}

	return hourStr + ":" + minuteStr
}

// print is a helper function for visualizing the clock (debugging)
func (c *ballClock) print() {
	fmt.Println("-------------------------------------------------------------------")
	c.tracks[minutePos].print(0)
	fmt.Println()
	c.tracks[fiveMinPos].print(0)
	fmt.Println()
	c.tracks[hrPos].print(1)
	fmt.Println()
	c.queue.print()
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println()
}
