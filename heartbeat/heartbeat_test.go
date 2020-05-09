package heartbeat

import (
	"context"
	"testing"
	"time"
)

func TestHeartbeat_CallAfterInterval(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	count := 0
	beatStream := Heartbeat(ctx, 3*time.Millisecond, func() interface{} {
		count++
		return nil
	})
	go func() {
		// range over it so that we don't block write to the channel
		for range beatStream {
		}
	}()
	time.Sleep(10 * time.Millisecond)
	if count != 3 {
		t.Errorf("Expected the function to invoke 3 time instead %d", count)
	}
}
func TestHeartbeat_ShouldCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	beatStream := Heartbeat(ctx, 1*time.Millisecond, func() interface{} {
		return nil
	})
	cancel()
	_, ok := <-beatStream
	if ok {
		t.Errorf("No work function should fire if cancel")
	}

}
