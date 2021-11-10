package leetcode

import (
	"reflect"
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	t.Run("new empty tree", func(t *testing.T) {
		assertTreeSame(t, nil, NewTree())
	})

	t.Run("tree with root node only", func(t *testing.T) {
		want := &TreeNode{Val: 7788}
		got := NewTree(7788)
		assertTreeSame(t, want, got)
	})

	t.Run("[1, 5, 7]", func(t *testing.T) {
		var got = NewTree(1, 5, 7)
		var want = &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		}
		assertTreeSame(t, want, got)
	})

	t.Run("[1, nil, 2, 3]", func(t *testing.T) {
		var got = NewTree(1, nil, 2, 3)
		var want = &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		}
		assertTreeSame(t, want, got)
	})

}

func assertTreeSame(t *testing.T, want, got *TreeNode) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("tree not equal")
		t.Logf("want: %v", want)
		t.Logf(" got: %v", got)
	}
}
