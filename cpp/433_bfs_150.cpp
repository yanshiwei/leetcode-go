class Solution {
    /*
    https://leetcode.cn/problems/minimum-genetic-mutation/solutions/1474617/by-fuxuemingzhu-t1mv/?envType=study-plan-v2&envId=top-interview-150
    给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数
    最少次数，单点BFS
    */
public:
    int minMutation(string startGene, string endGene, vector<string>& bank) {
        unordered_set<string>bank_set(bank.begin(),bank.end());
        queue<string>qu;//base deque list
        unordered_set<string>visited;
        qu.push(startGene);
        visited.insert(startGene);
        int step=0;
        while(!qu.empty()){
            int cur_size=qu.size();
            for(int i=0;i<cur_size;i++){
                string cur=qu.front();
                qu.pop();
                if(cur==endGene){
                    return step;
                }
                // 尝试4个字符
                for(char c: "ACGT"){
                    for(int j=0;j<cur.size();j++){
                        string next=cur;
                        next[j]=c;
                        // 在字典里且还没有被访问
                        if(bank_set.count(next)&&!visited.count(next)){
                            visited.insert(next);
                            qu.push(next);
                        }
                    }
                }
            }
            step++;
        }
        return -1;
    }
};
