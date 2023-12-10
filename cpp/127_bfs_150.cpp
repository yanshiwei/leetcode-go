class Solution {
    /*
    字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
每一对相邻的单词只差一个字母。
 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
https://leetcode.cn/problems/word-ladder/solutions/276923/yan-du-you-xian-bian-li-shuang-xiang-yan-du-you-2/
    */
    // 从beginWord开始转换，直到endWord，最终构成构成一个无向有环图。
    //                 hit
    //                  |
    //                 hot
    //                /   \
    //              dot -- lot
    //             /       |
    //            dog      |
    //            |  \    /
    //            |   log
    //            |  /
    //            cog
    // 无向图中两个顶点之间的最短路径的长度，可以通过广度优先遍历得到
    // 已知目标顶点的情况下，可以分别从起点和目标顶点（终点）执行广度优先遍历，
    // 直到遍历的部分有交集，这是双向广度优先遍历的思想。
public:
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        unordered_set<string>dict(wordList.begin(),wordList.end());
        queue<string>qu;//base deque list
        unordered_set<string>visited;
        qu.push(beginWord);
        visited.insert(beginWord);
        int step=1;// not 0，求的单词个数，不是转换次数，故最少是1，也就是beginWord
        while(!qu.empty()){
            int cur_size=qu.size();
            for(int i=0;i<cur_size;i++){
                string cur=qu.front();
                qu.pop();
                if(cur==endWord){
                    return step;
                }
                // 尝试26个字符
                for(char k=0;k<26;k++){
                    for(int j=0;j<cur.size();j++){
                        string next=cur;
                        next[j]=char('a'+k);
                        if(dict.count(next)&&!visited.count(next)){
                            qu.push(next);
                            visited.insert(next);
                        }
                    }
                }
            }
            step++;
        }
        return 0;
    }
};
