class Solution {
    /*
    给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效
    s 仅由括号 '()[]{}' 组成
    */
public:
    bool isValid(string s) {
        if(s.size()<2){
            return false;
        }
        stack<char>st;
        for(int i=0;i<s.size();i++){
            if (s[i]=='('||s[i]=='['||s[i]=='{'){
                st.push(s[i]);
            }else if(s[i]==')'){
                if(st.empty()){
                    return false;
                }
                if(st.top()!='('){
                    return false;
                }
                st.pop();
            }else if(s[i]==']'){
                if(st.empty()){
                    return false;
                }
                if(st.top()!='['){
                    return false;
                }
                st.pop();
            }else if(s[i]=='}'){
                if(st.empty()){
                    return false;
                }
                if(st.top()!='{'){
                    return false;
                }
                st.pop();
            }
        }
        return st.empty()==true;
    }
};
