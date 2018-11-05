package main

import (
    "fmt"
    "math/rand"
)

func main() {
    ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
    switch v := ia[rand.Intn(4) :4]; interface{}(v).(type) {
    case []interface{} :
	fmt.Println("Case A.")
	fmt.Println(v)
 	break;
    case byte :
	fmt.Println("Case B.")
 	break;
    default:
      	fmt.Println("Unknown!")
	break;
    }
}
