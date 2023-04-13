class Solution {
    /*
    题目：
    给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
    转化为是否可以用 wordDict 中的词组合成 s，由于转化为是否可以用 wordDict 中的词组合成 s，完全背包问题，所以是完全背包问题
    // 考虑顺序，组合不强调顺序、排列强调顺序！如果强调排列，我们应该先遍历背包，再遍历物品，只有这样才能让物品按一定顺序放入背包中
    // 根据完全背包问题模版：外层循环为 target ，内层循环为选择池 wordDict.https://mp.weixin.qq.com/s/nke1OjkhKACaONx1opk8AA
    //  dp[i] 表示长度i的字符串是否可以被 wordDict 中组合而成
    // s的一定长度作为target，且s的这部分长度要等于当前 wordDict 中的词组合
    //  wordDict 中的词作为包
    */
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        vector<bool>dp(s.size()+1);
        dp[0]=true;//空串且合法
        for(int i=1;i<=s.size();i++){
            for(auto&word:wordDict){
                int size=word.size();
                if(i>=size&&s.substr(i-size,size)==word){
                    dp[i]=dp[i]||dp[i-size];
                }
            }
        }
        return dp[s.size()];
    }
};
