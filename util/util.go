package util

import (
	"fmt"
)

func IsEmpty(i interface{}) bool {
	return i == nil || (i != nil && fmt.Sprintf("%v", i) == "")
}
