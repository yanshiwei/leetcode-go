func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    /*
    切完后蛋糕的最大高度及宽度，那么就可以求得最大面积
    1⃣最大高度的求法：
        1.将 horizontalCuts 排序；
        2.求第一条水平线与上侧形成的高度h0；最后一条水平线与下册形成的高度hl
        3.求horizontalCuts间高度差
    最大宽度求法类似
    */
    sort.Ints(horizontalCuts)
    sort.Ints(verticalCuts)
    var maxH,maxW int
    //求第一条水平线与上侧形成的高度h0；最后一条水平线与下侧形成的高度hl
    maxH=max(horizontalCuts[0],h-horizontalCuts[len(horizontalCuts)-1])
    //求horizontalCuts间高度差
    for i:=1;i<len(horizontalCuts);i++{
        maxH=max(maxH,horizontalCuts[i]-horizontalCuts[i-1])
    }

    maxW=max(verticalCuts[0],w-verticalCuts[len(verticalCuts)-1])
    for i:=1;i<len(verticalCuts);i++{
        maxW=max(maxW,verticalCuts[i]-verticalCuts[i-1])
    }
    return (maxH%1000000007)*(maxW%1000000007)%1000000007
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
