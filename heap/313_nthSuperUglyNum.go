func nthSuperUglyNumber(n int, primes []int) int {
    /*
    var minHeap[]int
    minHeap=append(minHeap,1)
    for i:=1;i<n;i++ {
        cur:=minHeap[0]
        minHeap=minHeap[1:]
        buildHeap(minHeap)
        for len(minHeap)>0&&minHeap[0]==cur{
            minHeap=minHeap[1:]
            buildHeap(minHeap)
        }
        for j:=range primes{
            if cur*primes[j]<INTMAX {
                minHeap=append(minHeap,cur*primes[j])
                buildHeap(minHeap)
            }
        }
    }
    return minHeap[0]
    */
    //way like 264 is time limited use db
    //db数组存储的树丑树序列如[1,2,4,7,8,13...] index数组与Primes数组一一对应，长度一致，
    //即index[j]代表的树primes[j]应该与丑数序列中的第index[j]相乘 
    var dp=make([]int,n)
    var index=make([]int,len(primes))
    dp[0]=1
    for i:=1;i<n;i++ {
        minV:=INTMAX
        for j:=0;j<len(primes);j++{
            if minV>dp[index[j]]*primes[j] {
                minV=dp[index[j]]*primes[j]
            }
        }
        dp[i]=minV
        for j:=0;j<len(primes);j++ {
            if minV==dp[index[j]]*primes[j] {
                //去除重复的影响 取最近一次
                index[j]++
            }
        }
    }
    return dp[n-1]
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
const INTMAX=int(^uint(0) >> 1)
func buildHeap(heap []int) {
	//half before
	for i:=len(heap)/2-1;i>=0;i-- {
		adjustHeap(heap,i,len(heap))
	}
}
