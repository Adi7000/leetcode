package leetcode

func candy(ratings []int) int {
	if len(ratings) <= 0 {
		return 0
	}
	// Initialize with index 0
	total := 1
	localMaximaIndex := 0
	lastRating := ratings[0]
	lastCandy := 1
	gapIndicies := []int{}
	gapVals := []int{}

	for i := 1; i < len(ratings); i++ {
		currentRating := ratings[i]
		if currentRating > lastRating {
			lastCandy++
			total += lastCandy
			localMaximaIndex = i
			gapIndicies = []int{}
			gapVals = []int{}
		} else if currentRating == lastRating {
			lastCandy = 1
			total += lastCandy
			localMaximaIndex = i
			gapIndicies = []int{}
			gapVals = []int{}
		} else { //currentRating < lastRating
			if lastCandy > 1 {
				if lastCandy > 2 {
					gapIndicies = append(gapIndicies, i)
					gapVals = append(gapVals, lastCandy-2)
				}
				lastCandy = 1
				total++
			} else {
				numGaps := len(gapVals)
				if numGaps == 0 {
					//Award each of the previous children upto local maxima 1 more candy
					total += i - localMaximaIndex + 1
				} else {
					//Award each of the previous children upto last gap 1 more candy
					total += i - gapIndicies[numGaps-1] + 1
					if gapVals[numGaps-1] > 1 {
						gapVals[numGaps-1]--
					} else {
						gapVals = gapVals[:numGaps-1]
						gapIndicies = gapIndicies[:numGaps-1]
					}
				}
			}
		}
		lastRating = currentRating
	}
	return total
}
