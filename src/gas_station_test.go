package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCompleteCircuitExample1(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	actual := canCompleteCircuit(gas, cost)
	assert.Equal(t, 3, actual)
}
func TestCanCompleteCircuitExample2(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	actual := canCompleteCircuit(gas, cost)
	assert.Equal(t, -1, actual)
}
func TestCanCompleteCircuitLastStationAnswer(t *testing.T) {
	gas := []int{1, 1, 9}
	cost := []int{3, 4, 3}
	actual := canCompleteCircuit(gas, cost)
	assert.Equal(t, 2, actual)
}
func TestCanCompleteCircuitOneReachableLocation(t *testing.T) {
	gas := []int{2}
	cost := []int{2}
	actual := canCompleteCircuit(gas, cost)
	assert.Equal(t, 0, actual)
}
func TestCanCompleteCircuitOneUnreachableLocation(t *testing.T) {
	gas := []int{1}
	cost := []int{2}
	actual := canCompleteCircuit(gas, cost)
	assert.Equal(t, -1, actual)
}
