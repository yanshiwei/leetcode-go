type RandomizedSet struct {
    ValueMapIndex map[int]int
    List []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    m:=RandomizedSet{
        ValueMapIndex:make(map[int]int),
        List:make([]int,0),
    }
    return m
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _,ok:=this.ValueMapIndex[val];ok==false {
        this.List=append(this.List,val)
        this.ValueMapIndex[val]=len(this.List)-1
        return true
    }
    return false
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    //fmt.Println(this.List,this.ValueMapIndex)
    if idx,ok:=this.ValueMapIndex[val];ok==true {
        lastV:=this.List[len(this.List)-1]
        this.List[idx]=lastV
        this.ValueMapIndex[lastV]=idx
        this.List=this.List[0:len(this.List)-1]//pop
        delete(this.ValueMapIndex,val)
        return true
    }
    return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    var m =len(this.List)
    if m==1 {
        return this.List[0]
    }
    rand.Seed(time.Now().UnixNano()) 
    x := rand.Intn(m)   //生成0-m随机整数
    return this.List[x]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
