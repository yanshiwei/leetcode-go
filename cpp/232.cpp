class MyQueue {
    /*
    仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作
    */
public:
    stack<int>st1;
    stack<int>st2;
    MyQueue() {

    }
    
    void push(int x) {
        st1.push(x);
    }
    
    int pop() {
        if(st2.empty()){
            // 从st1 pop
            while(st1.empty()==false){
                int cur=st1.top();
                st1.pop();
                st2.push(cur);
            }
        }
        int cur=st2.top();
        st2.pop();
        return cur;
    }
    
    int peek() {
        if(st2.empty()){
            // 从st1 pop
            while(st1.empty()==false){
                int cur=st1.top();
                st1.pop();
                st2.push(cur);
            }
        }
        int cur=st2.top();
        return cur;        
    }
    
    bool empty() {
        return st1.empty()&&st2.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
