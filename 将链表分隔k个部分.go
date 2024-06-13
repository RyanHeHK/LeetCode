package main

func splitListToParts(head *ListNode, k int) []*ListNode {
	// Step 1: Get the length of the linked list
	len := 0
	curr := head
	for curr != nil {
		len++
		curr = curr.Next
	}

	// Step 2: Calculate base length and extra parts
	baseLength := len / k
	extraParts := len % k

	// Step 3: Initialize the result array
	result := make([]*ListNode, k)

	// Step 4: Initialize the current node
	curr = head

	// Step 5: Split the linked list into parts
	for i := 0; i < k; i++ {
		// Set the head node of the current part
		result[i] = curr

		// Calculate the length of the current part
		partLength := baseLength
		if i < extraParts {
			partLength++
		}

		// Move curr to the end of the current part
		for j := 0; j < partLength-1 && curr != nil; j++ {
			curr = curr.Next
		}

		// Set the next node of the current part's tail
		if curr != nil {
			nextNode := curr.Next
			curr.Next = nil
			curr = nextNode
		}
	}

	// Step 6: Return the result array
	return result
}
