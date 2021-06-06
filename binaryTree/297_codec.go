/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strconv"
type Codec struct {
    str []string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func  serializeStr(root *TreeNode, str string)string{
    //DFS之先序遍历
    if root == nil {
        str += "null,"
    } else {
        str += strconv.Itoa(root.Val) + ","
        str = serializeStr(root.Left, str)
        str = serializeStr(root.Right, str)
    }
    return str
}
func (this *Codec) serialize(root *TreeNode) string {
    return serializeStr(root,"")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeStr() *TreeNode{
    if this.str[0]=="null"{
        this.str=this.str[1:]
        return nil
    }
    //root
    val,_:=strconv.Atoi(this.str[0])
    root:=&TreeNode{Val:val}
    this.str=this.str[1:]
    root.Left=this.deserializeStr()
    root.Right=this.deserializeStr()
    return root
}
func (this *Codec) deserialize(data string) *TreeNode {    
    l:=strings.Split(data,",")
    for i:=0;i<len(l);i++{
        if l[i]!=""{
            this.str=append(this.str,l[i])
        }
    }
    return this.deserializeStr()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
