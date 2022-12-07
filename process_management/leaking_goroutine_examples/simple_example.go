package process_mgmt

import (
	"fmt"
)

func leakingGoroutine(ch chan string) {
	val := <-ch
	fmt.Println(val)
}

func sample_handler() {
	ch := make(chan string)

	go leakingGoroutine(ch)
	return
}
