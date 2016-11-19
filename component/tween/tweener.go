package tween

import (
	"math"
	"time"
)

type Tweener struct {
	duration  time.Duration
	startTime time.Time
	update    func(float64)
	repeat    bool
	Done      bool
}

func (tweener *Tweener) Tween() {
	elapsedTime := time.Now()
	progress := float64(elapsedTime.UnixNano()-tweener.startTime.UnixNano()) / float64(tweener.duration.Nanoseconds())
	if progress > 1 {
		if tweener.repeat {
			progress = math.Mod(progress, float64(1))
		} else {
			tweener.Done = true
		}

	}
	tweener.update(math.Min(progress, 1))
}

func NewTweener(duration time.Duration, update func(float64), repeat bool) Tweener {
	return Tweener{
		duration:  duration,
		startTime: time.Now(),
		update:    update,
		repeat:    repeat,
	}
}
