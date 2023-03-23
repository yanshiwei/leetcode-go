class Solution {
    /*
    如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：
本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
    */
    const int intmax=0x7FFFFFFF;
    const int intmin=(-intmax-1);
public:
    bool isdigit(char c){
        if(c>='0'&&c<='9'){
            return true;
        }
        return false;
    }
    int myAtoi(string s) {
        int i=0;
        int flag=1;// 默认正数
        long res=0;
        while(s[i]==' '){
            i++;
        }
        if(s[i]=='-'){
            flag=-1;
        }
        if(s[i]=='-'||s[i]=='+'){
            i++;
        }
        // 正负号之后不是数字的则直接非法，返回0，如"words and 987"
        while(i<s.size()&&isdigit(s[i])){
                // 空格 正负号之后的非数字都忽略
                res=res*10+(s[i]-'0');
                if(res>=intmax&&flag==1){
                    return intmax;
                }
                if(res>intmax&&flag==-1){
                    return intmin;
                }
            i++;
        }
        return flag*res;
    }
};
