package main

import "fmt"

type TreeNode struct {
	val   int ;
	left, right *TreeNode
}

// 指针接受者！  用于修改数据
func (node *TreeNode) setVal(val int) {
	node.val = val ;
}
// 值接受者(copy本质)！ 创建structure treeNode的print方法
func (node TreeNode) print(){
	fmt.Printf("current value: %d\n", node.val)
}
// 跟上面效果一样 但是本质差距很大
func print(node TreeNode){
	fmt.Printf("current value: %d\n", node.val)
}

// 创建二叉树
func (node *TreeNode) createTree(val int){

}

func printTree(root *TreeNode, level int){
	if root == nil{
		return
	} else {
		level++
	}
	printTree(root.left, level) ;
	fmt.Printf("level %d, value %d\n", level, root.val);
	printTree(root.right, level) ;

}

// 创建工场方法
func createTreeNode(pv_val int)*TreeNode {
	return &TreeNode{val : pv_val} ;
}

func main(){

	root :=  TreeNode{0, nil, nil } ;
	root.print()
	print(root)

	// 批量初始化 Node列表
	Nodes := []TreeNode{
		{10, nil, nil},
		{},
		{20, nil, &root},
	}
	fmt.Println(Nodes)

	// 打造二叉树
	left  := new(TreeNode);
	left.val = - 10
	left.setVal(-10000)  // 设置值 与上面相同
	root.left  = left ;
	root.left.left  = &TreeNode{val:-20, left:nil, right:nil} ;
	root.left.right = &TreeNode{val:-25, left:nil, right:nil} ;
	root.right = &TreeNode{20, nil, nil } ;
	// 使用工场方法
	root.right.left = createTreeNode(15)
	root.right.left = createTreeNode(25)
	printTree(&root, 0)

	/*
	pRoot := &root ; // root值也会变化
	pRoot.print()
	pRoot.setVal(-1000)
	pRoot.print()
	root.print()
	*/

}
