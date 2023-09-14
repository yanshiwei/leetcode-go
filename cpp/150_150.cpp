class Solution {
public:
    // tokens[i] 是一个算符（"+"、"-"、"*" 或 "/"），或是在范围 [-200, 200] 内的一个整数
    bool isNumber(string &token){
        return !(token == "+" || token == "-" || token == "*" || token == "/");
    }
    int evalRPN(vector<string>& tokens) {
        stack<long>st;
        int n=tokens.size();
        for(int i=0;i<n;i++){
            string t=tokens[i];
            if(isNumber(t)){
                st.push(atoi(t.c_str()));
            }else{
                // op
                long num2=st.top();
                st.pop();
                long num1=st.top();
                st.pop();
                switch(t[0]){
                    case '+':
                        st.push(num1+num2);
                        break;
                    case '-':
                        st.push(num1-num2);
                        break;
                    case '*':
                        st.push(num1*num2);
                        break;
                    case '/':
                        st.push(num1/num2);
                        break;
                }
            }
        }
        return int(st.top());
    }
};
