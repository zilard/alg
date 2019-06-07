package main

import (
	"fmt"
)


type S struct {
	a int
}


func main() {

    s := &S{}

    fmt.Printf("s %v\n", s)

    s.a++

    fmt.Printf("s %v\n", s)


}
