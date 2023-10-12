package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Instructions on https://adventofcode.com/2021/day/13
*/

func riempiColonna(slice [][]string, maxYInput int, maxXInput int) [][]string {
	for i := 0; i < maxXInput; i++ {
		if len(slice[i]) != maxYInput {
			difference := maxYInput - len(slice[i])
			for j := 0; j < difference; j++ {
				slice[i] = append(slice[i], ".")
			}
		}
	}
	return slice
}

func main() {
	mapDots := make([][]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	maxXInput := -1
	maxYInput := -1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "fold") {

		} else {
			lineSplitted := strings.Split(line, ",")
			xValue, err := strconv.Atoi(lineSplitted[0])
			yValue, err2 := strconv.Atoi(lineSplitted[1])
			if err != nil || err2 != nil {
				fmt.Println("Error during input data, closing...")
				return
			}
			if xValue > maxXInput {
				//fmt.Println(xValue, maxXInput, xValue > maxXInput)
				for i := 0; i < xValue-maxXInput; i++ {
					mapDots = append(mapDots, []string{})
					//fmt.Println(mapDots)
				}
				maxXInput = xValue
				mapDots = riempiColonna(mapDots, maxYInput, maxXInput)
			}
			if yValue > maxYInput {
				//fmt.Println(yValue, maxYInput, yValue > maxYInput)
				for i := 0; i <= maxXInput; i++ {
					for j := 0; j < yValue-maxYInput; j++ {
						mapDots[i] = append(mapDots[i], ".")
						//fmt.Println(mapDots)
					}
				}
				maxYInput = yValue
				mapDots = riempiColonna(mapDots, maxYInput, maxXInput)
			}
		}
		fmt.Println(mapDots)
	}
}

/*
if xValue > maxXInput {
	fmt.Println(xValue, maxXInput, xValue > maxXInput)
	for i := 0; i < xValue-maxXInput; i++ {
		mapDots = append(mapDots, []string{})
		fmt.Println(mapDots)
	}
	maxXInput = xValue
}
fmt.Println()
if yValue > maxYInput {
	fmt.Println(yValue, maxYInput, yValue > maxYInput)
	for i := 0; i <= maxXInput; i++ {
		for j := 0; j < yValue-maxYInput; j++ {
			mapDots[i] = append(mapDots[i], ".")
			fmt.Println(mapDots)
		}
	}
	maxYInput = yValue
}
*/
