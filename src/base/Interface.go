package main

import "fmt"

type Animal interface {
    Grow()
	Move(string) string
}

type Cat struct {
    Name string
    Age  uint8
    Address string
}

func(myCat *Cat) Grow() {
    
}

func(myCat *Cat) Move(string) string {
    return ""
}

func main() {
    myCat := Cat{"Little C", 2, "In the house"}
    animal, ok := interface{}(&myCat).(Animal)
    fmt.Printf("%v, %v\n", ok, animal)
}
