package thread

import (
	"sync"
	"time"
)

type Limiter struct {
	rate       time.Duration
	burst      int
	count      int
	mu         sync.Mutex
	resetTime  time.Time
	waiting    int
	waitingCh  chan struct{}
	isBlocking bool
}
