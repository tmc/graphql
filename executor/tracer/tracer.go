// Package tracer provides an interface to attach execution metadata to a context.
package tracer

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
)

type Tracer struct {
	ID                 uint64
	startTime, endTime time.Time
	Duration           time.Duration `json:"-"`
	DurationMicros     int64
	DurationMillis     int64
	Extra              map[string]interface{} `json:",omitempty"`

	mu      sync.Mutex
	Queries int
}

func New(id uint64) *Tracer {
	return &Tracer{
		ID:        id,
		startTime: time.Now(),
	}
}

func FromRequest(r *http.Request) (*Tracer, error) {
	tID := r.Header.Get("X-Trace-ID")
	if tID == "" {
		return nil, fmt.Errorf("tracer: missing X-Trace-ID header")
	}
	traceId, err := strconv.ParseUint(tID, 10, 0)
	if err != nil {
		return nil, err
	}
	return New(traceId), nil
}

func (t *Tracer) IncQueries(n int) int {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Queries += n
	return t.Queries
}

func (t *Tracer) WithLock(fn func(*Tracer)) {
	t.mu.Lock()
	defer t.mu.Unlock()
	fn(t)
}

func (t *Tracer) Done() {
	t.endTime = time.Now()
	t.Duration = t.endTime.Sub(t.startTime)
	t.DurationMicros = t.Duration.Nanoseconds() / 1000
	t.DurationMillis = t.Duration.Nanoseconds() / 1000000
}

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// tracerKey is the context key for the tracers attached to a context.
const tracerKey = 0

// NewContext returns a new Context carrying tracer.
func NewContext(ctx context.Context, t *Tracer) context.Context {
	return context.WithValue(ctx, tracerKey, t)
}

// FromContext extracts the tracer from ctx, if present.
func FromContext(ctx context.Context) (*Tracer, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns ok=false for nil.
	t, ok := ctx.Value(tracerKey).(*Tracer)
	return t, ok
}
