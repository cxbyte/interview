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
	var finalPalindromeLength = 0
	var finalPalindrome = ``
	var lastStringIndex = len(s) - 1

	if lastStringIndex == 0 {
		return s
	}

	for i := 0; i < lastStringIndex; i++ {
		palindrome := checkPalindrome(s, i, lastStringIndex)
		palindromeLength := len(palindrome)
		if palindromeLength > finalPalindromeLength {
			finalPalindromeLength = palindromeLength
			finalPalindrome = palindrome
		}
	}

	return finalPalindrome
}

func checkPalindrome(str string, currentIndex int, lastStringIndex int) string {
	currentChar := str[currentIndex : currentIndex+1]
	palindrome := currentChar
	leftIndex := currentIndex - 1
	rightIndex := currentIndex + 1

	for rightIndex <= lastStringIndex {
		var nextChar = str[rightIndex : rightIndex+1]
		if currentChar == nextChar {
			palindrome = palindrome + nextChar
			rightIndex++
		} else {
			break
		}
	}

	for leftIndex >= 0 && rightIndex <= lastStringIndex {
		leftChar := str[leftIndex : leftIndex+1]
		rightChar := str[rightIndex : rightIndex+1]
		if leftChar == rightChar {
			palindrome = str[leftIndex : rightIndex+1]
		} else {
			break
		}

		leftIndex--
		rightIndex++
	}

	return palindrome
}
