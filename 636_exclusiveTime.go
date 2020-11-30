type Info struct {
    Id int
    Data int
}
func exclusiveTime(n int, logs []string) []int {
    var res=make([]int,n)
    var stack []Info
    if n<1 {
        return res
    }

    for _,one:=range logs {
        logList:=strings.Split(one,":")
        id,_:=strconv.Atoi(logList[0])
        typeName:=logList[1]
        timeData,_:=strconv.Atoi(logList[2])
        info:=Info{Id:id,Data:timeData}
        if typeName=="start" {
            stack=append(stack,info)
        }else {
            //end for pop
            head:=stack[len(stack)-1]
            delta:=timeData-head.Data+1
            res[head.Id]+=delta
            stack=stack[0:len(stack)-1]//pop
            if len(stack)>0 {
                res[stack[len(stack)-1].Id]-=delta//minus middle res
            }
        }
    }
    return res
}
