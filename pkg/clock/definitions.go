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

const (
	minutePos  int = 0
	fiveMinPos int = 1
	hrPos          = 2
)

type clockBall int

type ballTrack struct {
	name  string
	balls []clockBall
	max   int
}

type ballQueue struct {
	balls []clockBall
}

type ballClock struct {
	tracks []ballTrack
	queue  ballQueue
}

type clockJSON struct {
	OneMinTrack  []clockBall `json:"Min"`
	FiveMinTrack []clockBall `json:"FiveMin"`
	HourTrack    []clockBall `json:"Hour"`
	Queue        []clockBall `json:"Main"`
}
