class Solution {
    // 题目：给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。
    // 连接词 定义为：首先是words里，然后一个完全由给定数组中的至少两个较短单词（不一定是不同的两个单词）组成的字符串。
    // https://leetcode.cn/problems/concatenated-words/solutions/1180231/rustgolangjava-dpjie-by-kyushu-5fzz/
    // dp[i]表示word从[0, i)之间的字符串是否可以分解成Set集合已经包含的较短的字符串，dp[n]就表示可以完全分解掉；
    // dp[i]找出由dp[j]表示的转换公式，j<i
    // 也就对应着DP数组的两部分：word[0, j - 1]和word[j, i-1],
    // word[0, j - 1]对应dp[j], word[j, i-1]需要判断下Set集合中是否包含word[j, i-1]，如果包含则dp[i] = true
    // 其中0<=j<i
    // dp[j]不行时，dp[i]肯定也不行；
    // dp[j]可以时，只要找到一种word[j, i-1]满足在集合要求，那么整体dp[i]就满足，其中0<=j<i。
public:
    vector<string> findAllConcatenatedWordsInADict(vector<string>& words) {
        unordered_set<string>set;
        for(auto& word:words){
            set.insert(word);// 去重
        }
        vector<string>res;
        for(auto&word:words){
            if(set.count("")){
                // 空串不处理
                continue;
            }
            // 对单词数组words进行遍历，肯定先要把自己word移除掉，因为不能自己构成自己
            set.erase(word);
            if(canBreak(word,set)){
                // 是否能由目前Set(以不包含word)中的元素构成当前word
                res.push_back(word);
            }
            // 处理完word后又放回，进行下一轮处理
            set.insert(word);
        }
        return res;
    }
private:
    // 是否能由目前Set(以不包含word)中的元素构成当前word
    bool canBreak(string word,unordered_set<string>&set){
        int n=word.size();
        vector<bool>dp(n+1,false);// dp[i]表示word从[0, i)之间的前i字符串是否可以分解成Set集合已经包含的较短的字符串
        dp[0]=true;// 空字符串可以分解成Set集合已经包含的较短的字符串，“”
        for(int i=1;i<=n;i++){
            for(int j=0;j<i;j++){
                if(dp[j]==false){
                    // dp[j]不行时，dp[i]肯定也不行；
                    continue;
                }
                // word[j, i-1]需要判断Set集合中是否包含word[j, i-1]
                if(set.count(word.substr(j,i-j))){
                    dp[i]=true;
                    break;//只要找到一种word[j, i-1]满足在集合要求，那么整体dp[i]就满足，其中0<=j<i
                }
            }
        }
        return dp[n];
    }
};
