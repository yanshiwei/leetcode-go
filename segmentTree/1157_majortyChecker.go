type MajorityChecker struct {
    arr []int
    ma map[int][]int
    t []*Info
}
/*
2 * threshold > right - left + 1 可知出现次数大于一半的数字最多只可能有1个，如果有多个一样次数的，肯定就不是我们要找的结果。如果有 2 个数量相等，则肯定不是我们需要的，取任意一个都可以。
*/

func Constructor(arr []int) MajorityChecker {
    mc:=MajorityChecker{}
    mc.arr=arr
    mc.ma=make(map[int][]int)
    for i:=range arr {
        mc.ma[arr[i]]=append(mc.ma[arr[i]],i)
    }
    n:=len(arr)
    mc.t=make([]*Info,calTreeSize(n))
    build(arr,mc.t,0,n-1,0)
    for i:=range mc.t{
        if mc.t[i]==nil{
            continue
        }
        fmt.Println(mc.t[i].Value,mc.t[i].Cnt)
    }
    return mc
}

func calTreeSize(arrLen int) (size int) {
	for i := 1; i>>1 < arrLen; i <<= 1 {
		size += i
	}
	return
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
    res:=query(this.t,0,len(this.arr)-1,left,right,0)
    if res==nil || res.Value==0{
        return -1
    }
    //Search函数采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i。
    l:=sort.Search(len(this.ma[res.Value]),func(i int)bool{
        return left<=this.ma[res.Value][i]
    })
    r:=sort.Search(len(this.ma[res.Value]),func(i int)bool{
        return right<this.ma[res.Value][i]
    })
    if r-l<threshold{
        return -1
    }
    return res.Value
}

type Info struct{
    Value int
    Cnt int
}

func build(arr []int, t[]*Info, left,right,idx int){
    if left==right{
        t[idx]=&Info{Value:arr[left],Cnt:1}
        return
    }
    var mid=left+(right-left)/2
    childLeft:=idx<<1+1
    childRight:=childLeft+1
    build(arr,t,left,mid,childLeft)
    build(arr,t,mid+1,right,childRight)
    t[idx]=addSegTree(t[childLeft],t[childRight])
}

func addSegTree(left,right *Info)*Info{
    if left==nil && right==nil{
        return nil
    }
    if left==nil{
        return right
    }
    if right==nil{
        return left
    }
    var val,cnt int
    if left.Value==right.Value{
        val,cnt=left.Value,left.Cnt+right.Cnt
    }else if left.Cnt<right.Cnt{
        val,cnt=right.Value,right.Cnt-left.Cnt
    }else{
        val,cnt=left.Value,left.Cnt-right.Cnt
    }
    //这里减去，并不是找到当前段出现次数最多的数字，而是找到最有可能满足要求的值。比如[1 1 2 2 3]， 左边[1 1 2 2 ]计算出的结果为[1 0],右边为[3 1]，最后结合是[3 1]，3 并不是出现次数最多的数。这里主要是因为 1 的数量和2相同，所以如果有任意不是 1 和 2 的数字，那么 1 和 2 都不可能成为大于等于一半的数，所以可以不用考虑
    return &Info{
        Value:val,
        Cnt:cnt,
    }
}

func query(t []*Info,left,right,queryLeft,queryRight,idx int)*Info{
    if queryLeft>right||left>queryRight{
        return nil
    }
    if queryLeft<=left&&right<=queryRight{
        ////如果本区间完全在操作区间[L,R]以内
        return t[idx]
    }
    var mid=left+(right-left)>>1
    childLeft:=idx<<1+1
    childRight:=childLeft+1
    return addSegTree(query(t,left,mid,queryLeft,queryRight,childLeft),query(t,mid+1,right,queryLeft,queryRight,childRight))
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
