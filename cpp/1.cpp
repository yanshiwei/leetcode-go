class Solution {
    /*
    map:map<int, string> my_map;  
    insert:my_map.insert(pair<int, string>(1, "one"));  
    erase:my_map.erase(it);      
    find(2)
    count(2)
    empty size
    swap
    按照key 自小到大排序
    */
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> hashtable;
        for(int i=0;i<nums.size();i++){
            auto it=hashtable.find(target-nums[i]);
            if(it!=hashtable.end()){
                return {it->second,i};
            }
            hashtable[nums[i]]=i;
        }
        return {};
    }
};
