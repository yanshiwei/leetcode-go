class Solution {
public:
/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
*/
    void maxHeapUpdate(vector<int>&nums,int parent, int last_node){
        // 临时保存要下沉的元素
        int tmp=nums[parent];
        //定位左孩子节点的位置
        int child=2*parent+1;
        // 开始下沉
        while(child<=last_node){
            // 如果右孩子节点比左孩子大，则定位到右孩子
            if (child+1<=last_node&&nums[child]<nums[child+1]){
                child++;
            }
            // 如果孩子节点小于或等于父节点，则下沉结束
            if(nums[child]<tmp){
                break;
            }
            // 父节点进行下沉
            nums[parent]=nums[child];
            parent=child;
            child=2*parent+1;
        }
        nums[parent]=tmp;
    }
    void buildMAxHeap(vector<int>&nums,int heapSize){
        int parent=(heapSize)/2-1;
        for(int i=parent;i>=0;i--){
            maxHeapUpdate(nums,i,heapSize-1);
        }
    }
    int findKthLargest(vector<int>& nums, int k) {
        int heapSize=nums.size();
        buildMAxHeap(nums,heapSize);
        // 建立一个大根堆，做 k - 1次删除操作后堆顶元素就是我们要找的答案
        for(int i=0;i<k-1;i++){
            int t=nums[0];
            nums[0]=nums[nums.size()-1-i];
            nums[nums.size()-1-i]=t;
            maxHeapUpdate(nums,0,nums.size()-1-i-1); // 堆排序0到length-1-i-1
        }
        return nums[0];
    }
};
