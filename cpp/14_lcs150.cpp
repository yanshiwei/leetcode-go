class Solution {
public:
    string lcs(string&s1,string&s2){
        int length=min(s1.size(),s2.size());
        int i;
        for(i=0;i<length;i++){
            if(s1[i]!=s2[i]){
                break;
            }
        }
        return s1.substr(0,i);
    }
    string longestCommonPrefix(vector<string>& strs) {
        string res="";
        if(strs.size()<1){
            return res;
        }
        res=strs[0];
        for(int i=1;i<strs.size();i++){
            res=lcs(res,strs[i]);
            if(res.size()<1){
                break;
            }
        }
        return res;
    }
};
