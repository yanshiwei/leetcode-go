class Solution {
    /*
    给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
    */
public:
    int lengthOfLongestSubstring(string s) {
        if(s.size()<2){
            return s.size();
        }
        int res=0;
        // 滑动窗口
        unordered_set<char>table;
        int left=0;
        int right=0;
        while(right<s.size()){
            while(table.count(s[right])){
                //不满足要求，所以，我们要移动这个窗口
                table.erase(s[left]);
                left++;
            }
            res=max(res,right-left+1);
            table.insert(s[right]);
            right++;
        }
        return res;
    }
};
