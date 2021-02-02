type RandomizedCollection struct {
    ValueMapIndex map[int]map[int]struct{}//key is data value is data index set
    List []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
    m:=RandomizedCollection{
        ValueMapIndex:make(map[int]map[int]struct{}),
    }
    return m
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
    var flag bool
    if _,ok:=this.ValueMapIndex[val];ok==false {
        flag=true
        indexSet:=make(map[int]struct{})
        this.ValueMapIndex[val]=indexSet
    }
    this.List=append(this.List,val)
    this.ValueMapIndex[val][len(this.List)-1]=struct{}{}
    return flag
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
   var flag bool
   if indexSet,ok:=this.ValueMapIndex[val];ok==true {
       flag=true
       last:=len(this.List)-1//最后一个元素的索引
       var idx int//待删除元素索引，随机选择一个
       for id:=range indexSet {
           idx=id
           break
       }
       this.List[idx]=this.List[last]//用最后一个元素覆盖idx位置
       delete(indexSet,idx)//删除待删除元素索引
       delete(this.ValueMapIndex[this.List[idx]],last)//删除新的元素（即最后一个元素）对应的原索引
       if len(indexSet)<1 {
           //若待删除元素弹出之后，其索引集合为空，则删除其索引集合
           delete(this.ValueMapIndex,val)
       }
       if idx<len(this.List)-1 {
           //添加新的元素（即最后一个元素）对应的新索引
            if _,ok:=this.ValueMapIndex[this.List[idx]];ok==false {
            indexSet:=make(map[int]struct{})
            this.ValueMapIndex[this.List[idx]]=indexSet
            }
           this.ValueMapIndex[this.List[idx]][idx]=struct{}{}
       }
       this.List=this.List[0:len(this.List)-1]//pop
   }
   return flag
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
    var m =len(this.List)
    if m==1 {
        return this.List[0]
    }
    rand.Seed(time.Now().UnixNano()) 
    x := rand.Intn(m)   //生成0-m随机整数
    return this.List[x]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
