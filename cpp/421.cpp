class Trie{
    /*
    题目：
    给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 
    */
    private:
        Trie*next[2]={nullptr};
    public:
        Trie(){}
        // 在前缀树中插入值x
        void insert(int x){
            Trie*root=this;
            // 高位存储来Trie的前面，所以我们从左向右存储
            for(int i=30;i>=0;i--){
                // 取第i位的数字，30...0
                int u=(x>>i)&1;
                // 若第u位为空，则创建一个新节点，然后root移动到下一个节点
                if(root->next[u]==nullptr){
                    root->next[u]=new Trie();
                }
                root=root->next[u];
            }
        }
        // 在前缀树中寻找 x 的最大异或值
        int search(int x){
            Trie*root=this;
            int res=0;
            for(int i=30;i>=0;i--){
                // 取第i位的数字，30...0
                int u=(x>>i)&1;
                // 若 x 的第 u 位存在，我们走到相反的方向去故root->next[!u]，因为异或总是相反才取最大值的
                if(root->next[!u]){
                    root=root->next[!u];
                    res=res*2+!u;//高位存储来Trie的前面，故每处理一次，res*2表示左移一位，左移后开始值为0，+!u表示加上当前的最低位数字
                }else{
                    root=root->next[u];
                    res=res*2+u;
                }
            }
            res^=x;
            return res;
        }
};
class Solution {
public:
    int findMaximumXOR(vector<int>& nums) {
        Trie*root=new Trie();
        for(auto x:nums){
            root->insert(x);
        }
        int res=0;
        // 数组每一个都去搜索，遍历获得最大值
        for(auto x:nums){
            res=max(res,root->search(x));
        }
        return res;
    }
};
