/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var path[]string
var sliceStr= []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o",
		"p","q","r","s","t","u","v","w","x","y","z"}
func smallestFromLeaf(root *TreeNode) string {
    path=nil
    dfs(root,"")
    //按小到大排序
    sort.Strings(path)
    return path[0]//字典序最小
}

func dfs(root *TreeNode, str string){
    if root==nil{
        return
    }
    str+=sliceStr[root.Val]
    if root.Left==nil&&root.Right==nil{
        //leave
        str=reverse(str)//题目要求从叶节点到根节点，因此反转
        path=append(path,str)
        return
    }
    dfs(root.Left,str)
    dfs(root.Right,str)
    return
}

func reverse(str string)string{
    var res string
    for i:=0;i<len(str);i++{
        res+=fmt.Sprintf("%c",str[len(str)-1-i])
    }
    return res
}
