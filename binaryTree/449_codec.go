/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    str []string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    return serializeStr(root,"")
}

func serializeStr(root *TreeNode,str string)string{
    //DFS之先序遍历
    if root==nil{
        str+="null,"
    }else{
        str+=strconv.Itoa(root.Val)+","
        str=serializeStr(root.Left,str)
        str=serializeStr(root.Right,str)
    }
    return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    res:=strings.Split(data,",")
    for i:=0;i<len(res);i++{
        if res[i]!=""{
            this.str=append(this.str,res[i])
        }
    }
    return this.deserializeStr()
}

func (this *Codec)deserializeStr()*TreeNode{
    if this.str[0]=="null"{
        this.str=this.str[1:]
        return nil
    }
    val,_:=strconv.Atoi(this.str[0])
    root:=&TreeNode{Val:val}
    this.str=this.str[1:]
    root.Left=this.deserializeStr()
    root.Right=this.deserializeStr()
    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
