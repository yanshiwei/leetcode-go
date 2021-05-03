func getNumberOfBacklogOrders(orders [][]int) int {
    var sellMinHeap=&MinHeap{}
    var buyMaxHeap=&MaxHeap{}
    heap.Init(sellMinHeap)
    heap.Init(buyMaxHeap)
    for i:=range orders{
        price:=orders[i][0]
        cnt:=orders[i][1]
        ctype:=orders[i][2]
        if ctype==0{
            //buy，找最小的sell
            for cnt>0&&sellMinHeap.Len()>0{
                if sellMinHeap.Top().Price<=price{
                    //sell的价格小。能买到
                    if cnt>sellMinHeap.Top().Cnt{
                        //购买的数目大于销售的数目
                        cnt-=sellMinHeap.Top().Cnt
                        //卖完
                        heap.Pop(sellMinHeap)
                    }else if cnt==sellMinHeap.Top().Cnt{
                        //刚好支出平衡
                        cnt=0
                        //卖完
                        heap.Pop(sellMinHeap)
                    }else{
                        //买完，堆里还能剩下sell
                        resCnt:=sellMinHeap.Top().Cnt-cnt
                        cnt=0
                        newInfo:=sellMinHeap.Top()
                        newInfo.Cnt=resCnt
                        heap.Pop(sellMinHeap)
                        heap.Push(sellMinHeap,newInfo)
                    }
                }else{
                    break
                }
            }
            if cnt>0{
                //队列里没有匹配的销售sell，buy入队列
                heap.Push(buyMaxHeap,Info{Price:price,Cnt:cnt,Ctype:ctype})
            }
        }else{
            //sell，找最大的buy
            for cnt>0&&buyMaxHeap.Len()>0{
                if buyMaxHeap.Top().Price>=price{
                    //buy的价格大，能买到
                    if cnt>buyMaxHeap.Top().Cnt{
                        //销售的数目大于购买的数目
                        cnt-=buyMaxHeap.Top().Cnt
                        heap.Pop(buyMaxHeap)
                    }else if cnt==buyMaxHeap.Top().Cnt{
                        //支出平衡
                        cnt=0
                        heap.Pop(buyMaxHeap)
                    }else{
                        //销售的数目小于购买的收入
                        resCnt:=buyMaxHeap.Top().Cnt-cnt
                        cnt=0
                        newInfo:=buyMaxHeap.Top()
                        newInfo.Cnt=resCnt
                        heap.Pop(buyMaxHeap)
                        heap.Push(buyMaxHeap,newInfo)
                    }
                }else{
                    break
                }
            }
            if cnt>0{
                //队列里没有匹配的buy，sell入队列
                heap.Push(sellMinHeap,Info{Price:price,Cnt:cnt,Ctype:ctype})
            }
        }
    }
    var res int
    for sellMinHeap.Len()>0{
        res+=sellMinHeap.Top().Cnt
        res%=1000000007
        heap.Pop(sellMinHeap)
    }
    for buyMaxHeap.Len()>0{
        res+=buyMaxHeap.Top().Cnt
        res%=1000000007
        heap.Pop(buyMaxHeap)
    }
    return res
}

type Info struct {
    Price int
    Cnt int
    Ctype int
}

type MaxHeap []Info
func (h MaxHeap)Len()int{return len(h)}
func (h MaxHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h MaxHeap)Less(i,j int)bool{return h[i].Price>h[j].Price}
func (h MaxHeap)Top()Info{return h[0]}
func (h*MaxHeap)Push(x interface{}){*h=append(*h,x.(Info))}
func (h*MaxHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}

type MinHeap []Info
func (h MinHeap)Len()int{return len(h)}
func (h MinHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h MinHeap)Less(i,j int)bool{return h[i].Price<h[j].Price}
func (h MinHeap)Top()Info{return h[0]}
func (h*MinHeap)Push(x interface{}){*h=append(*h,x.(Info))}
func (h*MinHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
