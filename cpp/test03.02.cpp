class MinStack {
    /*请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)
    */
public:
    /** initialize your data structure here. */
    MinStack() {

    }
    
    void push(int x) {
        data.push(x);
        // 当前元素与之前最小元素比较，若辅助栈空或者当前元素更小则直接插入辅助栈
        // 如当前元素更大，则当前入到数据栈对应的最小值仍是之前元素
        if(min.empty()||x<min.top()){
            min.push(x);
        }else{
            min.push(min.top());
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
 * obj->push(x);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
