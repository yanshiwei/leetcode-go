class Solution {
    /*
    根据 逆波兰表示法，求表达式的值。有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
    */
public:
    int evalRPN(vector<string>& tokens) {
        stack<long>st;
        int n=tokens.size();
        for(int i=0;i<n;i++){
            string& token=tokens[i];
            if(isNumber(token)){
                st.push(atoi(token.c_str()));
            }else{
                // 普通字符
                long num2=st.top();
                st.pop();
                long num1=st.top();
                st.pop();
                switch(token[0]){
                    case '+':
                      st.push(num2+num1);
                      break;
                    case '-':
                      st.push(num1-num2);
                      break;
                    case '*':
                      //  -128*-128*-128*-128=2^31 ,这个数大于int最大范围2^31-1。最后乘以-1后的值-2^31是合法的，但是倒数第二步的值是非法的，所以中间计算过程要改为double，最后在转为int。
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
    // tokens[i] 是一个算符（"+"、"-"、"*" 或 "/"），或是在范围 [-200, 200] 内的一个整数
    bool isNumber(string &token){
        return !(token == "+" || token == "-" || token == "*" || token == "/");
    }
};
