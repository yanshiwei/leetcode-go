class BrowserHistory {
    /*
    你有一个只支持单个标签页的 浏览器 ，最开始你浏览的网页是 homepage ，你可以访问其他的网站 url ，也可以在浏览历史中后退 steps 步或前进 steps 步
    */
public:
    BrowserHistory(string homepage) {
        his=new string[101];
        his[++idx]=homepage;
        end++;
    }
    ~BrowserHistory(){
        delete []his;
    }
    void visit(string url) {
        //当前页跳转访问 url 对应的页面。执行此操作会把浏览历史前进的记录全部删除
        his[++idx]=url; // 覆盖idx之后的
        end=idx;
    }
    
    string back(int steps) {
        // 在浏览历史中后退 steps 步
        idx=idx>steps?idx-steps:0;
        return his[idx];
    }
    
    string forward(int steps) {
        idx=idx+steps<end?idx+steps:end;
        return his[idx];
    }
private:
    string *his;
    int idx=-1;
    int end=-1;
};

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * BrowserHistory* obj = new BrowserHistory(homepage);
 * obj->visit(url);
 * string param_2 = obj->back(steps);
 * string param_3 = obj->forward(steps);
 */
