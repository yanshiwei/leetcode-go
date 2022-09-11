class MyStack {
    /*
    请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）
    */
public:
    //使用队列实现栈时，应满足队列前端的元素是最后入栈的元素
    queue<int>q1; // 存储栈内元素
    queue<int>q2; // 入栈操作的辅助,总是首先存储最后入的元素 保证队列前端的元素是最后入的元素
    MyStack() {

    }
    // 入栈时，当前元素先入q2，然后从非空q1 从q1头部出队列到q2此时q2前端元素即新入栈的元素
    // 然后交换q1 q2  则q1的元素即为栈内的元素，q1的前端和后端分别对应栈顶和栈底。
    // key：每次入栈时，确保q1的前端元素为最近的元素即 栈顶元素
    void push(int x) {
        q2.push(x);// tail of q2
        while(q1.empty()==false){
            q2.push(q1.front());// head of q1
            q1.pop();// head of q1
        }
        swap(q1,q2);// store ele of stack
    }
    // 出栈时，假设q1非空，则从q1队头取元素 这个元素就是栈顶元素；
    int pop() {
        int cur=q1.front();
        q1.pop();
        return cur;
    }
    
    int top() {
        int cur=q1.front();
        return cur;
    }
    // 由于q1存储栈内的元素，判断栈是否为空时，只需要判断q1是否为空即可
    bool empty() {
        return q1.empty();
    }
};

/**
 * Your MyStack object will be instantiated and called as such:
 * MyStack* obj = new MyStack();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->top();
 * bool param_4 = obj->empty();
 */
