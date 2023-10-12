package main

import (
	"fmt"
	"pkg/convert"
	"pkg/fileparser"
	"pkg/sets"
	"pkg/slices"
	"strconv"
	"strings"
)

func main() {
	data := fileparser.ReadLines("day13/input.txt")
	dots, folds := ParseInstructionsLines(data)

	count := 0
	for i, fold := range folds {
		count++
		fold.Apply(dots)
		if count == 1 {
			fmt.Printf("[Part 1] Dots after %s (fold %d): %d\n", folds[i].Name, count, len(dots))
		}
	}
	fmt.Printf("[Part 2] Final paper image\n\n")
	PrintDots(dots)
}

func ParseInstructionsLines(data []string) (sets.Set[Coord], []Fold) {
	isDotsLine := func(t string) bool { return !strings.HasPrefix(t, "fold") }
	dotsData, foldsData := slices.Divide(data, isDotsLine)
	dotsData = slices.TrimEnd(dotsData, 1) // Get rid of the last blank line
	dotsSet := sets.NewSetFromSlice(slices.Map(dotsData, NewCoord))
	folds := slices.Map(foldsData, NewFold)
	return dotsSet, folds
}

type Coord struct{ X, Y int }

func NewCoord(line string) Coord {
	parts := strings.Split(line, ",")
	return Coord{
		X: convert.Apply[int](parts[0]),
		Y: convert.Apply[int](parts[1]),
	}
}

type Reflector func(pos Coord) Coord

type Fold struct {
	Name    string
	axis    string
	lineVal int
	reflect Reflector
}

// ReflectXFunc generates a function that will reflect a coord around the defined x line
func ReflectXFunc(xr int) Reflector {
	return func(pos Coord) Coord { return Coord{X: 2*xr - pos.X, Y: pos.Y} }
}

// ReflectYFunc generates a function that will reflect a coord around the defined y line
func ReflectYFunc(yr int) Reflector {
	return func(pos Coord) Coord { return Coord{X: pos.X, Y: 2*yr - pos.Y} }
}

func NewFold(line string) Fold {
	foldStr := strings.Split(line, " ")[2]
	refLine := strings.Split(foldStr, "=")
	refAxis := refLine[0]
	lineNumber, _ := strconv.Atoi(refLine[1])
	var reflect Reflector
	switch refAxis {
	case "x":
		reflect = ReflectXFunc(lineNumber)
	case "y":
		reflect = ReflectYFunc(lineNumber)
	default:
		panic("unrecognised reflect line")
	}
	return Fold{Name: foldStr, axis: refAxis, lineVal: lineNumber, reflect: reflect}
}

func (f Fold) ShouldReflect(coord Coord) bool {
	// Only reflect coords that are above or the left of the reflect line
	switch f.axis {
	case "x":
		return coord.X > f.lineVal
	case "y":
		return coord.Y > f.lineVal
	default:
		panic("unrecognised reflect line")
	}
}

func (f Fold) Apply(dots sets.Set[Coord]) {
	for _, coord := range dots.ToSlice() {
		if f.ShouldReflect(coord) {
			dots.Remove(coord)
			dots.Add(f.reflect(coord))
		}
	}
}

func PrintDots(dots sets.Set[Coord]) {
	xSelectFunc := func(c Coord) int { return c.X }
	ySelectFunc := func(c Coord) int { return c.Y }
	maxX := slices.Max(slices.Map(dots.ToSlice(), xSelectFunc))
	maxY := slices.Max(slices.Map(dots.ToSlice(), ySelectFunc))
	for j := 0; j <= maxY; j++ {
		for i := 0; i <= maxX; i++ {
			if dots.IsMember(Coord{i, j}) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
///*
//Instructions on https://adventofcode.com/2021/day/13
//*/
//
//func riempiColonna(slice [][]string, maxYInput int, maxXInput int) [][]string {
//	for i := 0; i < maxXInput; i++ {
//		if len(slice[i]) != maxYInput {
//			difference := maxYInput - len(slice[i])
//			for j := 0; j < difference; j++ {
//				slice[i] = append(slice[i], ".")
//			}
//		}
//	}
//	return slice
//}
//
//func main() {
//	mapDots := make([][]string, 0)
//	scanner := bufio.NewScanner(os.Stdin)
//	maxXInput := -1
//	maxYInput := -1
//	for scanner.Scan() {
//		line := scanner.Text()
//		if strings.Contains(line, "fold") {
//
//		} else {
//			lineSplitted := strings.Split(line, ",")
//			xValue, err := strconv.Atoi(lineSplitted[0])
//			yValue, err2 := strconv.Atoi(lineSplitted[1])
//			if err != nil || err2 != nil {
//				fmt.Println("Error during input data, closing...")
//				return
//			}
//			if xValue > maxXInput {
//				//fmt.Println(xValue, maxXInput, xValue > maxXInput)
//				for i := 0; i < xValue-maxXInput; i++ {
//					mapDots = append(mapDots, []string{})
//					//fmt.Println(mapDots)
//				}
//				maxXInput = xValue
//				mapDots = riempiColonna(mapDots, maxYInput, maxXInput)
//			}
//			if yValue > maxYInput {
//				//fmt.Println(yValue, maxYInput, yValue > maxYInput)
//				for i := 0; i <= maxXInput; i++ {
//					for j := 0; j < yValue-maxYInput; j++ {
//						mapDots[i] = append(mapDots[i], ".")
//						//fmt.Println(mapDots)
//					}
//				}
//				maxYInput = yValue
//				mapDots = riempiColonna(mapDots, maxYInput, maxXInput)
//			}
//		}
//		fmt.Println(mapDots)
//	}
//}

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
