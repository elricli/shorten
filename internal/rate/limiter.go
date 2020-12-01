package rate

import (
	"sync"
	"time"
)

type Limiter struct {
	mu     sync.Mutex
	last   time.Time
	rate   float64
	burst  int
	tokens float64
}

func NewLimiter(rate float64, b int) *Limiter {
	return &Limiter{
		rate:  rate,
		burst: b,
	}
}

func (l *Limiter) Allow() bool {
	return l.AllowN(time.Now(), 1)
}

func (l *Limiter) AllowN(now time.Time, n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	last := l.last
	if now.Before(last) {
		last = now
	}
	// Calculate the generation time of used tokens
	d := (float64(l.burst) - l.tokens) / l.rate
	maxElapsed := time.Duration(d*1e9) * time.Nanosecond
	elapsed := now.Sub(last)
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}
	// Calculate the number of tokens generated during elapsed time
	tokens := float64(elapsed/time.Second) * l.rate
	tokens += float64(elapsed%time.Second) * l.rate / 1e9
	tokens += l.tokens
	if burst := float64(l.burst); tokens > burst {
		tokens = burst
	}
	tokens -= float64(n)
	if tokens >= 0 {
		l.last = now
		l.tokens = tokens
		return true
	}
	l.last = last
	return false
}
