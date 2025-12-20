package main

type NumArray struct {
	prefix []int //Prefix Sums Array
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	// Formula: p[right+1] - p[left] == sum(left:right)
	return this.prefix[right+1] - this.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
