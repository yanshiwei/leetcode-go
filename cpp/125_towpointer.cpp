class Solution {
    /*
    如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 
    */
public:
    bool isPalindrome(string s){
        // build new str wthiout black or big
        string newStr="";
        for(char ch:s){
            if(isalnum(ch)){
                newStr+=tolower(ch);
            }
        }
        int n=newStr.size();
        if(n<1){
            return true;
        }
        // 相向指针
        int left=0;
        int right=n-1;
        while(left<right){
            if(newStr[left]!=newStr[right]){
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
/*
    bool isPalindrome(string s) {
        // C++判断字母数字isalnum 转小写tolower
        string newStr="";
        for(char ch:s){
            if(isalnum(ch)){
                // 全部小写
                newStr+=tolower(ch);
            }
        }
        int n=newStr.size();
        if(n<1){
            //在移除非字母数字字符之后，s 是一个空字符串 "" 定位为true
            return true;
        }
        //相向双指针
        int left=0,right=n-1;
        while(left<right){
            if(newStr[left]!=newStr[right]){
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
    */
};
