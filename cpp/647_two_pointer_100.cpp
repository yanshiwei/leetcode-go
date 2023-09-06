class Solution {
    /*
    给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
    */
public:
    int countSubstrings(string s) {
        int res=0;
        // 背向双指针
        //遍历每一个下标，以这个下标为中心，利用「回文串」中心对称的特点，往两边扩散，看最多能扩散多远
        for(int i=0;i<s.size();i++){
            //回文串的长度有可能是奇数，也有可能是偶数 故s[i]可能是奇数长度核心，也可能是偶数长度的一个边界，所以都需要计算
           //s = "aaa"
           //i=0: (0,0)->a;(0,1)->aa;
           //i=1: (1,1)->a/aaa;(1,2)->aa;
           //i=3: (2,2)->a;(2,3)->null;
           res+=palindrome(s,i,i);
           res+=palindrome(s,i,i+1);
        }
        return res;
    }
    int palindrome(string s,int left,int right){
        int res=0;
        while(left>=0&&right<s.size()&&s[left]==s[right]){
            left--;
            right++;
            res++;
        }
        return res;
    }
};
