class Solution {
    /*
    给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表
    字母异位词:指的是排序后都一样如"ate","eat","tea"排序后都是aet
    */
public:
// map:insert erase fdin count size empty swap
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> res; 
        if(strs.size()<1){
            return res;
        }
        unordered_map<string,int>table;// 每个同源字符串在res数组中的下标,res[i]对应这个同源词的各种排序链表
        for(int i=0;i<strs.size();i++){
            auto sorted_str=strs[i];// 保留原始副本要存储都结果
            sort(sorted_str.begin(),sorted_str.end());
            if(table.find(sorted_str)==table.end()){
                // 第一次出现
                table[sorted_str]=table.size();
                res.push_back(vector<string>());// 分配初始化一个新子数组
            }
            res[table[sorted_str]].push_back(strs[i]);
        }
        return res;
    }
};
