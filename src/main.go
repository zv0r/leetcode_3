package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbbabcdefgh"), lengthOfLongestSubstring("abcabcbbabcdefgh") == 8)
	fmt.Println(lengthOfLongestSubstring("abcabcbb"), lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb"), lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew"), lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("dvdf"), lengthOfLongestSubstring("dvdf") == 3)
	fmt.Println(lengthOfLongestSubstring("au"), lengthOfLongestSubstring("au") == 2)
	fmt.Println(lengthOfLongestSubstring("a"), lengthOfLongestSubstring("a") == 1)
	fmt.Println(lengthOfLongestSubstring(""), lengthOfLongestSubstring("") == 0)
}

// func lengthOfLongestSubstring(s string) int {
// 	max_unique_length := 0
// 	if len(s) > 1 {
// 	check_unique_substr:
// 		for i := 0; i < len(s); i++ {
// 			unique_substr := make(map[byte]bool)
// 			curr_unique_length := 0
// 			for j := i; j < len(s); j++ {
// 				symbol := s[j]
// 				_, symbol_exists := unique_substr[symbol]
// 				if symbol_exists {
// 					break
// 				} else {
// 					unique_substr[symbol] = true
// 					curr_unique_length++
// 					if curr_unique_length > max_unique_length {
// 						max_unique_length = curr_unique_length
// 					}
// 					if j == len(s)-1 {
// 						break check_unique_substr
// 					}
// 				}
// 			}
// 		}
// 	} else {
// 		max_unique_length = len(s)
// 	}
// 	return max_unique_length
// }

func lengthOfLongestSubstring(s string) int {
	max_unique_length := 0
	lastOccurred := make(map[rune]int)
	start := 0

	for i, ch := range s {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > max_unique_length {
			max_unique_length = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return max_unique_length
}
