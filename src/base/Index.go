package main

import "fmt"

func main() {
    var numbers3 = [5]int{1, 2, 3, 4, 5}
    slice3 := numbers3[2 : 3:4]
    capacity := cap(slice3)
    fmt.Printf("%d\n",capacity)
}
