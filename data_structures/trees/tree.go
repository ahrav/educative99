package trees

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

type TreeNode[T constraints.Integer] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func DFS[T constraints.Integer](root *TreeNode[T]) []T {
	var result []T
	if root == nil {
		return result
	}
	result = append(result, root.Data)
	result = append(result, DFS(root.Left)...)
	result = append(result, DFS(root.Right)...)
	return result
}

func BFS[T constraints.Integer](root *TreeNode[T]) []T {
	var result []T
	if root == nil {
		return result
	}
	var queue []*TreeNode[T]
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Data)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func InOrder[T constraints.Integer](root *TreeNode[T]) []T {
	var result []T
	if root == nil {
		return result
	}
	result = append(result, InOrder(root.Left)...)
	result = append(result, root.Data)
	result = append(result, InOrder(root.Right)...)
	return result
}

func PreOrder[T constraints.Integer](root *TreeNode[T]) []T {
	var result []T
	if root == nil {
		return result
	}
	result = append(result, root.Data)
	result = append(result, PreOrder(root.Left)...)
	result = append(result, PreOrder(root.Right)...)
	return result
}

func SerializePreOrder[T constraints.Integer](root *TreeNode[T]) []string {
	var result []string
	if root == nil {
		return result
	}
	result = append(result, strconv.Itoa(int(root.Data)))
	result = append(result, SerializePreOrder(root.Left)...)
	result = append(result, SerializePreOrder(root.Right)...)
	return result
}

func DeserializePreOrder[T constraints.Integer](data *[]T) *TreeNode[T] {
	if len(*data) == 0 {
		return nil
	}
	node := &TreeNode[T]{Data: (*data)[0]}
	*data = (*data)[1:]
	node.Left = DeserializePreOrder(data)
	node.Right = DeserializePreOrder(data)
	return node
}

func PostOrder[T constraints.Integer](root *TreeNode[T]) []T {
	var result []T
	if root == nil {
		return result
	}
	result = append(result, PostOrder(root.Left)...)
	result = append(result, PostOrder(root.Right)...)
	result = append(result, root.Data)
	return result
}

func Diameter[T constraints.Integer](root *TreeNode[T]) int {
	maxDiameter := 0

	var height func(node *TreeNode[T]) int
	height = func(node *TreeNode[T]) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.Left)
		rightHeight := height(node.Right)

		maxDiameter = max(maxDiameter, leftHeight+rightHeight)

		return 1 + max(leftHeight, rightHeight)
	}

	height(root)
	return maxDiameter
}

func MaxPathSum[T constraints.Integer](root *TreeNode[T]) T {
	maxSum := root.Data

	var maxPathSum func(node *TreeNode[T]) T
	maxPathSum = func(node *TreeNode[T]) T {
		if node == nil {
			return 0
		}

		leftSum := max(maxPathSum(node.Left), 0)
		rightSum := max(maxPathSum(node.Right), 0)

		maxSum = max(maxSum, node.Data+leftSum+rightSum)

		return node.Data + max(leftSum, rightSum)
	}

	maxPathSum(root)
	return maxSum
}

func Height[T constraints.Integer](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}

	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

func Invert[T constraints.Integer](root *TreeNode[T]) *TreeNode[T] {
	if root == nil {
		return nil
	}

	root.Left, root.Right = Invert(root.Right), Invert(root.Left)
	return root
}
