class Solution {
    /*
    题目：
    给你一个整数数组 nums ，其中 nums[i] 表示第 i 个袋子里球的数目。同时给你一个整数 maxOperations 。
你可以进行如下操作至多 maxOperations 次：
选择任意一个袋子，并将袋子里的球分到 2 个新的袋子中，每个袋子里都有 正整数 个球。
比方说，一个袋子里有 5 个球，你可以把它们分到两个新袋子里，分别有 1 个和 4 个球，或者分别有 2 个和 3 个球。
你的开销是单个袋子里球数目的 最大值 ，你想要 最小化 开销。
请你返回进行上述操作后的最小开销。
输入：nums = [2,4,8,2], maxOperations = 4
输出：2
解释：
- 将装有 8 个球的袋子分成装有 4 个和 4 个球的袋子。[2,4,8,2] -> [2,4,4,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,4,4,4,2] -> [2,2,2,4,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,4,4,2] -> [2,2,2,2,2,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2] 。
装有最多球的袋子里装有 2 个球，所以开销为 2 并返回 2 。
    思路：
    最终目的是经过最多maxOperations步后，每个袋子里球数目的最大值的最小化。二分！
    1.确定边界，袋子里球数最小是1个，最多是nums最大值，因为我们在一个连续区间不断查找，并且存在在某个临界点使得之后的数字都可以在规定的操作次数内完成，所以可以采用二分查找法。
    2.确定好了边界后，每次二分搜索时需要判断当前计算值是否满足条件，这里我们引入 check 函数，对当前计算出的开销最大值进行验证。
验证过程也即判断，在maxOperations的操作次数内，能否使所有袋子中的球类数目最大值控制为当前二分计算出的mid
    参考：https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/solution/python-er-fen-fa-mei-yi-ge-bu-zou-xiang-t3at8/
    */
public:
    int minimumSize(vector<int>& nums, int maxOperations) {
        // 最小化最大值，二分,初始的最小开销是1 最大开销是数组中的最大值
        long left=1,right=1e9;// 正常应该是数组最大值，首先得遍历一遍数组才能找到最大值，而题目说了数组长度最大可能到十万，那么何必花费一个这么大的开销来找最大值呢直接取所有数组最大值1e9即可。
        while(left<right){
            long mid=left+(right-left)/2;// 当前尝试的开销
            if (check(nums,mid,maxOperations)){
                // 说明可以将nums中所有的数变成mid，为了最小化mid,故可以减小mid以获得最小开销,
                right=mid;// 这里mid仍然满足条件故保留
            }else{
                left=mid+1;// mid不满足故肯定+1
            }
        }
        return left;
    }
private:
    // 验证过程在maxOperations的操作次数内，能否使所有袋子中的球类数目最大值控制为当前二分计算出的mid
    bool check(vector<int>& nums,long mid,int maxOperations){
        int curOps=0;
        for(auto &d:nums){
            if(d%mid==0){
                // 能够整除如d=14 mid=4 要想达到单个袋子里球数目的最大值为4 则应该16->8,8，然后8->4,4;8->4,4一共3次。这里是除法结果-1，因为，最后一次分，可以是两个结果都为mid，故少分一次
                curOps+=d/mid-1;
            }else{
                // 不整除，d=17 mid=4 要想达到单个袋子里球数目的最大值为4 则应该17->8,9，然后8->4,4;9->4,5；5->4,1一共4次，这里需要次数是除法结果，因为分的最后结果一个大于mid
                curOps+=d/mid;
            }
        }
        // 返回可不可以maxOperations步骤内实现把所有元素变成mid
        return curOps<=maxOperations;
    }
};
