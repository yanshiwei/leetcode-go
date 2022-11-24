class Solution {
    // 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
    // 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
public:
    int lengthOfLastWord(string s) {
        if(s.size()<1){
            return 0;
        }
        int idx=s.size()-1;
        // ignore space
        while(idx>=0){
            if(s[idx]!=' '){
                break;
            }
            idx--;
        }
        int cnt=0;
        // cnt statistics
        while(idx>=0){
            if(s[idx]==' '){
                break;
            }
            idx--;
            cnt++;
        }
        return cnt;
    }
};
