package process_mgmt

import "fmt"

func danglingReceiver(ch chan int) {
	val := <-ch
	fmt.Println(val)
}

func receiver_handler() {
	ch := make(chan int)
	go danglingReceiver(ch)
}
