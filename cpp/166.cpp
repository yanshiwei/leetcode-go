class Solution {
    /*
    题目：
    给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
如果小数部分为循环小数，则将循环的部分括在括号内。
如果存在多个答案，只需返回 任意一个 。
对于所有给定的输入，保证 答案字符串的长度小于 104 。
    题解：
    首先明确 像PI这样的无理数不能表示成分数；除了无限不循环小数以外的实数统称有理数，有理数包括整数和分数。
    其次，在计算余数时，小数部分如果余数出现两次就表示该小数是循环小数了
    */
public:
    string fractionToDecimal(int numerator, int denominator) {
        if(denominator==0){
            //边界条件，分母为0
            return "";
        }
        if(numerator==0){
            //边界条件，分子为0
            return "0";
        }
        string res="";
        //因为内部计算时要乘以10，故转换为longlong防止溢出
        long long num=static_cast<long long>(numerator);
        long long denom=static_cast<long long>(denominator);
        //处理正负号，一正一负取负号
        if((num>0)^(denom>0)){
            //异或
            res.push_back('-');
        }
        //分子分母全部转换为正数
        num=abs(num);
        denom=abs(denom);
        //处理整数部分
        res+=(to_string(num/denom));
        //处理小数部分
        num=num%denom;//获得余数
        if(num==0){
            //余数为0，表示整除了，直接返回结果
            return res;
        }
        //余数不为0，添加小数点
        res.push_back('.');
        int idx=res.size()-1;//获得小数点的下标
        unordered_map<int,int>hash;//map用来记录出现重复数的下标，然后将'('插入到重复数前面就好了
        while(num&&hash.count(num)==0){
            hash[num]=++idx;
            num*=10;//余数扩大10倍，然后求商，和草稿本上运算方法是一样的
            res+=to_string(num/denom);
            num=num%denom;
        }
        if(hash.count(num)>0){
            //出现循环余数，我们直接在重复数字前面添加'(',字符串末尾添加')
            res.insert(res.begin()+hash[num],'(');
            res.push_back(')');
        }
        return res;
    }
};
