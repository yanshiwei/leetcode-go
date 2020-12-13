func simplifyPath(path string) string {
    pathList:=strings.Split(path,"/")
    var st[]string
    for i,one:=range pathList {
        if i==0 {
            st=append(st,one)
        }else {
            if one=="." {
                continue
            }else if one ==".." {
                if len(st)>0 {
                    if len(st)==1 {
                        //root path
                        continue
                    }
                    st=st[0:len(st)-1]//pop
                }
            }else {
                if one!=""{
                    st=append(st,one)
                }
            }
        }
    }
    if len(st)==1 {
        // only ////
        return "/"
    }
    return strings.Join(st,"/")
}
