func corpFlightBookings(bookings [][]int, n int) []int {
    var fre=make([]int,n)
    for i:=range bookings{
        oneBook:=bookings[i]
        start:=oneBook[0]
        end:=oneBook[1]
        num:=oneBook[2]
        for j:=start;j<=end;j++{
            fre[j-1]+=num
        }
    }
    return fre
}
