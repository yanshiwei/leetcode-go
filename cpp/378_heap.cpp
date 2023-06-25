class Solution {
    // 给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
    // 请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
    // 第k小，大顶堆
public:
    int kthSmallest(vector<vector<int>>& matrix, int k) {
        priority_queue<int>maxHeap;//默认大顶堆
        for(int i=0;i<matrix.size();i++){
            for(int j=0;j<matrix[0].size();j++){
                maxHeap.push(matrix[i][j]);
                // 维持堆大小k，最终保留TOPK个最小
                if(maxHeap.size()>k){
                    maxHeap.pop();
                }
            }
        }
        // TOPK个最小 且大顶堆堆顶是第k个最小
        return maxHeap.top();
    }
};
