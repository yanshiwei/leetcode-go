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
        unordered_set<string> hash(wordList.begin(),wordList.end());//insert erase counf
        if(hash.empty()||hash.count(endWord)==0){
            return 0;
        }
        // bfs
        hash.erase(beginWord);// 遍历完
        queue<string> que;// push pop front back
        que.push(beginWord);
        unordered_set<string> visited;
        visited.insert(beginWord);
        int step=1;
        while(que.empty()==false){
            // 找到没有被访问过, 而且能够由当前单词转换而成的单词
            int n=que.size();
            // 每一轮(每一层step加个1)
            while(n--){
                string curWord=que.front();
                que.pop();
                // 当前单词的每个字符都替换成其他的25个字符, 然后在单词表中查询是不是包含转换后的单词
                // 这里千万不能遍历单词表, 因为单词表很长, 而哈希表使用的红黑树查询效率比遍历单词表高很多
                for(int i=0;i<curWord.size();i++){
                    char originalChar =curWord[i];// 当前单词的每个字符
                    // 当前单词的每个字符都替换成其他的25个字符
                    for(int j=0;j<26;j++){
                        if(char('a'+j)==originalChar){
                            continue;
                        }
                        curWord[i]=char('a'+j);
                        if(hash.count(curWord)&&visited.count(curWord)==0){
                            if(curWord==endWord){
                                return step+1;
                            }else{
                                que.push(curWord);
                                visited.insert(curWord);
                            }
                        }
                    }
                    curWord[i]=originalChar;// recovery
                }
            }
            step++;
        }
        return 0;
    }
};
