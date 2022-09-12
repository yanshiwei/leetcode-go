class Solution {
public:
/*
你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
*/
    bool uniqueOccurrences(vector<int>& arr) {
        unordered_map<int,int>fre;
        for(int i=0;i<arr.size();i++){
            fre[arr[i]]++;
        }
        unordered_set<int>times;
        for(const auto&x:fre){
            times.insert(x.second);
        }
        return fre.size()==times.size();
    }
};
