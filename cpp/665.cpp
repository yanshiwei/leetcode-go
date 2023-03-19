class Solution {
    /*
    题目：
    给你一个长度为 n 的整数数组 nums ，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
    题解：
    这道题极易出现一种错误：仅仅判断是否出现了一次下降，见下方错误代码。此种解法会在[3,4,2,3]处报错。因此需要仔细思考你的贪心策略在各种情况下，是否仍然是最优解。
    问题： 在遍历比较数组中nums[i]和nums[i-1]两个数的大小时，如果nums[i]<nums[i-1]，无法判断是nums[i]变为nums[i-1]还是nums[i-1]变为nums[i]。因此需要引入第三个数。
    1. 如果nums[i]>=nums[i-2]，则nums[i-1]=nums[i]或nums[i]=nums[i-1]都可，，但是未来贪心的实现单调递增，我们希望前面的越小越好，nums[i-1]=nums[i]
    2. 如果nums[i]<nums[i-2]，由于我们之前贪心的策略所以这里nums[i-1]>=nums[i-2]的，
    故nums[i]=nums[i-1]
    */
public:
    bool checkPossibility(vector<int>& nums) {
        int cnt=0;
        for(int i=1;i<nums.size();i++){
            if(nums[i]<nums[i-1]){
                if(i==1){
                    nums[i-1]=nums[i];
                }else if(nums[i]>=nums[i-2]){
                    nums[i-1]=nums[i];
                }else{
                    nums[i]=nums[i-1];
                }
                cnt++;
            }
        }
        return cnt<=1;
    }
};
