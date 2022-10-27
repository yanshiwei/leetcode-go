class Solution {
    /*
    给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
    */
public:
    string minWindow(string s, string t) {
        // hash+滑动窗口
        unordered_map<char,int>hs,ht;// hs哈希表维护的是s字符串中滑动窗口中各个字符出现多少次，ht哈希表维护的是t字符串各个字符出现多少次
        //如果hs哈希表中包含ht哈希表中的所有字符，并且对应的个数都不小于ht哈希表中各个字符的个数，
        // 那么说明当前的窗口是可行的，可行中的长度最短的滑动窗口就是答案
        // 1 遍历t字符串，用ht哈希表记录t字符串各个字符出现的次数
        for(auto c:t){
            ht[c]++;
        }
        // 2 区间[left,right]表示当前滑动窗口。首先让left和right指针都指向字符串s开头，然后枚举整个字符串s 
        // 枚举过程中，不断增加right使滑动窗口增大，相当于向右扩展滑动窗口。
        string res="";
        int cnt=0;//s字符串[left,right]区间中满足t字符串的元素的个数
        int left=0,right=0;
        while(right<s.size()){
            hs[s[right]]++;//每次向右扩展滑动窗口一步，将s[i]加入滑动窗口中，而新加入了s[i]，相当于滑动窗口维护的字符数加一，即hs[s[right]]++
            if(hs[s[right]]<=ht[s[right]]){
                // hs[s[i]] <= ht[s[i]]，说明当前新加入的字符s[i]是必需的，且还未到达字符串t所要求的数量
                cnt++;//新加入的字符s[i]必需，则cnt++
            }
            while(hs[s[left]]>ht[s[left]]){
                //hs哈希表中s[left]的数量多于ht哈希表中s[left]的数量，此时我们就需要向右收缩
                hs[s[left]]--;
                left++;
            }
            if(cnt==t.size()){
                //滑动窗口包含符串 t 的全部字符
                if(res.empty()||right-left+1<res.size()){
                    res=s.substr(left,right-left+1);
                }
            }
            right++;
        }
        return res;
    }
};
