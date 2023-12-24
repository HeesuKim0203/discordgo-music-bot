package util

type Event struct {
	skipEvent chan bool // True : when a Skip event occurs
	stopEvent chan bool // True : when a Stop event occurs
}

func NewEvent() *Event {
	return &Event{
		skipEvent: make(chan bool, 1),
		stopEvent: make(chan bool, 1),
	}
}

func (e *Event) Skip() {
	e.skipEvent <- true
}

func (e *Event) Stop() {
	e.stopEvent <- true
}

func (e *Event) GetSkipEvent() chan bool {
	return e.skipEvent
}

func (e *Event) GetStopEvent() chan bool {
	return e.stopEvent
}
