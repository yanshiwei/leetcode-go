func canArrange(arr []int, k int) bool {
    /*
    两个数 x 和 y的和能被 k整除，当且仅当这两个数对 k取模的结果xk
和yk的和就能被 k整除(由于xk,yk取值均小于k,故xk+yk=k)。这里我们规定取模的结果大于等于 0，无论 x 和 y 的正负性。因此，我们将数组arr 中的每个数 xx对 k进行取模，并将得到的余数xk进行配对,配对要求:
     11：如果 xk = 0 ，那么需要找到另一个满足 yk = 0的y 进行配对；
     22：如果 xk > 0，那么需要找到另一个满足 yk = k - xk的 y 进行配对。
     我们可以使用哈希映射（HashMap）统计取模的结果。对于哈希映射中的每个键值对，键表示一个余数，值表示这个余数出现的次数。
     统计要求 11：哈希映射中键 0 对应的值为偶数，参照第一条配对要求；
     统计要求 22：哈希映射中键 t (t>0) 和键 k−t 对应的值相等，参照第二条配对要求
     将负数 x 对一个正数 k 进行取模操作，得到的结果小于等于 0（即在[−k+1,0] 的范围内）。我们可以通过
     xk=(x%k+k)%k得到在 [0,k−1] 的范围内的余数。由于哈希映射中的键的范围为 [0, k-1]，因此我们可以使用一个长度为 kk 的数组代替哈希表，减少编码难度。
    */
    var mod=make(map[int]int)
    for i:=range arr{
        mod[(arr[i]%k+k)%k]++
    }
    for t,cnt:=range mod{
        if t>0{
            if cntAno,ok:=mod[k-t];ok==false{
                return false
            }else{
                if cnt!=cntAno{
                    return false
                }
            }
        }
    }
    if mod[0]%2!=0{
        return false
    }
    return true
}
