package graph

import (
	"testing"
	"bytes"
)

func TestPrintBTreeBFS(t *testing.T) {
	tree := NewBTree([]interface{}{3,9,20,nil,nil,15,7})

	var buf bytes.Buffer
	PrintBTreeBFS(&buf, tree)

	expectedBFS := "[3 9 20 <nil> <nil> 15 7]\n"
	if buf.String() != expectedBFS {
		t.Errorf("PrintBTreeBFS: Expect to have %s, got %s", expectedBFS, buf.String())
	}
}

func TestPrintBTreeBFSEmpty(t *testing.T) {
	tree := NewBTree([]interface{}{})

	var buf bytes.Buffer
	PrintBTreeBFS(&buf, tree)

	expectedBFS := "[]\n"
	if buf.String() != expectedBFS {
		t.Errorf("PrintBTreeBFS: Expect to have %s, got %s", expectedBFS, buf.String())
	}
}

func TestPrintBTreeDFS(t *testing.T) {
	tree := NewBTree([]interface{}{3,9,20,nil,nil,15,7})

	var buf bytes.Buffer
	PrintBTreeDFS(&buf, tree)

	expectedDFS := `Current node's value:  3
3 have a left child
Current node's value:  9
3 have a right child
Current node's value:  20
20 have a left child
Current node's value:  15
20 have a right child
Current node's value:  7
`
	if buf.String() != expectedDFS {
		t.Errorf("PrintBTreeDFS: Expect to have %s, got %s", expectedDFS, buf.String())
	}
}

func TestPrintBTreeDFSEmpty(t *testing.T) {
	tree := NewBTree([]interface{}{})

	var buf bytes.Buffer
	PrintBTreeDFS(&buf, tree)

	expectedDFS := "Empty tree\n"
	if buf.String() != expectedDFS {
		t.Errorf("PrintBTreeDFS: Expect to have %s, got %s", expectedDFS, buf.String())
	}
}
