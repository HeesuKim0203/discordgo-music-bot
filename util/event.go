package util

type Event struct {
	skipEvent chan bool // True : when a Skip event occurs
	exitEvent chan bool // True : when a Stop event occurs

	exit bool
}

func NewEvent() *Event {
	return &Event{
		skipEvent: make(chan bool, 1),
		exitEvent: make(chan bool, 1),
	}
}

func (e *Event) Skip() {
	e.skipEvent <- true
}

func (e *Event) Stop() {
	e.exit = true
	e.exitEvent <- true
}

func (e *Event) GetSkipEvent() chan bool {
	return e.skipEvent
}

func (e *Event) GetExitEvent() chan bool {
	return e.exitEvent
}

func (e *Event) GetExitState() bool {
	return e.exit
}
