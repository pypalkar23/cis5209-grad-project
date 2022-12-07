package process_mgmt

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func danglingSenderAsync(ch chan string) {
	//simulated async function
	val := "Brian"
	time.Sleep(15 * time.Millisecond)
	ch <- val
}

func handlerWithContext() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	ch := make(chan string)
	go danglingSenderAsync(ch)

	select {
	case val := <-ch:
		{
			fmt.Printf("Value Received :%s", val)
		}
	case <-ctx.Done():
		{
			return errors.New("Timeout! Exiting")
		}
	}
	return nil
}
