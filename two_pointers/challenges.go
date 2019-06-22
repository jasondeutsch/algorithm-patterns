package two_pointers

import "fmt"

/*
Problem: Target Sum Pair
*/
func TargetSumPair(nums []int, target int) []int {
	return []int{}
}

func TargetSumPairs(nums []int, target int) [][]int {
	var pair [][]int
	return pair
}

// not two pointers
// alternative solution worth mentioning
func TargetSumPairHash(nums []int, target int) []int {
	sums := make(map[int]int) // num: index
	for i, n := range nums {
		if idx, ok := sums[target-n]; ok {
			return []int{idx, i}
		} else {
			sums[n] = i
		}
	}
	return []int{}
}

/*
Problem: Triplet Sum
*/

// TODO check for improvements
func TargetSumTriplets(nums []int, target int) [][]int {
	triplets := [][]int{}
	numsMap := make(map[int]int)
	left, right := 0, len(nums)-1
	for left < right {
		PrintTwoPointers(left, right, len(nums))
		cur := nums[left] + nums[right]
		if idx, ok := numsMap[target-cur]; ok {
			triplets = append(triplets, []int{nums[left], nums[right], nums[idx]})
			numsMap[nums[right]] = right
			numsMap[nums[left]] = left
			right--
			left++
		} else if target-cur > right {
			numsMap[nums[right]] = right
			right--
		} else {
			numsMap[nums[left]] = left
			left++
		}
	}
	return triplets
}

/*
package helpers
*/
func PrintTwoPointers(a, b, length int) {
	arr := "\n"
	for i := 0; i < length; i++ {
		if i == a || i == b {
			arr += "✅"
		} else {
			arr += "⭕"
		}
	}
	fmt.Println(arr)
}
