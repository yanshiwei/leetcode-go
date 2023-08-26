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
    dp[i][j]=fun(dp[i-1][j],dp[i-1][j-k*t]+k*v),0<k*t<=j
    转为dp[i][j]=fun(dp[i-1][j],dp[i][j-t]+v)
    降维dp[j]=fun(dp[j],dp[j-t]+v)
    // 考虑顺序，组合不强调顺序、排列强调顺序！如果强调排列，我们应该先遍历背包，再遍历物品，只有这样才能让物品按一定顺序放入背包中
    // 本题讲究顺序，故先额度，再物品。
    // 背包里的总额是目标字符串长度，物品是wordDict中的单词，物品的重量是单词长度，没有价值
    // 根据完全背包问题模版：外层循环为 target ，内层循环为选择池 wordDict.https://mp.weixin.qq.com/s/nke1OjkhKACaONx1opk8AA
    //  dp[i] 表示长度i的字符串是否可以被 wordDict 中组合而成
    //  s的一定长度作为target，且s的这部分长度要等于当前 wordDict 中的词组合
    //  wordDict 中的词作为包
    */
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        vector<bool>dp(s.size()+1,false);//dp[i] 表示长度i的字符串是否可以被 wordDict 中组合而成
        dp[0]=true;//s空串且合法
        // 本题讲究顺序，故先背包，再物品
        //原始dp[j]=fun(dp[j],dp[j-t]+v)
        //这里存在性dp[j]=dp[j]||dp[j-t],v=0
        for(int j=1;j<=s.size();j++){
            for(string word:wordDict){
                // 物品的重量是单词长度
                int w=word.size();
                //. 若word没有参与情况下直接看之dp[j];
                //. 若word参与情况下，则需要j>=size&&s.substr(j-size,size)==word，此时就要看dp[j-size]
                if(j>=w&&s.substr(j-w,w)==word){
                    dp[j]=dp[j]||dp[j-w];
                }
            }
        }
        return dp[s.size()];
    }
};
