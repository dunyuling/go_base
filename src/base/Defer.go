package main

import (
    "fmt"
)

func main() {
    f := func(i int) int {
        a := fibonacci(i)
        fmt.Printf("%d ",a)
        return a
    }
    for i := 0; i < 10; i++ {
	defer fmt.Printf("%d ", f(i))
    }
}

func fibonacci(num int) int {
    if num == 0 {
	return 0
    }
    if num < 2 {
	return 1
    }
    return fibonacci(num-1) + fibonacci(num-2)
}
