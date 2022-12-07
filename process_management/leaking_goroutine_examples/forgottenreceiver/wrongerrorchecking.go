package process_mgmt

import (
	"errors"
	"fmt"
)

func danglingReceiverEC(ch chan []string) {
	val := <-ch
	fmt.Println(val)
}
func someDataReturningFunc() ([]string, error) {
	return nil, errors.New("Some error occured")
}

func handlerwithErrorCheck() error {
	ch := make(chan []string)
	go danglingReceiverEC(ch)

	data, err := someDataReturningFunc()
	if err != nil {
		return errors.New("Fatal Error, Exiting..!")
	}

	ch <- data

	return nil

}
