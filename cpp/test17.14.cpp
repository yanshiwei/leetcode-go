class Solution {
public:
    vector<int> smallestK(vector<int>& arr, int k) {
        /*
        sort+k
        sort(arr.begin(),arr.end());
        return {arr.begin(),arr.begin()+k};
        */
        if(arr.size()<1||k<1){
            return {};
        }
        priority_queue<int,vector<int>,less<int>>q;// big heap push/pop top size empty swap
        for(int i=0;i<arr.size();i++){
            if(q.size()==k){
                if(q.top()>arr[i]){
                    q.pop();
                    q.push(arr[i]);
                }
            }else{
                q.push(arr[i]);
            }
        }
        vector<int>res;
        while(q.empty()==false){
            res.push_back(q.top());
            q.pop();
        }
        return res;
    }
};
