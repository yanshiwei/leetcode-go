func averageWaitingTime(customers [][]int) float64 {
    var res float64
    var cookStartTime =0
    for i:=range customers{
        one:=customers[i]
        avrrive:=one[0]
        time:=one[1]
        if cookStartTime==0 ||cookStartTime<avrrive{
            cookStartTime=avrrive
        }
        cookStartTime=cookStartTime+time
        waitTime:=cookStartTime-avrrive
        //fmt.Println(cookStartTime,avrrive)
        res+=float64(waitTime)
    }
    return res/float64(len(customers))
}
