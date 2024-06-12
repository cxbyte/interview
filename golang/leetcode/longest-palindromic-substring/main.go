package longest_palindromic_substring

/**
5. Longest Palindromic Substring
https://leetcode.com/problems/longest-palindromic-substring/description/

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	var stringLen = len(s)
	var firstIndex = 0
	var lastIndex = stringLen - 1
	var palindrome []string
	for {
		for {
			checkString := s[firstIndex : lastIndex+1]
			if checkPalindrome(checkString) {
				palindrome = append(palindrome, checkString)
			}

			if lastIndex-firstIndex < 1 {
				break
			}
			lastIndex--
		}
		firstIndex++
		lastIndex = stringLen - 1
		if lastIndex-firstIndex < 1 {
			break
		}
	}

	if len(palindrome) == 0 {
		return s[0:1]
	}

	var longestPalindrome = palindrome[0]
	for _, p := range palindrome {
		if len(p) > len(longestPalindrome) {
			longestPalindrome = p
		}
	}

	return longestPalindrome
}

func checkPalindrome(s string) bool {
	var stringLen = len(s)
	if stringLen == 1 {
		return true
	}
	var firstIndex = 0
	var lastIndex = stringLen - 1

	for {
		if s[firstIndex:firstIndex+1] != s[lastIndex:lastIndex+1] {
			return false
		}

		if lastIndex-firstIndex < 1 {
			return true
		}

		lastIndex--
		firstIndex++
	}
}
