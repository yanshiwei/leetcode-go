/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    quue []int
}

func push_list(nestedList []*NestedInteger)[]int {
    var res []int
    for _,one:=range nestedList {
        if one.IsInteger() {
            res=append(res,one.GetInteger())
        }else {
            res=append(res,push_list(one.GetList())...)
        }
    }
    return res
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    var st =&NestedIterator{quue:make([]int,0)}
    st.quue=push_list(nestedList)
    return st
}

func (this *NestedIterator) Next() int {
    var final=this.quue[0]
    this.quue=this.quue[1:len(this.quue)]
    return final
}

func (this *NestedIterator) HasNext() bool {
    if len(this.quue)>0 {
        return true
    }
    return false
}
