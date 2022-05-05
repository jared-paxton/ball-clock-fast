package clock

const (
	oneMinTrackMax  int = 4
	fiveMinTrackMax int = 11
	hourTrackMax    int = 11
)

const (
	maxBalls uint8 = 127
	minBalls uint8 = 27
	// First time at which 27 balls will repeat
	minMinutesToRepeat int = 21600
)

const (
	oneMinTrackName  string = "Min"
	fiveMinTrackName string = "FiveMin"
	hourTrackName    string = "Hour"
)

type ballTrack struct {
	name  string
	balls []uint8
	max   int
}

type ballClock struct {
	minTrack     ballTrack
	fiveMinTrack ballTrack
	hrTrack      ballTrack
	queue        []uint8
}

type clockJSON struct {
	OneMinTrack  []uint8 `json:"Min"`
	FiveMinTrack []uint8 `json:"FiveMin"`
	HourTrack    []uint8 `json:"Hour"`
	Queue        []uint8 `json:"Main"`
}
