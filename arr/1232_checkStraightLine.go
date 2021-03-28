func checkStraightLine(coordinates [][]int) bool {
    if len(coordinates)<2{
        return false
    }
    if len(coordinates)==2{
        return true
    }
    deltaDemoX:=coordinates[1][0]-coordinates[0][0]
    deltaDemoY:=coordinates[1][1]-coordinates[0][1]
    for i:=2;i<len(coordinates);i++{
        deltaX:=coordinates[i][0]-coordinates[i-1][0]
        if deltaDemoX==0&&deltaX!=deltaDemoX{
            return false
        }else if deltaDemoX!=0{
            deltaY:=coordinates[i][1]-coordinates[i-1][1]
            if deltaY*deltaDemoX!=deltaX*deltaDemoY{
                //fmt.Println("t",deltaY,deltaDemoY)
                return false
            }
        }   
    }
    return true
}
