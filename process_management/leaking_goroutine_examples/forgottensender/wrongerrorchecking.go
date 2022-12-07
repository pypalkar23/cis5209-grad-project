package process_mgmt

import (
	"errors"
	"fmt"
)

func danglingSenderEC(ch chan string) {
	val := "Brian"
	ch <- val
}

func validateSomeData() error {
	//simulating a error throwing function
	return errors.New("Some error occured")
}
func HandlerWithErrorChecking() error {
	ch := make(chan string)
	go danglingSenderEC(ch)

	err := validateSomeData()

	if err != nil {
		return errors.New("Validation error.Exiting")
	}

	data := <-ch
	fmt.Println(data)
	return nil
}
