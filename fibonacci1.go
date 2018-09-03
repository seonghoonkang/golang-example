package main

import "fmt"

func fibonacci() func(int) int {
    return func(n int) int{
        return fib(n);
    }
}

func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}

func main() {
    f := fibonacci()
    for i := 0; i <= 21; i++ {
        fmt.Println(f(i))
    }
}
