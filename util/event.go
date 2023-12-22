package util

type Event struct {
	SkipEvent chan bool // True : when a Skip event occurs
	StopEvent chan bool // True : when a Stop event occurs
}
