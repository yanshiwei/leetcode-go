class Solution {
    /*
    题目：
    按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
每对相邻的单词之间仅有单个字母不同。
转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
sk == endWord
给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。
题解：
https://leetcode.cn/problems/word-ladder-ii/solutions/277612/yan-du-you-xian-bian-li-shuang-xiang-yan-du-you--2/
    */
public:
    //存放最短路径
    unordered_map<string,int>table;
    //存放通过相邻单词构建的有向图
    unordered_map<string,unordered_set<string>>graph;
    vector<vector<string>>ret;
    //存储回溯算法中的单条路径，用来加入ret，其中元素的顺序为倒序（从endWord到beginWord）
    vector<string>path;
    //用全局变量存储beginWord，减少传参个数
    string target;
    //思路：BFS找最短路同时建图存储节点信息+DFS回溯算法找到所有命中起终点的最短路径
    void dfs(string s){
        if(s==target){
            //因为path中元素为倒序，所以想要颠倒过来需要用rbegin, rend
            ret.emplace_back(path.rbegin(),path.rend());
            return;
        }
        for(auto node:graph[s]){
            path.emplace_back(node);
            dfs(node);
            path.pop_back();
        }
    }
    vector<vector<string>> findLadders(string beginWord, string endWord, vector<string>& wordList) {
        unordered_set<string>dict(wordList.begin(),wordList.end());
        if(dict.empty()||dict.count(endWord)==0){
            return ret;
        }
        dict.erase(beginWord);
        queue<string>qu;
        qu.push(beginWord);
        //在开始需要给beginWord设置初值0
        table[beginWord]=0;
        int n=beginWord.size();//字符串长度
        while(qu.empty()==false){
            auto p=qu.front();
            qu.pop();
            //对每个单词的每个字母进行循环，对每个字母都尝试a-z所有组合，来查找是否符合本题条件
            for(int i=0;i<n;i++){
                auto nextWord=p;
                //让p作为当前节点，nextWord作为下一个节点去探索新的字符组合
                for(auto c='a';c<='z';c++){
                    nextWord[i]=c;
                    //对于每一种组合，存在于wordList并且从未遍历过或者跟当前单词相邻（为了找到所有相邻节点）
                    if(dict.count(nextWord)&&(!table.count(nextWord)||table[nextWord]==table[p]+1)){
                        //graph反向插入，有向边的顺序是从endword到beginword
                        graph[nextWord].emplace(p);
                    }
                    //如果是存在于wordList并且从未遍历，那么就在最短路径中添加当前节点的相邻节点信息
                    if(dict.count(nextWord)&&!table.count(nextWord)){
                        table[nextWord]=table[p]+1;
                        //如果正好遍历到了endWord，那么就终止循环剪枝
                        if(nextWord==endWord){
                            break;
                        }
                        qu.push(nextWord);
                    }
                }
            }
        }
        //如果没找到，返回空vector
        if(!table.count(endWord)){
            return ret;
        }
        //开始dfs寻找命中起终点的最短路径，记住要先把endWord加进去（因为需要从尾到头找）
        path.emplace_back(endWord);
        target=beginWord;
        dfs(endWord);
        return ret;
    }
};
