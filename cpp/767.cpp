typedef pair<char,int>pp;
class Solution {
    /*
    给定一个字符串 s ，检查是否能重新排布其中的字母，使得两相邻的字符不同。返回 s 的任意可能的重新排列。若不可行，返回空字符串 "" 。
    思路：
    哈希表+堆+贪心
    1. 哈希表存储每个次数出现的次数
    2. 大顶堆存储每个字符及当前剩余的次数，pop出堆顶后 下一个top就是第二多的
    3. 贪心算法：每次使用剩余次数最多的元素（也就是大顶堆堆顶元素），如果与上一次使用的字符相同则使用次数第二多的（也就是pop出堆顶后 下一个top）如不存在则无法构造 返回空直到堆为空。
    */
public:
    string reorganizeString(string s) {
        auto cmp=[](pp a,pp b){
            // 大根堆 less
            return a.second<b.second;
        };
        // heap push pop top size empty
        priority_queue<pp,vector<pp>,decltype(cmp)>q(cmp);// 构造大根堆，存储的是s中出现的字符和剩余的个数
        vector<int>cnt(26);// 统计字符出现次数
        for(auto &ch:s){
            cnt[ch-'a']++;
        }
        for(int i=0;i<26;i++){
            if(cnt[i]>0){
                q.push(make_pair(i+'a',cnt[i]));
            }
        }
        string res="";
        while(q.empty()==false){
            auto p=q.top();
            q.pop();
            if(res.size()<1||p.first!=res.back()){
                // 第一次操作或剩余最多的元素与刚贴上去的字符不同
                res.push_back(p.first);
                // 消耗了一次更新次数
                if(p.second-->1){
                    q.push(p);
                }
            }else{
                // 本次最多的元素和上一次是同一个字符，那么就用第二多的字符
                if(q.empty()){
                    return "";
                }
                auto p1=q.top();//第二多的字符
                q.pop();
                res.push_back(p1.first);
                // 消耗了一次更新次数
                if(p1.second-->1){
                    q.push(p1);
                }
                // p没有使用重新入堆
                q.push(p);
            }
        }
        return res;
    }
};
