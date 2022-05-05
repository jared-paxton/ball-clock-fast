package clock

import (
	"testing"
)

func TestAddBAll(t *testing.T) {
	tests := []struct {
		name       string
		queue      ballQueue
		wantQueue  ballQueue
		wantReturn uint8
	}{
		{
			name: "should remove first ball from queue",
			queue: ballQueue{
				balls: []uint8{55, 3, 29, 100},
				max:   4,
				end:   3,
				start: 0,
			},
			wantQueue: ballQueue{
				balls: []uint8{55, 3, 29, 100},
				max:   4,
				end:   3,
				start: 1,
			},
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
			if !test.wantQueue.equals(&test.queue) {
				t.Errorf("\"%v\" failed - got %v; want %v\n",
					test.name, test.queue, test.wantQueue)
			}
		})
	}
}
