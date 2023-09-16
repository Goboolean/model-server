package context

import (
	"context"
	"time"
)

// A struct customContext implements every function of the context.Context interface
// It is an expansion of the golang context, which provides SetDeadline() and RemoveDeadline() additionally

// 포인터를 잘 활용해서 ctx 갱신이 잘 되어야 한다 아직은 어렵다
type customContext struct {
	ctx context.Context
	channel chan struct{}
	cancelFunc context.CancelFunc

}


func New(ctx context.Context) (context.Context, context.CancelFunc) {
	
	channel := make(chan struct{})

	c := &customContext{
		ctx: ctx,
		channel: channel,
	}

	go func() {
		<- ctx.Done()
	}()
	
	return c, nil
}

// Set deadline of ctx during runtime
func (c *customContext) SetDeadline() {
	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	c.ctx = ctx
	c.cancelFunc = cancel
}

// Remove deadline of ctx during runtime
func (c *customContext) RemoveDeadline() {
	ctx := context.Background()
	c.ctx = ctx
}


func (c *customContext) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *customContext) Done() <- chan struct{} {
	return c.ctx.Done()
}

func (c *customContext) Err() error {
	return c.ctx.Err()
}

func (c *customContext) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}