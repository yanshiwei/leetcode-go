func isOneBitCharacter(bits []int) bool {
    if len(bits)<1{
        return false
    }
    for i:=0;i<len(bits);{
        if i==len(bits)-1&&bits[i]==0{
            return true
        }
        if bits[i]==0 {
            i++
        }else{
            i+=2
        }
    }
    return false
}
