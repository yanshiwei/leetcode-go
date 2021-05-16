/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
var tail,res *Node
func flatten(root *Node) *Node {
    /*
    链表看成二叉树，child是左子树，next是右子树，前序遍历
    */
    if root==nil{
        return root
    }
    tail=nil//一定要初始化 否则缓存有历史结果
    res=nil
    preDfs(root)
    return res
}

func preDfs(node *Node){
    if node==nil{
        return
    }
    //相当于前序遍历的主体，把遍历到的当前节点放入新的链表里
    if tail==nil{
        res=&Node{}
        res.Val=node.Val
        tail=res
    }else{
        tmp:=&Node{}
        tmp.Val=node.Val
        tail.Next=tmp
        tmp.Prev=tail
        tail=tmp
    }
    preDfs(node.Child)
    preDfs(node.Next)
}
