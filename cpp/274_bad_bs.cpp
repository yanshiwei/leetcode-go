class Solution {
    // 二分https://leetcode.cn/problems/h-index/solutions/870465/er-fen-cai-lun-wen-pian-shu-java-by-liwe-zoh7/
public:
    int hIndex(vector<int>& citations) {
        //1 3 | 5 7 10 13
        // 引入次数>=4的有4篇，其余2篇引用次数不大于4，故h=4
        // 可以看到h=4时将数组划分2段，满足二分特性
        // 使用二分确定范围：
        // 1 最差情况下，所有的论文被引用次数都为 0；
        // 2 最好情况下，所有的论文的引用次数 >= 总共论文篇数
        // 找先猜一个论文篇数 int mid = (left + right + 1) / 2
        // 如果大于等于 mid 的论文篇数大于等于 mid ，说明 h 指数至少是 mid。例如 [0,1,|3,5,6]，引用次数大于等于 2 的论文有 3 篇，说明答案至少是 2，最终答案落在区间 [mid..right] 里，此时设置 left = mid；
        // 反面的情况不用思考，因为只要上面分析对了，最终答案落在区间 [mid..right] 的反面区间 [left..mid - 1] 里，此时设置 right = mid - 1
        int n=citations.size();
        int left=0;//最差情况下，所有的论文被引用次数都为 0
        int right=n;//最好情况下，所有的论文的引用次数 >= 总共论文篇数
        while(left<right){
            // 若等于，left==right。走到1步骤时候会循环
            int mid=left+(right-left)/2+1;//这里+1，否则[1]会循环
            if(check(citations,mid)){
                // 1 in the right
                left=mid;
            }else{
                // 2 in the left
                right=mid-1;
            }
        }
        return left;
    }
private:
    bool check(vector<int>&citations,int mid){
        int cnt=0;
        for(auto &n:citations){
            if(n>=mid){
                cnt++;
            }
        }
        return cnt>=mid;
    }
};
