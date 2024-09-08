package main

import (
	"fmt"
)

func solution(nums []int, target int) []int {
    type aInfo struct {
        idx int
    }
    bDic := make(map[int]*aInfo)
    
    ans := make([]int, 2)
   
    for idx, num := range nums {
        if bDic[num] != nil {
            ans[0] = bDic[num].idx
            ans[1] = idx
            break
        } else {
            bDic[target-num] = &aInfo{idx: idx}
        }

    }
    return ans
}

func main() {
    //var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
    fmt.Println(solution([]int{3,2,4}, 6))
//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
//     [[0, 4], [0, 1], [2, 3]] => 2
	
}
