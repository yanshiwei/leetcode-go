class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        if(s.size()<10){
            return {};
        }
        unordered_set<string>hash;
        unordered_set<string>res;
        for(int i=0;i<=s.size()-10;i++){
            string key=s.substr(i,10);
            if(hash.count(key)>0){
                res.insert(key);
            }else{
                hash.insert(key);
            }
        }
        return vector<string>(res.begin(),res.end());
    }
};
