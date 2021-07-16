/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
 参考：https://leetcode-cn.com/problems/merge-bsts-to-create-single-bst/solution/fen-xiang-xia-si-kao-guo-cheng-by-time-l-6o84/
 条件一：叶子节点的值不能重复。如果叶子节点有重复，必然无法构造出二叉搜索树
 条件二：设 S 为叶子节点的值的集合，则有且仅有一个根节点的值不在 S 内。
    2.1 当有多个根节点的值不在 S 内时，意味着有多棵树无法合并到其他树的叶子节点，则必然无法合成一棵树；
    2.2 当所有根节点的值都在 SS 内时，意味着有出现了合并的回路
 */
 var fre map[int]*TreeNode// key is val value is node
 var res []int
func canMerge(trees []*TreeNode) *TreeNode {
    //check case 1,叶子节点的值不能重复，重复了就不能是BST
    fre=make(map[int]*TreeNode)
    var set =make(map[int]struct{})
    for i:=range trees{
        t:=trees[i]
        if t.Left!=nil{
            //只有三个节点，故这就是叶子
            _,ok:=set[t.Left.Val]
            if ok==false{
                set[t.Left.Val]=struct{}{}
            }else{
                //exist
                return nil
            }
        }
        if t.Right!=nil{
            _,ok:=set[t.Right.Val]
            if ok==false{
                set[t.Right.Val]=struct{}{}
            }else{
                //exit
                return nil
            }
        }
    }
    //check case 2, 设 S 为叶子节点的值的集合，则有且仅有一个根节点的值不在 S 内。
    var include int
    var final *TreeNode
    for i:=range trees{
        t:=trees[i]
        if _,ok:=set[t.Val];ok{
            include++
        }else{
            final=t
        }
    }
    if include!=len(trees)-1{
        return nil
    }
    // 构造 node->val 到 node 的映射
    for i:=range trees{
        t:=trees[i]
        if t!=final{
            fre[t.Val]=t
        }
    }
    // 开始合并
    dfsForMerge(final)
    if len(fre)>0{
        //仍存在单独的二叉树
        return nil
    }
    // 中序遍历检查bst
    res=nil
    dfsForBST(final)
    for i:=1;i<len(res);i++{
        if res[i-1]>=res[i]{
            return nil
        }
    }
    return final
}

func dfsForMerge(final *TreeNode){
    if final==nil{
        return
    }
    if final.Left==nil&&final.Right==nil{
        //leaf
        it,ok:=fre[final.Val]
        if ok{
            //值相等
            final.Left=it.Left
            final.Right=it.Right
            delete(fre,it.Val)
        }
    }
    dfsForMerge(final.Left)
    dfsForMerge(final.Right)
}

func dfsForBST(root *TreeNode){
    if root==nil{
        return
    }
    dfsForBST(root.Left)
    res=append(res,root.Val)
    dfsForBST(root.Right)
}
