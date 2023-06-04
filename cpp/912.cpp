class Solution {
public:
    vector<int> sortArray(vector<int>& nums) {
        // solve 1
        // heapSort(nums);
        // solve 2
        make_heap(nums.begin(),nums.end());// default is big heap
        sort_heap(nums.begin(),nums.end());
        return nums;
    }
private:
    void constructHeap(vector<int>&nums){
        // 从最后一个非叶子节点len/2-1开始调整，从下到上
        for(int i=nums.size()/2-1;i>=0;i--){
            adjustHeap(nums,i,nums.size());
        }
    }
    void adjustHeap(vector<int>&nums,int start,int end){
        // 每次调整使得满足堆定义
        int child=0;
        int i=start;
        while(1){
            child=2*i+1;
            if(child>end-1){
                break;
            }
            if(child+1<=end-1&&nums[child]<nums[child+1]){
                // find bigger child
                child+=1;
            }
            if(nums[i]<nums[child]){
                swap(nums[i],nums[child]);
                i=child; //交换后，以child为根的子树不一定满足堆定义，所以从child处开始调整
            }else{
                break;
            }
        }
    }
    void heapSort(vector<int>&nums){
        // 构建大顶堆
        constructHeap(nums);
        // 把根节点跟最后一个元素交换位置，调整剩下的n-1个节点，即可排好序
        for(int i=nums.size()-1;i>=0;i--){
            swap(nums[0],nums[i]);
            adjustHeap(nums,0,i);
        }
    }
};
