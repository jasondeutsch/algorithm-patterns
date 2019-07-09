package two_pointers

import "fmt"

/*
Problem: Target Sum Pair
*/
func TargetSumPairSolution(nums []int, target int) []int {
	var pair []int
	i, j := 0, len(nums)-1
	for i != j {
		PrintTwoPointers(i, j, len(nums))
		sum := nums[i] + nums[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			pair = []int{i, j}
			break
		}
	}
	return pair
}

func TargetSumPairsSolution(nums []int, target int) [][]int {
	var pair [][]int
	i, j := 0, len(nums)-1
	for i != j {
		PrintTwoPointers(i, j, len(nums))
		sum := nums[i] + nums[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			fmt.Println("found it ^^^")
			pair = append(pair, []int{i, j})
			i++
		}
	}
	return pair
}

// not two pointers
func TargetSumPairHashSolution(nums []int, target int) []int {
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
func TargetSumTripletsSolution(nums []int, target int) [][]int {
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
