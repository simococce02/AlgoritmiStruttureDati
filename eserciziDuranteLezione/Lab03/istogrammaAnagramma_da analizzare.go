package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

func quali(riga string) map[string]int {
	lettere := make(map[string]int)
	for _, char := range riga {
		if unicode.IsLetter(char) {
			lettere[string(char)]++
		}
	}
	return lettere
}

func anagrammi(s1, s2 string) bool {
	return reflect.DeepEqual(quali(s1), quali(s2))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	riga := strings.ToLower(scanner.Text())
	scanner.Scan()
	riga2 := strings.ToLower(scanner.Text())
	lettere := quali(riga)
	fmt.Println(anagrammi(riga, riga2))
	lettere2 := []string{}
	for key := range lettere {
		lettere2 = append(lettere2, key)
	}
	sort.Strings(lettere2)
	for _, lett := range lettere2 {
		fmt.Print(lett, " ")
		fmt.Println(strings.Repeat("*", lettere[lett]))
	}
}
