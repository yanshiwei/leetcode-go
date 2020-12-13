func validateStackSequences(pushed []int, popped []int) bool {
    if len(pushed)!=len(popped) {
        return false
    }
    var stack []int
    for i:=range pushed {
        if len(popped)<1 {
            return false
        }
        stack=append(stack,pushed[i])
        for len(stack)>0&&len(popped)>0&&stack[len(stack)-1]==popped[0] {
            popped=popped[1:]//pop head
            stack=stack[0:len(stack)-1]
        }
    }
    if len(stack)>0 {
        return false
    }
    return true
}
