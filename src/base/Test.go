package main

import  (
    "fmt"
    "math/rand" 
    "unsafe"
)

func main() {
    fmt.Println("Hello, 世界")
    a := rand.Intn(3)
    fmt.Println(a)	
    var b = true
    fmt.Println()
    fmt.Println(unsafe.Sizeof(b))
    name := [5]byte{'i','m','o','o','c'}
    fmt.Println(name)
}  
