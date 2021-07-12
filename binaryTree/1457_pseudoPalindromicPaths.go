/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var fre []int
 var res int
func pseudoPalindromicPaths (root *TreeNode) int {
    /*
    1、对于二叉树路径使用dfs
    2、对于回文序列的判断，统计数字奇数个数即可（如果数字个数的奇数个数 <= 1，可以构成回文序列；否则不能构成回文序列）  
    */
    if root==nil{
        return 0
    }
    fre=make([]int,10)
    res=0
    dfs(root)
    return res
}
func dfs(root *TreeNode){
    if root==nil{
        return
    }
    fre[root.Val]++
    if root.Left==nil&&root.Right==nil{
        if check(){
            res++
        }
        fre[root.Val]--
        return
    }
    dfs(root.Left)
    dfs(root.Right)
    fre[root.Val]--
}

func check()bool{
    var res int
    for i:=0;i<10;i++{
        if fre[i]%2==1{
            //个数为奇数
            res++
        }
    }
    if res>1{
        return false
    }
    return true
}
