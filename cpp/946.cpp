class Solution {
    /*
    给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 
    */
public:
    bool validateStackSequences(vector<int>& pushed, vector<int>& popped) {
        if(pushed.size()<1){
            return popped.size()==0;
        }
        if(pushed.size()!=popped.size()){
            return false;
        }
        stack<int>st;
        int idx=0;
        for(int i=0;i<pushed.size();i++){
            st.push(pushed[i]);
            while(st.empty()==false&&idx<popped.size()&&st.top()==popped[idx]){
                st.pop();
                idx++;
            }
        }
        return st.empty();
    }
};
