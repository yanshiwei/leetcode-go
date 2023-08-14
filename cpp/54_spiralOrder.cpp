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
        int t=0;
        int b=m-1;
        int l=0;
        int r=n-1;
        vector<int> res;
        while(1){
            // 向右移动直到最右
            for(int i=l;i<=r;i++){
                res.push_back(matrix[t][i]);
            }
            t++;
            if(t>b){
                break;
            }
            // 向下移动直到最下
            for(int i=t;i<=b;i++){
                res.push_back(matrix[i][r]);
            }
            r--;
            if(l>r){
                break;
            }
            // 向左移动直到最左
            for(int i=r;i>=l;i--){
                res.push_back(matrix[b][i]);
            }
            b--;
            if(t>b){
                break;
            }
            // 向上移动直到最上
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
