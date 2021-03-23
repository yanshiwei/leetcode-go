func distanceBetweenBusStops(distance []int, start int, destination int) int {
    var m=len(distance)
    var leftDis,rightDis int
    var i =start
    for {
        leftDis+=distance[i]
        i=(i+1)%m
        if i==destination{
            break
        }
    }
    i=destination
    for {
        rightDis+=distance[i]
        i=(i+1)%m
        if i==start{
            break
        }
    }
    //fmt.Println(leftDis,rightDis)
    return min(leftDis,rightDis)
}

func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
