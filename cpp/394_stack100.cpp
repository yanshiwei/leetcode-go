class Solution {
    /*
    给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
    */
public:
    string decodeString(string s) {
        string res="";
        stack<int>nums;// 数字栈
        stack<string>strs;//符号栈
        int num=0;
        for(int i=0;i<s.size();i++){
            if(s[i]>='0'&&s[i]<='9'){
                num=num*10+s[i]-'0';
            }else if((s[i]>='a'&&s[i]<='z')||(s[i]>='A'&&s[i]<='Z')){
                res+=s[i];
            }else if(s[i]=='['){
                //入栈
                nums.push(num);
                num=0;//方便统计下一个
                strs.push(res);
                res="";
            }else{
                //遇到‘]’时，操作与之相配的‘[’之间的字符，使用分配律
                int cnt=nums.top();
                nums.pop();
                for(int j=0;j<cnt;j++){
                    strs.top()+=res;
                }
                res=strs.top();
                strs.pop();
            }
        }
        return res;
    }
};
