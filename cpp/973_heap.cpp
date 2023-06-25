class Solution {
    /*
    题目：
    给定一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点，并且是一个整数 k ，返回离原点 (0,0) 最近的 k 个点。
这里，平面上两点之间的距离是 欧几里德距离（ √(x1 - x2)2 + (y1 - y2)2 ）。
    题解：
    最小k个值，k大小的最大堆，当堆顶元素比当前元素大时，替换；当堆顶元素比当前元素小时跳过，最终堆里维护的k个最小
    */
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        if(k<1){
            return {};
        }
        priority_queue<pair<int,int>>maxHeap;// pair first is distace, second is idx, default is vector and less<int>
        for(int i=0;i<points.size();i++){
            if(i<k){
                // 构造大小k的堆
                maxHeap.push(make_pair(points[i][0]*points[i][0]+points[i][1]*points[i][1],i));
            }else{
                int distace=points[i][0]*points[i][0]+points[i][1]*points[i][1];
                if(maxHeap.top().first>distace){
                    maxHeap.pop();
                    maxHeap.push(make_pair(distace,i));
                }else{
                    continue;
                }
            }
        }
        vector<vector<int>>res;
        while(!maxHeap.empty()){
            int idx=maxHeap.top().second;
            res.push_back(points[idx]);
            maxHeap.pop();
        }
        return res;
    }
};
