func hIndex(citations []int) int {
    for i:=0;i<len(citations);i++{
        h:=len(citations)-i
        if h<=citations[i]{
            return h
        }
    }
    return 0
}
