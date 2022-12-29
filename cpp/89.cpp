class Solution {
    /*
    https://leetcode.cn/problems/gray-code/solutions/1198532/jin-dao-ge-lei-ma-shi-zhe-yang-wan-de-by-qrx4/?page=2
    总结规律：
    1位格雷码有两个码字
    (n+1)位格雷码中的前2^n个码字等于n位格雷码的码字，按顺序书写，加前缀0
    (n+1)位格雷码中的后2^n个码字等于n位格雷码的码字，按逆序书写，加前缀1
    n+1位格雷码的集合 = n位格雷码集合(顺序)加前缀0 + n位格雷码集合(逆序)加前缀1
    */
public:
    vector<int> grayCode(int n) {
        vector<int>res;
        res.push_back(0);
        res.push_back(1);
        int mask=1<<1; // 10
        for(int i=2;i<=n;i++){
            // 倒序访问 高位置1
            for(int j=res.size()-1;j>=0;j--){
                res.push_back(mask|res[j]);
            }
            mask=mask<<1;// 位置多1，就左移1
        }
        return res;
    }
};
