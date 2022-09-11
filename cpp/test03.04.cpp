class MyQueue {
public:
/*
实现一个MyQueue类，该类用两个栈来实现一个队列。
*/
    stack<int>st1;
    stack<int>st2;
    /** Initialize your data structure here. */
    MyQueue() {

    }
    
    /** Push element x to the back of queue. */
    void push(int x) {
        st1.push(x);
    }
    
    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        if(st2.empty()){
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
    
    /** Get the front element. */
    int peek() {
        if(st2.empty()){
            while(st1.empty()==false){
                int cur=st1.top();
                st1.pop();
                st2.push(cur);
            }
        }
        int cur=st2.top();
        return cur;
    }
    
    /** Returns whether the queue is empty. */
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
