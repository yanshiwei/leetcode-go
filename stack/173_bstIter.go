/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    //bst midlle travel is asc
    var bt =BSTIterator{stack: make([]*TreeNode,0)}
    var cur=root
    for cur!=nil {
        bt.stack=append(bt.stack,cur)
        cur=cur.Left
    }
    return bt
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    var res int
    cur:=this.stack[len(this.stack)-1]
    this.stack=this.stack[0:len(this.stack)-1]
    //pop head
    res=cur.Val
    cur=cur.Right
    for cur!=nil {
        this.stack=append(this.stack,cur)
        cur=cur.Left
    }
    return res
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    if len(this.stack)>0 {
        return true
    }
    return false
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
