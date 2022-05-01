package clock

const (
	oneMinTrackMax  int = 4
	fiveMinTrackMax int = 11
	hourTrackMax    int = 11
)

const (
	maxBalls int = 127
	minBalls int = 27
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
	balls []int
	max   int
}

type ballQueue struct {
	balls []int
}

type ballClock struct {
	minTrack     ballTrack
	fiveMinTrack ballTrack
	hrTrack      ballTrack
	queue        ballQueue
}

type clockJSON struct {
	OneMinTrack  []int `json:"Min"`
	FiveMinTrack []int `json:"FiveMin"`
	HourTrack    []int `json:"Hour"`
	Queue        []int `json:"Main"`
}
