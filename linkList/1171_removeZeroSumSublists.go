/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    // 首次遍历建立 节点处链表和<->节点 哈希表
    // 若同一和出现多次会覆盖，即记录该sum出现的最后一次节点
    var fre=make(map[int]*ListNode)
    var sum int
    var cur=head
    for cur!=nil{
        sum+=cur.Val
        fre[sum]=cur
        cur=cur.Next
    }
    // 第二遍遍历 若当前节点处sum在下一处出现了则表明两结点之间所有节点和为0 直接删除区间所有节点
    var dummy=&ListNode{
        Val:0,
        Next:head,
    }
    //引入dummy是防止head节点也会被处理
    sum=0
    cur=dummy
    for cur!=nil{
        sum+=cur.Val
        if anotherNode,ok:=fre[sum];ok{
            //如fre中存在该sum则说明当前节点处sum在下一处出现,则表明两结点之间所有节点和为0，如果sum只有唯一值，cur.Next=anotherNode.Next也不改变什么;如果sum出现了多次，fre中则出现的是最后一次的节点，直接清楚这两个节点所在的区间即可
            cur.Next=anotherNode.Next
        }
        cur=cur.Next
    }
    return dummy.Next
}
