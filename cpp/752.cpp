class Solution {
    /*
    题目：你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
题解：
从0000开始，可以想象成一棵树，每一组4位数字，都有8种变化如0000:
第一位：10000和90000
第二位：0100和0900
第三位：0010和0090
第四位：0001和0009
存在可能重复的 4 位数字，所以我们需要记录每组数字的访问情况
由于题目有一个「死亡数字」，所以我们遇到「死亡数字」，需要跳过
综上：我们可以很容易知道剪枝的情况，即：
访问「访问过的数字」时，跳过
访问「死亡数字」时，跳过
    https://leetcode.cn/problems/open-the-lock/solutions/1431738/by-lfool-9tbh/?page=2
    */
public:
    int openLock(vector<string>& deadends, string target) {
        // 记录死亡数字 方便快速判断
        unordered_set<string> hash;
        for(int i=0;i<deadends.size();i++){
            hash.insert(deadends[i]);
        }
        // 已知道起点0000和终点target，求最短次数BFS
        queue<string> qu;//push pop front back
        unordered_set<string> visited;//记录访问情况 避免出现重复数字
        int step=0;
        qu.push("0000");
        visited.insert("00000");
        while(qu.size()>0){
            int sz=qu.size();
            // 一层一层的遍历
            for(int i=0;i<sz;i++){
                auto cur=qu.front();
                qu.pop();
                // 如果是死亡数字 跳过
                if(hash.count(cur)){
                    continue;
                }
                // 找到目标，结束
                if(cur==target){
                    return step;
                }
                // 存储孩子节点到队列中，由于有4个故分4次循环，每个有两种情况一种是向上转+1，一种是向下转-1
                for(int j=0;j<4;j++){
                    // +1
                    auto addOne=increaseOne(cur,j);
                    // -1
                    auto minusOne=decreaseOne(cur,j);
                    // 如果没有访问过
                    if(visited.count(addOne)==0){
                        qu.push(addOne);
                        visited.insert(addOne);
                    }
                    if(visited.count(minusOne)==0){
                        qu.push(minusOne);
                        visited.insert(minusOne);
                    }                   
                }
            }
            step++;
        }
        return -1;
    }
    string increaseOne(string s, int i){
        string res(s);
        if (s[i]=='9'){
            res[i]='0';
        }else{
            res[i]+=1;
        }
        return res;
    }
    string decreaseOne(string s,int i){
        string res(s);
        if(s[i]=='0'){
            res[i]='9';
        }else{
            res[i]-=1;
        }
        return res;
    }
};
