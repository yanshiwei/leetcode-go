func romanToInt(s string) int {
    var dict=make(map[byte]int)
    dict['I']=1
    dict['V']=5
    dict['X']=10
    dict['L']=50
    dict['C']=100
    dict['D']=500
    dict['M']=1000
    var res int
    for i:=0;i<len(s);i++{
        val:=dict[s[i]]
        if i<len(s)-1&&val<dict[s[i+1]]{
            //IV case
            res-=val
        }else{
            res+=val
        }
    }
    return res
}
