package main

import (
	"fmt"

	"bst/entity"
)

// @Created 29/10/2021
// @Updated
func main() {
	var tree entity.TreeNode

	// Initialize binary tree
	tree.Insert(16)
	tree.Insert(7)
	tree.Insert(100)
	tree.Insert(10)
	tree.Insert(8)
	tree.Insert(14)

	fmt.Println("tree.Find(16)  => ",tree.Find(16))
	fmt.Println("tree.Find(7)  => ",tree.Find(7))
	fmt.Println("tree.Find(8)  => ",tree.Find(8))
	fmt.Println("tree.Find(9)  => ",tree.Find(9))
	fmt.Println("tree.Find(10) => ",tree.Find(10))
	fmt.Println("tree.Find(14) => ",tree.Find(14))
}