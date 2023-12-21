package medium

import "container/heap"

/*
Given an integer array 'nums' and an integer 'k', return the
kth largest element in the array.

Note that it is the 'kth' largest element in the sorted order,
not the distinct element.const

Can you solve it without sorting?

Example 1:
   Input: nums = [3, 2, 1, 5, 6, 4]; k = 2
   Output: 5

Example 2:
   Input: nums = [3, 2, 3, 1, 2, 4, 5, 5, 6]; k = 4
   Output: 4

Constraints:
-   1 <= k <= nums.Length <= 10^5
-   -10^4 <= nums[i] <= 10^4
*/
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKthLargest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

//second solution
/*
func FindKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func partition(nums []int, left, right, pivotIndex int) int {
	pivotValue := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	storeIndex := left

	for i := left; i < right; i++ {
		if nums[i] < pivotValue {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}

	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}

func quickSelect(nums []int, left, right, kthSmallest int) int {
	if left == right {
		return nums[left]
	}

	pivotIndex := left + (right-left)/2
	pivotIndex = partition(nums, left, right, pivotIndex)

	if kthSmallest == pivotIndex {
		return nums[kthSmallest]
	} else if kthSmallest < pivotIndex {
		return quickSelect(nums, left, pivotIndex-1, kthSmallest)
	} else {
		return quickSelect(nums, pivotIndex+1, right, kthSmallest)
	}
}

*/
