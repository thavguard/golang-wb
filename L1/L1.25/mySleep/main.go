package mysleep

import "time"

func MySleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func MySleepFor(duration time.Duration) {
	targetTime := time.Now().Add(duration)

	for {
		if ok := time.Now().Equal(targetTime); ok {
			return
		}
	}
}
