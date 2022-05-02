package clock

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// State prints the state of a ball clock with the given number of balls
// and ran for the specified minutes in a JSON format.
func State(numBalls uint8, minToRun int) error {
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
func CycleDays(numBalls uint8) error {
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

func determineClockState(numBalls uint8, minToRun int) *ballClock {
	clock := newClock(numBalls)
	clock.incrementMultipleMin(minToRun)
	return clock
}

func (c *ballClock) marshallJSON() ([]byte, error) {
	clock := clockJSON{
		OneMinTrack:  c.minTrack.balls,
		FiveMinTrack: c.fiveMinTrack.balls,
		HourTrack:    c.hrTrack.balls,
		Queue:        c.queue.balls,
	}

	js, err := json.Marshal(clock)
	if err != nil {
		return js, err
	}

	return js, nil
}

func determineCycleDays(numBalls uint8) int {
	c := newClock(numBalls)
	initialClock := *c

	// No need to check if the states are equal before the calculated minimum
	// clock.incrementMultipleMin(minMinutesToRepeat - 1)
	min := 0
	for min < minMinutesToRepeat {
		nextBall := c.queue.balls[0]
		c.queue.balls = c.queue.balls[1:]

		t := &c.minTrack
		//fmt.Printf("t (before): %v\n", *t)
		//fmt.Printf("c.minTrack address: %p\tt address: %p\n", &c.minTrack, t)
		if len(t.balls) == t.max {
			for j := t.max - 1; j >= 0; j-- {
				ball := t.balls[j]
				c.queue.balls = append(c.queue.balls, ball)
			}
			t.balls = t.balls[:0]
		} else {
			t.balls = append(t.balls, nextBall)
			min++
			continue
		}

		t = &c.fiveMinTrack
		if len(t.balls) == t.max {
			for j := t.max - 1; j >= 0; j-- {
				ball := t.balls[j]
				c.queue.balls = append(c.queue.balls, ball)
			}
			t.balls = t.balls[:0]
		} else {
			t.balls = append(t.balls, nextBall)
			min++
			continue
		}

		t = &c.hrTrack
		if len(t.balls) == t.max {
			for j := t.max - 1; j >= 0; j-- {
				ball := t.balls[j]
				c.queue.balls = append(c.queue.balls, ball)
			}
			t.balls = t.balls[:0]
		} else {
			t.balls = append(t.balls, nextBall)
			min++
			continue
		}

		c.queue.balls = append(c.queue.balls, nextBall)
		min++
	}

	min = minMinutesToRepeat
	for !c.equals(&initialClock) {
		nextBall := c.queue.balls[0]
		c.queue.balls = c.queue.balls[1:]

		t := &c.minTrack
		//fmt.Printf("t (before): %v\n", *t)
		//fmt.Printf("c.minTrack address: %p\tt address: %p\n", &c.minTrack, t)
		if len(t.balls) == t.max {
			for j := t.max - 1; j >= 0; j-- {
				ball := t.balls[j]
				c.queue.balls = append(c.queue.balls, ball)
			}
			t.balls = t.balls[:0]
		} else {
			t.balls = append(t.balls, nextBall)
			min++
			continue
		}

		t = &c.fiveMinTrack
		if len(t.balls) == t.max {
			for j := t.max - 1; j >= 0; j-- {
				ball := t.balls[j]
				c.queue.balls = append(c.queue.balls, ball)
			}
			t.balls = t.balls[:0]
		} else {
			t.balls = append(t.balls, nextBall)
			min++
			continue
		}

		t = &c.hrTrack
		if len(t.balls) == t.max {
			for j := t.max - 1; j >= 0; j-- {
				ball := t.balls[j]
				c.queue.balls = append(c.queue.balls, ball)
			}
			t.balls = t.balls[:0]
		} else {
			t.balls = append(t.balls, nextBall)
			min++
			continue
		}

		c.queue.balls = append(c.queue.balls, nextBall)
		min++
	}

	return minutesToDays(min)
}

func validateInput(numBalls uint8) error {
	if numBalls < minBalls || numBalls > maxBalls {
		return fmt.Errorf("number of balls must be between %d and %d", minBalls, maxBalls)
	}
	return nil
}

func newClock(numBalls uint8) *ballClock {
	return &ballClock{
		minTrack:     *newTrack(oneMinTrackName, oneMinTrackMax),
		fiveMinTrack: *newTrack(fiveMinTrackName, fiveMinTrackMax),
		hrTrack:      *newTrack(hourTrackName, hourTrackMax),
		queue:        *newQueue(numBalls),
	}
}

func (c *ballClock) incrementMultipleMin(minutes int) {
	for i := 0; i < minutes; i++ {
		c.incrementOneMin()
	}
}

func (c *ballClock) dropTrackBalls(t *ballTrack) {
	for j := t.max - 1; j >= 0; j-- {
		ball := t.getBall(j)
		c.queue.addBall(ball)
	}
	t.empty()
}

// incrementOneMin implements the core functionality of the ball clock
// simulation. It increments the clock by one minute, and modifies the
// state of the clock accordingly.
func (c *ballClock) incrementOneMin() {
	nextBall := c.queue.balls[0]
	c.queue.balls = c.queue.balls[1:]

	t := &c.minTrack
	//fmt.Printf("t (before): %v\n", *t)
	//fmt.Printf("c.minTrack address: %p\tt address: %p\n", &c.minTrack, t)
	if len(t.balls) == t.max {
		for j := t.max - 1; j >= 0; j-- {
			ball := t.balls[j]
			c.queue.balls = append(c.queue.balls, ball)
		}
		t.balls = t.balls[:0]
	} else {
		t.balls = append(t.balls, nextBall)
		return
	}

	t = &c.fiveMinTrack
	if len(t.balls) == t.max {
		for j := t.max - 1; j >= 0; j-- {
			ball := t.balls[j]
			c.queue.balls = append(c.queue.balls, ball)
		}
		t.balls = t.balls[:0]
	} else {
		t.balls = append(t.balls, nextBall)
		return
	}

	t = &c.hrTrack
	if len(t.balls) == t.max {
		for j := t.max - 1; j >= 0; j-- {
			ball := t.balls[j]
			c.queue.balls = append(c.queue.balls, ball)
		}
		t.balls = t.balls[:0]
	} else {
		t.balls = append(t.balls, nextBall)
		return
	}

	c.queue.balls = append(c.queue.balls, nextBall)
}

func (c *ballClock) equals(otherClock *ballClock) bool {
	return c.queue.equals(&otherClock.queue)
}

// time is a helper function for debugging and testing
func (c *ballClock) time() string {
	hour := len(c.hrTrack.balls) + 1
	fiveMin := len(c.fiveMinTrack.balls) * 5
	minute := len(c.minTrack.balls) + fiveMin

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
	c.minTrack.print(0)
	fmt.Println()
	c.fiveMinTrack.print(0)
	fmt.Println()
	c.hrTrack.print(1)
	fmt.Println()
	c.queue.print()
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println()
}
