package heartbeat

import (
	"context"
	"time"
)

// Heartbeat will get a function and run keep sending the heartbeat
func Heartbeat(ctx context.Context, interval time.Duration, work func() interface{}) <-chan interface{} {
	heartbeatStream := make(chan interface{})
	go func() {
		defer close(heartbeatStream)
		for {
			<-time.Tick(interval)
			select {
			case <-ctx.Done():
				return
			default:
				heartbeatStream <- work()
			}
		}
	}()
	return heartbeatStream
}
