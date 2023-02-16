package main

type TreeNode struct {
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
	Val   int       `json:"val"`
}

type ListNode struct {
	Val  int `json:"val"`
	Next *ListNode
}
