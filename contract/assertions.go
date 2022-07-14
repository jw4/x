package contract

import "fmt"

func AssertNotNil(params ...interface{}) {
	var indexes []int
	for ix, param := range params {
		if param == nil {
			indexes = append(indexes, ix)
		}
	}

	if len(indexes) > 0 {
		panic(fmt.Errorf("ASSERTION FAILED: parameters at index(es) %v are nil", indexes))
	}
}
