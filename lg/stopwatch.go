package lg

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start, stop time.Time
	logger      *Logger
}

func (s *Stopwatch) IsStopped() bool { return s.stop.After(s.start) }
func (s *Stopwatch) Stop() {
	s.stop = time.Now()
}

func (s *Stopwatch) ElapsedTime() time.Duration {
	if s.IsStopped() {
		return s.stop.Sub(s.start)
	}
	return time.Since(s.start)
}

func (s *Stopwatch) Dt() string {
	return fmt.Sprintf("%s", s.ElapsedTime())
}

func (s *Stopwatch) Log(msg string) {
	s.logger.Info("elapsed", "msg", msg, "dx", s.Dt())
}
