package main

func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	res := []int{}
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i == m {
		for ; j < n; j++ {
			res = append(res, nums2[j])
		}
	}
	if j == n {
		for ; i < m; i++ {
			res = append(res, nums1[i])
		}
	}
	copy(nums1, res)
}
