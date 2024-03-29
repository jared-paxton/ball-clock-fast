package clock

// func TestAddBall(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		track         ballTrack
// 		ballToAdd     int
// 		wantBallOrder []int
// 		wantReturn    []int
// 	}{
// 		{
// 			name: "should add ball to empty track",
// 			track: ballTrack{
// 				name:  "Minute",
// 				balls: []int{},
// 				max:   4,
// 			},
// 			ballToAdd:     1,
// 			wantBallOrder: []int{1},
// 			wantReturn:    nil,
// 		},
// 		{
// 			name: "should add ball to partially full track",
// 			track: ballTrack{
// 				name:  "Minute",
// 				balls: []int{32, 4},
// 				max:   4,
// 			},
// 			ballToAdd:     10,
// 			wantBallOrder: []int{32, 4, 10},
// 			wantReturn:    nil,
// 		},
// 		{
// 			name: "should trigger balls to \"fall\" from track",
// 			track: ballTrack{
// 				name:  "Minute",
// 				balls: []int{50, 1, 127, 43},
// 				max:   4,
// 			},
// 			ballToAdd:     10,
// 			wantBallOrder: []int{},
// 			wantReturn:    []int{43, 127, 1, 50},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			returnedBalls := test.track.addBall(test.ballToAdd)

// 			if returnedBalls != nil && test.wantReturn == nil {
// 				t.Errorf("\"%v\" failed - got %v; want %v\n", test.name, returnedBalls, test.wantReturn)
// 			}
// 			if returnedBalls != nil && test.wantReturn != nil {
// 				if len(test.wantReturn) != len(returnedBalls) {
// 					t.Errorf("\"%v\" failed - got %v; want %v\n", test.name, returnedBalls, test.wantReturn)
// 				} else {
// 					if !reflect.DeepEqual(returnedBalls, test.wantReturn) {
// 						t.Errorf("\"%v\" failed - got %v; want %v\n",
// 							test.name, returnedBalls, test.wantReturn)
// 					}
// 				}
// 			}
// 			if !reflect.DeepEqual(test.track.balls, test.wantBallOrder) {
// 				t.Errorf("\"%v\" failed - got %v; want %v\n",
// 					test.name, test.track.balls, test.wantBallOrder)
// 			}
// 		})
// 	}
//}
