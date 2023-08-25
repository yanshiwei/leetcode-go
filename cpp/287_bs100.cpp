class Solution {
    /*
    二分查找不是对原数组查找，而是对1,2,...n-1，这n-1个连续的数构成的升序数组进行查找
    依据什么属性查找？我们知道，一个数组在某个属性上有单调性，就可以进行二分查找（比如排序数组的数组值）。那升序的数组对于本题来说还有什么单调性呢。这个单调性就是: 数组的索引越大，数组值就越大，数组值越大，原数组中小于等于这个数组值的个数就越多。
    例如arr=[1,3,4,2,3]
    <=1的个数 1个
    <=2的个数 2个
    <=3的个数 4个
    <=4的个数 5个
    所以一个数越大，原数组中小于等于这个数就越多，我们记这个属性为cnt,则现在我们知道，对于【1，2，3,...n-1】这个数组值升序数组来说，数组的cnt属性值也是升序的
    对于目标target，它的cnt属性有什么特点呢。因为它在原数组中是重复的，那就说明一点：target的cnt属性值是大于target
    if num<target,则num的cnt属性值<=num
    if num>=target,则num的cnt属性值>num
    根据这个结论，我们就可以进行二分查找

    */
public:
    int findDuplicate(vector<int>& nums) {
        int n=nums.size();
        int left=1;
        int right=n-1;
        while(left<right){
            int mid=left+(right-left)/2;
            int cnt=countRange(nums,mid);
            if(cnt>mid){
                // 说明mid>=target，重复的元素会在[left,mid]
                right=mid;
            }else {
                // 说明mid<target，重复的元素会在[mid+1,right]
                left=mid+1;
            }
        }
        return left;
    }
private:
    int countRange(vector<int>& nums,int mid){
        int cnt=0;
        for(int num :nums){
            if(num<=mid){
                cnt++;
            }
        }
        return cnt;
    }
};
