package obs

import (
	"context"
	"sync"
	"time"
)

type Span struct {
	method    string
	startTime time.Time
	mux       sync.Mutex
	events    []string
}

type contextKeySpan struct{}

func CreateSpan(c context.Context, method string) (context.Context, *Span) {
	if c == nil {
		c = context.Background()
	}
	s := &Span{
		method:    method,
		startTime: time.Now(),
	}
	c = context.WithValue(c, contextKeySpan{}, s)
	return c, s
}
