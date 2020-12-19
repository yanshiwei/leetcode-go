const INTMAX=int(^uint(0) >> 1)
func nthUglyNumber(n int) int {

    var i int
    /*
    //暴力超时
    for index<n{
        if isUgly(num){
            index++
        }
    }
    return index
    */
    //因为丑数都是2,3,5的乘积，大的丑数可由小的丑数分别乘以2,3,5得到。所以可以使用一个最小堆不断更新，每第i次更新，得到对应第i大的丑数。每一次更新，从最小堆中pop出一个最小值，就是第i大的丑数（但是要注意处理掉重复的丑数）
    var heap []int
    heap=append(heap,1)
    for i=1;i<n;i++ {
        cur:=heap[0]
        heap=heap[1:]
        buildHeap(heap)
        for len(heap)>0&&cur==heap[0] {
            heap=heap[1:]
            buildHeap(heap)
        }
        if cur*2<INTMAX {
                heap=append(heap,cur*2)
                buildHeap(heap)
            }
        if cur*3<INTMAX {
                heap=append(heap,cur*3)
                buildHeap(heap)
        }
        if cur*5<INTMAX {
                heap=append(heap,cur*5)
                buildHeap(heap)
        }
    }
    return heap[0]
}

func adjustHeap(heap []int, start int, length int) {
	var i=start
	var child int
	for {
		child=2*i+1
		if child>=length-1 {
			break
		}
		if child+1<=length-1&&heap[child]>heap[child+1] {
			child+=1
		}
		if heap[i]>heap[child] {
			heap[i],heap[child]=heap[child],heap[i]
			i=child//child update
		}else {
			break
		}
	}
}
func buildHeap(heap []int) {
	//half before
	for i:=len(heap)/2-1;i>=0;i-- {
		adjustHeap(heap,i,len(heap))
	}
}
func isUgly(x int)bool {
    if x<=0 {
        return false
    }
    if x==1 {
        return true
    }
    //去掉2的因子
    for 0==x%2 {
        //能被2整除
        x/=2
    }
    //去掉3的因子
    for 0==x%3 {
        //能被3整除
        x/=3
    }
    //去掉5的因子
    for 0==x%5 {
        //能被5整除
        x/=5
    }
    if x==1 {
        return true
    }
    return false
}
