package main

import "fmt"

func fibonacci() func() (int64) {
    var fn0 , fn1 int64 = 0, 1
    return func() int64 {
        f :=  fn0
        fn1, fn0 = fn1 + fn0,  fn1
        return f
    }      
}

func main() {
    f := fibonacci()
    for i := 0; i < 50; i++ {
        fmt.Println(f())
    }
}
