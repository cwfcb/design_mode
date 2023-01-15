package template

import "fmt"

type methodToCallback func()

type Invoker struct {
}

func (i *Invoker) Process(callBack methodToCallback) {
	fmt.Println("before call back")
	callBack()
	fmt.Println("after call back")
}
