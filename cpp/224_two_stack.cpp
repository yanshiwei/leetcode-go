class Solution {
    /*
    题目：给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
    题解：
    字符串里只有—+和左右括号，非负的整数，空格
    两个栈，一个装数字st1，一个装操作符st2
    若是数字，直接入st1
    若是操作符，比较当前符号与栈顶元素比较：
    1 若小于或者等于栈顶元素，则将栈顶操作符出栈，并将数据栈st1的2个数据出栈计算然后结果入栈，重复进行；
    这里只有+-，故该条件一直成立，不会走到2
    2 若大于栈顶元素，直接入栈st2
    若是括号：
    1 左括号直接入栈
    2 右括号，从数据栈st1的2个数据出栈，从操作符st2出栈一个操作符计算结果入栈，重复直到遇到左括号
    注意-1，+1这种数据存在：
    1.由于第一个数可能是负数，为了减少边界判断。一个小技巧是先往 nums 添加一个 0
    2.为防止 () 内出现的首个字符为运算符，将所有的空格去掉，并将 (- 替换为 (0-，(+ 替换为 (0+
    比如"1-(     -2)"
    */
public:
    int calculate(string s) {
        // 存放所有的数字
        stack<int> nums;
        // 为了防止第一个数为负数，先往 nums 加个 0
        nums.push(0);
        // 将所有的空格去掉
        replace(s);
        // 存放所有的操作，包括 +/-
        stack<char> ops;
        int n = s.size();
        for(int i = 0; i < n; i++) {
            char c = s[i];
            if(c == '(')
                ops.push(c);
            else if(c == ')') {
                // 计算到最近一个左括号为止
                while(!ops.empty()) {
                    char op = ops.top();
                    if(op != '(')
                        calc(nums, ops);
                    else {
                        ops.pop();
                        break;
                    }
                }
            }else {
                if(isdigit(c)) {
                    int cur_num = 0;
                    int j = i;
                    // 将从 i 位置开始后面的连续数字整体取出，加入 nums
                    while(j <n && isdigit(s[j]))
                        cur_num = cur_num*10 + (s[j++] - '0');
                    // 注意上面的计算一定要有括号，否则有可能会溢出
                    nums.push(cur_num);
                    i = j-1;
                }
                else {
                    if (i > 0 && s[i - 1] == '(') {
                        nums.push(0);
                    }
                    // 有一个新操作要入栈时，先把栈内可以算的都算了
                    while(!ops.empty() && ops.top() != '(')
                        calc(nums, ops);
                    ops.push(c);
                }
            }
        }
        cout<<"a"<<nums.size()<<"."<<ops.size();
        //将栈中的其他元素继续运算
        while(!ops.empty())
            calc(nums, ops);
        return nums.top();
    }
private:
    void calc(stack<int>&nums,stack<char>&ops){
        if(nums.size() < 2 || ops.empty())
            return;
        int b = nums.top(); nums.pop();
        int a = nums.top(); nums.pop();
        char op = ops.top(); ops.pop();
        cout<<"res"<<a<<","<<b<<","<<ops.size()<<endl;
        nums.push(op == '+' ? a+b : a-b);
    }
    void replace(string &s){
        int pos=s.find(" ");
        while(pos!=-1){
            s.replace(pos,1,"");//replace(pos,len,str2);
            pos=s.find(" ");
        }
    }
};
