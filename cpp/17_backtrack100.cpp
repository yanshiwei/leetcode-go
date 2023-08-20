class Solution {
    /*
    给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
    */
public:
    vector<string>res;
    unordered_map<char,string>dict{
        {'2',"abc"},
        {'3', "def"},
        {'4', "ghi"},
        {'5', "jkl"},
        {'6', "mno"},
        {'7', "pqrs"},
        {'8', "tuv"},
        {'9', "wxyz"}
    };
    void dfs(string digits,int idx,const string&cur){
        if(idx==digits.size()&&cur.empty()==false){
            res.push_back(cur);
            return;
        }
        for(auto c:dict[digits[idx]]){
            dfs(digits,idx+1,cur+c);
        }
    }
    vector<string> letterCombinations(string digits) {
        dfs(digits,0,"");
        return res;
    }
};
