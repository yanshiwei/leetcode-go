class Solution {
public:
    vector<string>split(string&str,string&det){
        if(str.size()<1){
            return {};
        }
        vector<string>res;
        string s=str;
        size_t pos=s.find(det);
        while(pos!=string::npos){
            string tmp=s.substr(0,pos);
            res.push_back(tmp);
            s.erase(0,pos+det.size());//删除之前元素+det本身
            pos=s.find(det);
        }
        res.push_back(s);
        return res;
    }
    bool wordPattern(string pattern, string s) {
        string det=" ";
        vector<string>ss=split(s,det);
        if(pattern.size()!=ss.size()){
            return false;
        }
        unordered_map<char,string>c2s;
        unordered_map<string,char>s2c;
        for(int i=0;i<pattern.size();i++){
            char c=pattern[i];
            if(c2s.count(c)>0){
                // exist before
                if(c2s[c]!=ss[i]){
                    return false;
                }
            }else{
                c2s[c]=ss[i];
            }
        }
        for(int i=0;i<ss.size();i++){
            string s1=ss[i];
            if(s2c.count(s1)>0){
                // exist before
                if(s2c[s1]!=pattern[i]){
                    return false;
                }
            }else{
                s2c[s1]=pattern[i];
            }
        }
        return true;
    }
};
