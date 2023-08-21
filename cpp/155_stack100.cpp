class MinStack {
public:
    stack<int>minV;
    stack<int>data;
    MinStack() {

    }
    
    void push(int val) {
        data.push(val);
        if(minV.empty()||minV.top()>=val){
            // 注意要等于也要入，否则出错
            minV.push(val);
        }
    }
    
    void pop() {
        if(data.top()==minV.top()){
            minV.pop();
        }
        data.pop();
    }
    
    int top() {
        return data.top();
    }
    
    int getMin() {
        return minV.top();
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
