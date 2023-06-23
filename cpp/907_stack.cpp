class Solution {
    /*
    题目：
    给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
由于答案可能很大，因此 返回答案模 10^9 + 7 。
    题解：
    暴力方法：枚举所有子数组，然后逐个求最小值
    进一步，可以求最小值为arr[i]的区间个数，也就是遍历每一个值，求出以它为最小值的区间数目。
    如何求其个数呢，方法是找到左右两边最近的更小值的位置left,right
    那么在(left，i]选择区间左边，[i，right)选择区间右边，
    那么这时候组合的子序列肯定都是以arr[i]为最小值
    求arr[i]的下一个最小值么，单调递减栈.
    例如arr = [11,81,94,43,3]：
    1.对于11，左边没有更小值，右边更小值为3，那么，只能以11为起点，终点可以选11，81，94，43，共4个数组(左区间1个，右区间4个)，对答案的贡献就是4*11=44
    2.对于81，左边更小值为11，右边更小值为43，那么，可以以81为起点，81，94为终点，共两个数组(左区间1个，右区间2个)，对答案的贡献为2*81=162
    3.对于94，左边更小值为81，右边更小值为43，只能以自身为起点终点，共一个数组(左区间1个，右区间1个)，贡献94
    4.对于43，左边更小值为11，右边更小值为3，可以以81，94，43为起点，43为终点，共三个数组(左区间3个，右区间1个)，对答案的贡献为3*43=129
    4.对于3，左边没有更小值，右边没有更小值，可以以11，81，94，43，3为起点，3为终点，共五个数组(左区间5个，右区间1个)，对答案的贡献为5*3=15
最终结果就是累加起来，44+162+94+129+15=444
    总结：如果当前位置坐标为i，左边更小值位置为left，右边为right，则贡献就是(i-left)*(right-i)*arr[i]。特别地，如果左边没有更小值，left=-1，如果右边没有更小值，right=size
    类似84题，关键是找到arr[i]左右两边第一个小于本身的位置，对应与单调递减栈。
        其实就是栈顶元素，和栈顶元素的前一个元素（左边第一个小于的元素）以及将要入栈的元素（右边第一个小于的元素）组成
    */
public:
    const int int_min=0xffffffff;
    int sumSubarrayMins(vector<int>& arr) {
        arr.push_back(int_min);//为了确保全部计算，包括最后一个元素
        stack<int>st;
        long long res=0;
        long long mod=1e9+7;
        for(int i=0;i<arr.size();i++){
            if(st.empty()||arr[st.top()]<arr[i]){
                st.push(i);
            }else{
                while(!st.empty()&&arr[st.top()]>=arr[i]){
                    int cur=st.top();
                    st.pop();
                    int left=st.empty()?-1 : st.top();
                    int right=i;
                    res=(res+1LL*(cur-left)*(right-cur)*arr[cur])%mod;
                }
                st.push(i);
            }
        }
        return res;
    }
};