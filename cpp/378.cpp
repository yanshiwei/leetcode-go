class Solution {
    /*
    给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
    */
public:
// priority_queue push pop top size empty
    int kthSmallest(vector<vector<int>>& matrix, int k) {
        priority_queue<int,vector<int>,less<int>>q;// 大顶堆
        int n=matrix.size();
        for(int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                if(q.size()<k){
                    q.push(matrix[i][j]);
                }else{
                    auto val=q.top();
                    if(val>matrix[i][j]){
                        q.pop();
                        q.push(matrix[i][j]);
                    }else{
                        //基于已排序这个现实 如matrix[i][j]更大，则j后面的更大 更没有必要比较
                        break;
                    }
                }
            }
        }
        return q.top();// 第k小
    }
};
