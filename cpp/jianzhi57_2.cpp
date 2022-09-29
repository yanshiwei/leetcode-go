class Solution {
    /*
    输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
    */
public:
    vector<vector<int>> findContinuousSequence(int target) {
        vector<vector<int>> res;
        int small=1,big=2;
        int curSum=small+big;
        int mid=(target-1)/2;
        // 因为序列至少有两个数字，故small+small+1<=s
        while(small<big){
            // Loop 1
            if(curSum==target){
                vector<int> tmp;
                for(int i=small;i<=big;i++){
                    tmp.push_back(i);
                }
                res.push_back(tmp);
            }
            // curSum>target 说明序列过大，从序列去掉较小值 small++
            while(curSum>target){
                // Loop 2
                curSum-=small;
                small++;
                if(curSum==target&&small<=mid){
                    vector<int> tmp;
                    for(int i=small;i<=big;i++){
                        tmp.push_back(i);
                    }
                    res.push_back(tmp);
                }
            }
            // curSum<target 说明序列过小，从序列增加较大值 big++
            // 或者curSum==target 自动增加序列，big++
            big++;
            curSum+=big;
            // cout<<small<<" "<<big<<" "<<curSum<<endl;
        }// 一直到curSum-small>target 此时 target=9, small=9,big=10，curSum=19，会在Loop执行2次此后small=11,curSum=0 之后退出Loop2 big=11 
        // small=big=11 退出Loop1
        return res;
    }
};
