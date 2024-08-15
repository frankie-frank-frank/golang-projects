package main

import "fmt"

func findSuperimpositionIndexes(imageArr []string, patternArr []string) ([]int) {
	var initialRowIdx int
	var initialColIdx int
	var currPatternRowCount int = 0
	// var expectedSum int = len(patternArr) * len(patternArr[0])
	var activeSum int = 0
	
	if len(patternArr) == 0 { return []int{-1, -1} }
	
	for i := 0; i < len(imageArr); i++ { //col
		var currPatternColCount int = 0
		for j := 0; j < len(imageArr[i]); j++ { //row
			if imageArr[i][j] == patternArr[currPatternRowCount][currPatternColCount] {
				if currPatternRowCount == 0 && currPatternColCount == 0 {
					initialRowIdx = i
					initialColIdx = j
				}
				currPatternColCount++
				activeSum++
			} 
		} 
		//increment pattern row manually after each function if patternArr's length - 1 is > currPatternRow
		if len(patternArr[0]) == currPatternColCount {
			currPatternRowCount ++
		}
	}
	
	return []int{initialRowIdx, initialColIdx}
}

func main() {
	// // Example 1: 3*3 by 1*2
	// imgArr1 := []string{
	// 	"012",
	// 	"340",
	// 	"516",
	// }

	// patternArr1 := []string{
	// 	"16",
	// }

	// startingX1, startingY1 := findSuperimpositionIndexes(imgArr1, patternArr1)
	// fmt.Println(startingX1, startingY1)

	// // *************************************
	// // Example 2: 3*3 matched by 2*2
	// imgArr2 := []string{
	// 	"012",
	// 	"340",
	// 	"516",
	// }

	// patternArr2 := []string{
	// 	"40",
	// 	"16",
	// }

	// response2 := findSuperimpositionIndexes(imgArr2, patternArr2)
	// fmt.Println(response2)

	// *************************************
	// null check
	//Example 3: 3 * 3 matched by null
	// imgArr3 := []string{
	// 	"012",
	// 	"340",
	// 	"516",
	// }
	// patternArr3 := []string{}

	// response3 := findSuperimpositionIndexes(imgArr3, patternArr3)
	// fmt.Println(response3)
}