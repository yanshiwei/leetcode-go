func kidsWithCandies(candies []int, extraCandies int) []bool {
    var maxCnt int
    for i:=range candies{
        if maxCnt<candies[i]{
            maxCnt=candies[i]
        }
    }
    var res []bool
    for i:=range candies{
        if candies[i]+extraCandies>=maxCnt{
            res=append(res,true)
        }else{
            res=append(res,false)
        }
    }
    return res
}
