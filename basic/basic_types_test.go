package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	var aInt int
	of := reflect.TypeOf(aInt)
	fmt.Println(of)
}
