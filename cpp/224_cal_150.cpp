class Solution {
      /*
    题目：给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
    题解：
    字符串里只有—+和左右括号，非负的整数，空格
    引入两个栈，数字栈st1,操作符栈st2
    若是数字：直接入st1；
    若是操作符：比较当前操作符与栈顶元素：
    1. 若小于或者等于栈顶元素，则将栈顶操作符出栈，并将数据栈st1的2个数据出栈计算然后结果入栈，重复进行；
    本题只有+-，故该条件一直成立，不会走到2
    2. 若大于栈顶元素，直接入操作符栈st2
    若是括号：
    3. 左括号直接入操作符栈st2；
    4. 右括号，从数据栈st1的2个数据出栈，从操作符st2出栈一个操作符计算结果入栈，重复直到遇到左括号
    注意-1，+1这种数据存在：
    5. 由于第一个数可能是负数，为了减少边界判断。一个小技巧是先往 nums 添加一个 0
    6. 为防止 () 内出现的首个字符为运算符，将所有的空格去掉，并将 (- 替换为 (0-，(+ 替换为 (0+
    比如"1-(     -2)"
    */
public:
    int calculate(string s) {
        // 存放所有的数字
        stack<int>nums;
        // rule5
        // 为了防止第一个数为负数，先往 nums 加个 0
        nums.push(0);
        // rule6.1 将所有的空格去掉 方便处理
        replace(s);
        // 存放所有的操作，包括 +/- ()
        stack<char>ops;
        int n=s.size();
        for(int i=0;i<s.size();i++){
            char c=s[i];
            if(c=='('){
                // rule 3
                ops.push(c);
            }else if(c==')'){
                // rule 4
                // 计算到最近一个左括号为止
                while(!ops.empty()){
                    char op=ops.top();
                    if(op!='('){
                        cal(nums,ops);
                    }else{
                        ops.pop();
                        break;
                    }
                }
            }else {
                // digit or op
                if(isdigit(c)){
                    // digit
                    int cur_num=0;
                    int j=i;
                    // 将从 i 位置开始后面的连续数字整体取出，加入 nums
                    while(j<n&&isdigit(s[j])){
                        cur_num=cur_num*10+(s[j]-'0');
                        j++;
                    }
                    nums.push(cur_num);
                    i=j-1;//为了统一走for循环这里j++的结果又-1
                }else{
                    // common op
                    // rule6.2
                    if(i>0&&s[i-1]=='('){
                        nums.push(0);
                    }
                    // rule 1
                    // 有一个新操作要入栈时，先把栈内可以算的都算了
                    while(!ops.empty()&&ops.top()!='('){
                        cal(nums,ops);
                    }
                    ops.push(c);
                }
            }
        }
        //将栈中的其他元素继续运算
        while(!ops.empty()){
            cal(nums,ops);
        }
        return nums.top();
    }
private:
    void cal(stack<int>&nums,stack<char>&ops){
        if(nums.size()<2||ops.empty()){
            return;
        }
        int a2=nums.top();
        nums.pop();
        int a1=nums.top();
        nums.pop();
        char op=ops.top();
        ops.pop();
        nums.push(op=='+'?a1+a2:a1-a2);
    }
    void replace(string&str){
        size_t pos=str.find(" ");
        while(pos!=string::npos){
            str.replace(pos,1,"");//replace(pos,len,str2);
            pos=str.find(" ");
        }
    }
};
