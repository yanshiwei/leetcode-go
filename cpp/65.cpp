class Solution {
    //字符串含有数字/点/+-/eE 判断字符串是有效数字
    //Ae/EB
    //A是包含小数的，可以正负
    //B是整数，可以正负
public:
    bool isNumber(string s) {
        bool numFlag=false;
        bool dotFlag=false;
        bool eFlag=false;
        for(int i=0;i<s.size();i++){
            if('0'<=s[i]&&s[i]<='9'){
                numFlag=true;
            }else if(s[i]=='.'&&dotFlag==false&&eFlag==false){
                //之前没有出现. 且小数点肯定在e/E之前才合理
                dotFlag=true;
            }else if((s[i]=='e'||s[i]=='E')&&eFlag==false&&numFlag){
                // 之前没有出现e/E，且e/E前面肯定要有数字才合理
                // 同时确认是合理的e/E后，重置数字 由判断A重新判断B
                eFlag=true;
                numFlag=false;
            }else if((s[i]=='+'||s[i]=='-')&&(i==0||s[i-1]=='e'||s[i-1]=='E')){
                // 必须第一位或者e/E后第一位
                continue;
            }else{
                return false;
            }
        }
        return numFlag;//防止出现'.'这种 即必须要有数字
    }
};
