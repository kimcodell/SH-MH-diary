package utils

import "fmt"

type ErrorParams struct {
	Err     error
	Message string ""
}

func CatchError(err ErrorParams) {
	if err.Err != nil {
		fmt.Println(fmt.Errorf("err : %v\nmessage: %q", err.Err, err.Message))
		panic(err)
	}
}
