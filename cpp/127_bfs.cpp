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
    int ladderLength(string beginWord, string endWord, vector<string>& wordList){
        // 第 1 步：先将 wordList 放到哈希表里，便于判断某个单词是否在 wordList 里
        unordered_set<string>hash(wordList.begin(),wordList.end());
        if(hash.empty()||hash.count(endWord)==0){
            return 0;
        }
        hash.erase(beginWord);//删除起始结点
        // 第 2 步：图的广度优先遍历，初始化队列和以访问hash
        queue<string>qu;//push pop front back
        qu.push(beginWord);
        unordered_set<string>visited;
        visited.insert(beginWord);
        int step=1;//所需步骤
        // 第 3 步：开始广度优先遍历，包含起点，因此初始化的时候步数为 1
        while(!qu.empty()){
            int curSize=qu.size();
            for(int l=0;l<curSize;l++){
                string curWord=qu.front();
                qu.pop();
                // 当前单词的每个字符都替换成其他的25个字符, 然后在单词表中查询是不是包含转换后的单词
                for(int i=0;i<curWord.size();i++){
                    char originalChar=curWord[i];
                    // 当前单词的每个字符都替换成其他的25个字符
                    for(int j=0;j<26;j++){
                        if(originalChar==char('a'+j)){
                            continue;//本身
                        }
                        curWord[i]=char('a'+j);
                        if(hash.count(curWord)&&visited.count(curWord)==0){
                            // 命中
                            if(curWord==endWord){
                                return step+1;//+1表示endWord
                            }else{
                                qu.push(curWord);
                                visited.insert(curWord);
                            }
                        }
                    }
                    curWord[i]=originalChar;//恢复，保证不影响最终结果
                }
            }
            step++;//一轮结束就+1
        }
        return 0;
    }
};
