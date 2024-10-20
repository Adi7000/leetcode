package leetcode

func canCompleteCircuit(gas []int, cost []int) int {

	//Update gas slice with the effective amount of gas gained/lost at a station
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
	}
	//Try to start at the last station and test if circuit can be completed
	//if not, try to see if the second last station lets it be completed
	//and so on
	start := len(gas) - 1
	maxReachable := 0
	remainingGas := gas[start]
	for {
		if remainingGas >= 0 {
			if maxReachable == start {
				return start
			}
			remainingGas += gas[maxReachable]
			maxReachable++
		} else {
			if maxReachable == start {
				return -1
			}
			start--
			remainingGas += gas[start]
		}
	}
}
