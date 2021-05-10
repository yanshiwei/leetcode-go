/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    //map,golang中map要求key可比较如基本类型/结构体，切片不能比较
    var mapTo=make(map[*Node]*Node)//key是旧的节点，value是新的节点
    //第一步构造节点，并存储到map
    var cur=head
    for cur!=nil{
        mapTo[cur]=&Node{Val:cur.Val}
        cur=cur.Next
    }
    //第二步，构造next
    cur=head
    for cur!=nil{
        //新节点的next对于旧节点的next对应的节点
        mapTo[cur].Next=mapTo[cur.Next]
        cur=cur.Next
    }
    //第三步构造random
    cur=head
    for cur!=nil{
        //新节点的Random对于旧节点的Random对应的节点
        mapTo[cur].Random=mapTo[cur.Random]
        cur=cur.Next
    } 
    return mapTo[head]
}
