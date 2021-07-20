func verifyPostorder(postorder []int) bool {
    if len(postorder)==0||len(postorder)==1{
        return true
    }
    //后序序列数组的最后一个元素为根节点。前面的数组中小于根节点数值的为左子树，后一段大于根节点数值的为右子树。合法的是前一部分都小于根，剩下的后一部分都大于根
    rootV:=postorder[len(postorder)-1]
    var i,j int
    //找到第一个大于根节点
    for i=0;i<len(postorder)-1;i++{
        if postorder[i]>rootV{
            break
        }
    }
    //是否都大于根
    for j=i;j<len(postorder)-1;j++{
        if postorder[j]<rootV{
            return false
        }
    }
    //递归
    return verifyPostorder(postorder[0:i])&&verifyPostorder(postorder[i:len(postorder)-1])
}
