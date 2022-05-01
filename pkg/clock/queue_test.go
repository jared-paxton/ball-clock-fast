package clock

import (
	"reflect"
	"testing"
)

func TestAddBAll(t *testing.T) {
	tests := []struct {
		name       string
		queue      ballQueue
		wantBalls  []int
		wantReturn int
	}{
		{
			name: "should remove first ball from queue",
			queue: ballQueue{
				balls: []int{55, 3, 29, 100},
			},
			wantBalls:  []int{3, 29, 100},
			wantReturn: 55,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nextBall := test.queue.removeBall()
			if nextBall != test.wantReturn {
				t.Errorf("\"%v\" failed - got %d; want %d\n",
					test.name, nextBall, test.wantReturn)
			}
			if !reflect.DeepEqual(test.wantBalls, test.queue.balls) {
				t.Errorf("\"%v\" failed - got %v; want %v\n",
					test.name, test.queue.balls, test.wantBalls)
			}
		})
	}
}
