package tome

import (
	"fmt"
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	table := []struct {
		left int64
		right int64
		expected int64
	} {
		{0, 1, 0},
		{-1, 1, -1},
		{math.MaxInt64, math.MinInt64, math.MinInt64},
	}

	for _, tt := range table {
		testname := fmt.Sprintf("Testing Min(%d, %d)", tt.left, tt.right)
		t.Run(testname, func(t *testing.T) {
			ans := Min(tt.left, tt.right)
			if ans != tt.expected {
				t.Errorf("got %d, wanted %d", ans, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	table := []struct {
		left int64
		right int64
		expected int64
	} {
		{0, 1, 1},
		{-1, 1, 1},
		{math.MaxInt64, math.MinInt64, math.MaxInt64},
	}

	for _, tt := range table {
		testname := fmt.Sprintf("Testing Min(%d, %d)", tt.left, tt.right)
		t.Run(testname, func(t *testing.T) {
			ans := Max(tt.left, tt.right)
			if ans != tt.expected {
				t.Errorf("got %d, wanted %d", ans, tt.expected)
			}
		})
	}
}
