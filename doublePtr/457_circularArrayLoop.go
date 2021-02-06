
func circularArrayLoop(nums []int) bool {
    /*
    判断是否有环，快慢指针。快慢指针的思想是慢指针走一步，快指针走两步，若他俩相遇肯定是在环中的某一个节点上相遇，则证明存在环。假定慢指针为j快指针为k，那么一定会有j==k的时候。
    由于题目中说了 nums 中不会有0,故如果遇到已经访问过，被标记的地方nums[i]=0即可
    */
    for i:=0;i<len(nums);i++ {
        if nums[i]==0 {
            //已经访问过
            continue
        }
        slow:=i
        fast:=getNext(nums,slow)
        val:=nums[slow]
        for val*nums[fast]>0&&val*nums[getNext(nums,fast)]>0 {
            //由于环内都是同一个方向，且要求环长度大于1故：
            //val*nums[fast]表示慢指针指向的数要和快指针指向的数正负相同
            //val*nums[getNext(nums,fast)]并且慢指针指向的数还要跟快指针的下一个位置上的数符号相同
            if slow==fast {
                //有环
                //直到当快慢指针相遇的时候，就是环出现的时候，
                //即便快慢指针相遇了，也不同立马返回 true，因为题目中说了了环的长度必须大于1，所以我们要用慢指针指向的数和慢指针下一个位置上的数比较：
                //若相同，则说明环的长度为1，此时并不返回 false，，因为这只能说以i位置开始的链表无符合要求的环而已，后面可能还会出现符合要求的环，break 掉循环；
                //但是若二者不相同的话，则已经找到了符合
                if slow==getNext(nums,fast){
                    //长度为1
                    break
                }
                return true//存在长度大于1的环
            }
             //若快慢指针还不相同的，则分别更新
             slow=getNext(nums,slow)//一步
             fast=getNext(nums,getNext(nums,fast))//两步
        }
        //针对i的循环推出后，需要标记已经走过的点
        slow=0
        for val*nums[slow]>0{
            next:=getNext(nums,slow)
            nums[slow]=0
            slow=next
        }
    }
    return false
}
//计算 next 位置，
//对于会超出数组的长度的正数，我们可以通过对n取余
//但是对于负数，若这个负数远大于n的话，取余之前只加上一个n，可能是不够的，
//所以正确的方法是应该先对n取余，再加上n。
//为了同时把正数的情况也包含进来，最终我们的处理方法是先对n取余，
//再加上n，再对n取余，这样不管正数还是负数，大小如何，都可以成功的旋转跳跃了。
func getNext(nums[]int, start int)int {
    var m=len(nums)
    first:=(nums[start]+start)%m//先取余数
    second:=(first+m)%m//加上m再又取余数，应对负数情况
    return second
}
