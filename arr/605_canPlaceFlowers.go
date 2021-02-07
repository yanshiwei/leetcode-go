func canPlaceFlowers(flowerbed []int, n int) bool {
    var m=len(flowerbed)
    if m<1 {
        return false
    }
    if n==0 {
        return true
    }
    var last int
    flowerbed=append([]int{0},flowerbed...)
    flowerbed=append(flowerbed,0)
    for i:=1;i<=m;i++ {
        if flowerbed[i]==0 {
            if last==0&&flowerbed[i+1]==0 {
                flowerbed[i]=1
                n--
                if n<1 {
                    return true
                }
            }
        }
        last=flowerbed[i]
    }
    return false
}
