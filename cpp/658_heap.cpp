class Solution {
public:
    vector<int> findClosestElements(vector<int>& arr, int k, int x) {
        // build max heap
        priority_queue<pair<int,int>>maxheap;
        for(int i=0;i<arr.size();i++){
            // 维护大顶堆的的大小为K
            if(maxheap.size()<k){
                maxheap.push({abs(arr[i]-x),i});
            }else{
                if(abs(arr[i]-x)<maxheap.top().first){
                    maxheap.pop();
                    maxheap.push({abs(arr[i]-x),i});
                }
            }
        }
        vector<int> res;
        while(!maxheap.empty()){
            res.push_back(arr[maxheap.top().second]);
            maxheap.pop();
        }
        sort(res.begin(),res.end());// 返回的结果必须要是按升序排好的
        return res;
    }
};
