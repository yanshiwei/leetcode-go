class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        unordered_map<char,int>table;
        for(char c: magazine){
            table[c]++;
        }
        for(char c : ransomNote){
            if(table.count(c)<1){
                return false;
            }
            table[c]--;
            if(table[c]==0){
                table.erase(c);
            }
        }
        return true;
    }
};
