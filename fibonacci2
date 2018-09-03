package main

import "fmt"

func fibonacci() func() int {
    fn0 := 0
    fn1 := 0
    return func() int {
        fn2 := fn1 + fn0
        fn1, fn0 = fn2, fn1
        if fn2 == 0 {
            fn1 = 1
    	}
        return fn2
    }      
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
