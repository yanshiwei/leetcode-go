class Solution {
    /*
    给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
    */
public:
    int maxProfit(vector<int>& prices) {
        if(prices.size()<2){
            return 0;
        }
        int maxP=0x80000000;
        int minPrice=prices[0];
        for(int i=1;i<prices.size();i++){
            //第二天开始才能卖
            if(minPrice>prices[i]){
                minPrice=prices[i];
            }
            if(maxP<prices[i]-minPrice){
                maxP=prices[i]-minPrice;
            }
        }
        return maxP;
    }
};
