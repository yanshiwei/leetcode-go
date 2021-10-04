func lengthOfLongestSubstring(s string) int {
    if len(s)<1{
        return 0
    }
    var dict=make(map[byte]struct{})
    var maxLen=0
    var curLen=0
    var left=0
    for i:=range s{
        curLen++
        for {
            if _,ok:=dict[s[i]];ok==true{
                //remove windows left 
                delete(dict,s[left])
                left++
                curLen--
            }else{
                break
            }
        }
        if curLen>maxLen{
            maxLen=curLen
        }
        dict[s[i]]=struct{}{}
    }
    return maxLen
}
