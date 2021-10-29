package entity_test

import (
	// "fmt"
	"testing"

	"spenmo/entity"

	"github.com/stretchr/testify/assert"
)

// @Created 29/10/2021
// @Updated
func TestTreeNode_Insert(t *testing.T) {
	tree := generateTree()

	rootNode := tree.Val
	assert.Equal(t, rootNode, 16)
}

// @Created 29/10/2021
// @Updated
func TestTreeNode_Insert_Failed(t *testing.T) {
	tree := generateTree()

	rootNode := tree.Val
	assert.Equal(t, rootNode, 16)
}

// @Created 29/10/2021
// @Updated
func TestTreeNode_Find(t *testing.T) {
	tree := generateTree()

	find16 := tree.Find(16)
	assert.Equal(t, find16, true)
}

// @Created 29/10/2021
// @Updated
func TestTreeNode_Find_Failed(t *testing.T) {
	tree := generateTree()

	find8 := tree.Find(99)
	assert.Equal(t, find8, false)
}

// @Created 29/10/2021
// @Updated
// func TestTreeNode_Find_Failed_Default(t *testing.T) {
// 	var tree entity.TreeNode

// 	res := 
// }

// @Created 29/10/2021
// @Updated
func generateTree() *entity.TreeNode {
	var tree entity.TreeNode

	// Initialize binary tree
	tree.Insert(16)
	tree.Insert(7)
	tree.Insert(100)
	tree.Insert(10)
	tree.Insert(8)
	tree.Insert(14)

	return &tree
}