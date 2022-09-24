package decorator

import "fmt"

type Func2 func(int) int

func LogDecorate(fn Func2) Func2 {
	return func(i int) int {
		fmt.Println("start func i is", i)
		r := fn(i)
		fmt.Println("end func result is", r)
		return r
	}
}
