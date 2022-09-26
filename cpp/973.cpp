class Solution {
    /*
    给定一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点，并且是一个整数 k ，返回离原点 (0,0) 最近的 k 个点。
    */
public:
    static bool cmp(const vector<int>& u, const vector<int>& v){
        return u[0]*u[0]+u[1]*u[1]<v[0]*v[0]+v[1]*v[1];
    }
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        sort(points.begin(),points.end(),cmp);
        return {points.begin(),points.begin()+k};
    }

};
