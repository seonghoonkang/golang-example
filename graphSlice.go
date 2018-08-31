package main

import (
	"code.google.com/p/go-tour/pic"
    "fmt"
)

func Pic(dx, dy int) [][]uint8 {
    result := make([][]uint8, dy)
        
    for i := range result {
        result[i] = make([]uint8, dx)
        
        for j := range result[i] {
            result[i][j] = uint8(i*j)
        }
    }
    return result
}

func printSlice(s string, x [][]uint8) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}

func main() {
    pic.Show(Pic)
    //printSlice("Pic", Pic(256, 256))
}
