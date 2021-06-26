/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func minCameraCover(root *TreeNode) int {
    //参考https://leetcode-cn.com/problems/binary-tree-cameras/solution/968-jian-kong-er-cha-shu-di-gui-shang-de-zhuang-ta/
    res=0
    //我们要从底向上进行推导，因为尽量让叶子节点的父节点安装摄像头，这样摄像头的数量才是最少的 
    if postOrder(root)==0{
        res++
    }
    return res
}
//返回该节点状态 0 无覆盖；1 有覆盖且本身含有摄像头 2 有覆盖但本身无摄像头，左右一共9种情况
func postOrder(root *TreeNode)int{
    if root==nil{
        //空节点，因为我们希望让叶子节点的父节点安装摄像头，故状态是2，这样就可以在叶子节点的父节点放摄像头了
        return 2
    }
    left:=postOrder(root.Left)
    right:=postOrder(root.Right)
    //情况1，左右节点皆有覆盖且无摄像头 2 2这一种;
    if left==2&&right==2{
        //此时中间节点应该就是无覆
        return 0
    }
    //情况2，左右节点至少有一个无覆盖的情况,0 0;0 1; 0 2;1 0; 2 0;一共5种
    if left==0||right==0{
        res++//父节点加摄像头
        return 1
    }
    //情况3.左右节点皆有覆盖且至少有一个摄像头1 1; 1 2; 2 1;一共3种
    if left==1||right==1{
        return 2
    }
    return -1
}
