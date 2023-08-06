class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        int m=matrix.size();
        if(m<1){
            return {};
        }
        int n=matrix[0].size();
        if(n<1){
            return {};
        }
        // 初始化矩阵 左、右、上、下 四个边界 l , r , t , b ，用于打印的结果列表 res
        int l=0,r=n-1,t=0,b=m-1;
        vector<int>res;
        int idx=0;
        /*循环打印： “从左向右、从上向下、从右向左、从下向上” 四个方向循环，每个方向打印中做        以下三件事 （各方向的具体信息见下表） ；
            1 根据边界打印，即将元素按顺序添加至列表 res 尾部；
            2 边界向内收缩 1 （代表已被打印）；
            3 判断是否打印完毕（边界是否相遇），若打印完毕则跳出。
*/
        while(1){
            // 从左到右
            for(int j=l;j<=r;j++){
                res.push_back(matrix[t][j]);
            }
            t++;
            if(t>b){
                break;
            }
            // 从上到下
            for(int i=t;i<=b;i++){
                res.push_back(matrix[i][r]);
            }
            r--;
            if(r<l){
                break;
            }
            // 从右到左
            for(int j=r;j>=l;j--){
                res.push_back(matrix[b][j]);
            }
            b--;
            if(t>b){
                break;
            }
            // 从下到上
            for(int i=b;i>=t;i--){
                res.push_back(matrix[i][l]);
            }
            l++;
            if(l>r){
                break;
            }
        }
        return res;
    }
};
