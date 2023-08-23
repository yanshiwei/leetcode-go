class Solution {
    /*
    给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
https://leetcode.cn/problems/jump-game-ii/solutions/36035/45-by-ikaruga/?envType=study-plan-v2&envId=top-100-liked
    1.如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点。 可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离end不断更新。
    2. 如果从这个 起跳点 起跳叫做第 1 次 跳跃，那么从后面 3 个格子起跳 都 可以叫做第 2 次 跳跃。
    3. 所以，当一次 跳跃 结束时，从下一个格子开始，到现在 能跳到最远的距离，都 是下一次 跳跃 的 起跳点。

    */
public:
    int jump(vector<int>& nums) {
        if(nums.size()<1){
            return nums.size();
        }
        int cnt=0;
        int start=0;//下一次起跳点范围开始的格子
        int end=1;//下一次起跳点范围结束的格子，不包含
        while(end<nums.size()){
            int maxDis=0;
            // 在这个起跳范围内，不断尝试，并更新能跳到最远的距离maxDis
            for(int i=start;i<end;i++){
                // 能跳到最远的距离
                maxDis=max(maxDis,i+nums[i]);
            }
            // 更新下一次起跳点范围开始的格子
            start=end;
            // 更新下一次起跳点范围开始的格子
            end=maxDis+1;
            cnt++;
        }
        return cnt;
    }
};
