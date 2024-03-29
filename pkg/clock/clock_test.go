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
			clock := newClock(27)
			clock.incrementMultipleMin(test.minsToIncrement)

			if test.wantTime != clock.time() {
				t.Errorf("\"%v\" failed - got %s; want %s\n", test.name, clock.time(), test.wantTime)
			}
		})
	}
}

func TestDetermineClockState(t *testing.T) {
	wantClockState := ballClock{
		minTrack:     ballTrack{balls: []uint8{}},
		fiveMinTrack: ballTrack{balls: []uint8{22, 13, 25, 3, 7}},
		hrTrack:      ballTrack{balls: []uint8{6, 12, 17, 4, 15}},
		queue: ballQueue{
			balls: []uint8{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9},
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
		numBalls uint8
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			days := determineCycleDays(test.numBalls)

			if test.wantDays != days {
				t.Errorf("\"%v\" failed - got %v; want %v\n", test.name, days, test.wantDays)
			}
		})
	}
}
