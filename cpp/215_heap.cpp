class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        return findKthByMaxHeapSort(nums,k);
    }
private:
    // 构造k大小堆，反复比较
    int findKthByMinHeap(vector<int>& nums, int k){
        priority_queue<int,vector<int>,greater<int>>minHeap;
        for(int i=0;i<nums.size();i++){
            if(i<k){
                // 构造大小k的堆
                minHeap.push(nums[i]);
            }else{
                // 比较堆顶与当前元素大小
                if(minHeap.top()>nums[i]){
                    continue;
                }else{
                    minHeap.pop();
                    minHeap.push(nums[i]);
                }
            }
        }
        // min heap 保留了最大K个，其中堆顶就是第k个
        return minHeap.top();
    }
    // 堆排序，k次
    void sift(vector<int>&nums,int start,int end){
        int child;
        int i=start;
        while(1){
            child=2*i+1;
            if(child>=end){
                break;
            }
            if(child+1<end&&nums[child]<nums[child+1]){
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
    // 构造堆
    void buildHeap(vector<int>& nums){
        // 从第一个非叶子结点开始
        for(int i=nums.size()/2-1;i>=0;i--){
            sift(nums,i,nums.size());
        }
    }
    int findKthByMaxHeapSort(vector<int>& nums, int k){
        // 构建大顶堆
        buildHeap(nums);
        // 交换堆顶和末位元素k-1次后，堆顶元素就是第k大
        for(int i=0;i<k-1;i++){
            swap(nums[0],nums[nums.size()-1-i]);
            sift(nums,0,nums.size()-1-i);
        }
        return nums[0];
    }
};
