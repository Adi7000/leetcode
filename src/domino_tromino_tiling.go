package leetcode

import "math"

func numTilings(n int) int {
	for numPartitions := 1; numPartitions <= n; numPartitions++ {
		for numDominoSections := 0; numDominoSections <= int(math.Ceil(float64(numPartitions)/2.0)); numDominoSections++ {
			break
		}
	}
	return -1
}

func numTilingsDominos(n int, params ...*map[int]int) int {
	var nDominoConfigurationMap *map[int]int
	if len(params) == 0 {
		nDominoConfigurationMap = &map[int]int{1: 1, 2: 2}
	} else if len(params) == 1 {
		nDominoConfigurationMap = params[0]
	}

	numConfig, exists := (*nDominoConfigurationMap)[n]
	if exists {
		return numConfig
	} else {
		(*nDominoConfigurationMap)[n] = numTilingsDominos(n-1, nDominoConfigurationMap) + numTilingsDominos(n-2, nDominoConfigurationMap)
	}
	return (*nDominoConfigurationMap)[n]
}

// type extendablePattern {
//     var minimumLength int
// 	var extensionIncrement int
// }

// var twoTrominoConfigurationOne extendablePattern{minimumLength:3, extensionIncrement:1}
// var twoTrominoConfigurationTwo extendablePattern{minimumLength:3, extensionIncrement:1}
