class Solution {
    /*
有趣的切入点：
这道题难点在于理解。天际线这一情景十分抽象，我们不如将这道题转化为另一个情景--物种统治问题：假设[lefti, righti, heighti] 存放的是某一物种的出现、灭绝时间和能力值，某一时刻存活的能力值最高的物种统治世界。要求得每次统治的开始时间和统治物种的能力值。
这样就很容易思考了：首先用（time，+/-ability）来记录每个物种的出现、灭绝时间（正负号用于区分是出现时间还是灭绝时间），并按照时间顺序排序。维持一个大顶堆，当一个物种出现时，将对应的能力值丢入大顶堆；当一个物种灭绝时，将对应的能力从大顶堆中拿出。这样，大顶堆就时刻维持着当前统治的物种的能力值。进而求得每次统治的开始时间和对应能力值。
默认存在一个能力值为0，存在时间无限的憨憨物种兜底。
故这道题关键是如何排序然后一一处理。我们可以先记录下 buildings中所有的左右端点横坐标及高度，并根据端点横坐标进行从小到大排序，以buildings  [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8]为例子，
参考：https://leetcode.cn/problems/the-skyline-problem/solutions/873171/gong-shui-san-xie-sao-miao-xian-suan-fa-0z6xc/ 给图的扫描仪线，
根据 buildings 求出每个 building 的左上角和右上角坐标
将所有坐标按照 x 排序, 并标记当前坐标是左上角坐标还是右上角坐标
l(2,10) l(3,15) l(5,12) r(7,15) r(9,10) r(12,12) 
l(15,10) l(19,8) r(20,10) r(24,8)
PriorityQueue = {0}, preMax = 0

l(2,10) 将 10 加入优先队列
preMax = 0, PriorityQueue  = {0 10}
当前 PriorityQueue 的 max = 10, 相对于 preMax 更新了
将 (2,10) 加入到 res, res = {(2,10)}
更新 preMax = 10
    
l(3,15) 将 15 加入优先队列
preMax = 10, PriorityQueue  = {0 10 15}
当前 PriorityQueue 的 max = 15, 相对于 preMax 更新了（最大值更新了，说明称霸状态有了变化）
将 (3,15) 加入到 res, res = {(2,10) (3,15)}
更新 preMax = 15    
    
l(5,12) 将 12 加入优先队列
preMax = 15, PriorityQueue  = {0 10 15 12}
当前 PriorityQueue 的 max = 15, 相对于 preMax 没有更新（最大值没有更新了，说明称霸状态没有变化）
res 不变

r(7,15) , 遇到右上角坐标,说明该物种统计时间结束， 将 15 从优先队列删除
preMax = 15, PriorityQueue  = {0 10 12}
当前 PriorityQueue 的 max = 12, 相对于 preMax 更新了（最大值更新了，说明称霸状态有了变化）
将 (7,max) 即 (7,12) 加入到 res, res = {(2,10) (3,15) (7,12)}
更新 preMax = 12
    
r(9,10) , 遇到右上角坐标, 将 10 从优先队列删除
preMax = 12, PriorityQueue  = {0 12}
当前 PriorityQueue 的 max = 12, 相对于 preMax 没有更新（最大值没有更新了，说明称霸状态没有变化）
res 不变

r(12,12) , 遇到右上角坐标, 将 12 从优先队列删除
preMax = 12, PriorityQueue  = {0}
当前 PriorityQueue 的 max = 0, 相对于 preMax 更新了（最大值更新了，说明称霸状态有了变化）
将 (12,max) 即 (12,0) 加入到 res, res = {(2,10) (3,15) (7,12) (12,0)}
更新 preMax = 0
    
后边的同理，就不进行下去了。

    两个坐标比较时候,不等时候直接按照x从小到大x；坐标x相等时候有以下情况：
    1 两个坐标都是左上角的坐标x，我们要将高度高的y排在前面(物种出现时间相同时候就看能力大小，能力越大越容易称霸，故越靠前)；
    2 两个坐标都是右上角的坐标x，我们要将高度低的y排在前面（物种灭绝时间相同时候，能力越小越容易灭绝故放在前面）
    3 两个坐标一个是左上角，一个是右上角，我们要将左上角排在前面(例如(1,2,1),(2,3,2)这2个，左上角有l(1,1),l(2,2),右上角有r(2,1),r(3,2)，此时l(2,2)和r(2,1)x坐标相等，排序后应该是：l(1,1),r(2,1),l(2,2),r(3,2))
    为了汇总以上情况，存左上角坐标的时候， 将高度（y）存为负数。存右上角坐标的时候，将高度（y）存为正数。这样有2个好处：
    1 可以根据高度的正负数区分当前是左上角坐标还是右上角坐标
    2 比较器可以统一实现上面逻辑（对于上面的1和2容易理解，对于3：两个坐标一个是左上角，一个是右上角，我们要将左上角排在前面，由于此时左上角的高度为负值而右上角为正值，故当左上角排在前面时候也满足高度从小到大）
    */
public:
    vector<vector<int>> getSkyline(vector<vector<int>>& buildings) {
        if(buildings.size()<1){
            return {};
        }
        // 先记录buildings中所有点，并根据横坐标进行从小到大的排序
        vector<pair<int,int>>sortedBuilds;
        for(auto&building:buildings){
            int left=building[0];
            int right=building[1];
            int height=building[2];
            sortedBuilds.push_back(make_pair(left,-height));//正负号用于区分是出现时间还是灭绝时间
            sortedBuilds.push_back(make_pair(right,height));
        }
        sort(sortedBuilds.begin(),sortedBuilds.end(),[](pair<int,int>&a,pair<int,int>&b){
            int x1=a.first;
            int x2=b.first;
            // 两个坐标比较时候,不等时候直接按照x从小到大
            if(x1!=x2){
                return x1<x2;
            }
            int y1=a.second;
            int y2=b.second;
            // 坐标相等，则按照高度从小到大进行处理
            return y1<y2;
        });
        priority_queue<int>maxHeap;
        // 记录进行了删除操作的高度，以及删除次数，方便删除队列元素
        unordered_map<int,int>hash;
        int preMax=0;
        maxHeap.push(preMax);//因为题目本身要求我们把一个完整轮廓的「右下角」那个点也取到，所以需要先添加一个 0
        vector<vector<int>>res;
        for(auto &b:sortedBuilds){
            int x=b.first;
            int y=b.second;
            if(y<0){
                //如果是左端点，说明存在一条往右延伸的可记录的边，将高度存入优先队列
                maxHeap.push(-y);
            }else{
                // 如果是右端点，说明这条边结束了，将当前高度从队列中移除
                // maxHeap没有删除方法，使用「哈希表」记录「执行了删除操作的高度」及「删除次数」，在每次使用前先检查堆顶高度是否已经被标记删除，如果是则进行pop 操作，并更新删除次数，如此达到删除队列元素的目的。
                if(hash.count(y)>0){
                    hash[y]++;
                }else{
                    hash[y]=1;
                }
            }
            // 先经过 O(n) 的复杂度进行查找
            while(!maxHeap.empty()){
                int cur=maxHeap.top();
                if(hash.count(cur)){
                    // 要删除
                    if(hash[cur]==1){
                        hash.erase(cur);
                    }else{
                        hash[cur]--;
                    }
                    maxHeap.pop();
                }else{
                    break;
                }
            }
            
            int cur=maxHeap.top();
            if(cur!=preMax){
                vector<int>tmp;
                tmp.push_back(x);
                tmp.push_back(cur);
                res.push_back(tmp);
                preMax=cur;
            }
        }
        return res;
    }
};
