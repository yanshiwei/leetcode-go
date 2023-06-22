class Solution {
    /*
    寻找a[i]下一个更大值，使用递增栈（从栈低到栈顶从大到小）
    */
    const int int_max=0x7fffffff;
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        vector<int>res(temperatures.size(),0);
        stack<int>st;
        //temperatures.push_back(int_max);//方便处理最后一个元素 不用处理最后元素，没有找到更大的直接就是默认值0
        for(int i=0;i<temperatures.size();i++){
            if(st.empty()||temperatures[st.top()]>=temperatures[i]){
                st.push(i);
            }else{
                while(!st.empty()&&temperatures[st.top()]<temperatures[i]){
                    // 遇到第一个更大
                    int idx=st.top();
                    st.pop();
                    res[idx]=i-idx;
                }
                st.push(i);
            }
        }
        return res;
    }
};
