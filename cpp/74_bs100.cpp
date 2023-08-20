class Solution {
    //将这个二维矩阵看成一个一维有序数组，
    // 那么一维有序数组的边界就是0-- n*m-1
    //一维数组下标k=i*cols+j
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int n=matrix.size();
        if(n<1){
            return false;
        }
        int m=matrix[0].size();
        if(m<1){
            return false;
        }
        int left=0;
        int right=n*m-1;//区间左闭右闭
        while(left<=right){
            int mid=left+(right-left)/2;
            if(matrix[mid/m][mid%m]==target){
                return true;
            }else if(matrix[mid/m][mid%m]>target){
                // in [left,mid-1]
                right=mid-1;//因为区间左闭右闭，故mid-1才不会重复搜索
            }else{
                // in [mid+1,right]
                left=mid+1;
            }
        }
        return false;
    }
};
