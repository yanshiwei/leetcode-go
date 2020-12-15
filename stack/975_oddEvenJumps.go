type Info struct {
	Data int
	Idx  int
}

func getNextBigger(arr []int) []int {
	var res []int
	if len(arr) < 1 {
		return res
	}
	var stackIndex []int //单减栈
	res = make([]int, len(arr))
	for i := range arr {
		for len(stackIndex) > 0 && arr[i] > arr[stackIndex[len(stackIndex)-1]] {
			head := stackIndex[len(stackIndex)-1] //pop//单调减get right套路
			//fmt.Println("head",arr[head],arr[i])
			res[arr[head]]=arr[i]//res[head] = arr[i] 不用记录本身下标，记录输入的下标                           //单调减get right套路
			stackIndex=stackIndex[0:len(stackIndex)-1]//pop//单调减套路
		}
		stackIndex = append(stackIndex, i)
	}
	return res
}
func mergeSort(arr []Info) []Info {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func merge(left, right []Info) []Info {
	result := make([]Info, 0)
	m, n := 0, 0 // left和right的index位置
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m].Data > right[n].Data {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}
func QuickSort(arr []Info, l, r int) {
	if l < r {
		pivot := arr[r].Data
		i := l - 1
		for j := l; j < r; j++ {
			if arr[j].Data <= pivot {
				i++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		i++
		arr[r], arr[i] = arr[i], arr[r]
		QuickSort(arr, l, i-1)
		QuickSort(arr, i+1, r)
	}
}

func getSortBeforeIdx(arr []int) []int {
	var infoList []Info
	var indexRes []int
	for i := range arr {
		info := Info{Data: arr[i], Idx: i}
		infoList = append(infoList, info)
	}
    //不能用快排，因为是非稳定性排序
	//QuickSort(infoList, 0, len(infoList)-1)
    //选择排序、快速排序、希尔排序、堆排序不是稳定的排序算法，而冒泡排序、插入排序、归并排序和基数排序是稳定的排序算法
    infoList=mergeSort(infoList)
	for i := range infoList {
		info := infoList[i]
		indexRes = append(indexRes, info.Idx)
	}
	return indexRes
}

func oddEvenJumps(A []int) int {
    //奇数跳，a[i]<=a[j],先找到比它大或等于的数里面最小的，如果这样最小值有多个取最靠前的
	//非严格单减栈，找到最靠前的比它大或等的数
	//可以先对a从小到大排序，并得到排序前下标，然后通过单减栈获取最靠前的大于或等于的这样的最小值
	//偶数跳，a[i]>=a[j],先找到比它小或等于的数里面最大的，如果这样最大值有多个取最靠前的
	//乘以-1转换，则-a[i]<=-a[j],先找到比它大或等于的数里面最小的，如果这样最小值有多个取最靠前的
	//类似奇数跳做法

	indexListA := getSortBeforeIdx(A)//从小到大排序 并获取排序前下标
	biggerListA := getNextBigger(indexListA) //单减栈获取比该下标i第一个大的下标j（也就是最靠前），按照排序结果，下标j对应的a[j]>=a[i]，且a[j]肯定是最小的
	//A1=getNextBigger(A) A1[i]代表比A[i]大的第一个A[j]值
	fmt.Println("a",A, indexListA,biggerListA)
	var negativeA []int
	for i := range A {
		negativeA = append(negativeA, -1*A[i])
	}
	negativeIndexListA := getSortBeforeIdx(negativeA)
	biggerNegListA := getNextBigger(negativeIndexListA)
	fmt.Println("b",negativeA, negativeIndexListA,biggerNegListA)
	var oddJump = make([]int, len(A))  //奇数跳,oddJump[i] - if now it odd jumps from index i,
	// whether it could reach the end.
	var evenJump = make([]int, len(A)) //偶数跳
	oddJump[len(oddJump)-1] = 1            //末尾总能跳到
	evenJump[len(evenJump)-1] = 1          //末尾总能跳到
	for i := len(A)-2; i >= 0; i-- {
		evenJump[i] = oddJump[biggerListA[i]] //从i能否奇数跳到末尾，取决于它奇数跳的目标j也即j=biggerListA[i]能否偶数跳到末尾
		oddJump[i] = evenJump[biggerNegListA[i]]//从i能否偶数跳到末尾，取决于它偶数跳的目标j也即j=biggerNegListA[i]能否奇数跳到末尾
	}
	fmt.Println("res",evenJump,oddJump)
	var sum int
	for i := range evenJump {
		sum += evenJump[i]
	}
    return sum
}
