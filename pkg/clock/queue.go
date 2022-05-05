package clock

import "fmt"

func (bq *ballQueue) removeBall() uint8 {
	nextBall := bq.balls[bq.start]
	bq.start++
	if bq.start == bq.max {
		bq.start = 0
	}
	return nextBall
}

func (bq *ballQueue) addBall(ball uint8) {
	bq.end++
	if bq.end == bq.max {
		bq.end = 0
	}
	bq.balls[bq.end] = ball
}

func (bq *ballQueue) equals(otherQueue *ballQueue) bool {
	if bq.start != otherQueue.start || bq.end != otherQueue.end {
		return false
	}
	for i := 0; i < len(bq.balls); i++ {
		if bq.balls[i] != otherQueue.balls[i] {
			return false
		}
	}

	return true
}

func (bq *ballQueue) getList() []uint8 {
	list := make([]uint8, 0, bq.max)
	if bq.start < bq.end {
		for i := bq.start; i <= bq.end; i++ {
			list = append(list, bq.balls[i])
		}
	} else {
		for i := bq.start; i < bq.max; i++ {
			list = append(list, bq.balls[i])
		}
		for i := 0; i <= bq.end; i++ {
			list = append(list, bq.balls[i])
		}
	}

	return list
}

func newQueue(numBalls uint8) ballQueue {
	queueBalls := make([]uint8, 0, numBalls)

	var i uint8
	for i = 1; i <= uint8(numBalls); i++ {
		queueBalls = append(queueBalls, i)
	}

	return ballQueue{
		balls: queueBalls,
		end:   int(numBalls) - 1,
		start: 0,
		max:   int(numBalls),
	}
}

// print is a helper function for visualizing the queue (debugging)
func (bq *ballQueue) print() {
	fmt.Println("Queue:")
	for _, ball := range bq.balls {
		fmt.Printf("[%d] ", ball)
	}
	fmt.Println()
}
