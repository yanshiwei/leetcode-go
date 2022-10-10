class Solution {
    /*
    编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
    每行中的整数从左到右按升序排列。
    每行的第一个整数大于前一行的最后一个整数。
    */
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int n=matrix.size();
        int m=matrix[0].size();
        // 将这个二维矩阵看成一个一维有序数组，那么一维有序数组的边界就是0-- n*m-1
        // 一维有序数组下标k=i*cols+j;
        int left=0,right=n*m-1;
        while(left<=right){
            int mid=left+(right-left)/2;
            if(matrix[mid/m][mid%m]==target){
                return true;
            }else if(matrix[mid/m][mid%m]>target){
                right=mid-1;
            }else{
                left=mid+1;
            }
        }
        return false;
    }
};
