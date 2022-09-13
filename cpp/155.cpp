class MinStack {
    /*
    设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈
    */
public:
    MinStack() {

    }
    
    void push(int val) {
        data.push(val);
        // 当前元素与之前最小元素比较，若辅助栈空或者当前元素更小则直接插入辅助栈
        // 如当前元素更大，则当前入到数据栈对应的最小值仍是之前元素
        if(min.empty()||val<min.top()){
            min.push(val);
        }else{
            min.push(min.top()); // 与数据栈一一对应
        }
    }
    
    void pop() {
        data.pop();
        min.pop();
    }
    
    int top() {
        return data.top();
    }
    
    int getMin() {
        return min.top();
    }
private:
    stack<int> data;// 数据栈，存放栈的所有元素
    stack<int> min;// 辅助栈，存放栈的最小元素 每一位对应数据栈该元素的最小值
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
