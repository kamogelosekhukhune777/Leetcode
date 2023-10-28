package easy

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:
1.Open brackets must be closed by the same type of brackets.
2.Open brackets must be closed in the correct order.
3.Every close bracket has a corresponding open bracket of the same type.
-------------------------------------------------
Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
------------------------------------------------
Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func IsValid(s string) bool {
	array := []rune{}
	bucketMap := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, char := range s {
		if char == '{' || char == '(' || char == '[' {
			array = append(array, char)
		} else {
			if len(array) == 0 || array[len(array)-1] != bucketMap[char] {
				return false
			}
			array = array[:len(array)-1]
		}
	}
	return len(array) == 0
}
