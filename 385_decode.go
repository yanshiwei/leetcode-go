/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
    var res =new(NestedInteger)
    if len(s)<1 {
        return res
    }
    if s[0]!='[' {
        //only data
        intV,_:=strconv.Atoi(s)
        res.SetInteger(intV)
        return res
    }
    var stack []*NestedInteger
    var first =true
    var digitIndex int
    for i,one:=range s {
        if one=='['{
            //[ start a new NestedInteger
            newNest:=new(NestedInteger)
            stack=append(stack,newNest)
            if first {
                first=false//assigned to res
                res=stack[len(stack)-1]
            }
            digitIndex=i+1//must be next
        }else if one==',' {
            if digitIndex!=i {
                stack[len(stack)-1].Add(*deserialize(s[digitIndex:i]))//handle digital
            }
            digitIndex=i+1
        }else if one ==']' {
            head:=stack[len(stack)-1]
            if digitIndex!=i {
                head.Add(*deserialize(s[digitIndex:i]))//hanlde digital before ]
            }
            stack=stack[0:len(stack)-1]//pop to skip a []
            if len(stack)>0 {
                stack[len(stack)-1].Add(*head)//join a NestedInteger
            }
            digitIndex=i+1
        }
    }
    return res
}
