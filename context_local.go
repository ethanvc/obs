package obs

import (
	"context"
	"sync"
)

type contextLocal struct {
	mu  sync.Mutex
	kvs map[any]any
}

func (cl *contextLocal) Set(k, v any) *contextLocal {
	if cl == nil {
		return nil
	}
	cl.mu.Lock()
	cl.kvs[k] = v
	cl.mu.Unlock()
	return cl
}

func (cl *contextLocal) AddSliceItem(k, item any) *contextLocal {
	if cl == nil {
		return nil
	}
	cl.mu.Lock()
	v, _ := cl.kvs[k].([]any)
	v = append(v, item)
	cl.kvs[k] = v
	cl.mu.Unlock()
	return cl
}

func (cl *contextLocal) GetWithExistence(k any) (any, bool) {
	if cl == nil {
		return nil, false
	}
	cl.mu.Lock()
	v, ok := cl.kvs[k]
	cl.mu.Unlock()
	return v, ok
}

func (cl *contextLocal) Get(k any) any {
	v, _ := cl.GetWithExistence(k)
	return v
}

type contextKeyContextLocal struct{}

func CreateContextLocal(c context.Context) context.Context {
	if c == nil {
		c = context.Background()
	}
	return context.WithValue(c, contextKeyContextLocal{}, &contextLocal{})
}

func GetContextLocal(c context.Context) *contextLocal {
	if c == nil {
		return nil
	}
	v, _ := c.Value(contextKeyContextLocal{}).(*contextLocal)
	return v
}
