package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter digits: ")
	digits := readIntInput()

	fmt.Print("Enter palindromic: ")
	palindromic := readIntInput()

	largestPalindrome := getLargetPalindrome(digits, palindromic)

	fmt.Print("Resulted palindrome: ", largestPalindrome, "\n")
}

func readIntInput() (val int) {
	_, err := fmt.Scanf("%d", &val)

	if err != nil {
		panic(err)
	}

	return
}

func getLargetPalindrome(n int, k int) string {
	if k == 1 {
		return strings.Repeat("9", n)
	} else if k == 2 {
		if n <= 2 {
			return strings.Repeat("8", n)
		}
		return "8" + strings.Repeat("9", n-2) + "8"
	} else if k == 3 {
		return strings.Repeat("9", n)
	} else if k == 4 {
		if n <= 4 {
			return strings.Repeat("8", n)
		}
		return "88" + strings.Repeat("9", n-4) + "88"
	} else if k == 5 {
		if n <= 2 {
			return strings.Repeat("5", n)
		}
		return "5" + strings.Repeat("9", n-2) + "5"
	} else if k == 6 {
		if n <= 2 {
			return strings.Repeat("6", n)
		} else if n == 3 {
			return "888"
		} else if n == 4 {
			return "8778"
		}
		if n%2 == 1 {
			return "8" + strings.Repeat("9", (n/2)-1) + "8" + strings.Repeat("9", (n/2)-1) + "8"
		} else {
			return "8" + strings.Repeat("9", (n/2-2)) + "77" + strings.Repeat("9", (n/2)-2) + "8"
		}
	} else if k == 7 {
		dic := []string{
			"", "7", "77", "959", "9779", "99799", "999999", "9994999",
			"99944999", "999969999", "9999449999", "99999499999",
		}
		l := n / 12
		r := n % 12

		nines := strings.Builder{}
		for l > 0 {
			nines.WriteString("999999")
			l--
		}

		return nines.String() + dic[r] + nines.String()
	} else if k == 8 {
		if n <= 6 {
			return strings.Repeat("8", n)
		}
		return "888" + strings.Repeat("9", n-6) + "888"
	} else if k == 9 {
		return strings.Repeat("9", n)
	}
	return ""
}
