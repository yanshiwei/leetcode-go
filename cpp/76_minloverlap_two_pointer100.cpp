class Solution {
    /*
    给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
    */
public:
    //https://www.codetd.com/article/11582575
    string minWindow(string s, string t) {
        // hash+滑动窗口
        unordered_map<char,int>need,window;// hs哈希表维护的是s字符串中滑动窗口中各个字符出现多少次，ht哈希表维护的是t字符串各个字符出现多少次
        //如果hs哈希表中包含ht哈希表中的所有字符，并且对应的个数都不小于ht哈希表中各个字符的个数，
        // 那么说明当前的窗口是可行的，可行中的长度最短的滑动窗口就是答案
        // 1 遍历t字符串，用ht哈希表记录t字符串各个字符出现的次数
        for(auto c:t){
            need[c]++;
        }
        // 2 区间[left,right]表示当前滑动窗口。首先让left和right指针都指向字符串s开头，然后枚举整个字符串s 
        // 枚举过程中，不断增加right使滑动窗口增大，相当于向右扩展滑动窗口。
        string res="";
        int left=0;
        int right=0;
        int cnt=0;//s字符串[left,right]区间中满足t字符串的元素的个数
        // 记录最小覆盖子串的起始索引及长度
        int start = 0, len = INT_MAX;
        while(right<s.size()){
            // c 是将移入窗口的字符
            char c=s[right];
            right++;
            // 进行窗口内数据的一系列更新
            if(need.count(c)){
                window[c]++;
                if(window[c]==need[c]){
                    cnt++;
                }
            }
            // 关键点：判断左侧窗口是否要收缩
            while(cnt==need.size()){
                // 在这里更新最小覆盖子串
                if (right - left < len) {
                    start = left;
                    len = right - left;
                }
                // d 是将移出窗口的字符
                char d=s[left];
                left++;
                // 进行窗口内数据的一系列更新
                if(need.count(d)){
                    if(window[d]==need[d]){
                        cnt--;
                    }
                    window[d]--;
                }
            }
        }
        return len == INT_MAX ? "" : s.substr(start, len);
    }
};
