package main

import (
	"fmt"
)

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

func (TreeNode) InOrderPrint(n *TreeNode) {
	if n == nil {
		return
	}
	n.InOrderPrint(n.Left)
	fmt.Print(n.Value, " ")
	n.InOrderPrint(n.Right)
}

func (TreeNode) PreOrderPrint(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Print(n.Value, " ")
	n.PreOrderPrint(n.Left)
	n.PreOrderPrint(n.Right)
}

func (TreeNode) PostOrderPrint(n *TreeNode) {
	if n == nil {
		return
	}
	n.PostOrderPrint(n.Left)
	n.PostOrderPrint(n.Right)
	fmt.Print(n.Value, " ")
}

func main() {
	fmt.Println("Tree exercise")

	treeNode6 := TreeNode{Value: "6"}
	treeNode5 := TreeNode{Value: "5"}
	treeNode4 := TreeNode{Value: "4"}
	treeNode3 := TreeNode{Value: "3"}
	treeNode1 := TreeNode{"1", &treeNode5, &treeNode6}
	treeNode2 := TreeNode{"2", &treeNode3, &treeNode4}
	rootNode := TreeNode{"0", &treeNode1, &treeNode2}
	fmt.Println("PreOrderPrint:")
	rootNode.PreOrderPrint(&rootNode)
	fmt.Println()
	fmt.Println("InOrderPrint:")
	rootNode.InOrderPrint(&rootNode)
	fmt.Println()
	fmt.Println("PostOrderPrint:")
	rootNode.PostOrderPrint(&rootNode)
}
