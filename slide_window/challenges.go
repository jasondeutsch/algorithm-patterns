package slide_window

import "fmt"

func sumSubBrute(k int, arr []int) []int {
	var result []int
	for i := 0; i <= len(arr)-k; i++ {
		sum := 0
		for _, v := range arr[i:i+k] {
			sum += v
		}
		result = append(result, sum)
	}
	return result
}

func sumSub(k int, arr []int) []int {
	var result []int
	start, sum := 0, 0
	for end := 0; end < len(arr); end++ {
		sum += arr[end]
		if end >= k-1 {
			result = append(result, sum)
			sum -= arr[start]
			start++
		}
	}
	return result
}


func longestSubStringKDistinct(k int, str string) string {
	frequency := make(map[rune]int)
	max, distinct, start := "", 0, 0
	for i, c := range str {
		// frequency accounting
		frequency[c]++
		if frequency[c] == 1 {
			distinct++
		}
		// we have k distinct chars and we are at the end of the string or the next char is a 3rd distinct
		isSubStr := distinct == k && (i == len(str)-1 || frequency[rune(str[i+1])] == 0)
		//  handle substring
		//printWindow(start, i+1, len(str))
		if isSubStr {
			cur := str[start:i+1]
			if len(cur) > len(max) {
				max = cur
			}
			// move the  window start
			start += frequency[rune(cur[0])]
			// reset frequency and adjust distinct
			frequency[rune(cur[0])] = 0
			distinct--
		}
	}
	return string(max)
}

func longestUniqueSubString(str string) string {
	indexes := make(map[rune]int)
	longest,  start := "", 0

	for i, c := range str {
		if  _, ok := indexes[c]; ok {
			start = indexes[c]+1
		}

		indexes[c] = i

		current := str[start:i+1]
		if len(current) > len(longest) {
			longest = current
		}
	}
	return longest
}

func printWindow(start,end, length int) {
	window := "\n"
	for i := 0; i < length; i++ {
		in := i >= start && i <= end
		if in {
			window += "✅"
		} else {
			window += "⭕"
		}
	}
	fmt.Println(window)
}
