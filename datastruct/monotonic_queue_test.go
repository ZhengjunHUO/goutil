package datastruct

import (
	"testing"
	"reflect"
)

// return a list of local max of a sliding window of size k, moving from left to right in slice nums
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	mq := NewMonotonicQueue[int]()
	rslt := make([]int, n-k+1)

	// not needed, added for test coverage's sake
	if mq.Size() != 0 {
		return rslt
	}

	for i := range nums{
		if i < k - 1 {
			mq.Push(nums[i])
			continue
		}

		mq.Push(nums[i])
		rslt[i-k+1] = mq.Max()
		mq.Pop(nums[i-k+1])
	}

	return rslt
}

func TestMonotonicQueue(t *testing.T) {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	want := []int{3,3,5,5,6,7}
	rslt := maxSlidingWindow(nums, k)

	if !reflect.DeepEqual(rslt, want) {
		t.Errorf("Test monotonic queue's result: %v, expecting: %v", rslt, want)
	}
}
