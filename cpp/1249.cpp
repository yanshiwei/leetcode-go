class Solution {
    /*
    给你一个由 '('、')' 和小写字母组成的字符串 s
    */
public:
    string minRemoveToMakeValid(string s) {
        //两个栈记录无效括号的位置
        stack<int>left;
        stack<int>right;
        for(int i=0;i<s.size();i++){
            if(s[i]=='('){
                left.push(i);
            }else if(s[i]==')'){
                if(left.empty()==false){
                    left.pop();// 去掉对称的左括号
                }else{
                    // 没有对称的左括号
                    right.push(i);
                }
            }
        }
        while(right.empty()==false){
            s[right.top()]=' ';
            right.pop();
        }
        while(left.empty()==false){
            s[left.top()]=' ';
            left.pop();
        }
        string res="";
        for(int i=0;i<s.size();i++){
            if(s[i]==' '){
                continue;
            }
            res.push_back(s[i]);
        }
        return res;
    }
};
