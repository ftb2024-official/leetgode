/*
Write script to calculate the frequency of each word in a text file `words.txt`. For simplicity sake you may assume: 1) `words.txt` contains only
lowercase characters and space characters; 2) Each word must consist of lowercase characters only; 3) Words are separated by one or more whitespace
characters.

Example:
	Input: the day is sunny the the\nthe sunny is is
	Output:
		the 4
		is 3
		sunny 2
		day 1
*/

package arraysandhashing

import (
	"bufio"
	"fmt"
	"os"
)

func CountWordFrequency() map[string]int {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return nil
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := map[string]int{}

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		ans[scanner.Text()]++
	}

	return ans
}
