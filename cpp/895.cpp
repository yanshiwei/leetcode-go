class FreqStack {
    /*
    设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。
实现 FreqStack 类:
FreqStack() 构造一个空的堆栈。
void push(int val) 将一个整数 val 压入栈顶。
int pop() 删除并返回堆栈中出现频率最高的元素。
如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素
    */
public:
    FreqStack() {
        // vector push_back pop_back front back
    }
    
    void push(int val) {
        num_freq[val]+=1;
        freq_nums[num_freq[val]].push_back(val);
        max_fre=max(max_fre,num_freq[val]);
    }
    
    int pop() {
        int num=freq_nums[max_fre].back();// 最近
        // update num_freq
        num_freq[num]-=1;
        // update freq_nums
        freq_nums[max_fre].pop_back();
        // update max_fre
        if(freq_nums[max_fre].empty()){
            max_fre--;
        }
        return num;
    }
private:
    unordered_map<int,int>num_freq;// key is数值，value is 数值出现的次数
    unordered_map<int,vector<int>>freq_nums;// key is数值出现的次数，value is 出现这么多次的数的vector
    int max_fre=-1;// 出现最多的次数
};

/**
 * Your FreqStack object will be instantiated and called as such:
 * FreqStack* obj = new FreqStack();
 * obj->push(val);
 * int param_2 = obj->pop();
 */
