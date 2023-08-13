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
        int n=s.size();
        int m=p.size();
        vector<int>c1(26,0);
        vector<int>c2(26,0);
        for(auto &c:p){
            c2[c-'a']++;
        }
        int left=0;
        int right=0;
        vector<int>res;
        while(right<n){
            c1[s[right]-'a']++;
            if(right-left+1>m){
                //超过长度，去掉left
                c1[s[left]-'a']--;
                left++;
            }
            if(check(c1,c2)){
                res.push_back(left);
            }
            right++;
        }
        return res;
    }
};
