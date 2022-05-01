package clock

import "fmt"

func (bq *ballQueue) removeBall() uint8 {
	nextBall := bq.balls[0]
	bq.balls = bq.balls[1:]
	return nextBall
}

func (bq *ballQueue) addBalls(balls *[]uint8) {
	bq.balls = append(bq.balls, *balls...)
}

func (bq *ballQueue) addBall(ball uint8) {
	bq.balls = append(bq.balls, ball)
}

func (bq *ballQueue) equals(otherQueue *ballQueue) bool {
	if len(bq.balls) != len(otherQueue.balls) {
		return false
	}
	for i := 0; i < len(bq.balls); i++ {
		if bq.balls[i] != otherQueue.balls[i] {
			return false
		}
	}

	return true
}

func newQueue(numBalls uint8) *ballQueue {
	queueBalls := make([]uint8, 0, numBalls)

	var i uint8
	for i = 1; i <= numBalls; i++ {
		queueBalls = append(queueBalls, i)
	}

	return &ballQueue{
		balls: queueBalls,
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
