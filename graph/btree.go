package graph

import (
	"io"
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type TreeNode struct {
	// value stored in the node
	Val	any
	// left children
	Left	*TreeNode
	// right children
	Right	*TreeNode
}

// NewBTree is the entry point to build a complete binary tree from given a slice
func NewBTree(elems []any) *TreeNode {
	return newBTree(elems, 0)
}

// newBTree build a complete binary tree from given a slice
func newBTree(elems []any, index int) *TreeNode {
	if len(elems) == 0 {
		return nil
	}

        // 为叶子结点
        if 2*index + 1 > len(elems) - 1 {
		return &TreeNode{ elems[index], nil, nil, }
	}

        // 有子节点
	var l, r *TreeNode
        if 2*index + 1 < len(elems) && elems[2*index+1] != nil {
		l = newBTree(elems, 2*index+1)
	}

        if 2*index + 2 < len(elems) && elems[2*index+2] != nil {
		r = newBTree(elems, 2*index+2)
	}

	return &TreeNode{
		Val: elems[index],
		Left: l,
		Right: r,
        }
}

var count int
func printIndent(n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("  ")
    }
}

// PrintBTreeDFS print out the entire btree in DFS order.
func PrintBTreeDFS(w io.Writer, root *TreeNode) {
	if root == nil {
		fmt.Fprintln(w, "Empty tree")
		return
	}

	printIndent(count)
	fmt.Fprintln(w, "Current node's value: ", root.Val)
	if root.Left !=	nil {
		printIndent(count)
		fmt.Fprintf(w, "%v have a left child\n", root.Val)
		count++
		PrintBTreeDFS(w, root.Left)
		count--
	}
	if root.Right != nil {
		printIndent(count)
		fmt.Fprintf(w, "%v have a right child\n", root.Val)
		count++
		PrintBTreeDFS(w, root.Right)
		count--
	}
}

// PrintBTreeBFS print out the entire btree in BFS order.
func PrintBTreeBFS(w io.Writer, root *TreeNode) {
	rslt := []any{}
	if root == nil {
		fmt.Fprintln(w, rslt)
		return
	}

	q := datastruct.NewQueue([]*TreeNode{})
	q.Push(root)

	var emptyNode TreeNode
	size := 0

	loop: for !q.IsEmpty() {
		size = q.Size()
		emptyNum := 0
		for i:=0; i<size; i++ {
			node := q.Pop()
			if *node != emptyNode {
				rslt = append(rslt, node.Val)
			}else{
				emptyNum++
				rslt = append(rslt, nil)
			}

			if emptyNum == size {
				break loop
			}

			if node.Left != nil {
				q.Push(node.Left)
			}else{
				q.Push(&TreeNode{})
			}

			if node.Right != nil {
				q.Push(node.Right)
			}else{
				q.Push(&TreeNode{})
			}

		}
	}

	fmt.Fprintln(w, rslt[:len(rslt)-size])
}
