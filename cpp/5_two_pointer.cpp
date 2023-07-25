class Solution {
public:
    string longestPalindrome(string s){
        /*
        背向指针，遍历每一个下标，以这个为中心背向扩展，看最多扩展多远
        */
        string res="";
        for(int i=0;i<s.size();i++){
            //回文串的长度有可能是奇数，也有可能是偶数 故s[i]可能是奇数长度核心，也可能是偶数长度的一个边界，所以都需要计算
            string odd=palindrome(s,i,i);//奇数则以s[i]为中心扩散
            string even=palindrome(s,i,i+1);//偶数则以s[i]和s[i+1]为中心进行扩散
            if(res.size()<odd.size()){
                res=odd;
            }
            if(res.size()<even.size()){
                res=even;
            }            
        }
        return res;
    }
private:
    string palindrome(string s,int left,int right){
        while(left>=0&&right<s.size()&&s[left]==s[right]){
            left--;
            right++;
        }
        return s.substr(left+1,right-left-1);
    }
/*
    string longestPalindrome(string s) {
        // 背向双指针
        //遍历每一个下标，以这个下标为中心，利用「回文串」中心对称的特点，往两边扩散，看最多能扩散多远
        string res="";
        // left right为回文串的边界
        for(int i=0;i<s.size();i++){
            //回文串的长度有可能是奇数，也有可能是偶数 故s[i]可能是奇数长度核心，也可能是偶数长度的一个边界，所以都需要计算
            string odd=palindrome(s,i,i);//奇数则以s[i]为中心扩散
            string even=palindrome(s,i,i+1);//偶数则以s[i]和s[i+1]为中心进行扩散
            if(res.size()<odd.size()){
                res=odd;
            }
            if(res.size()<even.size()){
                res=even;
            }
        }
        return res;
    }
private:
    string palindrome(string s,int left,int right){
        while(left>=0&&right<s.size()&&s[left]==s[right]){
            left--;
            right++;
        }
        return s.substr(left+1,right-left-1);// substr(begin,len)
    }
    */
};
