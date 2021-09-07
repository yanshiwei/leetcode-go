func canCompleteCircuit(gas []int, cost []int) int {
    var n=len(gas)
    //each node to start
    for i:=0;i<n;i++{
        j:=i
        remain:=gas[i]
        //剩余的油能否到达下一点
        for remain-cost[j]>=0{
            remain=remain-cost[j]+gas[(j+1)%n]
            j=(j+1)%n
            if j==i{
                return i
            }
        }
    }
    return -1
}
