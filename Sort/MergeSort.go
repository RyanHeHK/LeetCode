package Sort

func mergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mid := len(nums) / 2
	left := make([]int, mid)
	right := make([]int, len(nums)-mid)

	copy(left, nums[:mid])
	copy(right, nums[mid:])

	mergeSort(left)
	mergeSort(right)

	merge(nums, left, right)
}

func merge(nums, left, right []int) {
	i, j, k := 0, 0, 0
	leftLen, rightLen := len(left), len(right)

	for i < leftLen && j < rightLen {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}

	for i < leftLen {
		nums[k] = left[i]
		i++
		k++
	}

	for j < rightLen {
		nums[k] = right[j]
		j++
		k++
	}
}
