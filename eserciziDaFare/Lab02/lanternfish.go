package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nDAYS = 80

func main() {
	var temp int
	pesci := map[int]int{8: 0, 7: 0, 6: 0, 5: 0, 4: 0, 3: 0, 2: 0, 1: 0, 0: 0}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text() //Gestione degli input
		inputSeparated := strings.Split(line, ",")
		for _, v := range inputSeparated {
			temp, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Input errato")
				return
			}
			pesci[temp]++
		}
	}

	for i := 0; i < nDAYS; i++ { //esegue le operazioni successive tante quanti sono i giorni
		temp = pesci[0]
		for j := 8; j >= 0; j-- { //riduce il timer dei pesci
			pesci[j], temp = temp, pesci[j]
		}
		pesci[6] += temp
	}
	count := 0
	for i := 0; i < 9; i++ {
		count += pesci[i]
	}
	fmt.Println(count)

}
