/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    rootIndex := &Index{len(postorder) - 1}
	return recursive(inorder, postorder, 0, len(postorder)-1, rootIndex)
}

type Index struct {
	index int
}

func recursive(inorder []int, postorder []int, start int, end int, rootIndex *Index) *TreeNode {
	if start > end {
		return nil
	}

	rootValue := postorder[rootIndex.index]
	root := &TreeNode{rootValue, nil, nil}
	rootIndex.index--

	if start == end {
		return root
	}

	index := 0
	for i := start; i <= end; i++ {
		if inorder[i] == rootValue {
			index = i
			break
		}
	}

	root.Right = recursive(inorder, postorder, index+1, end, rootIndex)
	root.Left = recursive(inorder, postorder, start, index-1, rootIndex)

	return root
}