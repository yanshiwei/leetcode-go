/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    var mapNew=make(map[*Node]*Node)//key是旧的节点，value是新的节点
    var cur=head
    //第一步构造节点，并存储到map
    for cur!=nil{
        mapNew[cur]=&Node{Val:cur.Val}
        cur=cur.Next
    }
    //第二步 构造next
    cur=head
    for cur!=nil{
        mapNew[cur].Next=mapNew[cur.Next]
        cur=cur.Next
    }
    //第二步 构造random
    cur=head
    for cur!=nil{
        mapNew[cur].Random=mapNew[cur.Random]
        cur=cur.Next
    }
    return mapNew[head]
}
