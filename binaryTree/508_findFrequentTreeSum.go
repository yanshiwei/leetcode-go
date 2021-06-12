/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var fre map[int]int//每个和出现的次数
var res []int
var maxCnt int
func findFrequentTreeSum(root *TreeNode) []int {
    /*
    一个结点的「子树元素和」:
    当前节点的值+其左子树的和+其右子树的和:
    int sum = root.val + findSum(root.left) +findSum(root.right);
    */
    fre=make(map[int]int)
    res=nil
    maxCnt=0
    if root==nil{
        return nil
    }
    findSum(root)
    for k,v:=range fre{
        if v==maxCnt{
            res=append(res,k)
        }
    }
    return res
}
//后序DFS，计算当前root开始的子树的和
func findSum(root *TreeNode)int {
    if root==nil{
        return 0
    }
    //计算左子树的和
    leftV:=findSum(root.Left)
    //计算右子树的和
    rightV:=findSum(root.Right)
    //计算当前子树的和
    sum:=root.Val+leftV+rightV
    if _,ok:=fre[sum];ok==false{
        fre[sum]=1
    }else{
        fre[sum]++
    }
    maxCnt=max(maxCnt,fre[sum])
    return sum
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
