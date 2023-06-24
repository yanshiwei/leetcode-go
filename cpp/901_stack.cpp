class StockSpanner {
    /*
    题目：
    设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。
当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。
例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。
    题解：
    股票价格小于或等于今天价格的最大连续日数
    也就是左侧第一个left大于当天i的间隔i-left
    单调递增栈
    */
    struct Info {
        int idx;
        int price;
    };
public:
    stack<Info>st;
    int idx;
    StockSpanner() {
        idx=0;
    }
    
    int next(int price) {
        int res=1;
        Info info=Info{
            idx:idx,
            price:price
        };
        if(st.empty()||st.top().price>price){
            st.push(info);
        }else{
            while(!st.empty()&&st.top().price<=price){
                Info cur=st.top();
                st.pop();
            }
            int fidx=st.empty()?-1:st.top().idx;
            res=idx-fidx;
            st.push(info);
        }
        idx++;
        return res;
    }
};

/**
 * Your StockSpanner object will be instantiated and called as such:
 * StockSpanner* obj = new StockSpanner();
 * int param_1 = obj->next(price);
 */
