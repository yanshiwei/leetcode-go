class Solution {
    /*
    给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
    */
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        vector<int>res;
        int rows=matrix.size();
        if(rows<1){
            return res;
        }
        int cols=matrix[0].size();
        if(cols<1){
            return res;
        }
        // 上下左右边界
        int u=0,d=rows-1,l=0,r=cols-1;
        while(1){
            // 向右移动直到最右
            for(int i=l;i<=r;i++){
                res.push_back(matrix[u][i]);
            }
            u++;// 更新上边界
            if(u>d){
                // 若上边界大于下边界，则遍历完成
                break;
            }
            // 向下移动直到最下
            for(int i=u;i<=d;i++){
                res.push_back(matrix[i][r]);
            }
            r--; //更新右边界
            if(r<l){
                // 若右边界小于左边界，则遍历完成
                break;  
            }
            // 向左移动直到最左
            for(int i=r;i>=l;i--){
                res.push_back(matrix[d][i]);
            }
            d--;// 更新下边界
            if(d<u){
                // 若下边界小于上边界，则遍历完成
                break;     
            }
            // 向上移动直到最上
            for(int i=d;i>=u;i--){
                res.push_back(matrix[i][l]);
            }
            l++; // 更新左边界
            if(l>r){
                // 若左边界大于右边界，则遍历完成
                break;                
            }
        }
        return res;
    }
};
