package slice

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	var txErr error

	defer func() {

		if txErr != nil {
			fmt.Println("txErr:", txErr)
		} else {
			fmt.Println("txErr: nil")

		}
	}()

	operation, txErr := dbOperation()

	if txErr != nil {

		fmt.Println(operation)
		return
	}

}

func dbOperation() (int, error) {

	return 0, errors.New("err happens")
}
