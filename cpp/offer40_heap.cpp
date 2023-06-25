class Solution {
    /*
    题目：
    输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
    题解：
    最小TOPK，维护大小K的最大堆，堆顶最大值，其他值进行比较，如果其他值更小，则放入，其他值更大则跳过，如此堆里就是最小的K个，其中堆顶的最大
    */
public:
    vector<int> getLeastNumbers(vector<int>& arr, int k) {
        if(k<1){
            return {};
        }
        priority_queue<int,vector<int>,less<int>>maxHeap;
        vector<int>res;
        for(int i=0;i<arr.size();i++){
            if(i<k){
                // 构造大小k的大顶堆
                maxHeap.push(arr[i]);
            }else{
                if(maxHeap.top()>arr[i]){
                    maxHeap.pop();
                    maxHeap.push(arr[i]);
                }else{
                    continue;
                }
            }
        }
        // maxheap store k 
        while(!maxHeap.empty()){
            res.push_back(maxHeap.top());
            maxHeap.pop();
        }
        return res;
    }
};
