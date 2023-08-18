class Solution {
public:
    vector<string>split(string&str,string&det){
        if(str.size()<1){
            return {};
        }
        string s=str;
        size_t pos=s.find(det);
        vector<string> res;
        while(pos!=string::npos){
            string tmp=s.substr(0,pos);
            res.push_back(tmp);
            s.erase(0,pos+det.size());
            pos=s.find(det);
        }
        res.push_back(s);
        return res;
    }
    bool wordPattern(string pattern, string s) {
        string det=" ";
        vector<string>ss=split(s,det);
        for(int i=0;i<ss.size();i++){
            cout<<ss[i]<<" ";
        }
        if(pattern.size()!=ss.size()){
            return false;
        }
        unordered_map<char,string>p2s;
        unordered_map<string,char>s2p;
        for(int i=0;i<pattern.size();i++){
            char c=pattern[i];
            if(p2s.count(c)>0){
                // exist
                if(ss[i]!=p2s[c]){
                    return false;
                }
            }else{
                p2s[c]=ss[i];
            }
        }
        for(int i=0;i<pattern.size();i++){
            string s1=ss[i];
            if(s2p.count(s1)>0){
                // exist
                if(pattern[i]!=s2p[s1]){
                    return false;
                }
            }else{
                s2p[s1]=pattern[i];
            }
        }
        return true;
    }
};
