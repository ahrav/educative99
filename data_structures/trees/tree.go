package trees

import (
	"fmt"
	"strconv"
	"strings"

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

func isSymmetric(root *TreeNode[int]) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode[int]
	if root.Left != nil && root.Right != nil {
		queue = append(queue, root.Left, root.Right)
	} else {
		return false
	}

	for len(queue) > 1 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Data != right.Data {
			return false
		}

		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}

func isSymmetricRecursive(root *TreeNode[int]) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode[int]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Data == right.Data && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func levelOrderTraversal(root *TreeNode[int]) string {
	if root == nil {
		return "None"
	}

	var results strings.Builder
	var queue []*TreeNode[int]
	var next []*TreeNode[int]
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}

		if len(queue) == 0 {
			if len(next) > 0 {
				results.WriteString(fmt.Sprintf("%s : ", strconv.Itoa(node.Data)))
				queue = append(queue, next...)
				next = next[:0]
				continue
			} else {
				results.WriteString(fmt.Sprintf("%s", strconv.Itoa(node.Data)))
				break
			}

		}
		results.WriteString(fmt.Sprintf("%s, ", strconv.Itoa(node.Data)))
	}

	return results.String()
}

func levelOrderTraversalNested(root *TreeNode[int]) string {
	if root == nil {
		return "None"
	}

	var results []string
	var queue []*TreeNode[int]
	queue = append(queue, root)
	for len(queue) > 0 {
		level := len(queue)
		var levelNodes []string

		for i := 0; i < level; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, strconv.Itoa(node.Data))

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		results = append(results, strings.Join(levelNodes, ", "))

	}

	return strings.Join(results, " : ")
}

func verticalOrder(root *TreeNode[int]) [][]int {
	type nodeInfo struct {
		idx  int
		node *TreeNode[int]
	}

	if root == nil {
		return make([][]int, 0)
	}

	// Mapping between the column index and all the node values that reside in that column.
	colIdxValues := make(map[int][]int)

	var queue []nodeInfo
	var minIdx, maxIdx int
	queue = append(queue, nodeInfo{idx: 0, node: root})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if _, ok := colIdxValues[node.idx]; ok {
			colIdxValues[node.idx] = append(colIdxValues[node.idx], node.node.Data)
		} else {
			colIdxValues[node.idx] = []int{node.node.Data}
		}

		if left := node.node.Left; left != nil {
			col := node.idx - 1
			minIdx = min(minIdx, col)
			queue = append(queue, nodeInfo{idx: col, node: left})
		}

		if right := node.node.Right; right != nil {
			col := node.idx + 1
			maxIdx = max(maxIdx, col)
			queue = append(queue, nodeInfo{idx: col, node: right})
		}
	}

	results := make([][]int, 0, len(colIdxValues))
	for i := minIdx; i <= maxIdx; i++ {
		results = append(results, colIdxValues[i])
	}

	return results
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

func min[T constraints.Integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func wordLadder(src string, dest string, words []string) int {
	set := make(map[string]struct{}, len(words))
	for _, w := range words {
		set[w] = struct{}{}
	}

	if _, ok := set[dest]; !ok {
		return 0
	}

	queue := make([]string, 0, len(words)+2)
	queue = append(queue, src)

	var cnt int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for w := range set {
			if differByOneChar(cur, w) {
				delete(set, w)
				queue = append(queue, w)
				cnt++
				if w == dest {
					return cnt
				}
			}
		}
	}

	return 0
}

func differByOneChar(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var cnt int
	for i := range a {
		if a[i] != b[i] {
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}

	return true
}
