package utils

import (
	"fmt"
	"log"
)

type ErrorParams struct {
	Err     error
	Message string ""
}

func CatchError(err ErrorParams) {
	if err.Err != nil {
		log.Fatal(fmt.Errorf("err : %v\nmessage: %q", err.Err, err.Message))
	}
}
