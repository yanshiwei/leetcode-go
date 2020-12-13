func isValidSerialization(preorder string) bool {
    inPutArr:=strings.Split(preorder,",")//never two ','
    for len(inPutArr)>0 {
        var deletedIndex=make(map[int]bool)
        var res[]string
        var beoreLen=len(inPutArr)
        //each loop delete # when 'data # #' exist
        for i:=0;i<len(inPutArr)-2;i++{
            if inPutArr[i]!="#" && inPutArr[i+1]=="#" && inPutArr[i+2]== "#" {
                inPutArr[i]="#"
                deletedIndex[i+1]=true
                deletedIndex[i+2]=true
            }
        }
        for i:=0;i<len(inPutArr);i++ {
            if v,ok:=deletedIndex[i];ok {
                if v==false {
                    res=append(res,inPutArr[i])
                }
            }else {
                res=append(res,inPutArr[i])
            }
        }
        inPutArr=res//update
        var afterLen=len(inPutArr)
        if beoreLen==afterLen {
            //no change so break
            break
        }
    }
    if len(inPutArr)==1&&inPutArr[0]=="#" {
        return true
    }else {
        return false
    }
}
