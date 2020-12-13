func minOperations(logs []string) int {
    var res int
    for _,one:=range logs {
        if one == "../" {
            if res>0 {
                res=res-1
            }
        }else if one == "./" {
            //
        }else {
            res=res+1
        }
    }
    return res
}
