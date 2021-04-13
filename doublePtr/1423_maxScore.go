func maxScore(cardPoints []int, k int) int {
    /*
    记数组cardPoints 的长度为 n，由于只能从开头和末尾拿 k 张卡牌，所以最后剩下的必然是连续的 n−k 张卡牌。
    我们可以通过求出剩余卡牌点数之和的最小值，来求出拿走卡牌点数之和的最大值。
    由于剩余卡牌是连续的，使用一个固定长度为 n-k 的滑动窗口对数组 ardPoints 进行遍历，求出滑动窗口最小值，然后用所有卡牌的点数之和减去该最小值，即得到了拿走卡牌点数之和的最大值
    */
    var n=len(cardPoints)
    // 滑动窗口大小为 n-k
    var windSize=n-k
    ///选前 n-k 个作为初始值
    var sum=calSum(cardPoints[0:windSize])
    var minSum=sum
    for i:=windSize;i<n;i++{
        // 滑动窗口每向右移动一格，增加从右侧进入窗口的元素值，并减少从左侧离开窗口的元素值
        sum+=cardPoints[i]-cardPoints[i-windSize]
        minSum=min(minSum,sum)
    }
    return calSum(cardPoints)-minSum
}

func calSum(cardPoints []int)int {
    var sum int
    for i:=range cardPoints{
        sum+=cardPoints[i]
    }
    return sum
}

func min(x, y int)int {
    if x<y {
        return x
    }
    return y
}
