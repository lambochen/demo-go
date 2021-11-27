package main

import "fmt"

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	//fmt.Println(math.Abs(-1))
	//
	//q, r := div(13, 3)
	//fmt.Println(q, r)
	//
	//s := make([]int, 10)
	//fmt.Println(s)

	a := adder(0)

	for i := 0; i < 10; i++ {
		var sum int
		sum, a = a(i)

		fmt.Println(sum)
	}

}

type iAdder func(int) (int, iAdder)

func adder(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder(base + v)
	}
}
