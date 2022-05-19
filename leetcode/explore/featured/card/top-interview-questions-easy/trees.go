package top_interview_questions_easy

import "math"

// TreeNode represents binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth Problem: Maximum Depth of Binary Tree.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/94/trees/555/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return visit(root, 0)
}

// visit visit node and return current depth.
func visit(node *TreeNode, currentDepth int) int {

	if node == nil {
		return currentDepth
	}

	currentDepth++
	leftDepth := visit(node.Left, currentDepth)
	rightDepth := visit(node.Right, currentDepth)

	if leftDepth >= rightDepth {
		return leftDepth
	}
	return rightDepth
}

// isValidBST Problem: Validate Binary Search Tree.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/94/trees/625/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	left := recursiveCheck(root.Left, math.MinInt32-1, root.Val)
	right := recursiveCheck(root.Right, root.Val, math.MaxInt32+1)

	return left && right
}

func recursiveCheck(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if !(node.Val > min && node.Val < max) {
		return false
	}

	left := recursiveCheck(node.Left, min, node.Val)
	right := recursiveCheck(node.Right, node.Val, max)

	return left && right
}

// levelOrder Problem: Binary Tree Level Order Traversal.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/94/trees/628/
func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}

	if root == nil {
		return arr
	}

	if root.Left == nil && root.Right == nil {
		arr = append(arr, []int{root.Val})
		return arr
	}

	q := &Queue{}
	q.Enqueue(&QueueValue{level: 0, node: root})
	for !q.IsEmpty() {

		qv := q.Dequeue()

		// Visiting node.
		currentLevel := qv.level
		currentNode := qv.node

		if len(arr) < currentLevel+1 {
			arr = append(arr, []int{})
		}

		arr[currentLevel] = append(arr[currentLevel], currentNode.Val)

		nextLevel := currentLevel + 1
		if currentNode.Left != nil {
			q.Enqueue(&QueueValue{level: nextLevel, node: currentNode.Left})
		}
		if currentNode.Right != nil {
			q.Enqueue(&QueueValue{level: nextLevel, node: currentNode.Right})
		}
	}

	return arr
}

type QueueValue struct {
	level int
	node  *TreeNode
}

type Queue struct {
	data []*QueueValue
}

func (q *Queue) Enqueue(value *QueueValue) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeue() *QueueValue {
	if q.IsEmpty() {
		return nil
	}
	ret := q.data[0]
	q.data = q.data[1:]
	return ret
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Top() *QueueValue {
	if q.IsEmpty() {
		return nil
	}
	return q.data[0]
}
