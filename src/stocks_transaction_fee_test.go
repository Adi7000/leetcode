package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit1(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	actual := maxProfit(prices, fee)
	assert.Equal(t, 8, actual)
}
func TestMaxProfit2(t *testing.T) {
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	actual := maxProfit(prices, fee)
	assert.Equal(t, 6, actual)
}
func TestMaxProfitNoData(t *testing.T) {
	prices := []int{}
	fee := 3
	actual := maxProfit(prices, fee)
	assert.Equal(t, 0, actual)
}
func TestMaxProfitOneDayData(t *testing.T) {
	prices := []int{1}
	fee := 4
	actual := maxProfit(prices, fee)
	assert.Equal(t, 0, actual)
}
func TestMaxProfitNoProfit(t *testing.T) {
	prices := []int{1, 5, 0, 4}
	fee := 4
	actual := maxProfit(prices, fee)
	assert.Equal(t, 0, actual)
}
