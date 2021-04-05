func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    var res int
    /*
    该做法 空间超了
    var mat=make([][]int,n)
    for i:=range mat{
        mat[i]=make([]int,10)
    }
    var m=len(reservedSeats)
    for i:=0;i<m;i++{
        point:=reservedSeats[i]
        mat[point[0]-1][point[1]-1]=1
    }
    for i:=0;i<n;i++{
        curCow:=mat[i]
        res+=getSegCnt(curCow)
    }
    */
    /*
    位运算，https://leetcode-cn.com/problems/cinema-seat-allocation/solution/an-pai-dian-ying-yuan-zuo-wei-by-leetcode-solution/
    对于一个家庭而言，只有以下三种给他们安排座位的方法：
    安排位置 2，3，4，5；
    安排位置 4，5，6，7；
    安排位置 6，7，8，9。
    因此每一排的位置 1 和位置 10 都是没有意义的，即使被预约了也对答案没有任何影响。从下面的叙述开始，我们忽略所有在位置 1 和位置 10 的预约。同时我们可以发现，如果一排位置没有被预约，那么恰好可以安排给两个家庭，即给一个家庭安排位置 2，3，4，5，给另一个家庭安排位置 6，7，8，9；如果一排位置被预约了至少一个座位，那么最多只能安排给一个家庭了。
    这样以来，我们可以使用 8 个二进制位表示一排座位的预约情况，这里的 8 即表示位置 2 到位置 9 的这些座位。
    第一排：预约了位置 2，3，8，那么二进制数为01000011
    第二排：预约了位置 6，那么二进制数为00010000
    第三排：预约了位置 1 和 10，那么二进制数00000000
    我们可以用哈希映射（HashMap）来存储每一排以及它们的二进制数。对于哈希映射中的每个键值对，键表示电影院中的一排，值表示这一排对应的二进制数。如果某一排没有任何位置被预约（例如上面的第三排），我们实际上知道了这一排一定可以安排给两个家庭，因此可以不必将这个键值对存放在哈希映射中。也就是说，只有某一排的某一座位被预约了，我们才将这一排放入哈希映射。
    在处理完了所有的预约之后，我们遍历哈希映射。对于一个键值对(row,bitmask)，我们如何知道row 这一排可以安排给几个家庭呢？根据之前的分析，被存储在哈希映射中的这些排最多只能安排给一个家庭，那么对于三种安排座位的方法：
    1.对于安排位置 2，3，4，5，如果bitmask 中第 0，1，2，3 个二进制位均为 0，那么就可以安排给一个家庭；也就是说，bitmask 和(11110000)的按位或值保持为(bitmask)不变；
    2.对于安排位置 4，5，6，7，如果bitmask 中第 2，3，4，5 个二进制位均为 0，那么就可以安排给一个家庭；也就是说，bitmask 和(11000011)的按位或值保持为(bitmask)不变；
    3.对于安排位置 6，7，8，9，如果bitmask 中第 4，5，6，7个二进制位均为 0，那么就可以安排给一个家庭；也就是说，bitmask 和(00001111)的按位或值保持为(bitmask)不变；
    */
    var left=240//11110000
    var mid=195//11000011
    var right=15//00001111
    var fre=make(map[int]int)//key is row value is bit mask
    //fre统计的是每一行 至少有一个已经预定
    for i:=range reservedSeats{
        one:=reservedSeats[i]
        if one[1]>1&&one[1]<10{
            fre[one[0]]|=(1<<(one[1]-2))
        }
    }
    //如果整行都没有被订，则可以安排2个家庭
    res+=(n-len(fre))*2
    for _,v:=range fre{
        if (v|left==left)||(v|mid==mid)||(v|right==right){
            res++
        }
    }
    return res
}

func getSegCnt(row []int)int {
    var res int
    sum1:=get4Sum(row[1:5])
    if sum1 <1 {
        res++
        sum2:=get4Sum(row[5:9])
        if sum2<1 {
            res++
        }
    }else{
        //[1-2]位置已经有
        sum2:=get4Sum(row[3:7])
        sum3:=get4Sum(row[5:9])
        if sum2<1 || sum3<1 {
            res++
        }
    }

    return res
}

func get4Sum(row []int)int{
    var res int
    for i:=0;i<4;i++{
        res+=row[i]
    }
    return res
}
