package comp2d

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

//Timer executes an event method when timer elapses
type Timer struct {
	elapsedCallback func()
	autoReset       bool
	running         bool
	interval        time.Duration
}

//Update updates timer
func (t *Timer) Update(delta int64) error {
	if t.autoReset {
		t.Start()
	}
	return nil
}

//Draw draws
func (t *Timer) Draw(screen *ebiten.Image) error {
	return nil
}

//Start starts the timer manually
func (t *Timer) Start() {
	if !t.running {
		t.running = true
		go func() {
			time.Sleep(t.interval)
			t.elapsedCallback()
			t.running = false
		}()
	}
}

//Reset resets the timer  manually
func (t *Timer) Reset() {
	t.Start()
}

//NewTimer creates a timer
func NewTimer(elapsedCallback func(), autoReset bool, interval time.Duration) *Timer {
	timer := Timer{
		elapsedCallback: elapsedCallback,
		autoReset:       autoReset,
		interval:        interval,
	}
	return &timer
}
