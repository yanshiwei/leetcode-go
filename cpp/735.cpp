class Solution {
    /*
    给定一个整数数组 asteroids，表示在同一行的行星。
对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。
找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。
    */
public:
    vector<int> asteroidCollision(vector<int>& asteroids) {
        if(asteroids.size()<2){
            return {};
        }
        /*
        1 当前行星大于0，直接入
        2 当前行星小于0，细分：
        2.1 空栈 直接入
        2.2 栈顶小于0，表示栈内元素均左，直接入栈
        2.3 栈顶大于0，与当前行星比较：
            2.3.1相等，均爆炸
            2.3.2栈顶大，当前行星爆炸
            2.3.3栈顶小，栈顶元素爆炸出栈，接着判断下一个栈顶元素
        2.3总而言之就是只有栈顶元素小当前行星绝对值时或为空时，当前行星才存活
        */
        stack<int>st;
        for(int i=0;i<asteroids.size();i++){
            if(asteroids[i]>0){
                // 1 当前行星大于0，直接入
                st.push(asteroids[i]);
            }else{
                // 2 当前行星小于0，细分：
                bool alive=true; // 当前行星是否存活
                while(alive&&st.empty()==false&&st.top()>0){
                    //只有栈顶元素小于当前元素的绝对值，当前行星才会存活
                    alive=st.top()<abs(asteroids[i]);
                    //相等或栈顶元素小于时，栈顶元素爆炸（出栈）
                    if(st.top()<=abs(asteroids[i])){
                        st.pop();
                    }
                }
                // 为空栈或者栈顶元素<0
                if(alive){
                    st.push(asteroids[i]);
                }

            }
        }
        vector<int>res(st.size());
        int idx=st.size()-1;
        while(st.empty()==false){
            res[idx--]=st.top();
            st.pop();
        }
        return res;
    }
private:
    int abs(int a){
        if(a<0){
            return -1*a;
        }
        return a;
    }
};
