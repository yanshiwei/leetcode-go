class MinStack {
    /*
    定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)
    */
public:
    /** initialize your data structure here. */
    MinStack() {

    }
    
    void push(int x) {
        // 新元素存储到数据栈
        data.push(x);
        // 当前元素与之前最小元素比较，若辅助栈空或者当前元素更小则直接插入辅助栈
        // 如当前元素更大，则当前入到数据栈对应的最小值仍是之前元素
        if(min_data.empty()||x<min_data.top()){
            min_data.push(x);
        }else{
            min_data.push(min_data.top());
        }
    }
    
    void pop() {
        data.pop();
        min_data.pop(); // 对应辅助栈也相同pop
    }
    
    int top() {
        return data.top();
    }
    
    int min() {
        return min_data.top();
    }
private:
    stack<int> data; // 数据栈，存放栈的所有元素
    stack<int> min_data; // 辅助栈，存放栈的最小元素
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(x);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->min();
 */
