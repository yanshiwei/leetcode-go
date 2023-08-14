class Solution {
    //数组从左到右和从上到下都是升序的，如果从右上角出发开始遍历
    // 向左，数字变小，向下数字变大，类似二分查找。二分查找树的话，是向左数字变小，向右数字变大。
    // 故以此类推：
    // 当前值与target比较：
    // 如果 target 的值大于当前值，那么就向下走
    // 如果 target 的值小于当前值，那么就向左走
    // 如果相等的话，直接返回 true
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int m=matrix.size();
        if(m<1){
            return false;
        }
        int n=matrix[0].size();
        if(n<1){
            return false;
        }
        int i=0;
        int j=n-1;
        // 右上角
        while(i<m&&j>=0){
            if(matrix[i][j]==target){
                return true;
            }else if(matrix[i][j]>target){
                // find in left
                j--;
            }else {
                // find in down
                i++;
            }
        }
        return false;
    }
};
