package process_mgmt

import "fmt"

func danglingReceiverUC(ch chan string) {
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("Done.. exiting")
}

func handlerWithUnclosedChannel(arr []string) {
	ch := make(chan string, len(arr))

	for _, data := range arr {
		ch <- data
	}

	go danglingReceiverUC(ch)

}
