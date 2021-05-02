func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    var res int
    for i:=range items{
        one:=items[i]
        if ruleKey=="type"&&one[0]==ruleValue{
            res++
        }
        if ruleKey=="color"&&one[1]==ruleValue{
            res++
        }
        if ruleKey=="name"&&one[2]==ruleValue{
            res++
        }
    }
    return res
}
