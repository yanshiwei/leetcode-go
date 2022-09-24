class RandomizedSet {
    // map+vector
    // map: insert erase count find empty size swap
    //vector push_back insert
    //.      pop_back insert
    //.      at. front back
    //.      empty size swap
public:
    RandomizedSet() {

    }
    
    bool insert(int val) {
        if(table.count(val)){
            return false;
        }
        table[val]=arr.size();
        arr.push_back(val);
        return true;
    }
    
    bool remove(int val) {
        if(table.count(val)==0){
            return false;
        }
        // 删除val，基于vector只能pop_back 交换arr最后一个元素，然后直接arr.pop_back
        int idx=table[val];// idx of vector
        int lastVal=arr.back();
        table[lastVal]=idx;
        arr[idx]=lastVal;
        table.erase(val);
        arr.pop_back();
        return true;
    }
    
    int getRandom() {
        return arr[rand()%arr.size()];
    }
private:
    unordered_map<int,int>table;// key is val value is idx of arr
    vector<int> arr;
};

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet* obj = new RandomizedSet();
 * bool param_1 = obj->insert(val);
 * bool param_2 = obj->remove(val);
 * int param_3 = obj->getRandom();
 */
