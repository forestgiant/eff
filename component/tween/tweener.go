package tween

import (
	"math"
	"time"
)

// EaseFunc function that takes a normalized value for input and returns an eased value based on that input
type EaseFunc func(float64) float64

// Tweener struct that changes a percentage value over time using
type Tweener struct {
	duration    time.Duration
	started     bool
	startTime   time.Time
	update      func(float64)
	complete    func()
	repeat      bool
	repeatCount uint64
	yoyo        bool
	yoyoDir     bool
	Done        bool
	ease        EaseFunc
}

// Tween each time this is called this pumps the tween, the more frequently it is called the smoother the tween
func (tweener *Tweener) Tween() {
	if !tweener.started {
		tweener.started = true
		tweener.startTime = time.Now()
	}
	elapsedTime := time.Now()
	progress := float64(elapsedTime.UnixNano()-tweener.startTime.UnixNano()) / float64(tweener.duration.Nanoseconds())
	progress -= float64(tweener.repeatCount)
	if tweener.yoyoDir {
		progress = 1 - progress
	}
	if progress > 1 || progress < 0 {
		if tweener.repeat {
			tweener.repeatCount++
			if tweener.yoyo {
				tweener.yoyoDir = !tweener.yoyoDir
			}
		} else {
			tweener.Done = true
		}

		if tweener.complete != nil {
			tweener.complete()
		}
	}

	progress = math.Min(progress, 1)
	progress = math.Max(progress, 0)

	if tweener.ease != nil {
		progress = tweener.ease(progress)
	}

	tweener.update(progress)
}

// Reset sets the tween state variables back to there starting points
func (tweener *Tweener) Reset() {
	tweener.started = false
	tweener.repeatCount = 0
}

// NewTweener creates a new tweener struct
func NewTweener(duration time.Duration, update func(float64), repeat bool, yoyo bool, complete func(), ease EaseFunc) Tweener {
	return Tweener{
		duration: duration,
		update:   update,
		repeat:   repeat,
		yoyo:     yoyo,
		complete: complete,
		ease:     ease,
	}
}
