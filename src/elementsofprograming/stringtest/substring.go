package stringtest

import "fmt"

func Test() {
	value := "Welcome, my friend"
	// Take substring of first word with runes.
	// ... This handles any kind of rune in the string.
	// ... Convert back into a string from rune slice.
	safeSubstring := value[0:]
	fmt.Println(" RUNE SUBSTRING:", safeSubstring)
}
