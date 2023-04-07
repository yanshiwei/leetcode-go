class Solution {
    /*
    假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ：
perm[i] 能够被 i 整除
i 能够被 perm[i] 整除
给你一个整数 n ，返回可以构造的 优美排列 的 数量 。
    */
public:
    unordered_set<int>visited;
    int ans=0;
    void dfs(int n,int idx){
        if(idx==n+1){
            //当找到n个符合要求的序列时，满足要求，ans++
            ans++;
        }
        for(int i=1;i<=n;i++){
            if(visited.count(i)<1&&(i%idx==0||idx%i==0)){
                visited.insert(i);//标记为遍历过
                dfs(n,idx+1);//寻找第index+1个满足要求的元素
                visited.erase(i);//回溯
            }
        }
    }
    int countArrangement(int n) {
        dfs(n,1);
        return ans;
    }
};
