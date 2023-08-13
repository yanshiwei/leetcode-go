class Solution {
public:
    bool check(vector<int>&c1,vector<int>&c2){
        for(int i=0;i<26;i++){
            if(c1[i]!=c2[i]){
                return false;
            }
        }
        return true;
    }
    /*
    题目：
    给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
    题解：
    这是一道使用双指针实现滑动窗口的裸题。
具体的，我们可以先创建一个大小为 26的数组 c2 来统计字符串 p 的词频，另外一个同等大小的数组 c1用来统计「滑动窗口」内的 s 的子串词频。
当两个数组所统计词频相等，说明找到了一个异位组，将窗口的左端点加入答案。
    */
    vector<int> findAnagrams(string s, string p) {
        // hash+滑动窗口
        unordered_map<char,int>hs,ht;// hs哈希表维护的是s字符串中滑动窗口中各个字符出现多少次，ht哈希表维护的是t字符串各个字符出现多少次
        //如果hs哈希表中包含ht哈希表中的所有字符，并且对应的个数都不小于ht哈希表中各个字符的个数，
        // 那么说明当前的窗口是可行的，可行中的长度最短的滑动窗口就是答案
        // 1 遍历t字符串，用ht哈希表记录t字符串各个字符出现的次数
        for(auto c:p){
            ht[c]++;
        }
        // 2 区间[left,right]表示当前滑动窗口。首先让left和right指针都指向字符串s开头，然后枚举整个字符串s 
        // 枚举过程中，不断增加right使滑动窗口增大，相当于向右扩展滑动窗口。
        vector<int>res;
        int left=0;
        int right=0;
        int cnt=0;//s字符串[left,right]区间中满足t字符串的元素的个数
        while(right<s.size()){
            hs[s[right]]++;//每次向右扩展滑动窗口一步，将s[i]加入滑动窗口中，而新加入了s[i]，相当于滑动窗口维护的字符数加一，即hs[s[right]]++
            while(hs[s[right]]>ht[s[right]]){
                //hs哈希表中s[left]的数量多于ht哈希表中s[left]的数量，此时我们就需要向右收缩
                hs[s[left]]--;
                left++;
            }
            if((right-left+1)==p.size()){
                //滑动窗口包含符串 t 的全部字符
                res.push_back(left);
            }
            right++;
        }
        return res;
    }
};
