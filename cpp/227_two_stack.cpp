class Solution {
    /*
     题目：给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格。 整数除法仅保留整数部分。
    题解：
     当s[i] 为数字数，提取连续的整数。存入nums
     当s[i] 为运算符时， 
        如果s[i]的优先级比ops栈顶运算符高，那么s[i]入ops栈。  
        如果不是，那么弹出ops栈顶运算符，和 nums栈顶的两个整数，计算，结果存入nums栈。
    */
public:
    int calculate(string s) {
        // 去除空格
        replace(s);
        stack<int>nums;
        stack<char>ops;
        int i=0;
        int n=s.size();
        while(i<n){
            char c=s[i];
            if(isdigit(c)){
                //取数字
                int cur=0;
                while(i<n&&isdigit(s[i])){
                    cur=cur*10+(s[i]-'0');
                    i++;
                }
                // 注意上面的计算一定要有括号，否则有可能会溢出
                nums.push(cur);
            }else{
                //取运算符
                while(!ops.empty()&&!higher(c,ops.top())){
                    calc(nums,ops);
                }
                ops.push(c);
                i++;
            }
        }
        //将栈中的其他元素继续运算
        while(ops.size()){
            calc(nums,ops);
        }
        return nums.top();
    }
private:
    void calc(stack<int>&nums,stack<char>&ops){
        if(nums.size()<2||ops.size()<1){
            return;
        }
        int b = nums.top(); nums.pop();
        int a = nums.top(); nums.pop();
        char op = ops.top(); ops.pop();
        int c=1;
        switch(op){
            case '+':c = a+b;break;
            case '-':c = a-b;break;
            case '*':c = a*b;break;
            case '/':c = a/b;break; 
            default: assert(0);
        }
        nums.push(c);
    }
    bool higher(char a,char b){
        if(a=='+'||a=='-'){
            return false;
        }
        // a==*/
        if(b=='+'||b=='-'){
            return true;
        }
        return false;
    }
    void replace(string &s){
        // 去除空格
        int pos=s.find(" ");
        while(pos!=-1){
            s.replace(pos,1,"");
            pos=s.find(" ");
        }
    }
};
