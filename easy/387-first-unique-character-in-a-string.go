package easy

/*
Given a string s, find the first non-repeating character
in it and return its index. if it does not exist, return -1

Example 1:
  Input: s = leetcode
  Output: 0

Example 2:
  Input: s = "loveleetcode"
  Output: 2

Example 3:
  Input: s = "aabb"
  Output: -1


Constraits:
-  1<=s.Length<=10^5
-  s consists of only lowercase English letters
*/

func firstUniqueChar(s string) int {
	charCount := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if charCount[s[i]] == 1 {
			return i
		}
	}

	return -1
}
