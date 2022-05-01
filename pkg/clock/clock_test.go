package clock

import (
	"fmt"
	"testing"
)

func TestClockIncrement(t *testing.T) {
	tests := []struct {
		name            string
		minsToIncrement int
		wantTime        string
	}{
		{
			name:            "should increment by 3 minutes successfully",
			minsToIncrement: 3,
			wantTime:        "01:03",
		},
		{
			name:            "should increment by 1 hour successfully",
			minsToIncrement: 60,
			wantTime:        "02:00",
		},
		{
			name:            "should increment by 12 hours successfully",
			minsToIncrement: 60 * 12,
			wantTime:        "01:00",
		},
		{
			name:            "should increment by 24 hours successfully",
			minsToIncrement: 60 * 24,
			wantTime:        "01:00",
		},
		{
			name:            "should increment by 47 hours 59 minutes successfully",
			minsToIncrement: 47*60 + 59,
			wantTime:        "12:59",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			clock := new(27)
			clock.incrementMultipleMin(test.minsToIncrement)

			if test.wantTime != clock.time() {
				t.Errorf("\"%v\" failed - got %s; want %s\n", test.name, clock.time(), test.wantTime)
			}
		})
	}
}

func TestDetermineClockState(t *testing.T) {
	var tracks []ballTrack
	tracks = append(tracks, ballTrack{balls: []int{}})
	tracks = append(tracks, ballTrack{balls: []int{22, 13, 25, 3, 7}})
	tracks = append(tracks, ballTrack{balls: []int{6, 12, 17, 4, 15}})
	wantClockState := ballClock{
		tracks: tracks,
		queue: ballQueue{
			balls: []int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9},
		},
	}

	clock := determineClockState(30, 325)
	if !clock.equals(&wantClockState) {
		t.Error("TestDetermineClockState failed. States were not the same:")
		json, _ := clock.marshallJSON()
		fmt.Println("Got:", string(json))
		json, _ = wantClockState.marshallJSON()
		fmt.Println("Wanted:", string(json))
	}
}

func TestDetermineCycleDays(t *testing.T) {
	tests := []struct {
		name     string
		numBalls int
		wantDays int
		wantErr  error
	}{
		{
			name:     "should successfully get cycle days with 30 balls",
			numBalls: 30,
			wantDays: 15,
			wantErr:  nil,
		},
		{
			name:     "should successfully get cycle days with 45 balls",
			numBalls: 45,
			wantDays: 378,
			wantErr:  nil,
		},
		{
			name:     "should return error with input of 26 balls",
			numBalls: 26,
			wantDays: 0,
			wantErr:  fmt.Errorf("number of balls must be between %d and %d", minBalls, maxBalls),
		},
		{
			name:     "should return error with input of 128 balls",
			numBalls: 128,
			wantDays: 0,
			wantErr:  fmt.Errorf("number of balls must be between %d and %d", minBalls, maxBalls),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			days, err := determineCycleDays(test.numBalls)

			if err == nil && test.wantErr != nil {
				t.Errorf("\"%v\" failed - got %v; want %v\n", test.name, err, test.wantErr)
			}
			if test.wantDays != days {
				t.Errorf("\"%v\" failed - got %v; want %v\n", test.name, days, test.wantDays)
			}
		})
	}
}
