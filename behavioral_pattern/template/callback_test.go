package template

import (
	"fmt"
	"testing"
)

func TestInvoker(t *testing.T) {
	invoker := &Invoker{}
	invoker.Process(func() {
		fmt.Println("call back todo")
	})

	invoker.Process(func() {
		fmt.Println("now call back")
	})
}
