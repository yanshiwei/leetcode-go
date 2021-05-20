type MyLinkedList struct {
    Head *Node
    CurLen int
}

type Node struct{
    Val int
    Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    var m=MyLinkedList{
        CurLen:0,
        Head:&Node{
            Val:0,
            Next:nil,
        },
    }
    return m
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index<0||index>=this.CurLen{
        return -1
    }
    var cur=this.Head
    for i:=0;i<index+1;i++{
        cur=cur.Next
    }
    return cur.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    this.AddAtIndex(0,val)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    this.AddAtIndex(this.CurLen,val)
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index>this.CurLen{
        //=代表插在最后一位
        return
    }
    if index<0{
        index=0
    }
    var pre=this.Head
    for i:=0;i<index;i++{
        pre=pre.Next
    }
    tmp:=&Node{
        Val:val,
        Next:nil,
    }
    tmp.Next=pre.Next
    pre.Next=tmp
    this.CurLen++
    // var cur=this.Head.Next
    // for cur!=nil{
    //     fmt.Printf("%d-%d",val,cur.Val)
    //     cur=cur.Next
    // }   
    // fmt.Println()
    return
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index<0||index>=this.CurLen{
        return
    }
    var pre=this.Head
    for i:=0;i<index;i++{
        pre=pre.Next
    }
    pre.Next=pre.Next.Next
    this.CurLen--     
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
