package trees

import "testing"

func BenchmarkIsSymmetric(b *testing.B) {
	// Create a larger symmetric tree
	root := &TreeNode[int]{
		Data: 1,
		Left: &TreeNode[int]{
			Data: 2,
			Left: &TreeNode[int]{
				Data:  3,
				Left:  &TreeNode[int]{Data: 4},
				Right: &TreeNode[int]{Data: 5},
			},
			Right: &TreeNode[int]{
				Data:  6,
				Left:  &TreeNode[int]{Data: 7},
				Right: &TreeNode[int]{Data: 8},
			},
		},
		Right: &TreeNode[int]{
			Data: 2,
			Left: &TreeNode[int]{
				Data:  6,
				Left:  &TreeNode[int]{Data: 8},
				Right: &TreeNode[int]{Data: 7},
			},
			Right: &TreeNode[int]{
				Data:  3,
				Left:  &TreeNode[int]{Data: 5},
				Right: &TreeNode[int]{Data: 4},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSymmetric(root)
	}
}

func BenchmarkIsSymmetricRecursive(b *testing.B) {
	// Create a larger symmetric tree
	root := &TreeNode[int]{
		Data: 1,
		Left: &TreeNode[int]{
			Data: 2,
			Left: &TreeNode[int]{
				Data:  3,
				Left:  &TreeNode[int]{Data: 4},
				Right: &TreeNode[int]{Data: 5},
			},
			Right: &TreeNode[int]{
				Data:  6,
				Left:  &TreeNode[int]{Data: 7},
				Right: &TreeNode[int]{Data: 8},
			},
		},
		Right: &TreeNode[int]{
			Data: 2,
			Left: &TreeNode[int]{
				Data:  6,
				Left:  &TreeNode[int]{Data: 8},
				Right: &TreeNode[int]{Data: 7},
			},
			Right: &TreeNode[int]{
				Data:  3,
				Left:  &TreeNode[int]{Data: 5},
				Right: &TreeNode[int]{Data: 4},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSymmetricRecursive(root)
	}
}

func BenchmarkLevelOrderTraversal(b *testing.B) {
	// Create a sample binary tree for benchmarking
	root := &TreeNode[int]{
		Data: 1,
		Left: &TreeNode[int]{
			Data:  2,
			Left:  &TreeNode[int]{Data: 4},
			Right: &TreeNode[int]{Data: 5},
		},
		Right: &TreeNode[int]{
			Data:  3,
			Left:  &TreeNode[int]{Data: 6},
			Right: &TreeNode[int]{Data: 7},
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		levelOrderTraversalNested(root)
	}
}
