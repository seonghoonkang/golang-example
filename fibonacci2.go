package main

import "fmt"

func fibonacci() func() int {
    fn1, fn0 := 1, 0
    return func() int {
        f := fn1 + fn0
        fn1, fn0 = f,  fn1
        return f
    }      
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
