package main

import "fmt"
const a = iota
const (
//    b = iota
    b = 1
    c = 3.14
    _
    d = iota
)

const (
    e = iota*2
    f 
    g 
)

const (
    h,i = iota,iota + 3
    j,k
    l = iota
)

func main() {
    //fmt.Print(iota)
    fmt.Print("a:")
    fmt.Println(a)
    fmt.Print("b:")
    fmt.Println(b)
    fmt.Print("c:")
    fmt.Println(c)
    fmt.Print("d:")
    fmt.Println(d)
    fmt.Print("e:")
    fmt.Println(e)
    fmt.Print("f:")
    fmt.Println(f)
    fmt.Print("g:")
    fmt.Println(g)
    fmt.Print("h:")
    fmt.Println(h)
    fmt.Print("i:")
    fmt.Println(i)
    fmt.Print("j:")
    fmt.Println(j)
    fmt.Print("k:")
    fmt.Println(k)
    fmt.Print("l:")
    fmt.Println(l)

}

