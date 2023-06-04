class SortedStack {
public:
    // 辅助栈st2可以保证每次插入st1新元素的时候s1都是有序的
    SortedStack() {
    }
    
    void push(int val) {
        // 直到st1栈顶元素>=val
        while(!st1.empty()&&st1.top()<val){
            st2.push(st1.top());
            st1.pop();
        }
        st1.push(val);
        // st2临时栈内容放入st1
        while(!st2.empty()){
            st1.push(st2.top());
            st2.pop();
        }
    }
    
    void pop() {
        if(!st1.empty()){
            st1.pop();
        }
    }
    
    int peek() {
        if(!st1.empty()){
            return st1.top();
        }
        return -1;
    }
    
    bool isEmpty() {
        return st1.empty();
    }
private:
    stack<int>st1;
    stack<int>st2;
};

/**
 * Your SortedStack object will be instantiated and called as such:
 * SortedStack* obj = new SortedStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->isEmpty();
 */
