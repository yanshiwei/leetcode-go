class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        if(strs.size()<1){
            return {};
        }
        vector<vector<string>> res;
        unordered_map<string,int>table;
        for(int i=0;i<strs.size();i++){
            string w=strs[i];
            sort(w.begin(),w.end());
            if(table.count(w)==0){
                table[w]=table.size();
                res.push_back(vector<string>{});//empty ready to fill
            }
            res[table[w]].push_back(strs[i]);
        }
        return res;
    }
};
