package main

import (
	"fmt"
)

func solution(height []int) int {
    start := 0
    end := len(height)-1
    max := 0
    for {
        if start >= end {
            break
        }

        startH := height[start]
        endH := height[end]
        x := end - start
        y := 0
        if startH > endH {
            y = endH
            end--
        } else {
            y = startH
            start++
        }
        if max < x * y {
            max = x*y
        }
    }
    return max
}

func main() {
    //var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
    fmt.Println(solution([]int{1,8,6,2,5,4,8,3,7}))
//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
//     [[0, 4], [0, 1], [2, 3]] => 2
	
}
