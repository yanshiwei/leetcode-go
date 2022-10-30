class Solution {
    /*
    给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
    */
public:
    int lengthOfLongestSubstring(string s) {
        // 滑动窗口
        if(s.size()<2){
            return s.size();
        }
        unordered_set<char>table;
        int maxLen=0;
        int left=0;
        for(int i=0;i<s.size();i++){
            while(table.find(s[i])!=table.end()){
                //不满足要求。所以，我们要移动这个队列
                table.erase(s[left]);
                left++;//把队列的左边的元素移出就行了，直到满足题目要求
            }
            maxLen=max(maxLen,i-left+1);
            table.insert(s[i]);
        }
        return maxLen;
    }
};
