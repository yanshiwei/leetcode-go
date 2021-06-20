/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
    root *TreeNode
}

/*
知识点：如果给完全二叉树的每个节点都标上序号，并令根节点的序号为1，则对于序号为n的节点，其左子节点的序号应为2n，右子节点的序号应为2n + 1，父节点的序号应为n/2
*/
func Constructor(root *TreeNode) CBTInserter {
    var m CBTInserter
    m.root=root
    return m
}


func (this *CBTInserter) Insert(v int) int {
    var queue []*TreeNode
    queue=append(queue,this.root)
    for len(queue)>0{
        cur:=queue[0]
        queue=queue[1:]
        if cur.Left==nil{
            cur.Left=&TreeNode{Val:v}
            return cur.Val
        }else{
            queue=append(queue,cur.Left)
        }
        if cur.Right==nil{
            cur.Right=&TreeNode{Val:v}
            return cur.Val
        }else{
            queue=append(queue,cur.Right)
        }
    }
    return 0
}


func (this *CBTInserter) Get_root() *TreeNode {
    return this.root
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
