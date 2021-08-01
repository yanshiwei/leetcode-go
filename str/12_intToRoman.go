func intToRoman(num int) string {
    var dict=make([]Info,0)
    dict=append(dict,Info{1000,"M"})
    dict=append(dict,Info{900,"CM"})
    dict=append(dict,Info{500,"D"})
    dict=append(dict,Info{400,"CD"})
    dict=append(dict,Info{400,"CD"})
    dict=append(dict,Info{100,"C"})
    dict=append(dict,Info{90,"XC"})
    dict=append(dict,Info{50,"L"})
    dict=append(dict,Info{40,"XL"})
    dict=append(dict,Info{10,"X"})
    dict=append(dict,Info{9,"IX"})
    dict=append(dict,Info{5,"V"})
    dict=append(dict,Info{4,"IV"})
    dict=append(dict,Info{1,"I"})
    var res string
    for _,v:=range dict{
        for num>=v.Int{
            num-=v.Int
            res+=v.Str
        }
        if num<=0{
            break
        }
    }
    return res
}

type Info struct {
    Int int
    Str string
}
