package irclone

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func CloneFromPIN(pin rpio.Pin, cb func([]time.Duration)) func() {
	var close = make(chan struct{}, 1)
	var timming []time.Duration
	go func() {
		var curstate rpio.State
		var elapsed time.Time
		for {
			select {
			case <-close:
				break
			case <-time.After(time.Microsecond):
				if curstate != pin.Read() {
					curstate = pin.Read()
					if elapsed.IsZero() || len(timming) == 0 {
						timming = append(timming, 0)
					} else {
						timming = append(timming, time.Since(elapsed))
					}
					elapsed = time.Now()
				} else if time.Since(elapsed) > time.Second && len(timming) > 0 {
					cb(timming)
					timming = make([]time.Duration, 0)
					elapsed = time.Now()
				}
			}
		}
	}()
	return func() {
		close <- struct{}{}
	}
}
