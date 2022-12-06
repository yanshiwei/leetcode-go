class Solution {
    //给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
public:
    string addBinary(string a, string b) {
        if(a.size()<1){
            return b;
        }
        if(b.size()<1){
            return a;
        }
        if(a.size()<b.size()){
            string tmp=b;
            b=a;
            a=tmp;
        }
        int i=a.size()-1;
        int j=b.size()-1;
        string res(a.size()+1,'0');
        int k=res.size()-1;
        int carry=0;
        while(i>=0&&j>=0){
            int ai=a[i]-'0';
            int bi=b[j]-'0';
            res[k]=((ai+bi+carry)%2+'0');
            carry=(ai+bi+carry)/2;
            i--;
            j--;
            k--;
        }
        while(i>=0){
            int ai=a[i]-'0';
            res[k]=((ai+carry)%2+'0');
            carry=(ai+carry)/2;
            i--;
            k--;
        }
        if(carry>0){
            res[k]=((carry)%2+'0');
        }
        if(res.size()>1&&res[0]=='0'){
            res.erase(0,1);
        }
        return res;
    }
};
