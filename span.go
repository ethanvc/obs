package obs

import (
	"context"
	"strings"
	"sync"
	"time"
)

type Span struct {
	clientSpan bool
	method     string
	startTime  time.Time
	mux        sync.Mutex
	events     []string
}

type contextKeySpan struct{}

func CreateSvrSpan(c context.Context, method string) (context.Context, *Span) {
	return createSpan(c, false, method)
}

func CreateClientSpan(c context.Context, method string) (context.Context, *Span) {
	return createSpan(c, true, method)
}

func createSpan(c context.Context, clientSpan bool, method string) (context.Context, *Span) {
	if c == nil {
		c = context.Background()
	}
	s := &Span{
		clientSpan: clientSpan,
		method:     method,
		startTime:  time.Now(),
	}
	c = context.WithValue(c, contextKeySpan{}, s)
	return c, s
}

func GetSpan(c context.Context) *Span {
	if c == nil {
		return nil
	}
	v, _ := c.Value(contextKeySpan{}).(*Span)
	return v
}

func (span *Span) Report(c context.Context, event string) *Span {
	if span == nil {
		return nil
	}
	span.mux.Lock()
	span.events = append(span.events, event)
	span.mux.Unlock()
	return span
}

func (span *Span) ReportErr(c context.Context, event string) *Span {
	if !strings.HasPrefix(event, "Err") {
		event = "Err" + event
	}
	return span.Report(c, event)
}
