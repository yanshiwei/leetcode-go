class Solution {
    /*
    两个数据下标之差小于等于indexDiff，可以引入滑动窗口，大小为indexDiff+1，
    在滑动窗口内，当前值v，历史值x，要求abs(v-x)<=valueDiff也就是需要找到x在区间[v-valueDiff,v+valueDiff]。
    也就是说要证明大于等于v-valueDiff数据集合中至少有一个小于等于valueDiff+v
    也即大于等于v-valueDiff的第一个数要小于等于valueDiff+v
    大于等于v-valueDiff的第一个数可以用<algorithm>里的lower_bound
    //lower_bound(begin,end,value)-->大于等于value的第一个
    //lower_bound(begin,end,value,cmp())-->第一个不符合cmp，cmp可以自定义或者函数类
    给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。
找出满足下述条件的下标对 (i, j)：   
    i != j,
    abs(i - j) <= indexDiff
    abs(nums[i] - nums[j]) <= valueDiff
    如果存在，返回 true ；否则，返回 false 。
    */
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int indexDiff, int valueDiff) {
        if(nums.size()<1||indexDiff<1||valueDiff<0){
            return false;
        }
        set<long long> st;
        int left = 0;
        for (int right = 0; right < nums.size(); right ++) {
            if (right - left > indexDiff) {
                st.erase(nums[left]);
                left ++;
            }
            auto a = st.lower_bound((long long) nums[right] - valueDiff);
            if (a != st.end() && *a<= (long long)valueDiff+nums[right]) {
                return true;
            }
            st.insert(nums[right]);
        }
        return false;
    }
};
