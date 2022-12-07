package process_mgmt

func danglingSender(ch chan string) {
	val := "Brian"
	ch <- val
}

func handler() {
	ch := make(chan string)

	go danglingSender(ch)
	return
}
