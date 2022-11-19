class Solution {
    /*
    给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
    */
public:
    string multiply(string num1, string num2) {
        if(num1=="0"||num2=="0"){
            return "0";
        }
        vector<int>tmp(num1.size()+num2.size());// num1[i]*num2[j] save sa tmp[i+j+1]
        for(int i=0;i<num1.size();i++){
            for(int j=0;j<num2.size();j++){
                tmp[i+j+1]+=int(num1[i]-'0')*int(num2[j]-'0');
            }
        }
        for(int i=tmp.size()-1;i>0;i--){
            tmp[i-1]+=tmp[i]/10;
            tmp[i]=tmp[i]%10;
        }
        string res="";
        for(int i=0;i<tmp.size();i++){
            if(i==0&&tmp[i]==0){
                continue;
            }
            res+=(tmp[i]+'0');
        }
        return res;
    }
};
