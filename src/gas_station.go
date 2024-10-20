package leetcode

func canCompleteCircuit(gas []int, cost []int) int {

	//Update gas slice with the effective amount of gas gained/lost at a station
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
	}
	if len(gas) == 1 && gas[0] >= 0 {
		return 0
	}
	//Solved by brute force
	for start := 0; start < len(gas)-1; start++ {
		remainingGas := gas[start]
		if remainingGas < 0 {
			continue
		}
		currentStation := start + 1
		for {
			remainingGas += gas[currentStation]
			if remainingGas < 0 {
				break
			}
			if currentStation != len(gas)-1 {
				currentStation++
			} else {
				currentStation = 0
			}
			if currentStation == start {
				return start
			}
		}
	}
	remainingGas := gas[len(gas)-1]
	currentStation := 0
	for {
		remainingGas += gas[currentStation]
		if remainingGas < 0 {
			break
		}
		currentStation++
		if currentStation == len(gas)-1 {
			return len(gas) - 1
		}
	}
	return -1
}
