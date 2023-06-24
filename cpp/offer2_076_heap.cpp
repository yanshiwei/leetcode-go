class Solution {
    // 方法1，小堆顶，priority_queue<int, vector<int>, greater<int>> heap;
    // 为什么是小顶堆求最大TOP K
    // 小顶堆大小K，且最小值是堆顶元素
    // 其他元素要入堆时，先和堆顶元素比较
    // 1 若当前元素比堆顶元素小，则跳过
    // 2 若当前元素比堆顶元素大，则替换堆顶元素并调整
    // 由于堆顶元素是最小值，而比它更小的不会插入，故堆里是最大TOPK
    // 方法2，大堆顶，自写堆排序
    // 方法3，小顶堆，C++自带方法实现或者
// make_heap()用于把一个可迭代容器变成一个堆，默认是大顶堆（less<T>）
    // push_heap()是在堆的基础上进行数据的插入操作，push_back之后操作，只会把最后一个元素（堆底）加入堆结构
    // pop_heap()用于将堆的第零个元素与最后一个元素交换位置，然后针对前n - 1个元素调用make_heap()函数，如果要删除这个元素，还需要对数组进行pop_back()
    // sort_heap()是将堆进行排序，排序后，序列将失去堆的特性
public:
    int findKthLargest(vector<int>& nums, int k) {
        return minHeapPriorityQueue(nums,k);
        // return maxHeapSortKth(nums,k);
    }
private:
    // 方法1
    int minHeapPriorityQueue(vector<int>& nums, int k){
        priority_queue<int,vector<int>,greater<int>>minHeap;
        for(int i=0;i<nums.size();i++){
            // 构造大小k的小顶堆
            if(i<k){
                minHeap.push(nums[i]);
            }else {
                // 大于堆顶则替换
                if(nums[i]>minHeap.top()){
                    minHeap.pop();
                    minHeap.push(nums[i]);
                }else{
                    continue;
                }
            }
        }
        return minHeap.top();//小顶堆存储的最大TOPK，其中堆顶生第K个
    }
    // 方法3
    void maxHeapUpdate(vector<int>&nums,int start, int end){
        int child;
        int i=start;
        while(1){
            child=2*i+1;
            if(child>=end){
                break;
            }
            if(child+1<end&&nums[child+1]>nums[child]){
                child+=1;
            }
            if(nums[i]<nums[child]){
                swap(nums[i],nums[child]);
                i=child;
            }else{
                break;
            }
        }
    }
    void maxHeapBuild(vector<int>&nums){
        // 从最后一个非叶子节点len/2-1开始调整，从下到上
        for(int i=nums.size()/2-1;i>=0;i--){
            maxHeapUpdate(nums,i,nums.size());
        }
    }
    int maxHeapSortKth(vector<int>&nums, int k){
        // 构建大顶堆
        maxHeapBuild(nums);
        // 把根节点跟最后一个元素交换位置，调整剩下的n-1个节点，即可排好序
        // k-1次这样操作就可以 堆顶就是第K大元素
        for(int i=0;i<k-1;i++){
            swap(nums[0],nums[nums.size()-1-i]);
            maxHeapUpdate(nums,0,nums.size()-1-i);
        }
        return nums[0];
    }
};
