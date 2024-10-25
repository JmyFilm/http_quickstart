package utils

import (
	"sync"
	"time"
)

var timings = struct {
	T map[int][]*time.Timer
	sync.Mutex
}{
	T:     make(map[int][]*time.Timer),
	Mutex: sync.Mutex{},
}

func AddTimer(id int, t time.Time, afterFunc ...func()) {
	if time.Now().Before(t) && !t.IsZero() {
		timings.Lock()
		defer timings.Unlock()
		for _, fn := range afterFunc {
			timings.T[id] = append(timings.T[id], time.AfterFunc(t.Sub(time.Now()), fn))
		}
	}
}

func DelTimer(ids ...int) {
	timings.Lock()
	defer timings.Unlock()

	for _, id := range ids {
		for _, t := range timings.T[id] {
			go t.Stop()
		}
		delete(timings.T, id)
	}
}

func DelAllTimer() {
	timings.Lock()
	defer timings.Unlock()

	for id, ts := range timings.T {
		for _, t := range ts {
			go t.Stop()
		}
		delete(timings.T, id)
	}
}
