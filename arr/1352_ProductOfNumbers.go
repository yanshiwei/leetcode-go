type ProductOfNumbers struct {
    List []int
}


func Constructor() ProductOfNumbers {
    var m=ProductOfNumbers{
        List:make([]int,0),
    }
    return m
}


func (this *ProductOfNumbers) Add(num int)  {
    this.List=append(this.List,num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
    var cnt int
    var product =1
    for i:=len(this.List)-1;i>=0;i--{
        product*=this.List[i]
        cnt++
        if cnt>=k{
            return product
        }
    }
    return product
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
