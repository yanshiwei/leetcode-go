class Solution {
public:
    int romanToInt(string s) {
        unordered_map<char,int>dict;
        dict['I']=1;
        dict['V']=5;
        dict['X']=10;
        dict['L']=50;
        dict['C']=100;
        dict['D']=500;
        dict['M']=1000;
        int res=0;
        for(int i=0;i<s.size();i++){
            char c=s[i];
            int v=dict[c];
            if((i<s.size()-1)&&(v<dict[s[i+1]])){
                // case IX
                res-=v;
            }else{
                res+=v;
            }
        }
        return res;
    }
};
