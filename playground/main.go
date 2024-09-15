package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int, target int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var ans [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
        }
		twoSumTarget, left, right := target-nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == twoSumTarget {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > twoSumTarget {
				right--
			} else if sum < twoSumTarget {
				left++
			}
		}   
    }
    return ans
}

func solution(nums []int, target int) [][]int {
    sort.Slice(nums, func(i int, j int) bool {
        return nums[i] < nums[j]
    })
    var ans [][]int
    for i:=0; i < len(nums)-3; i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        threeSums := threeSum(nums[i+1:], target-nums[i])
        for _, set := range threeSums {
            fourSum := []int{nums[i]}
            fourSum = append(fourSum, set...)
            ans = append(ans,fourSum)
        }
    }
    return ans
}


func main() {
	//var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
	fmt.Println(solution([]int{1,-2,-5,-4,-3,3,3,5}, -11))
	//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
	//     [[0, 4], [0, 1], [2, 3]] => 2

}
