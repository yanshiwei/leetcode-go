func deckRevealedIncreasing(deck []int) []int {
    /*
    如果从牌组中以 [0, 2, 4, ...] 的顺序取牌，我们只需要把最小的牌放在下标为 0 的地方，第二小的牌放在下标为 2 的地方，第三小的牌放在下标为 4 的地方，依次类推
    */
    var index=make([]int,len(deck))
    for i:=range deck{
        index[i]=i
    }
    sort.Ints(deck)
    var res=make([]int,len(deck))
    for i:=range deck {
        card:=deck[i]
        //pull first
        //index is 0 2 4 6...
        first:=index[0]
        index=index[1:]
        res[first]=card
        if len(index)>0 {
            //index 1 3 5...
            second:=index[0]
            index=index[1:]
            index=append(index,second)
        }
    }
    return res
}
