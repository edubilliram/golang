package main

import (
	"context"
	"fmt"
	"time"
)

func testCtx(ctx context.Context) {

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	return
}

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second+500*time.Millisecond)
	defer cancel()
	testCtx(ctx)

	testCtx(ctx)
}
